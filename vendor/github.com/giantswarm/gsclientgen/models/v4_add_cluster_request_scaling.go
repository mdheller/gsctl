// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4AddClusterRequestScaling Attributes specific to cluster node scaling. To have full control of
// the cluster size, min and max can be set to the same value. If only
// `min` or only `max` is specified, `min` and `max` will be set equally.
//
// swagger:model v4AddClusterRequestScaling
type V4AddClusterRequestScaling struct {

	// The maximum number of cluster nodes
	//
	Max int64 `json:"max,omitempty"`

	// The minimum number of cluster nodes
	//
	Min int64 `json:"min,omitempty"`
}

// Validate validates this v4 add cluster request scaling
func (m *V4AddClusterRequestScaling) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4AddClusterRequestScaling) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4AddClusterRequestScaling) UnmarshalBinary(b []byte) error {
	var res V4AddClusterRequestScaling
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}