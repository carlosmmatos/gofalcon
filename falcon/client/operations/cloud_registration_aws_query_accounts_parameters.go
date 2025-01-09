// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewCloudRegistrationAwsQueryAccountsParams creates a new CloudRegistrationAwsQueryAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudRegistrationAwsQueryAccountsParams() *CloudRegistrationAwsQueryAccountsParams {
	return &CloudRegistrationAwsQueryAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudRegistrationAwsQueryAccountsParamsWithTimeout creates a new CloudRegistrationAwsQueryAccountsParams object
// with the ability to set a timeout on a request.
func NewCloudRegistrationAwsQueryAccountsParamsWithTimeout(timeout time.Duration) *CloudRegistrationAwsQueryAccountsParams {
	return &CloudRegistrationAwsQueryAccountsParams{
		timeout: timeout,
	}
}

// NewCloudRegistrationAwsQueryAccountsParamsWithContext creates a new CloudRegistrationAwsQueryAccountsParams object
// with the ability to set a context for a request.
func NewCloudRegistrationAwsQueryAccountsParamsWithContext(ctx context.Context) *CloudRegistrationAwsQueryAccountsParams {
	return &CloudRegistrationAwsQueryAccountsParams{
		Context: ctx,
	}
}

// NewCloudRegistrationAwsQueryAccountsParamsWithHTTPClient creates a new CloudRegistrationAwsQueryAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudRegistrationAwsQueryAccountsParamsWithHTTPClient(client *http.Client) *CloudRegistrationAwsQueryAccountsParams {
	return &CloudRegistrationAwsQueryAccountsParams{
		HTTPClient: client,
	}
}

/*
CloudRegistrationAwsQueryAccountsParams contains all the parameters to send to the API endpoint

	for the cloud registration aws query accounts operation.

	Typically these are written to a http.Request.
*/
type CloudRegistrationAwsQueryAccountsParams struct {

	/* AccountStatus.

	   Account status to filter results by.
	*/
	AccountStatus *string

	/* Features.

	   Features registered for an account
	*/
	Features []string

	/* GroupBy.

	   Field to group by.
	*/
	GroupBy *string

	/* Limit.

	   The maximum number of items to return. When not specified or 0, 100 is used. When larger than 500, 500 is used.

	   Default: 100
	*/
	Limit *int64

	/* Offset.

	   The offset to start retrieving records from.
	*/
	Offset *int64

	/* OrganizationIds.

	   Organization IDs used to filter accounts
	*/
	OrganizationIds []string

	/* Products.

	   Products registered for an account
	*/
	Products []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cloud registration aws query accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudRegistrationAwsQueryAccountsParams) WithDefaults() *CloudRegistrationAwsQueryAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud registration aws query accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudRegistrationAwsQueryAccountsParams) SetDefaults() {
	var (
		limitDefault = int64(100)
	)

	val := CloudRegistrationAwsQueryAccountsParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) WithTimeout(timeout time.Duration) *CloudRegistrationAwsQueryAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) WithContext(ctx context.Context) *CloudRegistrationAwsQueryAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) WithHTTPClient(client *http.Client) *CloudRegistrationAwsQueryAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountStatus adds the accountStatus to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) WithAccountStatus(accountStatus *string) *CloudRegistrationAwsQueryAccountsParams {
	o.SetAccountStatus(accountStatus)
	return o
}

// SetAccountStatus adds the accountStatus to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) SetAccountStatus(accountStatus *string) {
	o.AccountStatus = accountStatus
}

// WithFeatures adds the features to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) WithFeatures(features []string) *CloudRegistrationAwsQueryAccountsParams {
	o.SetFeatures(features)
	return o
}

// SetFeatures adds the features to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) SetFeatures(features []string) {
	o.Features = features
}

// WithGroupBy adds the groupBy to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) WithGroupBy(groupBy *string) *CloudRegistrationAwsQueryAccountsParams {
	o.SetGroupBy(groupBy)
	return o
}

// SetGroupBy adds the groupBy to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) SetGroupBy(groupBy *string) {
	o.GroupBy = groupBy
}

