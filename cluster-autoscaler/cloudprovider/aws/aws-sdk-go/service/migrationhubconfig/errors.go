// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhubconfig

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeDryRunOperation for service response error code
	// "DryRunOperation".
	//
	// Exception raised to indicate that authorization of an action was successful,
	// when the DryRun flag is set to true.
	ErrCodeDryRunOperation = "DryRunOperation"

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// Exception raised when an internal, configuration, or dependency error is
	// encountered.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeInvalidInputException for service response error code
	// "InvalidInputException".
	//
	// Exception raised when the provided input violates a policy constraint or
	// is entered in the wrong format or data type.
	ErrCodeInvalidInputException = "InvalidInputException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// Exception raised when a request fails due to temporary unavailability of
	// the service.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":       newErrorAccessDeniedException,
	"DryRunOperation":             newErrorDryRunOperation,
	"InternalServerError":         newErrorInternalServerError,
	"InvalidInputException":       newErrorInvalidInputException,
	"ServiceUnavailableException": newErrorServiceUnavailableException,
	"ThrottlingException":         newErrorThrottlingException,
}
