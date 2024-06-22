
// GENERATED FILE -- DO NOT EDIT

package annotation

type FeatureStatus int

const (
	Alpha FeatureStatus = iota
	Beta
	Stable
)

func (s FeatureStatus) String() string {
	switch s {
	case Alpha:
		return "Alpha"
	case Beta:
		return "Beta"
	case Stable:
		return "Stable"
	}
	return "Unknown"
}

type ResourceTypes int

const (
	Unknown ResourceTypes = iota
    Any
    AuthorizationPolicy
    Ingress
    Namespace
    Pod
    Service
    WorkloadEntry
)

func (r ResourceTypes) String() string {
	switch r {
	case 1:
		return "Any"
	case 2:
		return "AuthorizationPolicy"
	case 3:
		return "Ingress"
	case 4:
		return "Namespace"
	case 5:
		return "Pod"
	case 6:
		return "Service"
	case 7:
		return "WorkloadEntry"
	}
	return "Unknown"
}

// Instance describes a single resource annotation
type Instance struct {
	// The name of the annotation.
	Name string

	// Description of the annotation.
	Description string

	// FeatureStatus of this annotation.
	FeatureStatus FeatureStatus

	// Hide the existence of this annotation when outputting usage information.
	Hidden bool

	// Mark this annotation as deprecated when generating usage information.
	Deprecated bool

	// The types of resources this annotation applies to.
	Resources []ResourceTypes
}

