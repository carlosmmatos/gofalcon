// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestAWSAccountValidated rest a w s account validated
//
// swagger:model rest.AWSAccountValidated
type RestAWSAccountValidated struct {

	// account id
	// Required: true
	AccountID *string `json:"account_id"`

	// conditions
	Conditions []*DomainCloudCondition `json:"conditions"`

	// Permissions status for each product returned via API.
	// Required: true
	IamServicePermissionsStatus []*DomainProductPermission `json:"iam_service_permissions_status"`

	// Permissions status for each product returned via API.
	// Required: true
	Status []*DomainProductFeaturesStatus `json:"status"`
}

// Validate validates this rest a w s account validated
func (m *RestAWSAccountValidated) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIamServicePermissionsStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestAWSAccountValidated) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *RestAWSAccountValidated) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestAWSAccountValidated) validateIamServicePermissionsStatus(formats strfmt.Registry) error {

	if err := validate.Required("iam_service_permissions_status", "body", m.IamServicePermissionsStatus); err != nil {
		return err
	}

	for i := 0; i < len(m.IamServicePermissionsStatus); i++ {
		if swag.IsZero(m.IamServicePermissionsStatus[i]) { // not required
			continue
		}

		if m.IamServicePermissionsStatus[i] != nil {
			if err := m.IamServicePermissionsStatus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("iam_service_permissions_status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("iam_service_permissions_status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestAWSAccountValidated) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	for i := 0; i < len(m.Status); i++ {
		if swag.IsZero(m.Status[i]) { // not required
			continue
		}

		if m.Status[i] != nil {
			if err := m.Status[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this rest a w s account validated based on the context it is used
func (m *RestAWSAccountValidated) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIamServicePermissionsStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestAWSAccountValidated) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conditions); i++ {

		if m.Conditions[i] != nil {

			if swag.IsZero(m.Conditions[i]) { // not required
				return nil
			}

			if err := m.Conditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestAWSAccountValidated) contextValidateIamServicePermissionsStatus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IamServicePermissionsStatus); i++ {

		if m.IamServicePermissionsStatus[i] != nil {

			if swag.IsZero(m.IamServicePermissionsStatus[i]) { // not required
				return nil
			}

			if err := m.IamServicePermissionsStatus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("iam_service_permissions_status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("iam_service_permissions_status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestAWSAccountValidated) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Status); i++ {

		if m.Status[i] != nil {

			if swag.IsZero(m.Status[i]) { // not required
				return nil
			}

			if err := m.Status[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestAWSAccountValidated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestAWSAccountValidated) UnmarshalBinary(b []byte) error {
	var res RestAWSAccountValidated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}