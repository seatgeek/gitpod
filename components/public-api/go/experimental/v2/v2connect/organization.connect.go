// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: gitpod/experimental/v2/organization.proto

package v2connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v2 "github.com/gitpod-io/gitpod/components/public-api/go/experimental/v2"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// OrganizationServiceName is the fully-qualified name of the OrganizationService service.
	OrganizationServiceName = "gitpod.experimental.v2.OrganizationService"
)

// OrganizationServiceClient is a client for the gitpod.experimental.v2.OrganizationService service.
type OrganizationServiceClient interface {
	// CreateOrganization creates a new Organization.
	CreateOrganization(context.Context, *connect_go.Request[v2.CreateOrganizationRequest]) (*connect_go.Response[v2.CreateOrganizationResponse], error)
	// GetOrganization retrieves a single Organization.
	GetOrganization(context.Context, *connect_go.Request[v2.GetOrganizationRequest]) (*connect_go.Response[v2.GetOrganizationResponse], error)
	// UpdateOrganization updates the properties of an Organization.
	UpdateOrganization(context.Context, *connect_go.Request[v2.UpdateOrganizationRequest]) (*connect_go.Response[v2.UpdateOrganizationResponse], error)
	// ListOrganizations lists all organization the caller has access to.
	ListOrganizations(context.Context, *connect_go.Request[v2.ListOrganizationsRequest]) (*connect_go.Response[v2.ListOrganizationsResponse], error)
	// DeleteOrganization deletes the specified organization.
	DeleteOrganization(context.Context, *connect_go.Request[v2.DeleteOrganizationRequest]) (*connect_go.Response[v2.DeleteOrganizationResponse], error)
	// GetOrganizationInvitation retrieves the invitation for a Organization.
	GetOrganizationInvitation(context.Context, *connect_go.Request[v2.GetOrganizationInvitationRequest]) (*connect_go.Response[v2.GetOrganizationInvitationResponse], error)
	// JoinOrganization makes the caller a OrganizationMember of the Organization.
	JoinOrganization(context.Context, *connect_go.Request[v2.JoinOrganizationRequest]) (*connect_go.Response[v2.JoinOrganizationResponse], error)
	// ResetOrganizationInvitation resets the invitation_id for a Organization.
	ResetOrganizationInvitation(context.Context, *connect_go.Request[v2.ResetOrganizationInvitationRequest]) (*connect_go.Response[v2.ResetOrganizationInvitationResponse], error)
	// ListOrganizationMembers lists the members of a Organization.
	ListOrganizationMembers(context.Context, *connect_go.Request[v2.ListOrganizationMembersRequest]) (*connect_go.Response[v2.ListOrganizationMembersResponse], error)
	// UpdateOrganizationMember updates organization membership properties.
	UpdateOrganizationMember(context.Context, *connect_go.Request[v2.UpdateOrganizationMemberRequest]) (*connect_go.Response[v2.UpdateOrganizationMemberResponse], error)
	// DeleteOrganizationMember removes a OrganizationMember from the Organization.
	DeleteOrganizationMember(context.Context, *connect_go.Request[v2.DeleteOrganizationMemberRequest]) (*connect_go.Response[v2.DeleteOrganizationMemberResponse], error)
	// GetOrganizationSettings retrieves the settings of a Organization.
	GetOrganizationSettings(context.Context, *connect_go.Request[v2.GetOrganizationSettingsRequest]) (*connect_go.Response[v2.GetOrganizationSettingsResponse], error)
	// UpdateOrganizationSettings updates the settings of a Organization.
	UpdateOrganizationSettings(context.Context, *connect_go.Request[v2.UpdateOrganizationSettingsRequest]) (*connect_go.Response[v2.UpdateOrganizationSettingsResponse], error)
}

