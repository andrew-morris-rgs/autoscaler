// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelsv2

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The action that you tried to perform couldn't be completed because the resource
	// is in a conflicting state. For example, deleting a bot that is in the CREATING
	// state. Try your request again.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// The service encountered an unexpected condition. Try your request again.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodePreconditionFailedException for service response error code
	// "PreconditionFailedException".
	//
	// Your request couldn't be completed because one or more request fields aren't
	// valid. Check the fields in your request and try again.
	ErrCodePreconditionFailedException = "PreconditionFailedException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// You asked to describe a resource that doesn't exist. Check the resource that
	// you are requesting and try again.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// You have reached a quota for your bot.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// Your request rate is too high. Reduce the frequency of requests.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// One of the input parameters in your request isn't valid. Check the parameters
	// and try your request again.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConflictException":             newErrorConflictException,
	"InternalServerException":       newErrorInternalServerException,
	"PreconditionFailedException":   newErrorPreconditionFailedException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"ThrottlingException":           newErrorThrottlingException,
	"ValidationException":           newErrorValidationException,
}
