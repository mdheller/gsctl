// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V4AppStatus v4 app status
// swagger:model v4AppStatus
type V4AppStatus struct {

	// Version of the installed app
	AppVersion string `json:"app_version,omitempty"`

	// release
	Release *V4AppStatusRelease `json:"release,omitempty"`

	// Version of the chart that was used to install this app
	Version string `json:"version,omitempty"`
}

// Validate validates this v4 app status
func (m *V4AppStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V4AppStatus) validateRelease(formats strfmt.Registry) error {

	if swag.IsZero(m.Release) { // not required
		return nil
	}

	if m.Release != nil {
		if err := m.Release.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("release")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4AppStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4AppStatus) UnmarshalBinary(b []byte) error {
	var res V4AppStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}