// NewOrganizationServiceClient constructs a client for the
// gitpod.experimental.v2.OrganizationService service. By default, it uses the Connect protocol with
// the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To use
// the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOrganizationServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) OrganizationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &organizationServiceClient{
		createOrganization: connect_go.NewClient[v2.CreateOrganizationRequest, v2.CreateOrganizationResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/CreateOrganization",
			opts...,
		),
		getOrganization: connect_go.NewClient[v2.GetOrganizationRequest, v2.GetOrganizationResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/GetOrganization",
			opts...,
		),
		updateOrganization: connect_go.NewClient[v2.UpdateOrganizationRequest, v2.UpdateOrganizationResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/UpdateOrganization",
			opts...,
		),
		listOrganizations: connect_go.NewClient[v2.ListOrganizationsRequest, v2.ListOrganizationsResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/ListOrganizations",
			opts...,
		),
		deleteOrganization: connect_go.NewClient[v2.DeleteOrganizationRequest, v2.DeleteOrganizationResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/DeleteOrganization",
			opts...,
		),
		getOrganizationInvitation: connect_go.NewClient[v2.GetOrganizationInvitationRequest, v2.GetOrganizationInvitationResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/GetOrganizationInvitation",
			opts...,
		),
		joinOrganization: connect_go.NewClient[v2.JoinOrganizationRequest, v2.JoinOrganizationResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/JoinOrganization",
			opts...,
		),
		resetOrganizationInvitation: connect_go.NewClient[v2.ResetOrganizationInvitationRequest, v2.ResetOrganizationInvitationResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/ResetOrganizationInvitation",
			opts...,
		),
		listOrganizationMembers: connect_go.NewClient[v2.ListOrganizationMembersRequest, v2.ListOrganizationMembersResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/ListOrganizationMembers",
			opts...,
		),
		updateOrganizationMember: connect_go.NewClient[v2.UpdateOrganizationMemberRequest, v2.UpdateOrganizationMemberResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/UpdateOrganizationMember",
			opts...,
		),
		deleteOrganizationMember: connect_go.NewClient[v2.DeleteOrganizationMemberRequest, v2.DeleteOrganizationMemberResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/DeleteOrganizationMember",
			opts...,
		),
		getOrganizationSettings: connect_go.NewClient[v2.GetOrganizationSettingsRequest, v2.GetOrganizationSettingsResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/GetOrganizationSettings",
			opts...,
		),
		updateOrganizationSettings: connect_go.NewClient[v2.UpdateOrganizationSettingsRequest, v2.UpdateOrganizationSettingsResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v2.OrganizationService/UpdateOrganizationSettings",
			opts...,
		),
	}
}

// organizationServiceClient implements OrganizationServiceClient.
type organizationServiceClient struct {
	createOrganization          *connect_go.Client[v2.CreateOrganizationRequest, v2.CreateOrganizationResponse]
	getOrganization             *connect_go.Client[v2.GetOrganizationRequest, v2.GetOrganizationResponse]
	updateOrganization          *connect_go.Client[v2.UpdateOrganizationRequest, v2.UpdateOrganizationResponse]
	listOrganizations           *connect_go.Client[v2.ListOrganizationsRequest, v2.ListOrganizationsResponse]
	deleteOrganization          *connect_go.Client[v2.DeleteOrganizationRequest, v2.DeleteOrganizationResponse]
	getOrganizationInvitation   *connect_go.Client[v2.GetOrganizationInvitationRequest, v2.GetOrganizationInvitationResponse]
	joinOrganization            *connect_go.Client[v2.JoinOrganizationRequest, v2.JoinOrganizationResponse]
	resetOrganizationInvitation *connect_go.Client[v2.ResetOrganizationInvitationRequest, v2.ResetOrganizationInvitationResponse]
	listOrganizationMembers     *connect_go.Client[v2.ListOrganizationMembersRequest, v2.ListOrganizationMembersResponse]
	updateOrganizationMember    *connect_go.Client[v2.UpdateOrganizationMemberRequest, v2.UpdateOrganizationMemberResponse]
	deleteOrganizationMember    *connect_go.Client[v2.DeleteOrganizationMemberRequest, v2.DeleteOrganizationMemberResponse]
	getOrganizationSettings     *connect_go.Client[v2.GetOrganizationSettingsRequest, v2.GetOrganizationSettingsResponse]
	updateOrganizationSettings  *connect_go.Client[v2.UpdateOrganizationSettingsRequest, v2.UpdateOrganizationSettingsResponse]
}

// CreateOrganization calls gitpod.experimental.v2.OrganizationService.CreateOrganization.
func (c *organizationServiceClient) CreateOrganization(ctx context.Context, req *connect_go.Request[v2.CreateOrganizationRequest]) (*connect_go.Response[v2.CreateOrganizationResponse], error) {
	return c.createOrganization.CallUnary(ctx, req)
}

