
// GENERATED FILE -- DO NOT EDIT

package label

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
    Namespace
    Node
    Pod
    Service
)

func (r ResourceTypes) String() string {
	switch r {
	case 1:
		return "Any"
	case 2:
		return "Namespace"
	case 3:
		return "Node"
	case 4:
		return "Pod"
	case 5:
		return "Service"
	}
	return "Unknown"
}

// Instance describes a single resource label
type Instance struct {
	// The name of the label.
	Name string

	// Description of the label.
	Description string

	// FeatureStatus of this label.
	FeatureStatus FeatureStatus

	// Hide the existence of this label when outputting usage information.
	Hidden bool

	// Mark this label as deprecated when generating usage information.
	Deprecated bool

	// The types of resources this label applies to.
	Resources []ResourceTypes
}

var (

	IoKhulnasoftRev = Instance {
		Name:          "khulnasoft.com/rev",
		Description:   "Khulnasoft control plane revision associated with the "+
                        "resource; e.g. `canary`",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Namespace,
		},
	}

	NetworkingGatewayPort = Instance {
		Name:          "networking.khulnasoft.com/gatewayPort",
		Description:   "KhulnasoftGatewayPortLabel overrides the default 15443 value "+
                        "to use for a multi-network gateway's port",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Service,
		},
	}

	OperatorComponent = Instance {
		Name:          "operator.khulnasoft.com/component",
		Description:   "Khulnasoft operator component name of the resource, e.g. "+
                        "`Pilot`",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	OperatorManaged = Instance {
		Name:          "operator.khulnasoft.com/managed",
		Description:   "Set to `Reconcile` if the Khulnasoft operator will reconcile "+
                        "the resource.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	OperatorVersion = Instance {
		Name:          "operator.khulnasoft.com/version",
		Description:   "The Khulnasoft operator version that installed the resource, "+
                        "e.g. `1.6.0`",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	SecurityTlsMode = Instance {
		Name:          "security.khulnasoft.com/tlsMode",
		Description:   "Specifies the TLS mode supported by a sidecar proxy. "+
                        "Valid values are 'khulnasoft', 'disabled'. When injecting "+
                        "sidecars into Pods, the sidecar injector will set the "+
                        "value of this label to 'khulnasoft' indicating that the "+
                        "sidecar is capable of supporting mTLS. Clients injected "+
                        "with sidecar proxies will opportunistically use this "+
                        "label to determine whether or not to secure the traffic "+
                        "to this workload using Khulnasoft mutual TLS.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	ServiceCanonicalName = Instance {
		Name:          "service.khulnasoft.com/canonical-name",
		Description:   "The name of the canonical service a workload belongs to",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	ServiceCanonicalRevision = Instance {
		Name:          "service.khulnasoft.com/canonical-revision",
		Description:   "The name of a revision within a canonical service that "+
                        "the workload belongs to",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarInject = Instance {
		Name:          "sidecar.khulnasoft.com/inject",
		Description:   "Specifies whether or not an Envoy sidecar should be "+
                        "automatically injected into the workload.",
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	TopologyCluster = Instance {
		Name:          "topology.khulnasoft.com/cluster",
		Description:   "This label is applied to a workload internally that "+
                        "identifies the Kubernetes cluster containing the "+
                        "workload. The cluster ID is specified during Khulnasoft "+
                        "installation for each cluster via "+
                        "`values.global.multiCluster.clusterName`. It should be "+
                        "noted that this is only used internally within Khulnasoft and "+
                        "is not an actual label on workload pods. If a pod "+
                        "contains this label, it will be overridden by Khulnasoft "+
                        "internally with the cluster ID specified during Khulnasoft "+
                        "installation. This label provides a way to select "+
                        "workloads by cluster when using DestinationRules. For "+
                        "example, a service owner could create a DestinationRule "+
                        "containing a subset per cluster and then use these "+
                        "subsets to control traffic flow to each cluster "+
                        "independently.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	TopologyNetwork = Instance {
		Name:          "topology.khulnasoft.com/network",
		Description:   `A label used to identify the network for one or more pods. This is used
internally by Khulnasoft to group pods resident in the same L3 domain/network.
Khulnasoft assumes that pods in the same network are directly reachable from
one another. When pods are in different networks, an Khulnasoft Gateway
(e.g. east-west gateway) is typically used to establish connectivity
(with AUTO_PASSTHROUGH mode). This label can be applied to the following
resources to help automate Khulnasoft's multi-network configuration.

* Khulnasoft System Namespace: Applying this label to the system namespace
  establishes a default network for pods managed by the control plane.
  This is typically configured during control plane installation using an
  admin-specified value.

* Pod: Applying this label to a pod allows overriding the default network
  on a per-pod basis. This is typically applied to the pod via webhook
  injection, but can also be manually specified on the pod by the service
  owner. The Khulnasoft installation in each cluster configures webhook injection
  using an admin-specified value.

* Gateway Service: Applying this label to the Service for an Khulnasoft Gateway,
  indicates that Khulnasoft should use this service as the gateway for the
  network, when configuring cross-network traffic. Khulnasoft will configure
  pods residing outside of the network to access the Gateway service
  via "spec.externalIPs", "status.loadBalancer.ingress[].ip", or in the case
  of a NodePort service, the Node's address. The label is configured when
  installing the gateway (e.g. east-west gateway) and should match either
  the default network for the control plane (as specified by the Khulnasoft System
  Namespace label) or the network of the targeted pods.`,
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Namespace,
			Pod,
			Service,
		},
	}

	TopologySubzone = Instance {
		Name:          "topology.khulnasoft.com/subzone",
		Description:   "User-provided node label for identifying the locality "+
                        "subzone of a workload. This allows admins to specify a "+
                        "more granular level of locality than what is offered by "+
                        "default with Kubernetes regions and zones.",
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Node,
		},
	}

)

func AllResourceLabels() []*Instance {
	return []*Instance {
		&IoKhulnasoftRev,
		&NetworkingGatewayPort,
		&OperatorComponent,
		&OperatorManaged,
		&OperatorVersion,
		&SecurityTlsMode,
		&ServiceCanonicalName,
		&ServiceCanonicalRevision,
		&SidecarInject,
		&TopologyCluster,
		&TopologyNetwork,
		&TopologySubzone,
	}
}

func AllResourceTypes() []string {
	return []string {
		"Any",
		"Namespace",
		"Node",
		"Pod",
		"Service",
	}
}
