// Package yidunerr represents API error interface accessors for the SDK.
package yidunerr

// An Error wraps lower level errors with code, message and an original error.
// The underlying concrete error type may also satisfy other interfaces which
// can be to used to obtain more specific information about the error.
//
// Calling Error() or String() will always include the full information about
// an error based on its underlying type.
//
// Example:
//
// output, err := s3manage.Upload(svc, input, opts)
// if err != nil {
//     if yidunErr, ok := err.(yidunerr.Error); ok {
//         // Get error details
//         log.Println("Error:", yidunerr.Code(), yidunerr.Message())

//         // Prints out full error message, including original error if there was one.
//         log.Println("Error:", yidunerr.Error())

//	        // Get original error
//	        if origErr := yidunerr.OrigErr(); origErr != nil {
//	            // operate on original error.
//	        }
//	    } else {
//	        fmt.Println(err.Error())
//	    }
//	}
type Error interface {
	// Satisfy the generic error interface.
	error

	// Returns the short phrase depicting the classification of the error.
	Code() string

	// Returns the error details message.
	Message() string

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErr() error
}

// BatchError is a batch of errors which also wraps lower level errors with
// code, message, and original errors. Calling Error() will include all errors
// that occurred in the batch.
//
// Deprecated: Replaced with BatchedErrors. Only defined for backwards
// compatibility.
type BatchError interface {
	// Satisfy the generic error interface.
	error

	// Returns the short phrase depicting the classification of the error.
	Code() string

	// Returns the error details message.
	Message() string

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErrs() []error
}

// BatchedErrors is a batch of errors which also wraps lower level errors with
// code, message, and original errors. Calling Error() will include all errors
// that occurred in the batch.
//
// Replaces BatchError
type BatchedErrors interface {
	// Satisfy the base Error interface.
	Error

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErrs() []error
}

// New returns an Error object described by the code, message, and origErr.
//
// If origErr satisfies the Error interface it will not be wrapped within a new
// Error object and will instead be returned.
func New(code, message string, origErr error) Error {
	var errs []error
	if origErr != nil {
		errs = append(errs, origErr)
	}
	return newBaseError(code, message, errs)
}

// NewBatchError returns an BatchedErrors with a collection of errors as an
// array of errors.
func NewBatchError(code, message string, errs []error) BatchedErrors {
	return newBaseError(code, message, errs)
}

// A RequestFailure is an interface to extract request failure information from
// an Error such as the request ID of the failed request returned by a service.
// prior to reaching the service such as a connection error.
//
// Example:
//
//	output, err := s3manage.Upload(svc, input, opts)
//	if err != nil {
//	    if reqerr, ok := err.(RequestFailure); ok {
//	        log.Println("Request failed", reqerr.Code(), reqerr.Message())
//	    } else {
//	        log.Println("Error:", err.Error())
//	    }
//	}
//
// Combined with yidunerr.Error:
//
//	output, err := s3manage.Upload(svc, input, opts)
//	if err != nil {
//	    if yidunErr, ok := err.(yidunerr.Error); ok {
//	        // Generic Yidun Error with Code, Message, and original error (if any)
//	        fmt.Println(yidunErr.Code(), yidunErr.Message(), yidunErr.OrigErr())
//
//	        if reqErr, ok := err.(yidunerr.RequestFailure); ok {
//	            // A service error occurred
//	            fmt.Println(reqErr.StatusCode())
//	        }
//	    } else {
//	        fmt.Println(err.Error())
//	    }
//	}
type RequestFailure interface {
	Error

	// The status code of the HTTP response.
	StatusCode() int
}

// NewRequestFailure returns a wrapped error with additional information for
// request status code.
//
// Should be used to wrap all request which involve service requests. Even if
// the request failed without a service response, but had an HTTP status code
// that may be meaningful.
func NewRequestFailure(err Error, statusCode int) RequestFailure {
	return newRequestError(err, statusCode)
}

// UnmarshalError provides the interface for the SDK failing to unmarshal data.
type UnmarshalError interface {
	yidunError
	Bytes() []byte
}

// NewUnmarshalError returns an initialized UnmarshalError error wrapper adding
// the bytes that fail to unmarshal to the error.
func NewUnmarshalError(err error, msg string, bytes []byte) UnmarshalError {
	return &unmarshalError{
		yidunError: New("UnmarshalError", msg, err),
		bytes:      bytes,
	}
}
