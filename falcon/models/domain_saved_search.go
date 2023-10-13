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

// DomainSavedSearch domain saved search
//
// swagger:model domain.SavedSearch
type DomainSavedSearch struct {

	// cid
	Cid string `json:"cid,omitempty"`

	// created by
	// Required: true
	CreatedBy *string `json:"created_by"`

	// created timestamp
	// Required: true
	// Format: date-time
	CreatedTimestamp *strfmt.DateTime `json:"created_timestamp"`

	// description
	// Required: true
	Description *string `json:"description"`

	// end
	End string `json:"end,omitempty"`

	// id of saved search
	// Required: true
	ID *string `json:"id"`

	// last modified by
	LastModifiedBy string `json:"last_modified_by,omitempty"`

	// last modified timestamp
	// Format: date-time
	LastModifiedTimestamp strfmt.DateTime `json:"last_modified_timestamp,omitempty"`

	// owner id
	OwnerID string `json:"owner_id,omitempty"`

	// owner kind
	OwnerKind string `json:"owner_kind,omitempty"`

	// permissions
	Permissions []string `json:"permissions"`

	// repo or view
	// Required: true
	RepoOrView *string `json:"repo_or_view"`

	// request schema
	RequestSchema string `json:"request_schema,omitempty"`

	// response schema
	ResponseSchema string `json:"response_schema,omitempty"`

	// search name
	// Required: true
	SearchName *string `json:"search_name"`

	// search query
	// Required: true
	SearchQuery *string `json:"search_query"`

	// search query args
	SearchQueryArgs map[string]string `json:"search_query_args,omitempty"`

	// start
	Start string `json:"start,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	// version
	// Required: true
	Version *string `json:"version"`

	// version numeric
	// Required: true
	VersionNumeric *int32 `json:"version_numeric"`

	// workflow id
	WorkflowID string `json:"workflow_id,omitempty"`

	// workflow meta
	WorkflowMeta *DomainWorkflowMetadata `json:"workflow_meta,omitempty"`
}

// Validate validates this domain saved search
func (m *DomainSavedSearch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepoOrView(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionNumeric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainSavedSearch) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("created_by", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *DomainSavedSearch) validateCreatedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("created_timestamp", "body", m.CreatedTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("created_timestamp", "body", "date-time", m.CreatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainSavedSearch) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *DomainSavedSearch) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainSavedSearch) validateLastModifiedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModifiedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("last_modified_timestamp", "body", "date-time", m.LastModifiedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainSavedSearch) validateRepoOrView(formats strfmt.Registry) error {

	if err := validate.Required("repo_or_view", "body", m.RepoOrView); err != nil {
		return err
	}

	return nil
}

func (m *DomainSavedSearch) validateSearchName(formats strfmt.Registry) error {

	if err := validate.Required("search_name", "body", m.SearchName); err != nil {
		return err
	}

	return nil
}

func (m *DomainSavedSearch) validateSearchQuery(formats strfmt.Registry) error {

	if err := validate.Required("search_query", "body", m.SearchQuery); err != nil {
		return err
	}

	return nil
}

func (m *DomainSavedSearch) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *DomainSavedSearch) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *DomainSavedSearch) validateVersionNumeric(formats strfmt.Registry) error {

	if err := validate.Required("version_numeric", "body", m.VersionNumeric); err != nil {
		return err
	}

	return nil
}

func (m *DomainSavedSearch) validateWorkflowMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkflowMeta) { // not required
		return nil
	}

	if m.WorkflowMeta != nil {
		if err := m.WorkflowMeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflow_meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workflow_meta")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this domain saved search based on the context it is used
func (m *DomainSavedSearch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkflowMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainSavedSearch) contextValidateWorkflowMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkflowMeta != nil {

		if swag.IsZero(m.WorkflowMeta) { // not required
			return nil
		}

		if err := m.WorkflowMeta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflow_meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workflow_meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainSavedSearch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainSavedSearch) UnmarshalBinary(b []byte) error {
	var res DomainSavedSearch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}