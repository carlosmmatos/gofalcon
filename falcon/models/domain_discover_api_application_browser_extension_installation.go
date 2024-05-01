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

// DomainDiscoverAPIApplicationBrowserExtensionInstallation domain discover API application browser extension installation
//
// swagger:model domain.DiscoverAPIApplicationBrowserExtensionInstallation
type DomainDiscoverAPIApplicationBrowserExtensionInstallation struct {

	// The browser profile ID of this installation
	BrowserProfileID string `json:"browser_profile_id,omitempty"`

	// The browser profile name of this installation
	BrowserProfileName string `json:"browser_profile_name,omitempty"`

	// The version of the browser running this extension
	BrowserVersion string `json:"browser_version,omitempty"`

	// Describes if the extension is enabled on this browser profile installation
	// Required: true
	Enabled *bool `json:"enabled"`

	// The method that was used to install the browser extension
	Method string `json:"method,omitempty"`

	// The file path location of the browser extension
	Path string `json:"path,omitempty"`

	// The role of the account that installed the extension
	Role string `json:"role,omitempty"`

	// The host user SID for which the extension was installed
	UserSid string `json:"user_sid,omitempty"`

	// The host username for which the extension was installed
	Username string `json:"username,omitempty"`
}

// Validate validates this domain discover API application browser extension installation
func (m *DomainDiscoverAPIApplicationBrowserExtensionInstallation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainDiscoverAPIApplicationBrowserExtensionInstallation) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain discover API application browser extension installation based on context it is used
func (m *DomainDiscoverAPIApplicationBrowserExtensionInstallation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainDiscoverAPIApplicationBrowserExtensionInstallation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDiscoverAPIApplicationBrowserExtensionInstallation) UnmarshalBinary(b []byte) error {
	var res DomainDiscoverAPIApplicationBrowserExtensionInstallation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}