package o_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/manifoldco/grafton/generated/connector/models"
)

// PostCredentialsReader is a Reader for the PostCredentials structure.
type PostCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostCredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostCredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCredentialsOK creates a PostCredentialsOK with default headers values
func NewPostCredentialsOK() *PostCredentialsOK {
	return &PostCredentialsOK{}
}

/*PostCredentialsOK handles this case with default header values.

The created OAuth 2.0 credential pair.
*/
type PostCredentialsOK struct {
	Payload *models.OAuthCredentialCreateResponse
}

func (o *PostCredentialsOK) Error() string {
	return fmt.Sprintf("[POST /credentials/][%d] postCredentialsOK  %+v", 200, o.Payload)
}

func (o *PostCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthCredentialCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCredentialsBadRequest creates a PostCredentialsBadRequest with default headers values
func NewPostCredentialsBadRequest() *PostCredentialsBadRequest {
	return &PostCredentialsBadRequest{}
}

/*PostCredentialsBadRequest handles this case with default header values.

Request denied due to invalid request body, path, or headers.
*/
type PostCredentialsBadRequest struct {
	Payload PostCredentialsBadRequestBody
}

func (o *PostCredentialsBadRequest) Error() string {
	return fmt.Sprintf("[POST /credentials/][%d] postCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *PostCredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCredentialsUnauthorized creates a PostCredentialsUnauthorized with default headers values
func NewPostCredentialsUnauthorized() *PostCredentialsUnauthorized {
	return &PostCredentialsUnauthorized{}
}

/*PostCredentialsUnauthorized handles this case with default header values.

Request denied as the provided credentials are no longer valid.
*/
type PostCredentialsUnauthorized struct {
	Payload PostCredentialsUnauthorizedBody
}

func (o *PostCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /credentials/][%d] postCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCredentialsNotFound creates a PostCredentialsNotFound with default headers values
func NewPostCredentialsNotFound() *PostCredentialsNotFound {
	return &PostCredentialsNotFound{}
}

/*PostCredentialsNotFound handles this case with default header values.

Request denied as the requested resource does not exist.
*/
type PostCredentialsNotFound struct {
	Payload PostCredentialsNotFoundBody
}

func (o *PostCredentialsNotFound) Error() string {
	return fmt.Sprintf("[POST /credentials/][%d] postCredentialsNotFound  %+v", 404, o.Payload)
}

func (o *PostCredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCredentialsInternalServerError creates a PostCredentialsInternalServerError with default headers values
func NewPostCredentialsInternalServerError() *PostCredentialsInternalServerError {
	return &PostCredentialsInternalServerError{}
}

/*PostCredentialsInternalServerError handles this case with default header values.

Request failed due to an internal server error.
*/
type PostCredentialsInternalServerError struct {
	Payload PostCredentialsInternalServerErrorBody
}

func (o *PostCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /credentials/][%d] postCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCredentialsBadRequestBody post credentials bad request body
swagger:model PostCredentialsBadRequestBody
*/
type PostCredentialsBadRequestBody struct {

	// Explanation of the errors
	// Required: true
	Message []string `json:"message"`

	// The error type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this post credentials bad request body
func (o *PostCredentialsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCredentialsBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("postCredentialsBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

var postCredentialsBadRequestBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bad_request","unauthorized","not_found","internal","invalid_grant","unsupported_grant_type"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postCredentialsBadRequestBodyTypeTypePropEnum = append(postCredentialsBadRequestBodyTypeTypePropEnum, v)
	}
}

const (
	// PostCredentialsBadRequestBodyTypeBadRequest captures enum value "bad_request"
	PostCredentialsBadRequestBodyTypeBadRequest string = "bad_request"
	// PostCredentialsBadRequestBodyTypeUnauthorized captures enum value "unauthorized"
	PostCredentialsBadRequestBodyTypeUnauthorized string = "unauthorized"
	// PostCredentialsBadRequestBodyTypeNotFound captures enum value "not_found"
	PostCredentialsBadRequestBodyTypeNotFound string = "not_found"
	// PostCredentialsBadRequestBodyTypeInternal captures enum value "internal"
	PostCredentialsBadRequestBodyTypeInternal string = "internal"
	// PostCredentialsBadRequestBodyTypeInvalidGrant captures enum value "invalid_grant"
	PostCredentialsBadRequestBodyTypeInvalidGrant string = "invalid_grant"
	// PostCredentialsBadRequestBodyTypeUnsupportedGrantType captures enum value "unsupported_grant_type"
	PostCredentialsBadRequestBodyTypeUnsupportedGrantType string = "unsupported_grant_type"
)

// prop value enum
func (o *PostCredentialsBadRequestBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postCredentialsBadRequestBodyTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PostCredentialsBadRequestBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("postCredentialsBadRequest"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("postCredentialsBadRequest"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

/*PostCredentialsInternalServerErrorBody post credentials internal server error body
swagger:model PostCredentialsInternalServerErrorBody
*/
type PostCredentialsInternalServerErrorBody struct {

	// Explanation of the errors
	// Required: true
	Message []string `json:"message"`

	// The error type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this post credentials internal server error body
func (o *PostCredentialsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCredentialsInternalServerErrorBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("postCredentialsInternalServerError"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

var postCredentialsInternalServerErrorBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bad_request","unauthorized","not_found","internal","invalid_grant","unsupported_grant_type"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postCredentialsInternalServerErrorBodyTypeTypePropEnum = append(postCredentialsInternalServerErrorBodyTypeTypePropEnum, v)
	}
}

const (
	// PostCredentialsInternalServerErrorBodyTypeBadRequest captures enum value "bad_request"
	PostCredentialsInternalServerErrorBodyTypeBadRequest string = "bad_request"
	// PostCredentialsInternalServerErrorBodyTypeUnauthorized captures enum value "unauthorized"
	PostCredentialsInternalServerErrorBodyTypeUnauthorized string = "unauthorized"
	// PostCredentialsInternalServerErrorBodyTypeNotFound captures enum value "not_found"
	PostCredentialsInternalServerErrorBodyTypeNotFound string = "not_found"
	// PostCredentialsInternalServerErrorBodyTypeInternal captures enum value "internal"
	PostCredentialsInternalServerErrorBodyTypeInternal string = "internal"
	// PostCredentialsInternalServerErrorBodyTypeInvalidGrant captures enum value "invalid_grant"
	PostCredentialsInternalServerErrorBodyTypeInvalidGrant string = "invalid_grant"
	// PostCredentialsInternalServerErrorBodyTypeUnsupportedGrantType captures enum value "unsupported_grant_type"
	PostCredentialsInternalServerErrorBodyTypeUnsupportedGrantType string = "unsupported_grant_type"
)

// prop value enum
func (o *PostCredentialsInternalServerErrorBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postCredentialsInternalServerErrorBodyTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PostCredentialsInternalServerErrorBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("postCredentialsInternalServerError"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("postCredentialsInternalServerError"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

/*PostCredentialsNotFoundBody post credentials not found body
swagger:model PostCredentialsNotFoundBody
*/
type PostCredentialsNotFoundBody struct {

	// Explanation of the errors
	// Required: true
	Message []string `json:"message"`

	// The error type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this post credentials not found body
func (o *PostCredentialsNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCredentialsNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("postCredentialsNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

var postCredentialsNotFoundBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bad_request","unauthorized","not_found","internal","invalid_grant","unsupported_grant_type"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postCredentialsNotFoundBodyTypeTypePropEnum = append(postCredentialsNotFoundBodyTypeTypePropEnum, v)
	}
}

const (
	// PostCredentialsNotFoundBodyTypeBadRequest captures enum value "bad_request"
	PostCredentialsNotFoundBodyTypeBadRequest string = "bad_request"
	// PostCredentialsNotFoundBodyTypeUnauthorized captures enum value "unauthorized"
	PostCredentialsNotFoundBodyTypeUnauthorized string = "unauthorized"
	// PostCredentialsNotFoundBodyTypeNotFound captures enum value "not_found"
	PostCredentialsNotFoundBodyTypeNotFound string = "not_found"
	// PostCredentialsNotFoundBodyTypeInternal captures enum value "internal"
	PostCredentialsNotFoundBodyTypeInternal string = "internal"
	// PostCredentialsNotFoundBodyTypeInvalidGrant captures enum value "invalid_grant"
	PostCredentialsNotFoundBodyTypeInvalidGrant string = "invalid_grant"
	// PostCredentialsNotFoundBodyTypeUnsupportedGrantType captures enum value "unsupported_grant_type"
	PostCredentialsNotFoundBodyTypeUnsupportedGrantType string = "unsupported_grant_type"
)

// prop value enum
func (o *PostCredentialsNotFoundBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postCredentialsNotFoundBodyTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PostCredentialsNotFoundBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("postCredentialsNotFound"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("postCredentialsNotFound"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

/*PostCredentialsUnauthorizedBody post credentials unauthorized body
swagger:model PostCredentialsUnauthorizedBody
*/
type PostCredentialsUnauthorizedBody struct {

	// Explanation of the errors
	// Required: true
	Message []string `json:"message"`

	// The error type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this post credentials unauthorized body
func (o *PostCredentialsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCredentialsUnauthorizedBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("postCredentialsUnauthorized"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

var postCredentialsUnauthorizedBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bad_request","unauthorized","not_found","internal","invalid_grant","unsupported_grant_type"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postCredentialsUnauthorizedBodyTypeTypePropEnum = append(postCredentialsUnauthorizedBodyTypeTypePropEnum, v)
	}
}

const (
	// PostCredentialsUnauthorizedBodyTypeBadRequest captures enum value "bad_request"
	PostCredentialsUnauthorizedBodyTypeBadRequest string = "bad_request"
	// PostCredentialsUnauthorizedBodyTypeUnauthorized captures enum value "unauthorized"
	PostCredentialsUnauthorizedBodyTypeUnauthorized string = "unauthorized"
	// PostCredentialsUnauthorizedBodyTypeNotFound captures enum value "not_found"
	PostCredentialsUnauthorizedBodyTypeNotFound string = "not_found"
	// PostCredentialsUnauthorizedBodyTypeInternal captures enum value "internal"
	PostCredentialsUnauthorizedBodyTypeInternal string = "internal"
	// PostCredentialsUnauthorizedBodyTypeInvalidGrant captures enum value "invalid_grant"
	PostCredentialsUnauthorizedBodyTypeInvalidGrant string = "invalid_grant"
	// PostCredentialsUnauthorizedBodyTypeUnsupportedGrantType captures enum value "unsupported_grant_type"
	PostCredentialsUnauthorizedBodyTypeUnsupportedGrantType string = "unsupported_grant_type"
)

// prop value enum
func (o *PostCredentialsUnauthorizedBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postCredentialsUnauthorizedBodyTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PostCredentialsUnauthorizedBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("postCredentialsUnauthorized"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("postCredentialsUnauthorized"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}