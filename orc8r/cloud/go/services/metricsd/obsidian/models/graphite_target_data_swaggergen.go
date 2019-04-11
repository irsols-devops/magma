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

// GraphiteTargetData graphite target data
// swagger:model graphite_target_data
type GraphiteTargetData struct {

	// datapoints
	// Required: true
	Datapoints MetricDatapoints `json:"datapoints"`

	// tags
	// Required: true
	Tags map[string]string `json:"tags"`

	// target
	// Required: true
	Target *string `json:"target"`
}

// Validate validates this graphite target data
func (m *GraphiteTargetData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatapoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphiteTargetData) validateDatapoints(formats strfmt.Registry) error {

	if err := validate.Required("datapoints", "body", m.Datapoints); err != nil {
		return err
	}

	if err := m.Datapoints.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("datapoints")
		}
		return err
	}

	return nil
}

func (m *GraphiteTargetData) validateTags(formats strfmt.Registry) error {

	return nil
}

func (m *GraphiteTargetData) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GraphiteTargetData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphiteTargetData) UnmarshalBinary(b []byte) error {
	var res GraphiteTargetData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