// WithLimit adds the limit to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) WithLimit(limit *int64) *CloudRegistrationAwsQueryAccountsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) WithOffset(offset *int64) *CloudRegistrationAwsQueryAccountsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrganizationIds adds the organizationIds to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) WithOrganizationIds(organizationIds []string) *CloudRegistrationAwsQueryAccountsParams {
	o.SetOrganizationIds(organizationIds)
	return o
}

// SetOrganizationIds adds the organizationIds to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) SetOrganizationIds(organizationIds []string) {
	o.OrganizationIds = organizationIds
}

// WithProducts adds the products to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) WithProducts(products []string) *CloudRegistrationAwsQueryAccountsParams {
	o.SetProducts(products)
	return o
}

// SetProducts adds the products to the cloud registration aws query accounts params
func (o *CloudRegistrationAwsQueryAccountsParams) SetProducts(products []string) {
	o.Products = products
}

// WriteToRequest writes these params to a swagger request
func (o *CloudRegistrationAwsQueryAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountStatus != nil {

		// query param account-status
		var qrAccountStatus string

		if o.AccountStatus != nil {
			qrAccountStatus = *o.AccountStatus
		}
		qAccountStatus := qrAccountStatus
		if qAccountStatus != "" {

			if err := r.SetQueryParam("account-status", qAccountStatus); err != nil {
				return err
			}
		}
	}

	if o.Features != nil {

		// binding items for features
		joinedFeatures := o.bindParamFeatures(reg)

		// query array param features
		if err := r.SetQueryParam("features", joinedFeatures...); err != nil {
			return err
		}
	}

	if o.GroupBy != nil {

		// query param group_by
		var qrGroupBy string

		if o.GroupBy != nil {
			qrGroupBy = *o.GroupBy
		}
		qGroupBy := qrGroupBy
		if qGroupBy != "" {

			if err := r.SetQueryParam("group_by", qGroupBy); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.OrganizationIds != nil {

		// binding items for organization-ids
		joinedOrganizationIds := o.bindParamOrganizationIds(reg)

		// query array param organization-ids
		if err := r.SetQueryParam("organization-ids", joinedOrganizationIds...); err != nil {
			return err
		}
	}

	if o.Products != nil {

		// binding items for products
		joinedProducts := o.bindParamProducts(reg)

		// query array param products
		if err := r.SetQueryParam("products", joinedProducts...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCloudRegistrationAwsQueryAccounts binds the parameter features
func (o *CloudRegistrationAwsQueryAccountsParams) bindParamFeatures(formats strfmt.Registry) []string {
	featuresIR := o.Features

	var featuresIC []string
	for _, featuresIIR := range featuresIR { // explode []string

		featuresIIV := featuresIIR // string as string
		featuresIC = append(featuresIC, featuresIIV)
	}

	// items.CollectionFormat: "multi"
	featuresIS := swag.JoinByFormat(featuresIC, "multi")

	return featuresIS
}

// bindParamCloudRegistrationAwsQueryAccounts binds the parameter organization-ids
func (o *CloudRegistrationAwsQueryAccountsParams) bindParamOrganizationIds(formats strfmt.Registry) []string {
	organizationIdsIR := o.OrganizationIds

	var organizationIdsIC []string
	for _, organizationIdsIIR := range organizationIdsIR { // explode []string

		organizationIdsIIV := organizationIdsIIR // string as string
		organizationIdsIC = append(organizationIdsIC, organizationIdsIIV)
	}

	// items.CollectionFormat: "multi"
	organizationIdsIS := swag.JoinByFormat(organizationIdsIC, "multi")

	return organizationIdsIS
}

// bindParamCloudRegistrationAwsQueryAccounts binds the parameter products
func (o *CloudRegistrationAwsQueryAccountsParams) bindParamProducts(formats strfmt.Registry) []string {
	productsIR := o.Products

	var productsIC []string
	for _, productsIIR := range productsIR { // explode []string

		productsIIV := productsIIR // string as string
		productsIC = append(productsIC, productsIIV)
	}

	// items.CollectionFormat: "multi"
	productsIS := swag.JoinByFormat(productsIC, "multi")

	return productsIS
}