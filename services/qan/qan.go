// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Package qan contains business logic of working with QAN and qan-agent on PMM Server node.
package qan

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/percona/kardianos-service"
	"github.com/percona/pmm/proto"
	"github.com/percona/pmm/proto/config"
	"github.com/pkg/errors"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/services/supervisor"
	"github.com/percona/pmm-managed/utils/logger"
)

type Service struct {
	baseDir    string
	supervisor *supervisor.Supervisor
	qanAPI     *http.Client
	qanURL     url.URL
}

func NewService(ctx context.Context, baseDir string, supervisor *supervisor.Supervisor) (*Service, error) {
	u, err := getQanURL(ctx)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		baseDir:    baseDir,
		supervisor: supervisor,
		qanAPI:     new(http.Client),
		qanURL:     *u,
	}
	return svc, nil
}

// EnsureAgentIsRegistered is registers qan-agent running on PMM Server node in QAN.
// It does nothing if agent is already registered.
func (svc *Service) EnsureAgentIsRegistered(ctx context.Context) error {
	// do nothing is qan-agent already registered
	path := filepath.Join(svc.baseDir, "config", "agent.conf")
	if _, err := os.Stat(path); err == nil {
		logger.Get(ctx).Debugf("qan-agent already registered (%s exists).", path)
		return nil
	}

	path = filepath.Join(svc.baseDir, "bin", "percona-qan-agent-installer")
	args := []string{"-debug", "-hostname=pmm-server"}
	if svc.qanURL.User != nil && svc.qanURL.User.Username() != "" {
		args = append(args, "-server-user="+svc.qanURL.User.Username())
		pass, _ := svc.qanURL.User.Password()
		args = append(args, "-server-pass="+pass)
	}
	args = append(args, svc.qanURL.String()) // full URL, with username and password (yes, again! that's how installer is written)
	cmd := exec.Command(path, args...)
	logger.Get(ctx).Debug(strings.Join(cmd.Args, " "))
	b, err := cmd.CombinedOutput()
	if err != nil {
		logger.Get(ctx).Infof("%s", b)
		return errors.Wrap(err, "failed to register qan-agent")
	}
	logger.Get(ctx).Debugf("%s", b)

	// make logging more verbose than default
	path = filepath.Join(svc.baseDir, "config", "log.conf")
	return ioutil.WriteFile(path, []byte(`{"Level":"debug","Offline":"false"}`), 0666)
}

// getAgentUUID returns agent UUID from the qan-agent configuration file.
func (svc *Service) getAgentUUID() (string, error) {
	path := filepath.Join(svc.baseDir, "config", "agent.conf")
	f, err := os.Open(path)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer f.Close()

	var cfg config.Agent
	if err = json.NewDecoder(f).Decode(&cfg); err != nil {
		return "", errors.WithStack(err)
	}
	if cfg.UUID == "" {
		err = errors.Errorf("missing agent UUID in configuration file %s", path)
	}
	return cfg.UUID, err
}

// getOSUUID returns OS UUID from the QAN API.
func (svc *Service) getOSUUID(ctx context.Context, agentUUID string) (string, error) {
	url := svc.qanURL
	url.Path = path.Join(url.Path, "instances", agentUUID)
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return "", errors.WithStack(err)
	}
	rb, _ := httputil.DumpRequestOut(req, true)
	logger.Get(ctx).Debugf("getOSUUID request:\n\n%s\n", rb)

	resp, err := svc.qanAPI.Do(req)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer resp.Body.Close()

	rb, _ = httputil.DumpResponse(resp, true)
	if resp.StatusCode != 200 {
		logger.Get(ctx).Errorf("getOSUUID response:\n\n%s\n", rb)
		return "", errors.Errorf("unexpected QAN response status code %d", resp.StatusCode)
	}
	logger.Get(ctx).Debugf("getOSUUID response:\n\n%s\n", rb)

	var instance proto.Instance
	if err = json.NewDecoder(resp.Body).Decode(&instance); err != nil {
		return "", errors.WithStack(err)
	}
	return instance.ParentUUID, nil
}

// addInstance adds instance to QAN API.
func (svc *Service) addInstance(ctx context.Context, instance *proto.Instance) (string, error) {
	b, err := json.Marshal(instance)
	if err != nil {
		return "", errors.WithStack(err)
	}

	url := svc.qanURL
	url.Path = path.Join(url.Path, "instances")
	req, err := http.NewRequest("POST", url.String(), bytes.NewReader(b))
	if err != nil {
		return "", errors.WithStack(err)
	}
	rb, _ := httputil.DumpRequestOut(req, true)
	logger.Get(ctx).Debugf("addInstance request:\n\n%s\n", rb)

	resp, err := svc.qanAPI.Post(url.String(), "application/json", bytes.NewReader(b))
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer resp.Body.Close()

	rb, _ = httputil.DumpResponse(resp, true)
	if resp.StatusCode != 201 {
		logger.Get(ctx).Errorf("addInstance response:\n\n%s\n", rb)
		return "", errors.Errorf("unexpected QAN response status code %d", resp.StatusCode)
	}
	logger.Get(ctx).Debugf("addInstance response:\n\n%s\n", rb)

	// Response Location header looks like this: http://127.0.0.1/qan-api/instances/6cea8824082d4ade682b94109664e6a9
	// Extract UUID directly from it instead of following it.
	parts := strings.Split(resp.Header.Get("Location"), "/")
	return parts[len(parts)-1], nil
}

