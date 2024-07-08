// Code generated by go-swagger; DO NOT EDIT.

package host_migration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// GetMigrationIDsV1Reader is a Reader for the GetMigrationIDsV1 structure.
type GetMigrationIDsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMigrationIDsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMigrationIDsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMigrationIDsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMigrationIDsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMigrationIDsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMigrationIDsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /host-migration/queries/migrations/v1] GetMigrationIDsV1", response, response.Code())
	}
}

// NewGetMigrationIDsV1OK creates a GetMigrationIDsV1OK with default headers values
func NewGetMigrationIDsV1OK() *GetMigrationIDsV1OK {
	return &GetMigrationIDsV1OK{}
}

/*
GetMigrationIDsV1OK describes a response with status code 200, with default header values.

OK
*/
type GetMigrationIDsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this get migration i ds v1 o k response has a 2xx status code
func (o *GetMigrationIDsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get migration i ds v1 o k response has a 3xx status code
func (o *GetMigrationIDsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration i ds v1 o k response has a 4xx status code
func (o *GetMigrationIDsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get migration i ds v1 o k response has a 5xx status code
func (o *GetMigrationIDsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get migration i ds v1 o k response a status code equal to that given
func (o *GetMigrationIDsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get migration i ds v1 o k response
func (o *GetMigrationIDsV1OK) Code() int {
	return 200
}

func (o *GetMigrationIDsV1OK) Error() string {
	return fmt.Sprintf("[GET /host-migration/queries/migrations/v1][%d] getMigrationIDsV1OK  %+v", 200, o.Payload)
}

func (o *GetMigrationIDsV1OK) String() string {
	return fmt.Sprintf("[GET /host-migration/queries/migrations/v1][%d] getMigrationIDsV1OK  %+v", 200, o.Payload)
}

func (o *GetMigrationIDsV1OK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *GetMigrationIDsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrationIDsV1BadRequest creates a GetMigrationIDsV1BadRequest with default headers values
func NewGetMigrationIDsV1BadRequest() *GetMigrationIDsV1BadRequest {
	return &GetMigrationIDsV1BadRequest{}
}

/*
GetMigrationIDsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetMigrationIDsV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this get migration i ds v1 bad request response has a 2xx status code
func (o *GetMigrationIDsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migration i ds v1 bad request response has a 3xx status code
func (o *GetMigrationIDsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration i ds v1 bad request response has a 4xx status code
func (o *GetMigrationIDsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migration i ds v1 bad request response has a 5xx status code
func (o *GetMigrationIDsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get migration i ds v1 bad request response a status code equal to that given
func (o *GetMigrationIDsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get migration i ds v1 bad request response
func (o *GetMigrationIDsV1BadRequest) Code() int {
	return 400
}

func (o *GetMigrationIDsV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /host-migration/queries/migrations/v1][%d] getMigrationIDsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetMigrationIDsV1BadRequest) String() string {
	return fmt.Sprintf("[GET /host-migration/queries/migrations/v1][%d] getMigrationIDsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetMigrationIDsV1BadRequest) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *GetMigrationIDsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrationIDsV1Forbidden creates a GetMigrationIDsV1Forbidden with default headers values
func NewGetMigrationIDsV1Forbidden() *GetMigrationIDsV1Forbidden {
	return &GetMigrationIDsV1Forbidden{}
}

/*
GetMigrationIDsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetMigrationIDsV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this get migration i ds v1 forbidden response has a 2xx status code
func (o *GetMigrationIDsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migration i ds v1 forbidden response has a 3xx status code
func (o *GetMigrationIDsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration i ds v1 forbidden response has a 4xx status code
func (o *GetMigrationIDsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migration i ds v1 forbidden response has a 5xx status code
func (o *GetMigrationIDsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get migration i ds v1 forbidden response a status code equal to that given
func (o *GetMigrationIDsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get migration i ds v1 forbidden response
func (o *GetMigrationIDsV1Forbidden) Code() int {
	return 403
}

func (o *GetMigrationIDsV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /host-migration/queries/migrations/v1][%d] getMigrationIDsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMigrationIDsV1Forbidden) String() string {
	return fmt.Sprintf("[GET /host-migration/queries/migrations/v1][%d] getMigrationIDsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMigrationIDsV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetMigrationIDsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrationIDsV1TooManyRequests creates a GetMigrationIDsV1TooManyRequests with default headers values
func NewGetMigrationIDsV1TooManyRequests() *GetMigrationIDsV1TooManyRequests {
	return &GetMigrationIDsV1TooManyRequests{}
}

/*
GetMigrationIDsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetMigrationIDsV1TooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this get migration i ds v1 too many requests response has a 2xx status code
func (o *GetMigrationIDsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migration i ds v1 too many requests response has a 3xx status code
func (o *GetMigrationIDsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration i ds v1 too many requests response has a 4xx status code
func (o *GetMigrationIDsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migration i ds v1 too many requests response has a 5xx status code
func (o *GetMigrationIDsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get migration i ds v1 too many requests response a status code equal to that given
func (o *GetMigrationIDsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get migration i ds v1 too many requests response
func (o *GetMigrationIDsV1TooManyRequests) Code() int {
	return 429
}

func (o *GetMigrationIDsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /host-migration/queries/migrations/v1][%d] getMigrationIDsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMigrationIDsV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /host-migration/queries/migrations/v1][%d] getMigrationIDsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMigrationIDsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetMigrationIDsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrationIDsV1InternalServerError creates a GetMigrationIDsV1InternalServerError with default headers values
func NewGetMigrationIDsV1InternalServerError() *GetMigrationIDsV1InternalServerError {
	return &GetMigrationIDsV1InternalServerError{}
}

/*
GetMigrationIDsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetMigrationIDsV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this get migration i ds v1 internal server error response has a 2xx status code
func (o *GetMigrationIDsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migration i ds v1 internal server error response has a 3xx status code
func (o *GetMigrationIDsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration i ds v1 internal server error response has a 4xx status code
func (o *GetMigrationIDsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get migration i ds v1 internal server error response has a 5xx status code
func (o *GetMigrationIDsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get migration i ds v1 internal server error response a status code equal to that given
func (o *GetMigrationIDsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get migration i ds v1 internal server error response
func (o *GetMigrationIDsV1InternalServerError) Code() int {
	return 500
}

func (o *GetMigrationIDsV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /host-migration/queries/migrations/v1][%d] getMigrationIDsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMigrationIDsV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /host-migration/queries/migrations/v1][%d] getMigrationIDsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMigrationIDsV1InternalServerError) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *GetMigrationIDsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}