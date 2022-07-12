// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You are not authorized to perform the action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeHumanLoopQuotaExceededException for service response error code
	// "HumanLoopQuotaExceededException".
	//
	// The number of in-progress human reviews you have has exceeded the number
	// allowed.
	ErrCodeHumanLoopQuotaExceededException = "HumanLoopQuotaExceededException"

	// ErrCodeIdempotentParameterMismatchException for service response error code
	// "IdempotentParameterMismatchException".
	//
	// A ClientRequestToken input parameter was reused with an operation, but at
	// least one of the other input parameters is different from the previous call
	// to the operation.
	ErrCodeIdempotentParameterMismatchException = "IdempotentParameterMismatchException"

	// ErrCodeImageTooLargeException for service response error code
	// "ImageTooLargeException".
	//
	// The input image size exceeds the allowed limit. If you are calling DetectProtectiveEquipment,
	// the image size or resolution exceeds the allowed limit. For more information,
	// see Guidelines and quotas in Amazon Rekognition in the Amazon Rekognition
	// Developer Guide.
	ErrCodeImageTooLargeException = "ImageTooLargeException"

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// Amazon Rekognition experienced a service issue. Try your call again.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeInvalidImageFormatException for service response error code
	// "InvalidImageFormatException".
	//
	// The provided image format is not supported.
	ErrCodeInvalidImageFormatException = "InvalidImageFormatException"

	// ErrCodeInvalidPaginationTokenException for service response error code
	// "InvalidPaginationTokenException".
	//
	// Pagination token in the request is not valid.
	ErrCodeInvalidPaginationTokenException = "InvalidPaginationTokenException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// Input parameter violated a constraint. Validate your parameter before calling
	// the API operation again.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidS3ObjectException for service response error code
	// "InvalidS3ObjectException".
	//
	// Amazon Rekognition is unable to access the S3 object specified in the request.
	ErrCodeInvalidS3ObjectException = "InvalidS3ObjectException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// An Amazon Rekognition service limit was exceeded. For example, if you start
	// too many Amazon Rekognition Video jobs concurrently, calls to start operations
	// (StartLabelDetection, for example) will raise a LimitExceededException exception
	// (HTTP status code: 400) until the number of concurrently running jobs is
	// below the Amazon Rekognition service limit.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeProvisionedThroughputExceededException for service response error code
	// "ProvisionedThroughputExceededException".
	//
	// The number of requests exceeded your throughput limit. If you want to increase
	// this limit, contact Amazon Rekognition.
	ErrCodeProvisionedThroughputExceededException = "ProvisionedThroughputExceededException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// A resource with the specified ID already exists.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The specified resource is already being used.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource specified in the request cannot be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceNotReadyException for service response error code
	// "ResourceNotReadyException".
	//
	// The requested resource isn't ready. For example, this exception occurs when
	// you call DetectCustomLabels with a model version that isn't deployed.
	ErrCodeResourceNotReadyException = "ResourceNotReadyException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// The size of the collection exceeds the allowed limit. For more information,
	// see Guidelines and quotas in Amazon Rekognition in the Amazon Rekognition
	// Developer Guide.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// Amazon Rekognition is temporarily unable to process the request. Try your
	// call again.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeVideoTooLargeException for service response error code
	// "VideoTooLargeException".
	//
	// The file size or duration of the supplied media is too large. The maximum
	// file size is 10GB. The maximum duration is 6 hours.
	ErrCodeVideoTooLargeException = "VideoTooLargeException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                  newErrorAccessDeniedException,
	"HumanLoopQuotaExceededException":        newErrorHumanLoopQuotaExceededException,
	"IdempotentParameterMismatchException":   newErrorIdempotentParameterMismatchException,
	"ImageTooLargeException":                 newErrorImageTooLargeException,
	"InternalServerError":                    newErrorInternalServerError,
	"InvalidImageFormatException":            newErrorInvalidImageFormatException,
	"InvalidPaginationTokenException":        newErrorInvalidPaginationTokenException,
	"InvalidParameterException":              newErrorInvalidParameterException,
	"InvalidS3ObjectException":               newErrorInvalidS3ObjectException,
	"LimitExceededException":                 newErrorLimitExceededException,
	"ProvisionedThroughputExceededException": newErrorProvisionedThroughputExceededException,
	"ResourceAlreadyExistsException":         newErrorResourceAlreadyExistsException,
	"ResourceInUseException":                 newErrorResourceInUseException,
	"ResourceNotFoundException":              newErrorResourceNotFoundException,
	"ResourceNotReadyException":              newErrorResourceNotReadyException,
	"ServiceQuotaExceededException":          newErrorServiceQuotaExceededException,
	"ThrottlingException":                    newErrorThrottlingException,
	"VideoTooLargeException":                 newErrorVideoTooLargeException,
}
