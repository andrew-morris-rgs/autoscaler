// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
)

// WaitUntilFlowActive uses the AWS MediaConnect API operation
// DescribeFlow to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaConnect) WaitUntilFlowActive(input *DescribeFlowInput) error {
	return c.WaitUntilFlowActiveWithContext(aws.BackgroundContext(), input)
}

// WaitUntilFlowActiveWithContext is an extended version of WaitUntilFlowActive.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaConnect) WaitUntilFlowActiveWithContext(ctx aws.Context, input *DescribeFlowInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilFlowActive",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Flow.Status",
				Expected: "ACTIVE",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Flow.Status",
				Expected: "STARTING",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Flow.Status",
				Expected: "UPDATING",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 503,
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Flow.Status",
				Expected: "ERROR",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeFlowInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeFlowRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilFlowDeleted uses the AWS MediaConnect API operation
// DescribeFlow to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaConnect) WaitUntilFlowDeleted(input *DescribeFlowInput) error {
	return c.WaitUntilFlowDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilFlowDeletedWithContext is an extended version of WaitUntilFlowDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaConnect) WaitUntilFlowDeletedWithContext(ctx aws.Context, input *DescribeFlowInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilFlowDeleted",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 404,
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Flow.Status",
				Expected: "DELETING",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 503,
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Flow.Status",
				Expected: "ERROR",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeFlowInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeFlowRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilFlowStandby uses the AWS MediaConnect API operation
// DescribeFlow to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaConnect) WaitUntilFlowStandby(input *DescribeFlowInput) error {
	return c.WaitUntilFlowStandbyWithContext(aws.BackgroundContext(), input)
}

// WaitUntilFlowStandbyWithContext is an extended version of WaitUntilFlowStandby.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaConnect) WaitUntilFlowStandbyWithContext(ctx aws.Context, input *DescribeFlowInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilFlowStandby",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Flow.Status",
				Expected: "STANDBY",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Flow.Status",
				Expected: "STOPPING",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 503,
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Flow.Status",
				Expected: "ERROR",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeFlowInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeFlowRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
