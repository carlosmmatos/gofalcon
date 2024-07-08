// Code generated by go-swagger; DO NOT EDIT.

package compliance_assessments

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
)

// NewExtAggregateImageAssessmentsParams creates a new ExtAggregateImageAssessmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtAggregateImageAssessmentsParams() *ExtAggregateImageAssessmentsParams {
	return &ExtAggregateImageAssessmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtAggregateImageAssessmentsParamsWithTimeout creates a new ExtAggregateImageAssessmentsParams object
// with the ability to set a timeout on a request.
func NewExtAggregateImageAssessmentsParamsWithTimeout(timeout time.Duration) *ExtAggregateImageAssessmentsParams {
	return &ExtAggregateImageAssessmentsParams{
		timeout: timeout,
	}
}

// NewExtAggregateImageAssessmentsParamsWithContext creates a new ExtAggregateImageAssessmentsParams object
// with the ability to set a context for a request.
func NewExtAggregateImageAssessmentsParamsWithContext(ctx context.Context) *ExtAggregateImageAssessmentsParams {
	return &ExtAggregateImageAssessmentsParams{
		Context: ctx,
	}
}

// NewExtAggregateImageAssessmentsParamsWithHTTPClient creates a new ExtAggregateImageAssessmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtAggregateImageAssessmentsParamsWithHTTPClient(client *http.Client) *ExtAggregateImageAssessmentsParams {
	return &ExtAggregateImageAssessmentsParams{
		HTTPClient: client,
	}
}

/*
ExtAggregateImageAssessmentsParams contains all the parameters to send to the API endpoint

	for the ext aggregate image assessments operation.

	Typically these are written to a http.Request.
*/
type ExtAggregateImageAssessmentsParams struct {

	/* After.

	   'after' value from the last response. Keep it empty for the first request.
	*/
	After *string

	/* Filter.

	     Filter results using a query in Falcon Query Language (FQL). Supported Filters:
	image_digest: Image digest (sha256 digest)
	cloud_info.cluster_name: Kubernetes cluster name
	compliance_finding.framework: Compliance finding framework (available values: CIS)
	cid: Customer ID
	compliance_finding.id: Compliance finding ID
	cloud_info.namespace: Kubernetes namespace
	image_repository: Image repository
	compliance_finding.severity: Compliance finding severity; available values: 4, 3, 2, 1 (4: critical, 3: high, 2: medium, 1:low)
	cloud_info.cloud_account_id: Cloud account ID
	image_registry: Image registry
	image_id: Image ID
	image_tag: Image tag
	compliance_finding.name: Compliance finding Name
	asset_type: asset type (container, image)
	cloud_info.cloud_provider: Cloud provider
	cloud_info.cloud_region: Cloud region

	*/
	Filter *string

	/* Limit.

	   number of images to return in the response after 'after' key. Keep it empty for the default number of 10000
	*/
	Limit *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ext aggregate image assessments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtAggregateImageAssessmentsParams) WithDefaults() *ExtAggregateImageAssessmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ext aggregate image assessments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtAggregateImageAssessmentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ext aggregate image assessments params
func (o *ExtAggregateImageAssessmentsParams) WithTimeout(timeout time.Duration) *ExtAggregateImageAssessmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ext aggregate image assessments params
func (o *ExtAggregateImageAssessmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ext aggregate image assessments params
func (o *ExtAggregateImageAssessmentsParams) WithContext(ctx context.Context) *ExtAggregateImageAssessmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ext aggregate image assessments params
func (o *ExtAggregateImageAssessmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ext aggregate image assessments params
func (o *ExtAggregateImageAssessmentsParams) WithHTTPClient(client *http.Client) *ExtAggregateImageAssessmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ext aggregate image assessments params
func (o *ExtAggregateImageAssessmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the ext aggregate image assessments params
func (o *ExtAggregateImageAssessmentsParams) WithAfter(after *string) *ExtAggregateImageAssessmentsParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the ext aggregate image assessments params
func (o *ExtAggregateImageAssessmentsParams) SetAfter(after *string) {
	o.After = after
}

// WithFilter adds the filter to the ext aggregate image assessments params
func (o *ExtAggregateImageAssessmentsParams) WithFilter(filter *string) *ExtAggregateImageAssessmentsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the ext aggregate image assessments params
func (o *ExtAggregateImageAssessmentsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the ext aggregate image assessments params
func (o *ExtAggregateImageAssessmentsParams) WithLimit(limit *string) *ExtAggregateImageAssessmentsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ext aggregate image assessments params
func (o *ExtAggregateImageAssessmentsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *ExtAggregateImageAssessmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}