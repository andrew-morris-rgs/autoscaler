// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package recyclebin provides the client and types for making API
// requests to Amazon Recycle Bin.
//
// This is the Recycle Bin API Reference. This documentation provides descriptions
// and syntax for each of the actions and data types in Recycle Bin.
//
// Recycle Bin is a resource recovery feature that enables you to restore accidentally
// deleted snapshots and EBS-backed AMIs. When using Recycle Bin, if your resources
// are deleted, they are retained in the Recycle Bin for a time period that
// you specify.
//
// You can restore a resource from the Recycle Bin at any time before its retention
// period expires. After you restore a resource from the Recycle Bin, the resource
// is removed from the Recycle Bin, and you can then use it in the same way
// you use any other resource of that type in your account. If the retention
// period expires and the resource is not restored, the resource is permanently
// deleted from the Recycle Bin and is no longer available for recovery. For
// more information about Recycle Bin, see Recycle Bin (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-recycle-bin.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// See https://docs.aws.amazon.com/goto/WebAPI/rbin-2021-06-15 for more information on this service.
//
// See recyclebin package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/recyclebin/
//
// Using the Client
//
// To contact Amazon Recycle Bin with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Recycle Bin client RecycleBin for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/recyclebin/#New
package recyclebin
