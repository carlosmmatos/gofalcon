// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PublicPrevalence public prevalence
//
// swagger:model public.Prevalence
type PublicPrevalence struct {

	// computed timestamp
	// Required: true
	ComputedTimestamp *string `json:"computed_timestamp"`

	// Possible values: PENDING, UNIQUE, LOW, COMMON.
	// Required: true
	Current *string `json:"current"`

	// key
	// Required: true
	Key *string `json:"key"`

	// Possible values: PENDING, UNIQUE, LOW, COMMON.
	// Required: true
	Reported *string `json:"reported"`
}

// Validate validates this public prevalence
func (m *PublicPrevalence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComputedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReported(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicPrevalence) validateComputedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("computed_timestamp", "body", m.ComputedTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *PublicPrevalence) validateCurrent(formats strfmt.Registry) error {

	if err := validate.Required("current", "body", m.Current); err != nil {
		return err
	}

	return nil
}

func (m *PublicPrevalence) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *PublicPrevalence) validateReported(formats strfmt.Registry) error {

	if err := validate.Required("reported", "body", m.Reported); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this public prevalence based on context it is used
func (m *PublicPrevalence) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PublicPrevalence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicPrevalence) UnmarshalBinary(b []byte) error {
	var res PublicPrevalence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}