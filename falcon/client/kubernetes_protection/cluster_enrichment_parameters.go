// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// NewClusterEnrichmentParams creates a new ClusterEnrichmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterEnrichmentParams() *ClusterEnrichmentParams {
	return &ClusterEnrichmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterEnrichmentParamsWithTimeout creates a new ClusterEnrichmentParams object
// with the ability to set a timeout on a request.
func NewClusterEnrichmentParamsWithTimeout(timeout time.Duration) *ClusterEnrichmentParams {
	return &ClusterEnrichmentParams{
		timeout: timeout,
	}
}

// NewClusterEnrichmentParamsWithContext creates a new ClusterEnrichmentParams object
// with the ability to set a context for a request.
func NewClusterEnrichmentParamsWithContext(ctx context.Context) *ClusterEnrichmentParams {
	return &ClusterEnrichmentParams{
		Context: ctx,
	}
}

// NewClusterEnrichmentParamsWithHTTPClient creates a new ClusterEnrichmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterEnrichmentParamsWithHTTPClient(client *http.Client) *ClusterEnrichmentParams {
	return &ClusterEnrichmentParams{
		HTTPClient: client,
	}
}

/*
ClusterEnrichmentParams contains all the parameters to send to the API endpoint

	for the cluster enrichment operation.

	Typically these are written to a http.Request.
*/
type ClusterEnrichmentParams struct {

	/* ClusterID.

	   One or more cluster ids for which to retrieve enrichment info
	*/
	ClusterID []string

	/* Filter.

	   Supported filters:  last_seen
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster enrichment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterEnrichmentParams) WithDefaults() *ClusterEnrichmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster enrichment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterEnrichmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster enrichment params
func (o *ClusterEnrichmentParams) WithTimeout(timeout time.Duration) *ClusterEnrichmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster enrichment params
func (o *ClusterEnrichmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster enrichment params
func (o *ClusterEnrichmentParams) WithContext(ctx context.Context) *ClusterEnrichmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster enrichment params
func (o *ClusterEnrichmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster enrichment params
func (o *ClusterEnrichmentParams) WithHTTPClient(client *http.Client) *ClusterEnrichmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster enrichment params
func (o *ClusterEnrichmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the cluster enrichment params
func (o *ClusterEnrichmentParams) WithClusterID(clusterID []string) *ClusterEnrichmentParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the cluster enrichment params
func (o *ClusterEnrichmentParams) SetClusterID(clusterID []string) {
	o.ClusterID = clusterID
}

// WithFilter adds the filter to the cluster enrichment params
func (o *ClusterEnrichmentParams) WithFilter(filter *string) *ClusterEnrichmentParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the cluster enrichment params
func (o *ClusterEnrichmentParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterEnrichmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterID != nil {

		// binding items for cluster_id
		joinedClusterID := o.bindParamClusterID(reg)

		// query array param cluster_id
		if err := r.SetQueryParam("cluster_id", joinedClusterID...); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamClusterEnrichment binds the parameter cluster_id
func (o *ClusterEnrichmentParams) bindParamClusterID(formats strfmt.Registry) []string {
	clusterIDIR := o.ClusterID

	var clusterIDIC []string
	for _, clusterIDIIR := range clusterIDIR { // explode []string

		clusterIDIIV := clusterIDIIR // string as string
		clusterIDIC = append(clusterIDIC, clusterIDIIV)
	}

	// items.CollectionFormat: "csv"
	clusterIDIS := swag.JoinByFormat(clusterIDIC, "csv")

	return clusterIDIS
}