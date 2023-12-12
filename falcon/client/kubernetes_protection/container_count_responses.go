// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// ContainerCountReader is a Reader for the ContainerCount structure.
type ContainerCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewContainerCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewContainerCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/containers/count/v1] ContainerCount", response, response.Code())
	}
}

// NewContainerCountOK creates a ContainerCountOK with default headers values
func NewContainerCountOK() *ContainerCountOK {
	return &ContainerCountOK{}
}

/*
ContainerCountOK describes a response with status code 200, with default header values.

OK
*/
type ContainerCountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CommonCountResponse
}

// IsSuccess returns true when this container count o k response has a 2xx status code
func (o *ContainerCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container count o k response has a 3xx status code
func (o *ContainerCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container count o k response has a 4xx status code
func (o *ContainerCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container count o k response has a 5xx status code
func (o *ContainerCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container count o k response a status code equal to that given
func (o *ContainerCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the container count o k response
func (o *ContainerCountOK) Code() int {
	return 200
}

func (o *ContainerCountOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count/v1][%d] containerCountOK  %+v", 200, o.Payload)
}

func (o *ContainerCountOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count/v1][%d] containerCountOK  %+v", 200, o.Payload)
}

func (o *ContainerCountOK) GetPayload() *models.CommonCountResponse {
	return o.Payload
}

func (o *ContainerCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CommonCountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerCountForbidden creates a ContainerCountForbidden with default headers values
func NewContainerCountForbidden() *ContainerCountForbidden {
	return &ContainerCountForbidden{}
}

/*
ContainerCountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ContainerCountForbidden struct {

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

// IsSuccess returns true when this container count forbidden response has a 2xx status code
func (o *ContainerCountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container count forbidden response has a 3xx status code
func (o *ContainerCountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container count forbidden response has a 4xx status code
func (o *ContainerCountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this container count forbidden response has a 5xx status code
func (o *ContainerCountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this container count forbidden response a status code equal to that given
func (o *ContainerCountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the container count forbidden response
func (o *ContainerCountForbidden) Code() int {
	return 403
}

func (o *ContainerCountForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count/v1][%d] containerCountForbidden  %+v", 403, o.Payload)
}

func (o *ContainerCountForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count/v1][%d] containerCountForbidden  %+v", 403, o.Payload)
}

func (o *ContainerCountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ContainerCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewContainerCountTooManyRequests creates a ContainerCountTooManyRequests with default headers values
func NewContainerCountTooManyRequests() *ContainerCountTooManyRequests {
	return &ContainerCountTooManyRequests{}
}

/*
ContainerCountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ContainerCountTooManyRequests struct {

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

// IsSuccess returns true when this container count too many requests response has a 2xx status code
func (o *ContainerCountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container count too many requests response has a 3xx status code
func (o *ContainerCountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container count too many requests response has a 4xx status code
func (o *ContainerCountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this container count too many requests response has a 5xx status code
func (o *ContainerCountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this container count too many requests response a status code equal to that given
func (o *ContainerCountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the container count too many requests response
func (o *ContainerCountTooManyRequests) Code() int {
	return 429
}

func (o *ContainerCountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count/v1][%d] containerCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ContainerCountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count/v1][%d] containerCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ContainerCountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ContainerCountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewContainerCountInternalServerError creates a ContainerCountInternalServerError with default headers values
func NewContainerCountInternalServerError() *ContainerCountInternalServerError {
	return &ContainerCountInternalServerError{}
}

/*
ContainerCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ContainerCountInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CoreEntitiesResponse
}

// IsSuccess returns true when this container count internal server error response has a 2xx status code
func (o *ContainerCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container count internal server error response has a 3xx status code
func (o *ContainerCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container count internal server error response has a 4xx status code
func (o *ContainerCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container count internal server error response has a 5xx status code
func (o *ContainerCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container count internal server error response a status code equal to that given
func (o *ContainerCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the container count internal server error response
func (o *ContainerCountInternalServerError) Code() int {
	return 500
}

func (o *ContainerCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count/v1][%d] containerCountInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerCountInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count/v1][%d] containerCountInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerCountInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *ContainerCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CoreEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}