// Code generated by codegen; DO NOT EDIT.

package core

import (
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Nodes() *schema.Table {
	return &schema.Table{
		Name:      "k8s_core_nodes",
		Resolver:  fetchCoreNodes,
		Multiplex: client.ContextMultiplex,
		Columns: []schema.Column{
			{
				Name:     "context",
				Type:     schema.TypeString,
				Resolver: client.ResolveContext,
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "spec_pod_cidr",
				Type:     schema.TypeCIDR,
				Resolver: client.StringToCidrPathResolver("Spec.PodCIDR"),
			},
			{
				Name:     "spec_pod_cidrs",
				Type:     schema.TypeCIDRArray,
				Resolver: schema.PathResolver("Spec.PodCIDRs"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "api_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("APIVersion"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Namespace"),
			},
			{
				Name:     "resource_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceVersion"),
			},
			{
				Name:     "generation",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Generation"),
			},
			{
				Name:     "deletion_grace_period_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DeletionGracePeriodSeconds"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "annotations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Annotations"),
			},
			{
				Name:     "owner_references",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OwnerReferences"),
			},
			{
				Name:     "finalizers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Finalizers"),
			},
			{
				Name:     "spec_provider_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.ProviderID"),
			},
			{
				Name:     "spec_unschedulable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.Unschedulable"),
			},
			{
				Name:     "spec_taints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Taints"),
			},
			{
				Name:     "spec_config_source",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.ConfigSource"),
			},
			{
				Name:     "status_capacity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.Capacity"),
			},
			{
				Name:     "status_allocatable",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.Allocatable"),
			},
			{
				Name:     "status_phase",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.Phase"),
			},
			{
				Name:     "status_conditions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.Conditions"),
			},
			{
				Name:     "status_addresses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.Addresses"),
			},
			{
				Name:     "status_daemon_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.DaemonEndpoints"),
			},
			{
				Name:     "status_node_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.NodeInfo"),
			},
			{
				Name:     "status_images",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.Images"),
			},
			{
				Name:     "status_volumes_in_use",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Status.VolumesInUse"),
			},
			{
				Name:     "status_volumes_attached",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.VolumesAttached"),
			},
			{
				Name:     "status_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.Config"),
			},
		},
	}
}
