// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Category category
// swagger:model Category
type Category struct {

	// category parent
	CategoryParent *Category `json:"category_parent,omitempty"`

	// category parent id
	CategoryParentID int64 `json:"category_parent_id,omitempty" gorm:"ForeignKey:CategoryID"`

	// id
	ID int64 `json:"id,omitempty" gorm:"primary_key"`

	// name
	Name string `json:"name,omitempty"`

	// status
	Status int64 `json:"status,omitempty"`
}

// Validate validates this category
func (m *Category) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryParent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Category) validateCategoryParent(formats strfmt.Registry) error {

	if swag.IsZero(m.CategoryParent) { // not required
		return nil
	}

	if m.CategoryParent != nil {
		if err := m.CategoryParent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category_parent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Category) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Category) UnmarshalBinary(b []byte) error {
	var res Category
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
