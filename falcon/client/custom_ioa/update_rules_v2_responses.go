// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// UpdateRulesV2Reader is a Reader for the UpdateRulesV2 structure.
type UpdateRulesV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRulesV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRulesV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateRulesV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRulesV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateRulesV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /ioarules/entities/rules/v2] update-rules-v2", response, response.Code())
	}
}

// NewUpdateRulesV2OK creates a UpdateRulesV2OK with default headers values
func NewUpdateRulesV2OK() *UpdateRulesV2OK {
	return &UpdateRulesV2OK{}
}

/*
UpdateRulesV2OK describes a response with status code 200, with default header values.

OK
*/
type UpdateRulesV2OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIRulesResponse
}

// IsSuccess returns true when this update rules v2 o k response has a 2xx status code
func (o *UpdateRulesV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update rules v2 o k response has a 3xx status code
func (o *UpdateRulesV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rules v2 o k response has a 4xx status code
func (o *UpdateRulesV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update rules v2 o k response has a 5xx status code
func (o *UpdateRulesV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this update rules v2 o k response a status code equal to that given
func (o *UpdateRulesV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update rules v2 o k response
func (o *UpdateRulesV2OK) Code() int {
	return 200
}

func (o *UpdateRulesV2OK) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v2][%d] updateRulesV2OK  %+v", 200, o.Payload)
}

func (o *UpdateRulesV2OK) String() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v2][%d] updateRulesV2OK  %+v", 200, o.Payload)
}

func (o *UpdateRulesV2OK) GetPayload() *models.APIRulesResponse {
	return o.Payload
}

func (o *UpdateRulesV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRulesV2Forbidden creates a UpdateRulesV2Forbidden with default headers values
func NewUpdateRulesV2Forbidden() *UpdateRulesV2Forbidden {
	return &UpdateRulesV2Forbidden{}
}

/*
UpdateRulesV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateRulesV2Forbidden struct {

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

// IsSuccess returns true when this update rules v2 forbidden response has a 2xx status code
func (o *UpdateRulesV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rules v2 forbidden response has a 3xx status code
func (o *UpdateRulesV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rules v2 forbidden response has a 4xx status code
func (o *UpdateRulesV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rules v2 forbidden response has a 5xx status code
func (o *UpdateRulesV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update rules v2 forbidden response a status code equal to that given
func (o *UpdateRulesV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update rules v2 forbidden response
func (o *UpdateRulesV2Forbidden) Code() int {
	return 403
}

func (o *UpdateRulesV2Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v2][%d] updateRulesV2Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateRulesV2Forbidden) String() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v2][%d] updateRulesV2Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateRulesV2Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRulesV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRulesV2NotFound creates a UpdateRulesV2NotFound with default headers values
func NewUpdateRulesV2NotFound() *UpdateRulesV2NotFound {
	return &UpdateRulesV2NotFound{}
}

/*
UpdateRulesV2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateRulesV2NotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this update rules v2 not found response has a 2xx status code
func (o *UpdateRulesV2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rules v2 not found response has a 3xx status code
func (o *UpdateRulesV2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rules v2 not found response has a 4xx status code
func (o *UpdateRulesV2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rules v2 not found response has a 5xx status code
func (o *UpdateRulesV2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update rules v2 not found response a status code equal to that given
func (o *UpdateRulesV2NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update rules v2 not found response
func (o *UpdateRulesV2NotFound) Code() int {
	return 404
}

func (o *UpdateRulesV2NotFound) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v2][%d] updateRulesV2NotFound  %+v", 404, o.Payload)
}

func (o *UpdateRulesV2NotFound) String() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v2][%d] updateRulesV2NotFound  %+v", 404, o.Payload)
}

func (o *UpdateRulesV2NotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *UpdateRulesV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRulesV2TooManyRequests creates a UpdateRulesV2TooManyRequests with default headers values
func NewUpdateRulesV2TooManyRequests() *UpdateRulesV2TooManyRequests {
	return &UpdateRulesV2TooManyRequests{}
}

/*
UpdateRulesV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateRulesV2TooManyRequests struct {

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

// IsSuccess returns true when this update rules v2 too many requests response has a 2xx status code
func (o *UpdateRulesV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rules v2 too many requests response has a 3xx status code
func (o *UpdateRulesV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rules v2 too many requests response has a 4xx status code
func (o *UpdateRulesV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rules v2 too many requests response has a 5xx status code
func (o *UpdateRulesV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update rules v2 too many requests response a status code equal to that given
func (o *UpdateRulesV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update rules v2 too many requests response
func (o *UpdateRulesV2TooManyRequests) Code() int {
	return 429
}

func (o *UpdateRulesV2TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v2][%d] updateRulesV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateRulesV2TooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v2][%d] updateRulesV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateRulesV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRulesV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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