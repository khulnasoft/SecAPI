// Code generated by protoc-gen-alias. DO NOT EDIT.
package v1

import "github.com/khulnasoft/secapi/networking/v1alpha3"

// `WorkloadGroup` enables specifying the properties of a single workload for bootstrap and
// provides a template for `WorkloadEntry`, similar to how `Deployment` specifies properties
// of workloads via `Pod` templates. A `WorkloadGroup` can have more than one `WorkloadEntry`.
// `WorkloadGroup` has no relationship to resources which control service registry like `ServiceEntry`
// and as such doesn't configure host name for these workloads.
//
// <!-- crd generation tags
// +cue-gen:WorkloadGroup:groupName:networking.khulnasoft.com
// +cue-gen:WorkloadGroup:versions:v1beta1,v1alpha3,v1
// +cue-gen:WorkloadGroup:labels:app=khulnasoft-pilot,chart=khulnasoft,heritage=Tiller,release=khulnasoft
// +cue-gen:WorkloadGroup:subresource:status
// +cue-gen:WorkloadGroup:scope:Namespaced
// +cue-gen:WorkloadGroup:resource:categories=khulnasoft-io,networking-khulnasoft-io,shortNames=wg,plural=workloadgroups
// +cue-gen:WorkloadGroup:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:WorkloadGroup:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.khulnasoft.com/v1alpha3
// +genclient
// +k8s:deepcopy-gen=true
// -->
type WorkloadGroup = v1alpha3.WorkloadGroup

// `ObjectMeta` describes metadata that will be attached to a `WorkloadEntry`.
// It is a subset of the supported Kubernetes metadata.
type WorkloadGroup_ObjectMeta = v1alpha3.WorkloadGroup_ObjectMeta
type ReadinessProbe = v1alpha3.ReadinessProbe

// `httpGet` is performed to a given endpoint
// and the status/able to connect determines health.
type ReadinessProbe_HttpGet = v1alpha3.ReadinessProbe_HttpGet

// Health is determined by if the proxy is able to connect.
type ReadinessProbe_TcpSocket = v1alpha3.ReadinessProbe_TcpSocket

// Health is determined by how the command that is executed exited.
type ReadinessProbe_Exec = v1alpha3.ReadinessProbe_Exec
type HTTPHealthCheckConfig = v1alpha3.HTTPHealthCheckConfig
type HTTPHeader = v1alpha3.HTTPHeader
type TCPHealthCheckConfig = v1alpha3.TCPHealthCheckConfig
type ExecHealthCheckConfig = v1alpha3.ExecHealthCheckConfig
