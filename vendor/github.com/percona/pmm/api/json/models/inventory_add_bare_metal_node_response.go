// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InventoryAddBareMetalNodeResponse inventory add bare metal node response
// swagger:model inventoryAddBareMetalNodeResponse
type InventoryAddBareMetalNodeResponse struct {

	// bare metal
	BareMetal *InventoryBareMetalNode `json:"bare_metal,omitempty"`
}

// Validate validates this inventory add bare metal node response
func (m *InventoryAddBareMetalNodeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBareMetal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryAddBareMetalNodeResponse) validateBareMetal(formats strfmt.Registry) error {

	if swag.IsZero(m.BareMetal) { // not required
		return nil
	}

	if m.BareMetal != nil {
		if err := m.BareMetal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bare_metal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryAddBareMetalNodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryAddBareMetalNodeResponse) UnmarshalBinary(b []byte) error {
	var res InventoryAddBareMetalNodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
