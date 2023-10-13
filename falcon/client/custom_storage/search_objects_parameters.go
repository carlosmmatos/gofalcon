// Code generated by go-swagger; DO NOT EDIT.

package custom_storage

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

// NewSearchObjectsParams creates a new SearchObjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchObjectsParams() *SearchObjectsParams {
	return &SearchObjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchObjectsParamsWithTimeout creates a new SearchObjectsParams object
// with the ability to set a timeout on a request.
func NewSearchObjectsParamsWithTimeout(timeout time.Duration) *SearchObjectsParams {
	return &SearchObjectsParams{
		timeout: timeout,
	}
}

// NewSearchObjectsParamsWithContext creates a new SearchObjectsParams object
// with the ability to set a context for a request.
func NewSearchObjectsParamsWithContext(ctx context.Context) *SearchObjectsParams {
	return &SearchObjectsParams{
		Context: ctx,
	}
}

// NewSearchObjectsParamsWithHTTPClient creates a new SearchObjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchObjectsParamsWithHTTPClient(client *http.Client) *SearchObjectsParams {
	return &SearchObjectsParams{
		HTTPClient: client,
	}
}

/*
SearchObjectsParams contains all the parameters to send to the API endpoint

	for the search objects operation.

	Typically these are written to a http.Request.
*/
type SearchObjectsParams struct {

	/* XCSADBNAMESPACE.

	   The name of the namespace the collection belongs to
	*/
	XCSADBNAMESPACE string

	/* XCSAPPID.

	   The id of the app the collection belongs to. This will map to the namespace of the collection
	*/
	XCSAPPID *string

	/* CollectionName.

	   The name of the collection
	*/
	CollectionName string

	/* Filter.

	   The filter to limit the returned results.
	*/
	Filter string

	/* Limit.

	   The limit of results to return
	*/
	Limit int64

	/* Offset.

	   The offset of results to return
	*/
	Offset int64

	/* Sort.

	   The sort order for the returned results.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchObjectsParams) WithDefaults() *SearchObjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchObjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search objects params
func (o *SearchObjectsParams) WithTimeout(timeout time.Duration) *SearchObjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search objects params
func (o *SearchObjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search objects params
func (o *SearchObjectsParams) WithContext(ctx context.Context) *SearchObjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search objects params
func (o *SearchObjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search objects params
func (o *SearchObjectsParams) WithHTTPClient(client *http.Client) *SearchObjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search objects params
func (o *SearchObjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSADBNAMESPACE adds the xCSADBNAMESPACE to the search objects params
func (o *SearchObjectsParams) WithXCSADBNAMESPACE(xCSADBNAMESPACE string) *SearchObjectsParams {
	o.SetXCSADBNAMESPACE(xCSADBNAMESPACE)
	return o
}

// SetXCSADBNAMESPACE adds the xCSADBNAMESPACE to the search objects params
func (o *SearchObjectsParams) SetXCSADBNAMESPACE(xCSADBNAMESPACE string) {
	o.XCSADBNAMESPACE = xCSADBNAMESPACE
}

// WithXCSAPPID adds the xCSAPPID to the search objects params
func (o *SearchObjectsParams) WithXCSAPPID(xCSAPPID *string) *SearchObjectsParams {
	o.SetXCSAPPID(xCSAPPID)
	return o
}

// SetXCSAPPID adds the xCSAPPId to the search objects params
func (o *SearchObjectsParams) SetXCSAPPID(xCSAPPID *string) {
	o.XCSAPPID = xCSAPPID
}

// WithCollectionName adds the collectionName to the search objects params
func (o *SearchObjectsParams) WithCollectionName(collectionName string) *SearchObjectsParams {
	o.SetCollectionName(collectionName)
	return o
}

// SetCollectionName adds the collectionName to the search objects params
func (o *SearchObjectsParams) SetCollectionName(collectionName string) {
	o.CollectionName = collectionName
}

// WithFilter adds the filter to the search objects params
func (o *SearchObjectsParams) WithFilter(filter string) *SearchObjectsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the search objects params
func (o *SearchObjectsParams) SetFilter(filter string) {
	o.Filter = filter
}

// WithLimit adds the limit to the search objects params
func (o *SearchObjectsParams) WithLimit(limit int64) *SearchObjectsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search objects params
func (o *SearchObjectsParams) SetLimit(limit int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search objects params
func (o *SearchObjectsParams) WithOffset(offset int64) *SearchObjectsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search objects params
func (o *SearchObjectsParams) SetOffset(offset int64) {
	o.Offset = offset
}

// WithSort adds the sort to the search objects params
func (o *SearchObjectsParams) WithSort(sort *string) *SearchObjectsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search objects params
func (o *SearchObjectsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *SearchObjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-CS-ADB-NAMESPACE
	if err := r.SetHeaderParam("X-CS-ADB-NAMESPACE", o.XCSADBNAMESPACE); err != nil {
		return err
	}

	if o.XCSAPPID != nil {

		// header param X-CS-APP-ID
		if err := r.SetHeaderParam("X-CS-APP-ID", *o.XCSAPPID); err != nil {
			return err
		}
	}

	// path param collection_name
	if err := r.SetPathParam("collection_name", o.CollectionName); err != nil {
		return err
	}

	// query param filter
	qrFilter := o.Filter
	qFilter := qrFilter
	if qFilter != "" {

		if err := r.SetQueryParam("filter", qFilter); err != nil {
			return err
		}
	}

	// query param limit
	qrLimit := o.Limit
	qLimit := swag.FormatInt64(qrLimit)

	if err := r.SetQueryParam("limit", qLimit); err != nil {
		return err
	}

	// query param offset
	qrOffset := o.Offset
	qOffset := swag.FormatInt64(qrOffset)

	if err := r.SetQueryParam("offset", qOffset); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
