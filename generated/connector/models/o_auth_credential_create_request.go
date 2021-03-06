package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"

	manifold "github.com/manifoldco/go-manifold"
)

// OAuthCredentialCreateRequest o auth credential create request
// swagger:model OAuthCredentialCreateRequest
type OAuthCredentialCreateRequest struct {

	// A human readable description of this credential pair.
	//
	// Required: true
	// Max Length: 256
	// Min Length: 3
	Description *string `json:"description"`

	// product id
	// Required: true
	ProductID manifold.ID `json:"product_id"`
}

// Validate validates this o auth credential create request
func (m *OAuthCredentialCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProductID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OAuthCredentialCreateRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 256); err != nil {
		return err
	}

	return nil
}

func (m *OAuthCredentialCreateRequest) validateProductID(formats strfmt.Registry) error {

	if err := m.ProductID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("product_id")
		}
		return err
	}

	return nil
}
