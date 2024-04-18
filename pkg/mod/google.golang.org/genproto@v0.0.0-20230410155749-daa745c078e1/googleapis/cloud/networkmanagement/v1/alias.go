// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package networkmanagement aliases all exported identifiers in package
// "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb".
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package networkmanagement

import (
	src "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
const (
	AbortInfo_CAUSE_UNSPECIFIED                                   = src.AbortInfo_CAUSE_UNSPECIFIED
	AbortInfo_DESTINATION_ENDPOINT_NOT_FOUND                      = src.AbortInfo_DESTINATION_ENDPOINT_NOT_FOUND
	AbortInfo_INTERNAL_ERROR                                      = src.AbortInfo_INTERNAL_ERROR
	AbortInfo_INVALID_ARGUMENT                                    = src.AbortInfo_INVALID_ARGUMENT
	AbortInfo_MISMATCHED_DESTINATION_NETWORK                      = src.AbortInfo_MISMATCHED_DESTINATION_NETWORK
	AbortInfo_MISMATCHED_SOURCE_NETWORK                           = src.AbortInfo_MISMATCHED_SOURCE_NETWORK
	AbortInfo_NO_EXTERNAL_IP                                      = src.AbortInfo_NO_EXTERNAL_IP
	AbortInfo_NO_SOURCE_LOCATION                                  = src.AbortInfo_NO_SOURCE_LOCATION
	AbortInfo_PERMISSION_DENIED                                   = src.AbortInfo_PERMISSION_DENIED
	AbortInfo_SOURCE_ENDPOINT_NOT_FOUND                           = src.AbortInfo_SOURCE_ENDPOINT_NOT_FOUND
	AbortInfo_TRACE_TOO_LONG                                      = src.AbortInfo_TRACE_TOO_LONG
	AbortInfo_UNINTENDED_DESTINATION                              = src.AbortInfo_UNINTENDED_DESTINATION
	AbortInfo_UNKNOWN_IP                                          = src.AbortInfo_UNKNOWN_IP
	AbortInfo_UNKNOWN_NETWORK                                     = src.AbortInfo_UNKNOWN_NETWORK
	AbortInfo_UNKNOWN_PROJECT                                     = src.AbortInfo_UNKNOWN_PROJECT
	AbortInfo_UNSUPPORTED                                         = src.AbortInfo_UNSUPPORTED
	DeliverInfo_CLOUD_SQL_INSTANCE                                = src.DeliverInfo_CLOUD_SQL_INSTANCE
	DeliverInfo_GKE_MASTER                                        = src.DeliverInfo_GKE_MASTER
	DeliverInfo_GOOGLE_API                                        = src.DeliverInfo_GOOGLE_API
	DeliverInfo_INSTANCE                                          = src.DeliverInfo_INSTANCE
	DeliverInfo_INTERNET                                          = src.DeliverInfo_INTERNET
	DeliverInfo_TARGET_UNSPECIFIED                                = src.DeliverInfo_TARGET_UNSPECIFIED
	DropInfo_CAUSE_UNSPECIFIED                                    = src.DropInfo_CAUSE_UNSPECIFIED
	DropInfo_CLOUD_SQL_INSTANCE_NO_IP_ADDRESS                     = src.DropInfo_CLOUD_SQL_INSTANCE_NO_IP_ADDRESS
	DropInfo_CLOUD_SQL_INSTANCE_UNAUTHORIZED_ACCESS               = src.DropInfo_CLOUD_SQL_INSTANCE_UNAUTHORIZED_ACCESS
	DropInfo_DROPPED_INSIDE_CLOUD_SQL_SERVICE                     = src.DropInfo_DROPPED_INSIDE_CLOUD_SQL_SERVICE
	DropInfo_DROPPED_INSIDE_GKE_SERVICE                           = src.DropInfo_DROPPED_INSIDE_GKE_SERVICE
	DropInfo_FIREWALL_BLOCKING_LOAD_BALANCER_BACKEND_HEALTH_CHECK = src.DropInfo_FIREWALL_BLOCKING_LOAD_BALANCER_BACKEND_HEALTH_CHECK
	DropInfo_FIREWALL_RULE                                        = src.DropInfo_FIREWALL_RULE
	DropInfo_FOREIGN_IP_DISALLOWED                                = src.DropInfo_FOREIGN_IP_DISALLOWED
	DropInfo_FORWARDING_RULE_MISMATCH                             = src.DropInfo_FORWARDING_RULE_MISMATCH
	DropInfo_FORWARDING_RULE_NO_INSTANCES                         = src.DropInfo_FORWARDING_RULE_NO_INSTANCES
	DropInfo_GKE_MASTER_UNAUTHORIZED_ACCESS                       = src.DropInfo_GKE_MASTER_UNAUTHORIZED_ACCESS
	DropInfo_GOOGLE_MANAGED_SERVICE_NO_PEERING                    = src.DropInfo_GOOGLE_MANAGED_SERVICE_NO_PEERING
	DropInfo_INSTANCE_NOT_RUNNING                                 = src.DropInfo_INSTANCE_NOT_RUNNING
	DropInfo_NO_EXTERNAL_ADDRESS                                  = src.DropInfo_NO_EXTERNAL_ADDRESS
	DropInfo_NO_ROUTE                                             = src.DropInfo_NO_ROUTE
	DropInfo_PRIVATE_GOOGLE_ACCESS_DISALLOWED                     = src.DropInfo_PRIVATE_GOOGLE_ACCESS_DISALLOWED
	DropInfo_PRIVATE_TRAFFIC_TO_INTERNET                          = src.DropInfo_PRIVATE_TRAFFIC_TO_INTERNET
	DropInfo_ROUTE_BLACKHOLE                                      = src.DropInfo_ROUTE_BLACKHOLE
	DropInfo_ROUTE_WRONG_NETWORK                                  = src.DropInfo_ROUTE_WRONG_NETWORK
	DropInfo_TRAFFIC_TYPE_BLOCKED                                 = src.DropInfo_TRAFFIC_TYPE_BLOCKED
	DropInfo_UNKNOWN_EXTERNAL_ADDRESS                             = src.DropInfo_UNKNOWN_EXTERNAL_ADDRESS
	DropInfo_UNKNOWN_INTERNAL_ADDRESS                             = src.DropInfo_UNKNOWN_INTERNAL_ADDRESS
	Endpoint_GCP_NETWORK                                          = src.Endpoint_GCP_NETWORK
	Endpoint_NETWORK_TYPE_UNSPECIFIED                             = src.Endpoint_NETWORK_TYPE_UNSPECIFIED
	Endpoint_NON_GCP_NETWORK                                      = src.Endpoint_NON_GCP_NETWORK
	FirewallInfo_FIREWALL_RULE_TYPE_UNSPECIFIED                   = src.FirewallInfo_FIREWALL_RULE_TYPE_UNSPECIFIED
	FirewallInfo_HIERARCHICAL_FIREWALL_POLICY_RULE                = src.FirewallInfo_HIERARCHICAL_FIREWALL_POLICY_RULE
	FirewallInfo_IMPLIED_VPC_FIREWALL_RULE                        = src.FirewallInfo_IMPLIED_VPC_FIREWALL_RULE
	FirewallInfo_VPC_FIREWALL_RULE                                = src.FirewallInfo_VPC_FIREWALL_RULE
	ForwardInfo_CLOUD_SQL_INSTANCE                                = src.ForwardInfo_CLOUD_SQL_INSTANCE
	ForwardInfo_GKE_MASTER                                        = src.ForwardInfo_GKE_MASTER
	ForwardInfo_IMPORTED_CUSTOM_ROUTE_NEXT_HOP                    = src.ForwardInfo_IMPORTED_CUSTOM_ROUTE_NEXT_HOP
	ForwardInfo_INTERCONNECT                                      = src.ForwardInfo_INTERCONNECT
	ForwardInfo_PEERING_VPC                                       = src.ForwardInfo_PEERING_VPC
	ForwardInfo_TARGET_UNSPECIFIED                                = src.ForwardInfo_TARGET_UNSPECIFIED
	ForwardInfo_VPN_GATEWAY                                       = src.ForwardInfo_VPN_GATEWAY
	LoadBalancerBackend_CONFIGURED                                = src.LoadBalancerBackend_CONFIGURED
	LoadBalancerBackend_HEALTH_CHECK_FIREWALL_STATE_UNSPECIFIED   = src.LoadBalancerBackend_HEALTH_CHECK_FIREWALL_STATE_UNSPECIFIED
	LoadBalancerBackend_MISCONFIGURED                             = src.LoadBalancerBackend_MISCONFIGURED
	LoadBalancerInfo_BACKEND_SERVICE                              = src.LoadBalancerInfo_BACKEND_SERVICE
	LoadBalancerInfo_BACKEND_TYPE_UNSPECIFIED                     = src.LoadBalancerInfo_BACKEND_TYPE_UNSPECIFIED
	LoadBalancerInfo_HTTP_PROXY                                   = src.LoadBalancerInfo_HTTP_PROXY
	LoadBalancerInfo_INTERNAL_TCP_UDP                             = src.LoadBalancerInfo_INTERNAL_TCP_UDP
	LoadBalancerInfo_LOAD_BALANCER_TYPE_UNSPECIFIED               = src.LoadBalancerInfo_LOAD_BALANCER_TYPE_UNSPECIFIED
	LoadBalancerInfo_NETWORK_TCP_UDP                              = src.LoadBalancerInfo_NETWORK_TCP_UDP
	LoadBalancerInfo_SSL_PROXY                                    = src.LoadBalancerInfo_SSL_PROXY
	LoadBalancerInfo_TARGET_POOL                                  = src.LoadBalancerInfo_TARGET_POOL
	LoadBalancerInfo_TCP_PROXY                                    = src.LoadBalancerInfo_TCP_PROXY
	ReachabilityDetails_AMBIGUOUS                                 = src.ReachabilityDetails_AMBIGUOUS
	ReachabilityDetails_REACHABLE                                 = src.ReachabilityDetails_REACHABLE
	ReachabilityDetails_RESULT_UNSPECIFIED                        = src.ReachabilityDetails_RESULT_UNSPECIFIED
	ReachabilityDetails_UNDETERMINED                              = src.ReachabilityDetails_UNDETERMINED
	ReachabilityDetails_UNREACHABLE                               = src.ReachabilityDetails_UNREACHABLE
	RouteInfo_DYNAMIC                                             = src.RouteInfo_DYNAMIC
	RouteInfo_NEXT_HOP_BLACKHOLE                                  = src.RouteInfo_NEXT_HOP_BLACKHOLE
	RouteInfo_NEXT_HOP_ILB                                        = src.RouteInfo_NEXT_HOP_ILB
	RouteInfo_NEXT_HOP_INSTANCE                                   = src.RouteInfo_NEXT_HOP_INSTANCE
	RouteInfo_NEXT_HOP_INTERCONNECT                               = src.RouteInfo_NEXT_HOP_INTERCONNECT
	RouteInfo_NEXT_HOP_INTERNET_GATEWAY                           = src.RouteInfo_NEXT_HOP_INTERNET_GATEWAY
	RouteInfo_NEXT_HOP_IP                                         = src.RouteInfo_NEXT_HOP_IP
	RouteInfo_NEXT_HOP_NETWORK                                    = src.RouteInfo_NEXT_HOP_NETWORK
	RouteInfo_NEXT_HOP_PEERING                                    = src.RouteInfo_NEXT_HOP_PEERING
	RouteInfo_NEXT_HOP_ROUTER_APPLIANCE                           = src.RouteInfo_NEXT_HOP_ROUTER_APPLIANCE
	RouteInfo_NEXT_HOP_TYPE_UNSPECIFIED                           = src.RouteInfo_NEXT_HOP_TYPE_UNSPECIFIED
	RouteInfo_NEXT_HOP_VPN_GATEWAY                                = src.RouteInfo_NEXT_HOP_VPN_GATEWAY
	RouteInfo_NEXT_HOP_VPN_TUNNEL                                 = src.RouteInfo_NEXT_HOP_VPN_TUNNEL
	RouteInfo_PEERING_DYNAMIC                                     = src.RouteInfo_PEERING_DYNAMIC
	RouteInfo_PEERING_STATIC                                      = src.RouteInfo_PEERING_STATIC
	RouteInfo_PEERING_SUBNET                                      = src.RouteInfo_PEERING_SUBNET
	RouteInfo_ROUTE_TYPE_UNSPECIFIED                              = src.RouteInfo_ROUTE_TYPE_UNSPECIFIED
	RouteInfo_STATIC                                              = src.RouteInfo_STATIC
	RouteInfo_SUBNET                                              = src.RouteInfo_SUBNET
	Step_ABORT                                                    = src.Step_ABORT
	Step_APPLY_EGRESS_FIREWALL_RULE                               = src.Step_APPLY_EGRESS_FIREWALL_RULE
	Step_APPLY_FORWARDING_RULE                                    = src.Step_APPLY_FORWARDING_RULE
	Step_APPLY_INGRESS_FIREWALL_RULE                              = src.Step_APPLY_INGRESS_FIREWALL_RULE
	Step_APPLY_ROUTE                                              = src.Step_APPLY_ROUTE
	Step_ARRIVE_AT_EXTERNAL_LOAD_BALANCER                         = src.Step_ARRIVE_AT_EXTERNAL_LOAD_BALANCER
	Step_ARRIVE_AT_INSTANCE                                       = src.Step_ARRIVE_AT_INSTANCE
	Step_ARRIVE_AT_INTERNAL_LOAD_BALANCER                         = src.Step_ARRIVE_AT_INTERNAL_LOAD_BALANCER
	Step_ARRIVE_AT_VPN_GATEWAY                                    = src.Step_ARRIVE_AT_VPN_GATEWAY
	Step_ARRIVE_AT_VPN_TUNNEL                                     = src.Step_ARRIVE_AT_VPN_TUNNEL
	Step_DELIVER                                                  = src.Step_DELIVER
	Step_DROP                                                     = src.Step_DROP
	Step_FORWARD                                                  = src.Step_FORWARD
	Step_NAT                                                      = src.Step_NAT
	Step_PROXY_CONNECTION                                         = src.Step_PROXY_CONNECTION
	Step_SPOOFING_APPROVED                                        = src.Step_SPOOFING_APPROVED
	Step_START_FROM_CLOUD_SQL_INSTANCE                            = src.Step_START_FROM_CLOUD_SQL_INSTANCE
	Step_START_FROM_GKE_MASTER                                    = src.Step_START_FROM_GKE_MASTER
	Step_START_FROM_INSTANCE                                      = src.Step_START_FROM_INSTANCE
	Step_START_FROM_INTERNET                                      = src.Step_START_FROM_INTERNET
	Step_START_FROM_PRIVATE_NETWORK                               = src.Step_START_FROM_PRIVATE_NETWORK
	Step_STATE_UNSPECIFIED                                        = src.Step_STATE_UNSPECIFIED
	Step_VIEWER_PERMISSION_MISSING                                = src.Step_VIEWER_PERMISSION_MISSING
	VpnTunnelInfo_DYNAMIC                                         = src.VpnTunnelInfo_DYNAMIC
	VpnTunnelInfo_POLICY_BASED                                    = src.VpnTunnelInfo_POLICY_BASED
	VpnTunnelInfo_ROUTE_BASED                                     = src.VpnTunnelInfo_ROUTE_BASED
	VpnTunnelInfo_ROUTING_TYPE_UNSPECIFIED                        = src.VpnTunnelInfo_ROUTING_TYPE_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
var (
	AbortInfo_Cause_name                                           = src.AbortInfo_Cause_name
	AbortInfo_Cause_value                                          = src.AbortInfo_Cause_value
	DeliverInfo_Target_name                                        = src.DeliverInfo_Target_name
	DeliverInfo_Target_value                                       = src.DeliverInfo_Target_value
	DropInfo_Cause_name                                            = src.DropInfo_Cause_name
	DropInfo_Cause_value                                           = src.DropInfo_Cause_value
	Endpoint_NetworkType_name                                      = src.Endpoint_NetworkType_name
	Endpoint_NetworkType_value                                     = src.Endpoint_NetworkType_value
	File_google_cloud_networkmanagement_v1_connectivity_test_proto = src.File_google_cloud_networkmanagement_v1_connectivity_test_proto
	File_google_cloud_networkmanagement_v1_reachability_proto      = src.File_google_cloud_networkmanagement_v1_reachability_proto
	File_google_cloud_networkmanagement_v1_trace_proto             = src.File_google_cloud_networkmanagement_v1_trace_proto
	FirewallInfo_FirewallRuleType_name                             = src.FirewallInfo_FirewallRuleType_name
	FirewallInfo_FirewallRuleType_value                            = src.FirewallInfo_FirewallRuleType_value
	ForwardInfo_Target_name                                        = src.ForwardInfo_Target_name
	ForwardInfo_Target_value                                       = src.ForwardInfo_Target_value
	LoadBalancerBackend_HealthCheckFirewallState_name              = src.LoadBalancerBackend_HealthCheckFirewallState_name
	LoadBalancerBackend_HealthCheckFirewallState_value             = src.LoadBalancerBackend_HealthCheckFirewallState_value
	LoadBalancerInfo_BackendType_name                              = src.LoadBalancerInfo_BackendType_name
	LoadBalancerInfo_BackendType_value                             = src.LoadBalancerInfo_BackendType_value
	LoadBalancerInfo_LoadBalancerType_name                         = src.LoadBalancerInfo_LoadBalancerType_name
	LoadBalancerInfo_LoadBalancerType_value                        = src.LoadBalancerInfo_LoadBalancerType_value
	ReachabilityDetails_Result_name                                = src.ReachabilityDetails_Result_name
	ReachabilityDetails_Result_value                               = src.ReachabilityDetails_Result_value
	RouteInfo_NextHopType_name                                     = src.RouteInfo_NextHopType_name
	RouteInfo_NextHopType_value                                    = src.RouteInfo_NextHopType_value
	RouteInfo_RouteType_name                                       = src.RouteInfo_RouteType_name
	RouteInfo_RouteType_value                                      = src.RouteInfo_RouteType_value
	Step_State_name                                                = src.Step_State_name
	Step_State_value                                               = src.Step_State_value
	VpnTunnelInfo_RoutingType_name                                 = src.VpnTunnelInfo_RoutingType_name
	VpnTunnelInfo_RoutingType_value                                = src.VpnTunnelInfo_RoutingType_value
)

// Details of the final state "abort" and associated resource.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type AbortInfo = src.AbortInfo

// Abort cause types:
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type AbortInfo_Cause = src.AbortInfo_Cause

// For display only. Metadata associated with a Cloud SQL instance.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type CloudSQLInstanceInfo = src.CloudSQLInstanceInfo

// A Connectivity Test for a network reachability analysis.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type ConnectivityTest = src.ConnectivityTest

// Request for the `CreateConnectivityTest` method.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type CreateConnectivityTestRequest = src.CreateConnectivityTestRequest

// Request for the `DeleteConnectivityTest` method.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type DeleteConnectivityTestRequest = src.DeleteConnectivityTestRequest

// Details of the final state "deliver" and associated resource.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type DeliverInfo = src.DeliverInfo

// Deliver target types:
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type DeliverInfo_Target = src.DeliverInfo_Target

// Details of the final state "drop" and associated resource.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type DropInfo = src.DropInfo

// Drop cause types:
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type DropInfo_Cause = src.DropInfo_Cause

// Source or destination of the Connectivity Test.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type Endpoint = src.Endpoint

// For display only. The specification of the endpoints for the test.
// EndpointInfo is derived from source and destination Endpoint and validated
// by the backend data plane model.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type EndpointInfo = src.EndpointInfo

// The type definition of an endpoint's network. Use one of the following
// choices:
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type Endpoint_NetworkType = src.Endpoint_NetworkType

// For display only. Metadata associated with a VPC firewall rule, an implied
// VPC firewall rule, or a hierarchical firewall policy rule.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type FirewallInfo = src.FirewallInfo

// The firewall rule's type.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type FirewallInfo_FirewallRuleType = src.FirewallInfo_FirewallRuleType

// Details of the final state "forward" and associated resource.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type ForwardInfo = src.ForwardInfo

// Forward target types.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type ForwardInfo_Target = src.ForwardInfo_Target

// For display only. Metadata associated with a Compute Engine forwarding
// rule.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type ForwardingRuleInfo = src.ForwardingRuleInfo

// For display only. Metadata associated with a Google Kubernetes Engine (GKE)
// cluster master.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type GKEMasterInfo = src.GKEMasterInfo

// Request for the `GetConnectivityTest` method.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type GetConnectivityTestRequest = src.GetConnectivityTestRequest

// For display only. Metadata associated with a Compute Engine instance.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type InstanceInfo = src.InstanceInfo

// Request for the `ListConnectivityTests` method.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type ListConnectivityTestsRequest = src.ListConnectivityTestsRequest

// Response for the `ListConnectivityTests` method.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type ListConnectivityTestsResponse = src.ListConnectivityTestsResponse

// For display only. Metadata associated with a specific load balancer
// backend.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type LoadBalancerBackend = src.LoadBalancerBackend

// State of a health check firewall configuration:
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type LoadBalancerBackend_HealthCheckFirewallState = src.LoadBalancerBackend_HealthCheckFirewallState

// For display only. Metadata associated with a load balancer.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type LoadBalancerInfo = src.LoadBalancerInfo

// The type definition for a load balancer backend configuration:
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type LoadBalancerInfo_BackendType = src.LoadBalancerInfo_BackendType

// The type definition for a load balancer:
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type LoadBalancerInfo_LoadBalancerType = src.LoadBalancerInfo_LoadBalancerType

// For display only. Metadata associated with a Compute Engine network.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type NetworkInfo = src.NetworkInfo

// Metadata describing an [Operation][google.longrunning.Operation]
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type OperationMetadata = src.OperationMetadata

// Results of the configuration analysis from the last run of the test.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type ReachabilityDetails = src.ReachabilityDetails

// The overall result of the test's configuration analysis.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type ReachabilityDetails_Result = src.ReachabilityDetails_Result

// ReachabilityServiceClient is the client API for ReachabilityService
// service. For semantics around ctx use and closing/ending streaming RPCs,
// please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type ReachabilityServiceClient = src.ReachabilityServiceClient

// ReachabilityServiceServer is the server API for ReachabilityService
// service.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type ReachabilityServiceServer = src.ReachabilityServiceServer

// Request for the `RerunConnectivityTest` method.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type RerunConnectivityTestRequest = src.RerunConnectivityTestRequest

// For display only. Metadata associated with a Compute Engine route.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type RouteInfo = src.RouteInfo

// Type of next hop:
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type RouteInfo_NextHopType = src.RouteInfo_NextHopType

// Type of route:
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type RouteInfo_RouteType = src.RouteInfo_RouteType

// A simulated forwarding path is composed of multiple steps. Each step has a
// well-defined state and an associated configuration.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type Step = src.Step
type Step_Abort = src.Step_Abort
type Step_CloudSqlInstance = src.Step_CloudSqlInstance
type Step_Deliver = src.Step_Deliver
type Step_Drop = src.Step_Drop
type Step_Endpoint = src.Step_Endpoint
type Step_Firewall = src.Step_Firewall
type Step_Forward = src.Step_Forward
type Step_ForwardingRule = src.Step_ForwardingRule
type Step_GkeMaster = src.Step_GkeMaster
type Step_Instance = src.Step_Instance
type Step_LoadBalancer = src.Step_LoadBalancer
type Step_Network = src.Step_Network
type Step_Route = src.Step_Route

// Type of states that are defined in the network state machine. Each step in
// the packet trace is in a specific state.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type Step_State = src.Step_State
type Step_VpnGateway = src.Step_VpnGateway
type Step_VpnTunnel = src.Step_VpnTunnel

// Trace represents one simulated packet forwarding path. - Each trace
// contains multiple ordered steps. - Each step is in a particular state with
// associated configuration. - State is categorized as final or non-final
// states. - Each final state has a reason associated. - Each trace must end
// with a final state (the last step). ```
// |---------------------Trace----------------------| Step1(State) Step2(State)
// --- StepN(State(final)) ```
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type Trace = src.Trace

// UnimplementedReachabilityServiceServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type UnimplementedReachabilityServiceServer = src.UnimplementedReachabilityServiceServer

// Request for the `UpdateConnectivityTest` method.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type UpdateConnectivityTestRequest = src.UpdateConnectivityTestRequest

// For display only. Metadata associated with a Compute Engine VPN gateway.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type VpnGatewayInfo = src.VpnGatewayInfo

// For display only. Metadata associated with a Compute Engine VPN tunnel.
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type VpnTunnelInfo = src.VpnTunnelInfo

// Types of VPN routing policy. For details, refer to [Networks and Tunnel
// routing](https://cloud.google.com/network-connectivity/docs/vpn/concepts/choosing-networks-routing/).
//
// Deprecated: Please use types in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
type VpnTunnelInfo_RoutingType = src.VpnTunnelInfo_RoutingType

// Deprecated: Please use funcs in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
func NewReachabilityServiceClient(cc grpc.ClientConnInterface) ReachabilityServiceClient {
	return src.NewReachabilityServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb
func RegisterReachabilityServiceServer(s *grpc.Server, srv ReachabilityServiceServer) {
	src.RegisterReachabilityServiceServer(s, srv)
}
