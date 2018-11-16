// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/percona/pmm/api/json/models"
)

// NewAddContainerNodeParams creates a new AddContainerNodeParams object
// with the default values initialized.
func NewAddContainerNodeParams() *AddContainerNodeParams {
	var ()
	return &AddContainerNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddContainerNodeParamsWithTimeout creates a new AddContainerNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddContainerNodeParamsWithTimeout(timeout time.Duration) *AddContainerNodeParams {
	var ()
	return &AddContainerNodeParams{

		timeout: timeout,
	}
}

// NewAddContainerNodeParamsWithContext creates a new AddContainerNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddContainerNodeParamsWithContext(ctx context.Context) *AddContainerNodeParams {
	var ()
	return &AddContainerNodeParams{

		Context: ctx,
	}
}

// NewAddContainerNodeParamsWithHTTPClient creates a new AddContainerNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddContainerNodeParamsWithHTTPClient(client *http.Client) *AddContainerNodeParams {
	var ()
	return &AddContainerNodeParams{
		HTTPClient: client,
	}
}

/*AddContainerNodeParams contains all the parameters to send to the API endpoint
for the add container node operation typically these are written to a http.Request
*/
type AddContainerNodeParams struct {

	/*Body*/
	Body *models.InventoryAddContainerNodeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add container node params
func (o *AddContainerNodeParams) WithTimeout(timeout time.Duration) *AddContainerNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add container node params
func (o *AddContainerNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add container node params
func (o *AddContainerNodeParams) WithContext(ctx context.Context) *AddContainerNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add container node params
func (o *AddContainerNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add container node params
func (o *AddContainerNodeParams) WithHTTPClient(client *http.Client) *AddContainerNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add container node params
func (o *AddContainerNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add container node params
func (o *AddContainerNodeParams) WithBody(body *models.InventoryAddContainerNodeRequest) *AddContainerNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add container node params
func (o *AddContainerNodeParams) SetBody(body *models.InventoryAddContainerNodeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddContainerNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
