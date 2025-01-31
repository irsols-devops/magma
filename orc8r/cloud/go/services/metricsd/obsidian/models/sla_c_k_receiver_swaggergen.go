// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SLACKReceiver slack receiver
// swagger:model slack_receiver
type SLACKReceiver struct {

	// api url
	// Required: true
	APIURL *string `json:"api_url"`

	// channel
	Channel string `json:"channel,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this slack receiver
func (m *SLACKReceiver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SLACKReceiver) validateAPIURL(formats strfmt.Registry) error {

	if err := validate.Required("api_url", "body", m.APIURL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SLACKReceiver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SLACKReceiver) UnmarshalBinary(b []byte) error {
	var res SLACKReceiver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
