// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package nimblestudio

import (
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
)

// WaitUntilLaunchProfileDeleted uses the AmazonNimbleStudio API operation
// GetLaunchProfile to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *NimbleStudio) WaitUntilLaunchProfileDeleted(input *GetLaunchProfileInput) error {
	return c.WaitUntilLaunchProfileDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilLaunchProfileDeletedWithContext is an extended version of WaitUntilLaunchProfileDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NimbleStudio) WaitUntilLaunchProfileDeletedWithContext(ctx aws.Context, input *GetLaunchProfileInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilLaunchProfileDeleted",
		MaxAttempts: 150,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "launchProfile.state",
				Expected: "DELETED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "launchProfile.state",
				Expected: "DELETE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetLaunchProfileInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetLaunchProfileRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilLaunchProfileReady uses the AmazonNimbleStudio API operation
// GetLaunchProfile to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *NimbleStudio) WaitUntilLaunchProfileReady(input *GetLaunchProfileInput) error {
	return c.WaitUntilLaunchProfileReadyWithContext(aws.BackgroundContext(), input)
}

// WaitUntilLaunchProfileReadyWithContext is an extended version of WaitUntilLaunchProfileReady.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NimbleStudio) WaitUntilLaunchProfileReadyWithContext(ctx aws.Context, input *GetLaunchProfileInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilLaunchProfileReady",
		MaxAttempts: 150,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "launchProfile.state",
				Expected: "READY",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "launchProfile.state",
				Expected: "CREATE_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "launchProfile.state",
				Expected: "UPDATE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetLaunchProfileInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetLaunchProfileRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStreamingImageDeleted uses the AmazonNimbleStudio API operation
