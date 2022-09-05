// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package outpostsiface provides an interface to enable mocking the AWS Outposts service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package outpostsiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/outposts"
)

// OutpostsAPI provides an interface to enable mocking the
// outposts.Outposts service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Outposts.
//    func myFunc(svc outpostsiface.OutpostsAPI) bool {
//        // Make svc.CancelOrder request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := outposts.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockOutpostsClient struct {
//        outpostsiface.OutpostsAPI
//    }
//    func (m *mockOutpostsClient) CancelOrder(input *outposts.CancelOrderInput) (*outposts.CancelOrderOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockOutpostsClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type OutpostsAPI interface {
	CancelOrder(*outposts.CancelOrderInput) (*outposts.CancelOrderOutput, error)
	CancelOrderWithContext(aws.Context, *outposts.CancelOrderInput, ...request.Option) (*outposts.CancelOrderOutput, error)
	CancelOrderRequest(*outposts.CancelOrderInput) (*request.Request, *outposts.CancelOrderOutput)

	CreateOrder(*outposts.CreateOrderInput) (*outposts.CreateOrderOutput, error)
	CreateOrderWithContext(aws.Context, *outposts.CreateOrderInput, ...request.Option) (*outposts.CreateOrderOutput, error)
	CreateOrderRequest(*outposts.CreateOrderInput) (*request.Request, *outposts.CreateOrderOutput)

	CreateOutpost(*outposts.CreateOutpostInput) (*outposts.CreateOutpostOutput, error)
	CreateOutpostWithContext(aws.Context, *outposts.CreateOutpostInput, ...request.Option) (*outposts.CreateOutpostOutput, error)
	CreateOutpostRequest(*outposts.CreateOutpostInput) (*request.Request, *outposts.CreateOutpostOutput)

	CreateSite(*outposts.CreateSiteInput) (*outposts.CreateSiteOutput, error)
	CreateSiteWithContext(aws.Context, *outposts.CreateSiteInput, ...request.Option) (*outposts.CreateSiteOutput, error)
	CreateSiteRequest(*outposts.CreateSiteInput) (*request.Request, *outposts.CreateSiteOutput)

	DeleteOutpost(*outposts.DeleteOutpostInput) (*outposts.DeleteOutpostOutput, error)
	DeleteOutpostWithContext(aws.Context, *outposts.DeleteOutpostInput, ...request.Option) (*outposts.DeleteOutpostOutput, error)
	DeleteOutpostRequest(*outposts.DeleteOutpostInput) (*request.Request, *outposts.DeleteOutpostOutput)

	DeleteSite(*outposts.DeleteSiteInput) (*outposts.DeleteSiteOutput, error)
	DeleteSiteWithContext(aws.Context, *outposts.DeleteSiteInput, ...request.Option) (*outposts.DeleteSiteOutput, error)
	DeleteSiteRequest(*outposts.DeleteSiteInput) (*request.Request, *outposts.DeleteSiteOutput)

	GetCatalogItem(*outposts.GetCatalogItemInput) (*outposts.GetCatalogItemOutput, error)
	GetCatalogItemWithContext(aws.Context, *outposts.GetCatalogItemInput, ...request.Option) (*outposts.GetCatalogItemOutput, error)
	GetCatalogItemRequest(*outposts.GetCatalogItemInput) (*request.Request, *outposts.GetCatalogItemOutput)

	GetOrder(*outposts.GetOrderInput) (*outposts.GetOrderOutput, error)
	GetOrderWithContext(aws.Context, *outposts.GetOrderInput, ...request.Option) (*outposts.GetOrderOutput, error)
	GetOrderRequest(*outposts.GetOrderInput) (*request.Request, *outposts.GetOrderOutput)

	GetOutpost(*outposts.GetOutpostInput) (*outposts.GetOutpostOutput, error)
	GetOutpostWithContext(aws.Context, *outposts.GetOutpostInput, ...request.Option) (*outposts.GetOutpostOutput, error)
	GetOutpostRequest(*outposts.GetOutpostInput) (*request.Request, *outposts.GetOutpostOutput)

	GetOutpostInstanceTypes(*outposts.GetOutpostInstanceTypesInput) (*outposts.GetOutpostInstanceTypesOutput, error)
	GetOutpostInstanceTypesWithContext(aws.Context, *outposts.GetOutpostInstanceTypesInput, ...request.Option) (*outposts.GetOutpostInstanceTypesOutput, error)
	GetOutpostInstanceTypesRequest(*outposts.GetOutpostInstanceTypesInput) (*request.Request, *outposts.GetOutpostInstanceTypesOutput)

	GetOutpostInstanceTypesPages(*outposts.GetOutpostInstanceTypesInput, func(*outposts.GetOutpostInstanceTypesOutput, bool) bool) error
	GetOutpostInstanceTypesPagesWithContext(aws.Context, *outposts.GetOutpostInstanceTypesInput, func(*outposts.GetOutpostInstanceTypesOutput, bool) bool, ...request.Option) error

	GetSite(*outposts.GetSiteInput) (*outposts.GetSiteOutput, error)
	GetSiteWithContext(aws.Context, *outposts.GetSiteInput, ...request.Option) (*outposts.GetSiteOutput, error)
	GetSiteRequest(*outposts.GetSiteInput) (*request.Request, *outposts.GetSiteOutput)

	GetSiteAddress(*outposts.GetSiteAddressInput) (*outposts.GetSiteAddressOutput, error)
	GetSiteAddressWithContext(aws.Context, *outposts.GetSiteAddressInput, ...request.Option) (*outposts.GetSiteAddressOutput, error)
	GetSiteAddressRequest(*outposts.GetSiteAddressInput) (*request.Request, *outposts.GetSiteAddressOutput)

	ListAssets(*outposts.ListAssetsInput) (*outposts.ListAssetsOutput, error)
	ListAssetsWithContext(aws.Context, *outposts.ListAssetsInput, ...request.Option) (*outposts.ListAssetsOutput, error)
	ListAssetsRequest(*outposts.ListAssetsInput) (*request.Request, *outposts.ListAssetsOutput)

	ListAssetsPages(*outposts.ListAssetsInput, func(*outposts.ListAssetsOutput, bool) bool) error
	ListAssetsPagesWithContext(aws.Context, *outposts.ListAssetsInput, func(*outposts.ListAssetsOutput, bool) bool, ...request.Option) error

	ListCatalogItems(*outposts.ListCatalogItemsInput) (*outposts.ListCatalogItemsOutput, error)
	ListCatalogItemsWithContext(aws.Context, *outposts.ListCatalogItemsInput, ...request.Option) (*outposts.ListCatalogItemsOutput, error)
	ListCatalogItemsRequest(*outposts.ListCatalogItemsInput) (*request.Request, *outposts.ListCatalogItemsOutput)

	ListCatalogItemsPages(*outposts.ListCatalogItemsInput, func(*outposts.ListCatalogItemsOutput, bool) bool) error
	ListCatalogItemsPagesWithContext(aws.Context, *outposts.ListCatalogItemsInput, func(*outposts.ListCatalogItemsOutput, bool) bool, ...request.Option) error

	ListOrders(*outposts.ListOrdersInput) (*outposts.ListOrdersOutput, error)
	ListOrdersWithContext(aws.Context, *outposts.ListOrdersInput, ...request.Option) (*outposts.ListOrdersOutput, error)
	ListOrdersRequest(*outposts.ListOrdersInput) (*request.Request, *outposts.ListOrdersOutput)

	ListOrdersPages(*outposts.ListOrdersInput, func(*outposts.ListOrdersOutput, bool) bool) error
	ListOrdersPagesWithContext(aws.Context, *outposts.ListOrdersInput, func(*outposts.ListOrdersOutput, bool) bool, ...request.Option) error

	ListOutposts(*outposts.ListOutpostsInput) (*outposts.ListOutpostsOutput, error)
	ListOutpostsWithContext(aws.Context, *outposts.ListOutpostsInput, ...request.Option) (*outposts.ListOutpostsOutput, error)
	ListOutpostsRequest(*outposts.ListOutpostsInput) (*request.Request, *outposts.ListOutpostsOutput)

	ListOutpostsPages(*outposts.ListOutpostsInput, func(*outposts.ListOutpostsOutput, bool) bool) error
	ListOutpostsPagesWithContext(aws.Context, *outposts.ListOutpostsInput, func(*outposts.ListOutpostsOutput, bool) bool, ...request.Option) error

	ListSites(*outposts.ListSitesInput) (*outposts.ListSitesOutput, error)
	ListSitesWithContext(aws.Context, *outposts.ListSitesInput, ...request.Option) (*outposts.ListSitesOutput, error)
	ListSitesRequest(*outposts.ListSitesInput) (*request.Request, *outposts.ListSitesOutput)

	ListSitesPages(*outposts.ListSitesInput, func(*outposts.ListSitesOutput, bool) bool) error
	ListSitesPagesWithContext(aws.Context, *outposts.ListSitesInput, func(*outposts.ListSitesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*outposts.ListTagsForResourceInput) (*outposts.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *outposts.ListTagsForResourceInput, ...request.Option) (*outposts.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*outposts.ListTagsForResourceInput) (*request.Request, *outposts.ListTagsForResourceOutput)

	TagResource(*outposts.TagResourceInput) (*outposts.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *outposts.TagResourceInput, ...request.Option) (*outposts.TagResourceOutput, error)
	TagResourceRequest(*outposts.TagResourceInput) (*request.Request, *outposts.TagResourceOutput)

	UntagResource(*outposts.UntagResourceInput) (*outposts.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *outposts.UntagResourceInput, ...request.Option) (*outposts.UntagResourceOutput, error)
	UntagResourceRequest(*outposts.UntagResourceInput) (*request.Request, *outposts.UntagResourceOutput)

	UpdateOutpost(*outposts.UpdateOutpostInput) (*outposts.UpdateOutpostOutput, error)
	UpdateOutpostWithContext(aws.Context, *outposts.UpdateOutpostInput, ...request.Option) (*outposts.UpdateOutpostOutput, error)
	UpdateOutpostRequest(*outposts.UpdateOutpostInput) (*request.Request, *outposts.UpdateOutpostOutput)

	UpdateSite(*outposts.UpdateSiteInput) (*outposts.UpdateSiteOutput, error)
	UpdateSiteWithContext(aws.Context, *outposts.UpdateSiteInput, ...request.Option) (*outposts.UpdateSiteOutput, error)
	UpdateSiteRequest(*outposts.UpdateSiteInput) (*request.Request, *outposts.UpdateSiteOutput)

	UpdateSiteAddress(*outposts.UpdateSiteAddressInput) (*outposts.UpdateSiteAddressOutput, error)
	UpdateSiteAddressWithContext(aws.Context, *outposts.UpdateSiteAddressInput, ...request.Option) (*outposts.UpdateSiteAddressOutput, error)
	UpdateSiteAddressRequest(*outposts.UpdateSiteAddressInput) (*request.Request, *outposts.UpdateSiteAddressOutput)

	UpdateSiteRackPhysicalProperties(*outposts.UpdateSiteRackPhysicalPropertiesInput) (*outposts.UpdateSiteRackPhysicalPropertiesOutput, error)
	UpdateSiteRackPhysicalPropertiesWithContext(aws.Context, *outposts.UpdateSiteRackPhysicalPropertiesInput, ...request.Option) (*outposts.UpdateSiteRackPhysicalPropertiesOutput, error)
	UpdateSiteRackPhysicalPropertiesRequest(*outposts.UpdateSiteRackPhysicalPropertiesInput) (*request.Request, *outposts.UpdateSiteRackPhysicalPropertiesOutput)
}

var _ OutpostsAPI = (*outposts.Outposts)(nil)
