// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3outposts

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/client"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/client/metadata"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/signer/v4"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol/restjson"
)

// S3Outposts provides the API operation methods for making requests to
// Amazon S3 on Outposts. See this package's package overview docs
// for details on the service.
//
// S3Outposts methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type S3Outposts struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "S3Outposts"  // Name of service.
	EndpointsID = "s3-outposts" // ID to lookup a service endpoint with.
	ServiceID   = "S3Outposts"  // ServiceID is a unique identifier of a specific service.
)

// New creates a new instance of the S3Outposts client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     mySession := session.Must(session.NewSession())
//
//     // Create a S3Outposts client from just a session.
//     svc := s3outposts.New(mySession)
//
//     // Create a S3Outposts client with additional configuration
//     svc := s3outposts.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *S3Outposts {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "s3-outposts"
	}
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName, c.ResolvedRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, partitionID, endpoint, signingRegion, signingName, resolvedRegion string) *S3Outposts {
	svc := &S3Outposts{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:    ServiceName,
				ServiceID:      ServiceID,
				SigningName:    signingName,
				SigningRegion:  signingRegion,
				PartitionID:    partitionID,
				Endpoint:       endpoint,
				APIVersion:     "2017-07-25",
				ResolvedRegion: resolvedRegion,
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(
		protocol.NewUnmarshalErrorHandler(restjson.NewUnmarshalTypedError(exceptionFromCode)).NamedHandler(),
	)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a S3Outposts operation and runs any
// custom request initialization.
func (c *S3Outposts) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
