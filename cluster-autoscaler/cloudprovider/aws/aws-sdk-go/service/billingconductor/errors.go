// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package billingconductor

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// You can cause an inconsistent state by updating or deleting a resource.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An unexpected error occurred while processing a request.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The request references a resource that doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceLimitExceededException for service response error code
	// "ServiceLimitExceededException".
	//
	// The request would cause a service limit to exceed.
	ErrCodeServiceLimitExceededException = "ServiceLimitExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The input doesn't match with the constraints specified by Amazon Web Services
	// services.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":         newErrorAccessDeniedException,
	"ConflictException":             newErrorConflictException,
	"InternalServerException":       newErrorInternalServerException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceLimitExceededException": newErrorServiceLimitExceededException,
	"ThrottlingException":           newErrorThrottlingException,
	"ValidationException":           newErrorValidationException,
}
