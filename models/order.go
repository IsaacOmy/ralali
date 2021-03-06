// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Order order
// swagger:model Order
type Order struct {

	// created date
	CreatedDate string `json:"created_date,omitempty"`

	// id
	ID int64 `json:"id,omitempty" gorm:"primary_key"`

	// item orders
	ItemOrders []*ItemOrder `json:"item_orders"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemOrders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Order) validateItemOrders(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemOrders) { // not required
		return nil
	}

	for i := 0; i < len(m.ItemOrders); i++ {
		if swag.IsZero(m.ItemOrders[i]) { // not required
			continue
		}

		if m.ItemOrders[i] != nil {
			if err := m.ItemOrders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("item_orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
