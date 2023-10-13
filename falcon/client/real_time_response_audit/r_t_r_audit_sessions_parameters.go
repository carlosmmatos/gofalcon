// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewRTRAuditSessionsParams creates a new RTRAuditSessionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRTRAuditSessionsParams() *RTRAuditSessionsParams {
	return &RTRAuditSessionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRTRAuditSessionsParamsWithTimeout creates a new RTRAuditSessionsParams object
// with the ability to set a timeout on a request.
func NewRTRAuditSessionsParamsWithTimeout(timeout time.Duration) *RTRAuditSessionsParams {
	return &RTRAuditSessionsParams{
		timeout: timeout,
	}
}

// NewRTRAuditSessionsParamsWithContext creates a new RTRAuditSessionsParams object
// with the ability to set a context for a request.
func NewRTRAuditSessionsParamsWithContext(ctx context.Context) *RTRAuditSessionsParams {
	return &RTRAuditSessionsParams{
		Context: ctx,
	}
}

// NewRTRAuditSessionsParamsWithHTTPClient creates a new RTRAuditSessionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRTRAuditSessionsParamsWithHTTPClient(client *http.Client) *RTRAuditSessionsParams {
	return &RTRAuditSessionsParams{
		HTTPClient: client,
	}
}

/*
RTRAuditSessionsParams contains all the parameters to send to the API endpoint

	for the r t r audit sessions operation.

	Typically these are written to a http.Request.
*/
type RTRAuditSessionsParams struct {

	/* Filter.

	   Optional filter criteria in the form of an FQL query. For more information about FQL queries, see our [FQL documentation in Falcon](https://falcon.crowdstrike.com/support/documentation/45/falcon-query-language-feature-guide).
	*/
	Filter *string

	/* Limit.

	   number of sessions to be returned
	*/
	Limit *string

	/* Offset.

	   offset value to be used for paginated results
	*/
	Offset *string

	/* Sort.

	   how to sort the session IDs. e.g. sort=created_at|desc will sort the results based on createdAt in descending order
	*/
	Sort *string

	/* WithCommandInfo.

	   get sessions with command info included; by default sessions are returned without command info which include cloud_request_ids and logs fields
	*/
	WithCommandInfo *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the r t r audit sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRAuditSessionsParams) WithDefaults() *RTRAuditSessionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the r t r audit sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRAuditSessionsParams) SetDefaults() {
	var (
		withCommandInfoDefault = bool(false)
	)

	val := RTRAuditSessionsParams{
		WithCommandInfo: &withCommandInfoDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the r t r audit sessions params
func (o *RTRAuditSessionsParams) WithTimeout(timeout time.Duration) *RTRAuditSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r audit sessions params
func (o *RTRAuditSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r audit sessions params
func (o *RTRAuditSessionsParams) WithContext(ctx context.Context) *RTRAuditSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r audit sessions params
func (o *RTRAuditSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r audit sessions params
func (o *RTRAuditSessionsParams) WithHTTPClient(client *http.Client) *RTRAuditSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r audit sessions params
func (o *RTRAuditSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the r t r audit sessions params
func (o *RTRAuditSessionsParams) WithFilter(filter *string) *RTRAuditSessionsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the r t r audit sessions params
func (o *RTRAuditSessionsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the r t r audit sessions params
func (o *RTRAuditSessionsParams) WithLimit(limit *string) *RTRAuditSessionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the r t r audit sessions params
func (o *RTRAuditSessionsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the r t r audit sessions params
func (o *RTRAuditSessionsParams) WithOffset(offset *string) *RTRAuditSessionsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the r t r audit sessions params
func (o *RTRAuditSessionsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithSort adds the sort to the r t r audit sessions params
func (o *RTRAuditSessionsParams) WithSort(sort *string) *RTRAuditSessionsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the r t r audit sessions params
func (o *RTRAuditSessionsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithWithCommandInfo adds the withCommandInfo to the r t r audit sessions params
func (o *RTRAuditSessionsParams) WithWithCommandInfo(withCommandInfo *bool) *RTRAuditSessionsParams {
	o.SetWithCommandInfo(withCommandInfo)
	return o
}

// SetWithCommandInfo adds the withCommandInfo to the r t r audit sessions params
func (o *RTRAuditSessionsParams) SetWithCommandInfo(withCommandInfo *bool) {
	o.WithCommandInfo = withCommandInfo
}

// WriteToRequest writes these params to a swagger request
func (o *RTRAuditSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.WithCommandInfo != nil {

		// query param with_command_info
		var qrWithCommandInfo bool

		if o.WithCommandInfo != nil {
			qrWithCommandInfo = *o.WithCommandInfo
		}
		qWithCommandInfo := swag.FormatBool(qrWithCommandInfo)
		if qWithCommandInfo != "" {

			if err := r.SetQueryParam("with_command_info", qWithCommandInfo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}