// GetOrganization calls gitpod.experimental.v2.OrganizationService.GetOrganization.
func (c *organizationServiceClient) GetOrganization(ctx context.Context, req *connect_go.Request[v2.GetOrganizationRequest]) (*connect_go.Response[v2.GetOrganizationResponse], error) {
	return c.getOrganization.CallUnary(ctx, req)
}

// UpdateOrganization calls gitpod.experimental.v2.OrganizationService.UpdateOrganization.
func (c *organizationServiceClient) UpdateOrganization(ctx context.Context, req *connect_go.Request[v2.UpdateOrganizationRequest]) (*connect_go.Response[v2.UpdateOrganizationResponse], error) {
	return c.updateOrganization.CallUnary(ctx, req)
}

// ListOrganizations calls gitpod.experimental.v2.OrganizationService.ListOrganizations.
func (c *organizationServiceClient) ListOrganizations(ctx context.Context, req *connect_go.Request[v2.ListOrganizationsRequest]) (*connect_go.Response[v2.ListOrganizationsResponse], error) {
	return c.listOrganizations.CallUnary(ctx, req)
}

// DeleteOrganization calls gitpod.experimental.v2.OrganizationService.DeleteOrganization.
func (c *organizationServiceClient) DeleteOrganization(ctx context.Context, req *connect_go.Request[v2.DeleteOrganizationRequest]) (*connect_go.Response[v2.DeleteOrganizationResponse], error) {
	return c.deleteOrganization.CallUnary(ctx, req)
}

// GetOrganizationInvitation calls
// gitpod.experimental.v2.OrganizationService.GetOrganizationInvitation.
func (c *organizationServiceClient) GetOrganizationInvitation(ctx context.Context, req *connect_go.Request[v2.GetOrganizationInvitationRequest]) (*connect_go.Response[v2.GetOrganizationInvitationResponse], error) {
	return c.getOrganizationInvitation.CallUnary(ctx, req)
}

// JoinOrganization calls gitpod.experimental.v2.OrganizationService.JoinOrganization.
func (c *organizationServiceClient) JoinOrganization(ctx context.Context, req *connect_go.Request[v2.JoinOrganizationRequest]) (*connect_go.Response[v2.JoinOrganizationResponse], error) {
	return c.joinOrganization.CallUnary(ctx, req)
}

// ResetOrganizationInvitation calls
// gitpod.experimental.v2.OrganizationService.ResetOrganizationInvitation.
func (c *organizationServiceClient) ResetOrganizationInvitation(ctx context.Context, req *connect_go.Request[v2.ResetOrganizationInvitationRequest]) (*connect_go.Response[v2.ResetOrganizationInvitationResponse], error) {
	return c.resetOrganizationInvitation.CallUnary(ctx, req)
}

// ListOrganizationMembers calls gitpod.experimental.v2.OrganizationService.ListOrganizationMembers.
func (c *organizationServiceClient) ListOrganizationMembers(ctx context.Context, req *connect_go.Request[v2.ListOrganizationMembersRequest]) (*connect_go.Response[v2.ListOrganizationMembersResponse], error) {
	return c.listOrganizationMembers.CallUnary(ctx, req)
}

// UpdateOrganizationMember calls
// gitpod.experimental.v2.OrganizationService.UpdateOrganizationMember.
func (c *organizationServiceClient) UpdateOrganizationMember(ctx context.Context, req *connect_go.Request[v2.UpdateOrganizationMemberRequest]) (*connect_go.Response[v2.UpdateOrganizationMemberResponse], error) {
	return c.updateOrganizationMember.CallUnary(ctx, req)
}

// DeleteOrganizationMember calls
// gitpod.experimental.v2.OrganizationService.DeleteOrganizationMember.
func (c *organizationServiceClient) DeleteOrganizationMember(ctx context.Context, req *connect_go.Request[v2.DeleteOrganizationMemberRequest]) (*connect_go.Response[v2.DeleteOrganizationMemberResponse], error) {
	return c.deleteOrganizationMember.CallUnary(ctx, req)
}

