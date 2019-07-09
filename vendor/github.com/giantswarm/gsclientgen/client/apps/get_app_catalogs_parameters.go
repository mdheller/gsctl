// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAppCatalogsParams creates a new GetAppCatalogsParams object
// with the default values initialized.
func NewGetAppCatalogsParams() *GetAppCatalogsParams {
	var ()
	return &GetAppCatalogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppCatalogsParamsWithTimeout creates a new GetAppCatalogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppCatalogsParamsWithTimeout(timeout time.Duration) *GetAppCatalogsParams {
	var ()
	return &GetAppCatalogsParams{

		timeout: timeout,
	}
}

// NewGetAppCatalogsParamsWithContext creates a new GetAppCatalogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppCatalogsParamsWithContext(ctx context.Context) *GetAppCatalogsParams {
	var ()
	return &GetAppCatalogsParams{

		Context: ctx,
	}
}

// NewGetAppCatalogsParamsWithHTTPClient creates a new GetAppCatalogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAppCatalogsParamsWithHTTPClient(client *http.Client) *GetAppCatalogsParams {
	var ()
	return &GetAppCatalogsParams{
		HTTPClient: client,
	}
}

/*GetAppCatalogsParams contains all the parameters to send to the API endpoint
for the get app catalogs operation typically these are written to a http.Request
*/
type GetAppCatalogsParams struct {

	/*Authorization
	  As described in the [authentication](#section/Authentication) section


	*/
	Authorization string
	/*XGiantSwarmActivity
	  Name of an activity to track, like "list-clusters". This allows to
	analyze several API requests sent in context and gives an idea on
	the purpose.


	*/
	XGiantSwarmActivity *string
	/*XGiantSwarmCmdLine
	  If activity has been issued by a CLI, this header can contain the
	command line


	*/
	XGiantSwarmCmdLine *string
	/*XRequestID
	  A randomly generated key that can be used to track a request throughout
	services of Giant Swarm.


	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get app catalogs params
func (o *GetAppCatalogsParams) WithTimeout(timeout time.Duration) *GetAppCatalogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app catalogs params
func (o *GetAppCatalogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app catalogs params
func (o *GetAppCatalogsParams) WithContext(ctx context.Context) *GetAppCatalogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app catalogs params
func (o *GetAppCatalogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app catalogs params
func (o *GetAppCatalogsParams) WithHTTPClient(client *http.Client) *GetAppCatalogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app catalogs params
func (o *GetAppCatalogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get app catalogs params
func (o *GetAppCatalogsParams) WithAuthorization(authorization string) *GetAppCatalogsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get app catalogs params
func (o *GetAppCatalogsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the get app catalogs params
func (o *GetAppCatalogsParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *GetAppCatalogsParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the get app catalogs params
func (o *GetAppCatalogsParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the get app catalogs params
func (o *GetAppCatalogsParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *GetAppCatalogsParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the get app catalogs params
func (o *GetAppCatalogsParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the get app catalogs params
func (o *GetAppCatalogsParams) WithXRequestID(xRequestID *string) *GetAppCatalogsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get app catalogs params
func (o *GetAppCatalogsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppCatalogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.XGiantSwarmActivity != nil {

		// header param X-Giant-Swarm-Activity
		if err := r.SetHeaderParam("X-Giant-Swarm-Activity", *o.XGiantSwarmActivity); err != nil {
			return err
		}

	}

	if o.XGiantSwarmCmdLine != nil {

		// header param X-Giant-Swarm-CmdLine
		if err := r.SetHeaderParam("X-Giant-Swarm-CmdLine", *o.XGiantSwarmCmdLine); err != nil {
			return err
		}

	}

	if o.XRequestID != nil {

		// header param X-Request-ID
		if err := r.SetHeaderParam("X-Request-ID", *o.XRequestID); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}