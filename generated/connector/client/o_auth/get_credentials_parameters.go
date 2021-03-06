package o_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCredentialsParams creates a new GetCredentialsParams object
// with the default values initialized.
func NewGetCredentialsParams() *GetCredentialsParams {
	var ()
	return &GetCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCredentialsParamsWithTimeout creates a new GetCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCredentialsParamsWithTimeout(timeout time.Duration) *GetCredentialsParams {
	var ()
	return &GetCredentialsParams{

		timeout: timeout,
	}
}

// NewGetCredentialsParamsWithContext creates a new GetCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCredentialsParamsWithContext(ctx context.Context) *GetCredentialsParams {
	var ()
	return &GetCredentialsParams{

		Context: ctx,
	}
}

// NewGetCredentialsParamsWithHTTPClient creates a new GetCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCredentialsParamsWithHTTPClient(client *http.Client) *GetCredentialsParams {
	var ()
	return &GetCredentialsParams{
		HTTPClient: client,
	}
}

/*GetCredentialsParams contains all the parameters to send to the API endpoint
for the get credentials operation typically these are written to a http.Request
*/
type GetCredentialsParams struct {

	/*ProductID
	  ID of the Product to filter Resources by, stored as a
	base32 encoded 18 byte identifier.


	*/
	ProductID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get credentials params
func (o *GetCredentialsParams) WithTimeout(timeout time.Duration) *GetCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get credentials params
func (o *GetCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get credentials params
func (o *GetCredentialsParams) WithContext(ctx context.Context) *GetCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get credentials params
func (o *GetCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get credentials params
func (o *GetCredentialsParams) WithHTTPClient(client *http.Client) *GetCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get credentials params
func (o *GetCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProductID adds the productID to the get credentials params
func (o *GetCredentialsParams) WithProductID(productID string) *GetCredentialsParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the get credentials params
func (o *GetCredentialsParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param product_id
	qrProductID := o.ProductID
	qProductID := qrProductID
	if qProductID != "" {
		if err := r.SetQueryParam("product_id", qProductID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