func (svc *Service) ensureAgentRuns(ctx context.Context, qanAgent *models.QanAgent) error {
	name := qanAgent.NameForSupervisor()
	err := svc.supervisor.Status(ctx, name)
	if err != nil {
		err = svc.supervisor.Stop(ctx, name)
		if err != nil {
			logger.Get(ctx).Warn(err)
		}

		config := &service.Config{
			Name:        name,
			DisplayName: name,
			Description: name,
			Executable:  filepath.Join(svc.baseDir, "bin", "percona-qan-agent"),
			Arguments: []string{
				fmt.Sprintf("-listen=127.0.0.1:%d", *qanAgent.ListenPort),
			},
		}
		err = svc.supervisor.Start(ctx, config)
	}
	return err
}

func (svc *Service) sendQANCommand(ctx context.Context, agentUUID string, command string, data []byte) error {
	cmd := proto.Cmd{
		User:      "pmm-managed",
		AgentUUID: agentUUID,
		Service:   "qan",
		Cmd:       command,
		Data:      data,
	}
	b, err := json.Marshal(cmd)
	if err != nil {
		return errors.WithStack(err)
	}

	// Send the command to the API which relays it to the agent, then relays the agent's reply back to here.
	// It takes a few seconds for agent to connect to QAN API once it is started via service manager.
	// QAN API fails to start/stop unconnected agent for QAN, so we retry the request when getting 404 response.
	const attempts = 10
	url := svc.qanURL
	url.Path = path.Join(url.Path, "agents", agentUUID, "cmd")
	for i := 0; i < attempts; i++ {
		req, err := http.NewRequest("PUT", url.String(), bytes.NewReader(b))
		if err != nil {
			return errors.WithStack(err)
		}
		req.Header.Set("Content-Type", "application/json")
		rb, _ := httputil.DumpRequestOut(req, true)
		logger.Get(ctx).Debugf("sendQANCommand request:\n\n%s\n", rb)

		resp, err := svc.qanAPI.Do(req)
		if err != nil {
			return errors.WithStack(err)
		}
		rb, _ = httputil.DumpResponse(resp, true)
		resp.Body.Close()

		if resp.StatusCode == 200 {
			logger.Get(ctx).Debugf("sendQANCommand response:\n\n%s\n", rb)
			return nil
		}
		if resp.StatusCode == 404 {
			logger.Get(ctx).Debugf("sendQANCommand response:\n\n%s\n", rb)
			time.Sleep(time.Second)
			continue
		}

		logger.Get(ctx).Errorf("sendQANCommand response:\n\n%s\n", rb)
		return errors.Errorf("%s: unexpected QAN API response status code %d", command, resp.StatusCode)
	}

	return errors.Errorf("%s: failed to send command after %d attempts", command, attempts)
}

func (svc *Service) AddMySQL(ctx context.Context, rdsNode *models.RDSNode, rdsService *models.RDSService, qanAgent *models.QanAgent) error {
	agentUUID, err := svc.getAgentUUID()
	if err != nil {
		return err
	}
	osUUID, err := svc.getOSUUID(ctx, agentUUID)
	if err != nil {
		return err
	}

	instance := &proto.Instance{
		Subsystem:  "mysql",
		ParentUUID: osUUID,
		Name:       rdsNode.Name,
		DSN:        sanitizeDSN(qanAgent.DSN(rdsService)),
		Version:    *rdsService.EngineVersion,
	}
	instanceUUID, err := svc.addInstance(ctx, instance)
	if err != nil {
		return err
	}

	// we need real DSN (with password) for qan-agent to work, and it seems to be the only way to pass it
	path := filepath.Join(svc.baseDir, "instance", fmt.Sprintf("%s.json", instanceUUID))
	instance.DSN = qanAgent.DSN(rdsService)
	b, err := json.MarshalIndent(instance, "", "    ")
	if err != nil {
		return errors.WithStack(err)
	}
	if err = ioutil.WriteFile(path, b, 0666); err != nil {
		return errors.WithStack(err)
	}

	if err = svc.ensureAgentRuns(ctx, qanAgent); err != nil {
		return err
	}

	command := "StartTool"
	config := map[string]interface{}{
		"UUID":           instanceUUID,
		"CollectFrom":    "perfschema",
		"Interval":       60,
		"ExampleQueries": true,
	}
	b, err = json.Marshal(config)
	if err != nil {
		return errors.WithStack(err)
	}
	logger.Get(ctx).Debugf("%s %s %s", agentUUID, command, b)
	return svc.sendQANCommand(ctx, agentUUID, command, b)
}
