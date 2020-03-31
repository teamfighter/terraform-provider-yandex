// Code generated by sdkgen. DO NOT EDIT.

//nolint
package vpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	vpc "github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

//revive:disable

// SecurityGroupServiceClient is a vpc.SecurityGroupServiceClient with
// lazy GRPC connection initialization.
type SecurityGroupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements vpc.SecurityGroupServiceClient
func (c *SecurityGroupServiceClient) Create(ctx context.Context, in *vpc.CreateSecurityGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSecurityGroupServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements vpc.SecurityGroupServiceClient
func (c *SecurityGroupServiceClient) Delete(ctx context.Context, in *vpc.DeleteSecurityGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSecurityGroupServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements vpc.SecurityGroupServiceClient
func (c *SecurityGroupServiceClient) Get(ctx context.Context, in *vpc.GetSecurityGroupRequest, opts ...grpc.CallOption) (*vpc.SecurityGroup, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSecurityGroupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements vpc.SecurityGroupServiceClient
func (c *SecurityGroupServiceClient) List(ctx context.Context, in *vpc.ListSecurityGroupsRequest, opts ...grpc.CallOption) (*vpc.ListSecurityGroupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSecurityGroupServiceClient(conn).List(ctx, in, opts...)
}

type SecurityGroupIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *SecurityGroupServiceClient
	request *vpc.ListSecurityGroupsRequest

	items []*vpc.SecurityGroup
}

func (c *SecurityGroupServiceClient) SecurityGroupIterator(ctx context.Context, folderId string, opts ...grpc.CallOption) *SecurityGroupIterator {
	return &SecurityGroupIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &vpc.ListSecurityGroupsRequest{
			FolderId: folderId,
			PageSize: 1000,
		},
	}
}

func (it *SecurityGroupIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.SecurityGroups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *SecurityGroupIterator) Value() *vpc.SecurityGroup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SecurityGroupIterator) Error() error {
	return it.err
}

// ListOperations implements vpc.SecurityGroupServiceClient
func (c *SecurityGroupServiceClient) ListOperations(ctx context.Context, in *vpc.ListSecurityGroupOperationsRequest, opts ...grpc.CallOption) (*vpc.ListSecurityGroupOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSecurityGroupServiceClient(conn).ListOperations(ctx, in, opts...)
}

type SecurityGroupOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *SecurityGroupServiceClient
	request *vpc.ListSecurityGroupOperationsRequest

	items []*operation.Operation
}

func (c *SecurityGroupServiceClient) SecurityGroupOperationsIterator(ctx context.Context, securityGroupId string, opts ...grpc.CallOption) *SecurityGroupOperationsIterator {
	return &SecurityGroupOperationsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &vpc.ListSecurityGroupOperationsRequest{
			SecurityGroupId: securityGroupId,
			PageSize:        1000,
		},
	}
}

func (it *SecurityGroupOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *SecurityGroupOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SecurityGroupOperationsIterator) Error() error {
	return it.err
}

// Move implements vpc.SecurityGroupServiceClient
func (c *SecurityGroupServiceClient) Move(ctx context.Context, in *vpc.MoveSecurityGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSecurityGroupServiceClient(conn).Move(ctx, in, opts...)
}

// Update implements vpc.SecurityGroupServiceClient
func (c *SecurityGroupServiceClient) Update(ctx context.Context, in *vpc.UpdateSecurityGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSecurityGroupServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateRule implements vpc.SecurityGroupServiceClient
func (c *SecurityGroupServiceClient) UpdateRule(ctx context.Context, in *vpc.UpdateSecurityGroupRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSecurityGroupServiceClient(conn).UpdateRule(ctx, in, opts...)
}

// UpdateRules implements vpc.SecurityGroupServiceClient
func (c *SecurityGroupServiceClient) UpdateRules(ctx context.Context, in *vpc.UpdateSecurityGroupRulesRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSecurityGroupServiceClient(conn).UpdateRules(ctx, in, opts...)
}
