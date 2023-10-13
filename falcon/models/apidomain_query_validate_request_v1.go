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

// ApidomainQueryValidateRequestV1 apidomain query validate request v1
//
// swagger:model apidomain.QueryValidateRequestV1
type ApidomainQueryValidateRequestV1 struct {

	// arguments
	// Required: true
	Arguments map[string]string `json:"arguments"`

	// query string
	// Required: true
	QueryString *string `json:"query_string"`
}

// Validate validates this apidomain query validate request v1
func (m *ApidomainQueryValidateRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryString(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApidomainQueryValidateRequestV1) validateArguments(formats strfmt.Registry) error {

	if err := validate.Required("arguments", "body", m.Arguments); err != nil {
		return err
	}

	return nil
}

func (m *ApidomainQueryValidateRequestV1) validateQueryString(formats strfmt.Registry) error {

	if err := validate.Required("query_string", "body", m.QueryString); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this apidomain query validate request v1 based on context it is used
func (m *ApidomainQueryValidateRequestV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApidomainQueryValidateRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApidomainQueryValidateRequestV1) UnmarshalBinary(b []byte) error {
	var res ApidomainQueryValidateRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}