// GetOrganizationSettings calls gitpod.experimental.v2.OrganizationService.GetOrganizationSettings.
func (c *organizationServiceClient) GetOrganizationSettings(ctx context.Context, req *connect_go.Request[v2.GetOrganizationSettingsRequest]) (*connect_go.Response[v2.GetOrganizationSettingsResponse], error) {
	return c.getOrganizationSettings.CallUnary(ctx, req)
}

// UpdateOrganizationSettings calls
// gitpod.experimental.v2.OrganizationService.UpdateOrganizationSettings.
func (c *organizationServiceClient) UpdateOrganizationSettings(ctx context.Context, req *connect_go.Request[v2.UpdateOrganizationSettingsRequest]) (*connect_go.Response[v2.UpdateOrganizationSettingsResponse], error) {
	return c.updateOrganizationSettings.CallUnary(ctx, req)
}

// OrganizationServiceHandler is an implementation of the gitpod.experimental.v2.OrganizationService
// service.
type OrganizationServiceHandler interface {
	// CreateOrganization creates a new Organization.
	CreateOrganization(context.Context, *connect_go.Request[v2.CreateOrganizationRequest]) (*connect_go.Response[v2.CreateOrganizationResponse], error)
	// GetOrganization retrieves a single Organization.
	GetOrganization(context.Context, *connect_go.Request[v2.GetOrganizationRequest]) (*connect_go.Response[v2.GetOrganizationResponse], error)
	// UpdateOrganization updates the properties of an Organization.
	UpdateOrganization(context.Context, *connect_go.Request[v2.UpdateOrganizationRequest]) (*connect_go.Response[v2.UpdateOrganizationResponse], error)
	// ListOrganizations lists all organization the caller has access to.
	ListOrganizations(context.Context, *connect_go.Request[v2.ListOrganizationsRequest]) (*connect_go.Response[v2.ListOrganizationsResponse], error)
	// DeleteOrganization deletes the specified organization.
	DeleteOrganization(context.Context, *connect_go.Request[v2.DeleteOrganizationRequest]) (*connect_go.Response[v2.DeleteOrganizationResponse], error)
	// GetOrganizationInvitation retrieves the invitation for a Organization.
	GetOrganizationInvitation(context.Context, *connect_go.Request[v2.GetOrganizationInvitationRequest]) (*connect_go.Response[v2.GetOrganizationInvitationResponse], error)
	// JoinOrganization makes the caller a OrganizationMember of the Organization.
	JoinOrganization(context.Context, *connect_go.Request[v2.JoinOrganizationRequest]) (*connect_go.Response[v2.JoinOrganizationResponse], error)
	// ResetOrganizationInvitation resets the invitation_id for a Organization.
	ResetOrganizationInvitation(context.Context, *connect_go.Request[v2.ResetOrganizationInvitationRequest]) (*connect_go.Response[v2.ResetOrganizationInvitationResponse], error)
	// ListOrganizationMembers lists the members of a Organization.
	ListOrganizationMembers(context.Context, *connect_go.Request[v2.ListOrganizationMembersRequest]) (*connect_go.Response[v2.ListOrganizationMembersResponse], error)
	// UpdateOrganizationMember updates organization membership properties.
	UpdateOrganizationMember(context.Context, *connect_go.Request[v2.UpdateOrganizationMemberRequest]) (*connect_go.Response[v2.UpdateOrganizationMemberResponse], error)
	// DeleteOrganizationMember removes a OrganizationMember from the Organization.
	DeleteOrganizationMember(context.Context, *connect_go.Request[v2.DeleteOrganizationMemberRequest]) (*connect_go.Response[v2.DeleteOrganizationMemberResponse], error)
	// GetOrganizationSettings retrieves the settings of a Organization.
	GetOrganizationSettings(context.Context, *connect_go.Request[v2.GetOrganizationSettingsRequest]) (*connect_go.Response[v2.GetOrganizationSettingsResponse], error)
	// UpdateOrganizationSettings updates the settings of a Organization.
	UpdateOrganizationSettings(context.Context, *connect_go.Request[v2.UpdateOrganizationSettingsRequest]) (*connect_go.Response[v2.UpdateOrganizationSettingsResponse], error)
}

// NewOrganizationServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOrganizationServiceHandler(svc OrganizationServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/gitpod.experimental.v2.OrganizationService/CreateOrganization", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/CreateOrganization",
		svc.CreateOrganization,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v2.OrganizationService/GetOrganization", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/GetOrganization",
		svc.GetOrganization,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v2.OrganizationService/UpdateOrganization", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/UpdateOrganization",
		svc.UpdateOrganization,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v2.OrganizationService/ListOrganizations", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/ListOrganizations",
		svc.ListOrganizations,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v2.OrganizationService/DeleteOrganization", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/DeleteOrganization",
		svc.DeleteOrganization,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v2.OrganizationService/GetOrganizationInvitation", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/GetOrganizationInvitation",
		svc.GetOrganizationInvitation,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v2.OrganizationService/JoinOrganization", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/JoinOrganization",
		svc.JoinOrganization,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v2.OrganizationService/ResetOrganizationInvitation", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/ResetOrganizationInvitation",
		svc.ResetOrganizationInvitation,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v2.OrganizationService/ListOrganizationMembers", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/ListOrganizationMembers",
		svc.ListOrganizationMembers,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v2.OrganizationService/UpdateOrganizationMember", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/UpdateOrganizationMember",
		svc.UpdateOrganizationMember,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v2.OrganizationService/DeleteOrganizationMember", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/DeleteOrganizationMember",
		svc.DeleteOrganizationMember,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v2.OrganizationService/GetOrganizationSettings", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/GetOrganizationSettings",
		svc.GetOrganizationSettings,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v2.OrganizationService/UpdateOrganizationSettings", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v2.OrganizationService/UpdateOrganizationSettings",
		svc.UpdateOrganizationSettings,
		opts...,
	))
	return "/gitpod.experimental.v2.OrganizationService/", mux
}

// UnimplementedOrganizationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOrganizationServiceHandler struct{}

func (UnimplementedOrganizationServiceHandler) CreateOrganization(context.Context, *connect_go.Request[v2.CreateOrganizationRequest]) (*connect_go.Response[v2.CreateOrganizationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.CreateOrganization is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) GetOrganization(context.Context, *connect_go.Request[v2.GetOrganizationRequest]) (*connect_go.Response[v2.GetOrganizationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.GetOrganization is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) UpdateOrganization(context.Context, *connect_go.Request[v2.UpdateOrganizationRequest]) (*connect_go.Response[v2.UpdateOrganizationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.UpdateOrganization is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) ListOrganizations(context.Context, *connect_go.Request[v2.ListOrganizationsRequest]) (*connect_go.Response[v2.ListOrganizationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.ListOrganizations is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) DeleteOrganization(context.Context, *connect_go.Request[v2.DeleteOrganizationRequest]) (*connect_go.Response[v2.DeleteOrganizationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.DeleteOrganization is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) GetOrganizationInvitation(context.Context, *connect_go.Request[v2.GetOrganizationInvitationRequest]) (*connect_go.Response[v2.GetOrganizationInvitationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.GetOrganizationInvitation is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) JoinOrganization(context.Context, *connect_go.Request[v2.JoinOrganizationRequest]) (*connect_go.Response[v2.JoinOrganizationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.JoinOrganization is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) ResetOrganizationInvitation(context.Context, *connect_go.Request[v2.ResetOrganizationInvitationRequest]) (*connect_go.Response[v2.ResetOrganizationInvitationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.ResetOrganizationInvitation is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) ListOrganizationMembers(context.Context, *connect_go.Request[v2.ListOrganizationMembersRequest]) (*connect_go.Response[v2.ListOrganizationMembersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.ListOrganizationMembers is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) UpdateOrganizationMember(context.Context, *connect_go.Request[v2.UpdateOrganizationMemberRequest]) (*connect_go.Response[v2.UpdateOrganizationMemberResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.UpdateOrganizationMember is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) DeleteOrganizationMember(context.Context, *connect_go.Request[v2.DeleteOrganizationMemberRequest]) (*connect_go.Response[v2.DeleteOrganizationMemberResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.DeleteOrganizationMember is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) GetOrganizationSettings(context.Context, *connect_go.Request[v2.GetOrganizationSettingsRequest]) (*connect_go.Response[v2.GetOrganizationSettingsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.GetOrganizationSettings is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) UpdateOrganizationSettings(context.Context, *connect_go.Request[v2.UpdateOrganizationSettingsRequest]) (*connect_go.Response[v2.UpdateOrganizationSettingsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v2.OrganizationService.UpdateOrganizationSettings is not implemented"))
}