// GetStreamingImage to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *NimbleStudio) WaitUntilStreamingImageDeleted(input *GetStreamingImageInput) error {
	return c.WaitUntilStreamingImageDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStreamingImageDeletedWithContext is an extended version of WaitUntilStreamingImageDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NimbleStudio) WaitUntilStreamingImageDeletedWithContext(ctx aws.Context, input *GetStreamingImageInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStreamingImageDeleted",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(2 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "streamingImage.state",
				Expected: "DELETED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "streamingImage.state",
				Expected: "DELETE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetStreamingImageInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetStreamingImageRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStreamingImageReady uses the AmazonNimbleStudio API operation
// GetStreamingImage to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *NimbleStudio) WaitUntilStreamingImageReady(input *GetStreamingImageInput) error {
	return c.WaitUntilStreamingImageReadyWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStreamingImageReadyWithContext is an extended version of WaitUntilStreamingImageReady.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NimbleStudio) WaitUntilStreamingImageReadyWithContext(ctx aws.Context, input *GetStreamingImageInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStreamingImageReady",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(2 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "streamingImage.state",
				Expected: "READY",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "streamingImage.state",
				Expected: "CREATE_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "streamingImage.state",
				Expected: "UPDATE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetStreamingImageInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetStreamingImageRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStreamingSessionDeleted uses the AmazonNimbleStudio API operation
// GetStreamingSession to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *NimbleStudio) WaitUntilStreamingSessionDeleted(input *GetStreamingSessionInput) error {
	return c.WaitUntilStreamingSessionDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStreamingSessionDeletedWithContext is an extended version of WaitUntilStreamingSessionDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NimbleStudio) WaitUntilStreamingSessionDeletedWithContext(ctx aws.Context, input *GetStreamingSessionInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStreamingSessionDeleted",
		MaxAttempts: 180,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "session.state",
				Expected: "DELETED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "session.state",
				Expected: "DELETE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetStreamingSessionInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetStreamingSessionRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStreamingSessionReady uses the AmazonNimbleStudio API operation
// GetStreamingSession to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *NimbleStudio) WaitUntilStreamingSessionReady(input *GetStreamingSessionInput) error {
	return c.WaitUntilStreamingSessionReadyWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStreamingSessionReadyWithContext is an extended version of WaitUntilStreamingSessionReady.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NimbleStudio) WaitUntilStreamingSessionReadyWithContext(ctx aws.Context, input *GetStreamingSessionInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStreamingSessionReady",
		MaxAttempts: 180,
		Delay:       request.ConstantWaiterDelay(10 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "session.state",
				Expected: "READY",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "session.state",
				Expected: "CREATE_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "session.state",
				Expected: "START_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetStreamingSessionInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetStreamingSessionRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStreamingSessionStopped uses the AmazonNimbleStudio API operation
// GetStreamingSession to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *NimbleStudio) WaitUntilStreamingSessionStopped(input *GetStreamingSessionInput) error {
	return c.WaitUntilStreamingSessionStoppedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStreamingSessionStoppedWithContext is an extended version of WaitUntilStreamingSessionStopped.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NimbleStudio) WaitUntilStreamingSessionStoppedWithContext(ctx aws.Context, input *GetStreamingSessionInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStreamingSessionStopped",
		MaxAttempts: 180,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "session.state",
				Expected: "STOPPED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "session.state",
				Expected: "STOP_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetStreamingSessionInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetStreamingSessionRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStreamingSessionStreamReady uses the AmazonNimbleStudio API operation
// GetStreamingSessionStream to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *NimbleStudio) WaitUntilStreamingSessionStreamReady(input *GetStreamingSessionStreamInput) error {
	return c.WaitUntilStreamingSessionStreamReadyWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStreamingSessionStreamReadyWithContext is an extended version of WaitUntilStreamingSessionStreamReady.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NimbleStudio) WaitUntilStreamingSessionStreamReadyWithContext(ctx aws.Context, input *GetStreamingSessionStreamInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStreamingSessionStreamReady",
		MaxAttempts: 30,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "stream.state",
				Expected: "READY",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "stream.state",
				Expected: "CREATE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetStreamingSessionStreamInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetStreamingSessionStreamRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStudioComponentDeleted uses the AmazonNimbleStudio API operation
// GetStudioComponent to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *NimbleStudio) WaitUntilStudioComponentDeleted(input *GetStudioComponentInput) error {
	return c.WaitUntilStudioComponentDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStudioComponentDeletedWithContext is an extended version of WaitUntilStudioComponentDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NimbleStudio) WaitUntilStudioComponentDeletedWithContext(ctx aws.Context, input *GetStudioComponentInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStudioComponentDeleted",
		MaxAttempts: 120,
		Delay:       request.ConstantWaiterDelay(1 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "studioComponent.state",
				Expected: "DELETED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "studioComponent.state",
				Expected: "DELETE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetStudioComponentInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetStudioComponentRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStudioComponentReady uses the AmazonNimbleStudio API operation
// GetStudioComponent to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *NimbleStudio) WaitUntilStudioComponentReady(input *GetStudioComponentInput) error {
	return c.WaitUntilStudioComponentReadyWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStudioComponentReadyWithContext is an extended version of WaitUntilStudioComponentReady.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NimbleStudio) WaitUntilStudioComponentReadyWithContext(ctx aws.Context, input *GetStudioComponentInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStudioComponentReady",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(2 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "studioComponent.state",
				Expected: "READY",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "studioComponent.state",
				Expected: "CREATE_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "studioComponent.state",
				Expected: "UPDATE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetStudioComponentInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetStudioComponentRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStudioDeleted uses the AmazonNimbleStudio API operation
// GetStudio to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *NimbleStudio) WaitUntilStudioDeleted(input *GetStudioInput) error {
	return c.WaitUntilStudioDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStudioDeletedWithContext is an extended version of WaitUntilStudioDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NimbleStudio) WaitUntilStudioDeletedWithContext(ctx aws.Context, input *GetStudioInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStudioDeleted",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(2 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "studio.state",
				Expected: "DELETED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "studio.state",
				Expected: "DELETE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetStudioInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetStudioRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStudioReady uses the AmazonNimbleStudio API operation
// GetStudio to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *NimbleStudio) WaitUntilStudioReady(input *GetStudioInput) error {
	return c.WaitUntilStudioReadyWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStudioReadyWithContext is an extended version of WaitUntilStudioReady.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NimbleStudio) WaitUntilStudioReadyWithContext(ctx aws.Context, input *GetStudioInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStudioReady",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(2 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "studio.state",
				Expected: "READY",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "studio.state",
				Expected: "CREATE_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "studio.state",
				Expected: "UPDATE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetStudioInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetStudioRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