var (

	AlphaCanonicalServiceAccounts = Instance {
		Name:          "alpha.khulnasoft.com/canonical-serviceaccounts",
		Description:   "Specifies the non-Kubernetes service accounts that are "+
                        "allowed to run this service.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Service,
		},
	}

	AlphaIdentity = Instance {
		Name:          "alpha.khulnasoft.com/identity",
		Description:   "Identity for the workload.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	AlphaKubernetesServiceAccounts = Instance {
		Name:          "alpha.khulnasoft.com/kubernetes-serviceaccounts",
		Description:   "Specifies the Kubernetes service accounts that are "+
                        "allowed to run this service on the VMs.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Service,
		},
	}

	GalleyAnalyzeSuppress = Instance {
		Name:          "galley.khulnasoft.com/analyze-suppress",
		Description:   "A comma separated list of configuration analysis message "+
                        "codes to suppress when Khulnasoft analyzers are run. For "+
                        "example, to suppress reporting of IST0103 "+
                        "(PodMissingProxy) and IST0108 (UnknownAnnotation) on a "+
                        "resource, apply the annotation "+
                        "'galley.khulnasoft.com/analyze-suppress=IST0108,IST0103'. If "+
                        "the value is '*', then all configuration analysis "+
                        "messages are suppressed.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	GatewayControllerVersion = Instance {
		Name:          "gateway.khulnasoft.com/controller-version",
		Description:   "A version added to the Gateway by the controller "+
                        "specifying the `controller version`.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	InjectTemplates = Instance {
		Name:          "inject.khulnasoft.com/templates",
		Description:   "The name of the inject template(s) to use, as a comma "+
                        "separate list. See "+
                        "https://khulnasoft.com/latest/docs/setup/additional-setup/sidecar-injection/#custom-templates-experimental "+
                        "for more information.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	OperatorInstallChartOwner = Instance {
		Name:          "install.operator.khulnasoft.com/chart-owner",
		Description:   "Represents the name of the chart used to create this "+
                        "resource.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	OperatorInstallOwnerGeneration = Instance {
		Name:          "install.operator.khulnasoft.com/owner-generation",
		Description:   "Represents the generation to which the resource was last "+
                        "reconciled.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	OperatorInstallVersion = Instance {
		Name:          "install.operator.khulnasoft.com/version",
		Description:   "Represents the Khulnasoft version associated with the resource",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	IoKhulnasoftAutoRegistrationGroup = Instance {
		Name:          "khulnasoft.com/autoRegistrationGroup",
		Description:   "On a WorkloadEntry stores the associated WorkloadGroup.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			WorkloadEntry,
		},
	}

	IoKhulnasoftConnectedAt = Instance {
		Name:          "khulnasoft.com/connectedAt",
		Description:   "On a WorkloadEntry stores the time in nanoseconds when "+
                        "the associated workload connected to a Pilot instance.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			WorkloadEntry,
		},
	}

	IoKhulnasoftDisconnectedAt = Instance {
		Name:          "khulnasoft.com/disconnectedAt",
		Description:   "On a WorkloadEntry stores the time in nanoseconds when "+
                        "the associated workload disconnected from a Pilot "+
                        "instance.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			WorkloadEntry,
		},
	}

	IoKhulnasoftDryRun = Instance {
		Name:          "khulnasoft.com/dry-run",
		Description:   "Specifies whether or not the given resource is in dry-run "+
                        "mode. See "+
                        "https://khulnasoft.com/latest/docs/tasks/security/authorization/authz-dry-run/ "+
                        "for more information.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			AuthorizationPolicy,
		},
	}

	IoKhulnasoftRev = Instance {
		Name:          "khulnasoft.com/rev",
		Description:   "Specifies a control plane revision to which a given proxy "+
                        "is connected. This annotation is added automatically, not "+
                        "set by a user. In contrary to the label khulnasoft.com/rev, it "+
                        "represents the actual revision, not the requested "+
                        "revision.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	IoKhulnasoftWorkloadController = Instance {
		Name:          "khulnasoft.com/workloadController",
		Description:   "On a WorkloadEntry should store the current/last pilot "+
                        "instance connected to the workload for XDS.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			WorkloadEntry,
		},
	}

	IoKubernetesIngressClass = Instance {
		Name:          "kubernetes.io/ingress.class",
		Description:   "Annotation on an Ingress resources denoting the class of "+
                        "controllers responsible for it.",
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Ingress,
		},
	}

	NetworkingExportTo = Instance {
		Name:          "networking.khulnasoft.com/exportTo",
		Description:   "Specifies the namespaces to which this service should be "+
                        "exported to. A value of '*' indicates it is reachable "+
                        "within the mesh '.' indicates it is reachable within its "+
                        "namespace.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Service,
		},
	}

	PrometheusMergeMetrics = Instance {
		Name:          "prometheus.khulnasoft.com/merge-metrics",
		Description:   "Specifies if application Prometheus metric will be merged "+
                        "with Envoy metrics for this workload.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	ProxyConfig = Instance {
		Name:          "proxy.khulnasoft.com/config",
		Description:   "Overrides for the proxy configuration for this specific "+
                        "proxy. Available options can be found at "+
                        "https://khulnasoft.com/docs/reference/config/khulnasoft.mesh.v1alpha1/#ProxyConfig.",
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	ProxyOverrides = Instance {
		Name:          "proxy.khulnasoft.com/overrides",
		Description:   "Used internally to indicate user-specified overrides in "+
                        "the proxy container of the pod during injection.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatusReadinessApplicationPorts = Instance {
		Name:          "readiness.status.sidecar.khulnasoft.com/applicationPorts",
		Description:   "Specifies the list of ports exposed by the application "+
                        "container. Used by the Envoy sidecar readiness probe to "+
                        "determine that Envoy is configured and ready to receive "+
                        "traffic.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatusReadinessFailureThreshold = Instance {
		Name:          "readiness.status.sidecar.khulnasoft.com/failureThreshold",
		Description:   "Specifies the failure threshold for the Envoy sidecar "+
                        "readiness probe.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatusReadinessInitialDelaySeconds = Instance {
		Name:          "readiness.status.sidecar.khulnasoft.com/initialDelaySeconds",
		Description:   "Specifies the initial delay (in seconds) for the Envoy "+
                        "sidecar readiness probe.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatusReadinessPeriodSeconds = Instance {
		Name:          "readiness.status.sidecar.khulnasoft.com/periodSeconds",
		Description:   "Specifies the period (in seconds) for the Envoy sidecar "+
                        "readiness probe.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarAgentLogLevel = Instance {
		Name:          "sidecar.khulnasoft.com/agentLogLevel",
		Description:   "Specifies the log output level for pilot-agent.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarBootstrapOverride = Instance {
		Name:          "sidecar.khulnasoft.com/bootstrapOverride",
		Description:   "Specifies an alternative Envoy bootstrap configuration "+
                        "file.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarComponentLogLevel = Instance {
		Name:          "sidecar.khulnasoft.com/componentLogLevel",
		Description:   "Specifies the component log level for Envoy.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarControlPlaneAuthPolicy = Instance {
		Name:          "sidecar.khulnasoft.com/controlPlaneAuthPolicy",
		Description:   "Specifies the auth policy used by the Khulnasoft control "+
                        "plane. If NONE, traffic will not be encrypted. If "+
                        "MUTUAL_TLS, traffic between Envoy sidecar will be wrapped "+
                        "into mutual TLS connections.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarDiscoveryAddress = Instance {
		Name:          "sidecar.khulnasoft.com/discoveryAddress",
		Description:   "Specifies the XDS discovery address to be used by the "+
                        "Envoy sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarEnableCoreDump = Instance {
		Name:          "sidecar.khulnasoft.com/enableCoreDump",
		Description:   "Specifies whether or not an Envoy sidecar should enable "+
                        "core dump.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarExtraStatTags = Instance {
		Name:          "sidecar.khulnasoft.com/extraStatTags",
		Description:   "An additional list of tags to extract from the in-proxy "+
                        "Khulnasoft Wasm telemetry. Each additional tag needs to be "+
                        "present in this list.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarInject = Instance {
		Name:          "sidecar.khulnasoft.com/inject",
		Description:   "Specifies whether or not an Envoy sidecar should be "+
                        "automatically injected into the workload. Deprecated in "+
                        "favor of `sidecar.khulnasoft.com/inject` label.",
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarInterceptionMode = Instance {
		Name:          "sidecar.khulnasoft.com/interceptionMode",
		Description:   "Specifies the mode used to redirect inbound connections "+
                        "to Envoy (REDIRECT or TPROXY).",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarLogLevel = Instance {
		Name:          "sidecar.khulnasoft.com/logLevel",
		Description:   "Specifies the log level for Envoy.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarProxyCPU = Instance {
		Name:          "sidecar.khulnasoft.com/proxyCPU",
		Description:   "Specifies the requested CPU setting for the Envoy "+
                        "sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarProxyCPULimit = Instance {
		Name:          "sidecar.khulnasoft.com/proxyCPULimit",
		Description:   "Specifies the CPU limit for the Envoy sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarProxyImage = Instance {
		Name:          "sidecar.khulnasoft.com/proxyImage",
		Description:   "Specifies the Docker image to be used by the Envoy "+
                        "sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarProxyImageType = Instance {
		Name:          "sidecar.khulnasoft.com/proxyImageType",
		Description:   "Specifies the Docker image type to be used by the Envoy "+
                        "sidecar. Khulnasoft publishes debug and distroless image types "+
                        "for every release tag.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarProxyMemory = Instance {
		Name:          "sidecar.khulnasoft.com/proxyMemory",
		Description:   "Specifies the requested memory setting for the Envoy "+
                        "sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarProxyMemoryLimit = Instance {
		Name:          "sidecar.khulnasoft.com/proxyMemoryLimit",
		Description:   "Specifies the memory limit for the Envoy sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarRewriteAppHTTPProbers = Instance {
		Name:          "sidecar.khulnasoft.com/rewriteAppHTTPProbers",
		Description:   "Rewrite HTTP readiness and liveness probes to be "+
                        "redirected to the Envoy sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatsHistogramBuckets = Instance {
		Name:          "sidecar.khulnasoft.com/statsHistogramBuckets",
		Description:   "Specifies the custom histogram buckets with a prefix "+
                        "matcher to separate the Khulnasoft mesh metrics from the Envoy "+
                        "stats, e.g. "+
                        "`{`khulnasoftcustom`:[1,5,10,50,100,500,1000,5000,10000],`cluster.xds-grpc`:[1,5,10,25,50,100,250,500,1000,2500,5000,10000]}`. "+
                        "Default buckets are "+
                        "`[0.5,1,5,10,25,50,100,250,500,1000,2500,5000,10000,30000,60000,300000,600000,1800000,3600000]`.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatsInclusionPrefixes = Instance {
		Name:          "sidecar.khulnasoft.com/statsInclusionPrefixes",
		Description:   "Specifies the comma separated list of prefixes of the "+
                        "stats to be emitted by Envoy.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatsInclusionRegexps = Instance {
		Name:          "sidecar.khulnasoft.com/statsInclusionRegexps",
		Description:   "Specifies the comma separated list of regexes the stats "+
                        "should match to be emitted by Envoy.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatsInclusionSuffixes = Instance {
		Name:          "sidecar.khulnasoft.com/statsInclusionSuffixes",
		Description:   "Specifies the comma separated list of suffixes of the "+
                        "stats to be emitted by Envoy.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatus = Instance {
		Name:          "sidecar.khulnasoft.com/status",
		Description:   "Generated by Envoy sidecar injection that indicates the "+
                        "status of the operation. Includes a version hash of the "+
                        "executed template, as well as names of injected "+
                        "resources.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarUserVolume = Instance {
		Name:          "sidecar.khulnasoft.com/userVolume",
		Description:   "Specifies one or more user volumes (as a JSON array) to "+
                        "be added to the Envoy sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarUserVolumeMount = Instance {
		Name:          "sidecar.khulnasoft.com/userVolumeMount",
		Description:   "Specifies one or more user volume mounts (as a JSON "+
                        "array) to be added to the Envoy sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatusPort = Instance {
		Name:          "status.sidecar.khulnasoft.com/port",
		Description:   "Specifies the HTTP status Port for the Envoy sidecar. If "+
                        "zero, the sidecar will not provide status.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	TopologyControlPlaneClusters = Instance {
		Name:          "topology.khulnasoft.com/controlPlaneClusters",
		Description:   "A comma-separated list of clusters (or * for any) running "+
                        "khulnasoftd that should attempt leader election for a remote "+
                        "cluster thats system namespace includes this annotation. "+
                        "Khulnasoftd will not attempt to lead unannotated remote "+
                        "clusters.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Namespace,
		},
	}

	TrafficNodeSelector = Instance {
		Name:          "traffic.khulnasoft.com/nodeSelector",
		Description:   "This annotation is a set of node-labels "+
                        "(key1=value,key2=value). If the annotated Service is of "+
                        "type NodePort and is a multi-network gateway (see "+
                        "topology.khulnasoft.com/network), the addresses for selected "+
                        "nodes will be used for cross-network communication.",
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Service,
		},
	}

	SidecarTrafficExcludeInboundPorts = Instance {
		Name:          "traffic.sidecar.khulnasoft.com/excludeInboundPorts",
		Description:   "A comma separated list of inbound ports to be excluded "+
                        "from redirection to Envoy. Only applies when all inbound "+
                        "traffic (i.e. '*') is being redirected.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficExcludeInterfaces = Instance {
		Name:          "traffic.sidecar.khulnasoft.com/excludeInterfaces",
		Description:   "A comma separated list of interfaces to be excluded from "+
                        "Khulnasoft traffic capture",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficExcludeOutboundIPRanges = Instance {
		Name:          "traffic.sidecar.khulnasoft.com/excludeOutboundIPRanges",
		Description:   "A comma separated list of IP ranges in CIDR form to be "+
                        "excluded from redirection. Only applies when all outbound "+
                        "traffic (i.e. '*') is being redirected.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficExcludeOutboundPorts = Instance {
		Name:          "traffic.sidecar.khulnasoft.com/excludeOutboundPorts",
		Description:   "A comma separated list of outbound ports to be excluded "+
                        "from redirection to Envoy.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficIncludeInboundPorts = Instance {
		Name:          "traffic.sidecar.khulnasoft.com/includeInboundPorts",
		Description:   "A comma separated list of inbound ports for which traffic "+
                        "is to be redirected to Envoy. The wildcard character '*' "+
                        "can be used to configure redirection for all ports. An "+
                        "empty list will disable all inbound redirection.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficIncludeOutboundIPRanges = Instance {
		Name:          "traffic.sidecar.khulnasoft.com/includeOutboundIPRanges",
		Description:   "A comma separated list of IP ranges in CIDR form to "+
                        "redirect to Envoy (optional). The wildcard character '*' "+
                        "can be used to redirect all outbound traffic. An empty "+
                        "list will disable all outbound redirection.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficIncludeOutboundPorts = Instance {
		Name:          "traffic.sidecar.khulnasoft.com/includeOutboundPorts",
		Description:   "A comma separated list of outbound ports for which "+
                        "traffic is to be redirected to Envoy, regardless of the "+
                        "destination IP.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficKubevirtInterfaces = Instance {
		Name:          "traffic.sidecar.khulnasoft.com/kubevirtInterfaces",
		Description:   "A comma separated list of virtual interfaces whose "+
                        "inbound traffic (from VM) will be treated as outbound.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

)

func AllResourceAnnotations() []*Instance {
	return []*Instance {
		&AlphaCanonicalServiceAccounts,
		&AlphaIdentity,
		&AlphaKubernetesServiceAccounts,
		&GalleyAnalyzeSuppress,
		&GatewayControllerVersion,
		&InjectTemplates,
		&OperatorInstallChartOwner,
		&OperatorInstallOwnerGeneration,
		&OperatorInstallVersion,
		&IoKhulnasoftAutoRegistrationGroup,
		&IoKhulnasoftConnectedAt,
		&IoKhulnasoftDisconnectedAt,
		&IoKhulnasoftDryRun,
		&IoKhulnasoftRev,
		&IoKhulnasoftWorkloadController,
		&IoKubernetesIngressClass,
		&NetworkingExportTo,
		&PrometheusMergeMetrics,
		&ProxyConfig,
		&ProxyOverrides,
		&SidecarStatusReadinessApplicationPorts,
		&SidecarStatusReadinessFailureThreshold,
		&SidecarStatusReadinessInitialDelaySeconds,
		&SidecarStatusReadinessPeriodSeconds,
		&SidecarAgentLogLevel,
		&SidecarBootstrapOverride,
		&SidecarComponentLogLevel,
		&SidecarControlPlaneAuthPolicy,
		&SidecarDiscoveryAddress,
		&SidecarEnableCoreDump,
		&SidecarExtraStatTags,
		&SidecarInject,
		&SidecarInterceptionMode,
		&SidecarLogLevel,
		&SidecarProxyCPU,
		&SidecarProxyCPULimit,
		&SidecarProxyImage,
		&SidecarProxyImageType,
		&SidecarProxyMemory,
		&SidecarProxyMemoryLimit,
		&SidecarRewriteAppHTTPProbers,
		&SidecarStatsHistogramBuckets,
		&SidecarStatsInclusionPrefixes,
		&SidecarStatsInclusionRegexps,
		&SidecarStatsInclusionSuffixes,
		&SidecarStatus,
		&SidecarUserVolume,
		&SidecarUserVolumeMount,
		&SidecarStatusPort,
		&TopologyControlPlaneClusters,
		&TrafficNodeSelector,
		&SidecarTrafficExcludeInboundPorts,
		&SidecarTrafficExcludeInterfaces,
		&SidecarTrafficExcludeOutboundIPRanges,
		&SidecarTrafficExcludeOutboundPorts,
		&SidecarTrafficIncludeInboundPorts,
		&SidecarTrafficIncludeOutboundIPRanges,
		&SidecarTrafficIncludeOutboundPorts,
		&SidecarTrafficKubevirtInterfaces,
	}
}

func AllResourceTypes() []string {
	return []string {
		"Any",
		"AuthorizationPolicy",
		"Ingress",
		"Namespace",
		"Pod",
		"Service",
		"WorkloadEntry",
	}
}
