// Code generated by codegen; DO NOT EDIT.

package core

import (
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Pods() *schema.Table {
	return &schema.Table{
		Name:      "k8s_core_pods",
		Resolver:  fetchCorePods,
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
				Name:     "status_host_ip",
				Type:     schema.TypeInet,
				Resolver: client.StringToInetPathResolver("Status.HostIP"),
			},
			{
				Name:     "status_pod_ip",
				Type:     schema.TypeInet,
				Resolver: client.StringToInetPathResolver("Status.PodIP"),
			},
			{
				Name:     "status_pod_ips",
				Type:     schema.TypeInetArray,
				Resolver: resolveCorePodPodIPs,
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
				Name:     "spec_volumes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Volumes"),
			},
			{
				Name:     "spec_init_containers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.InitContainers"),
			},
			{
				Name:     "spec_containers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Containers"),
			},
			{
				Name:     "spec_ephemeral_containers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.EphemeralContainers"),
			},
			{
				Name:     "spec_restart_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.RestartPolicy"),
			},
			{
				Name:     "spec_termination_grace_period_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.TerminationGracePeriodSeconds"),
			},
			{
				Name:     "spec_active_deadline_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.ActiveDeadlineSeconds"),
			},
			{
				Name:     "spec_dns_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.DNSPolicy"),
			},
			{
				Name:     "spec_node_selector",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.NodeSelector"),
			},
			{
				Name:     "spec_service_account_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.ServiceAccountName"),
			},
			{
				Name:     "spec_automount_service_account_token",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.AutomountServiceAccountToken"),
			},
			{
				Name:     "spec_node_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.NodeName"),
			},
			{
				Name:     "spec_host_network",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.HostNetwork"),
			},
			{
				Name:     "spec_host_pid",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.HostPID"),
			},
			{
				Name:     "spec_host_ipc",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.HostIPC"),
			},
			{
				Name:     "spec_share_process_namespace",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.ShareProcessNamespace"),
			},
			{
				Name:     "spec_security_context",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.SecurityContext"),
			},
			{
				Name:     "spec_image_pull_secrets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.ImagePullSecrets"),
			},
			{
				Name:     "spec_hostname",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Hostname"),
			},
			{
				Name:     "spec_subdomain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Subdomain"),
			},
			{
				Name:     "spec_affinity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Affinity"),
			},
			{
				Name:     "spec_scheduler_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.SchedulerName"),
			},
			{
				Name:     "spec_tolerations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Tolerations"),
			},
			{
				Name:     "spec_host_aliases",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.HostAliases"),
			},
			{
				Name:     "spec_priority_class_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.PriorityClassName"),
			},
			{
				Name:     "spec_priority",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.Priority"),
			},
			{
				Name:     "spec_dns_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.DNSConfig"),
			},
			{
				Name:     "spec_readiness_gates",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.ReadinessGates"),
			},
			{
				Name:     "spec_runtime_class_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.RuntimeClassName"),
			},
			{
				Name:     "spec_enable_service_links",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.EnableServiceLinks"),
			},
			{
				Name:     "spec_preemption_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.PreemptionPolicy"),
			},
			{
				Name:     "spec_overhead",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Overhead"),
			},
			{
				Name:     "spec_topology_spread_constraints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.TopologySpreadConstraints"),
			},
			{
				Name:     "spec_set_hostname_as_fqdn",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.SetHostnameAsFQDN"),
			},
			{
				Name:     "spec_os",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.OS"),
			},
			{
				Name:     "spec_host_users",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.HostUsers"),
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
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.Message"),
			},
			{
				Name:     "status_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.Reason"),
			},
			{
				Name:     "status_nominated_node_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.NominatedNodeName"),
			},
			{
				Name:     "status_start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Status.StartTime"),
			},
			{
				Name:     "status_init_container_statuses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.InitContainerStatuses"),
			},
			{
				Name:     "status_container_statuses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.ContainerStatuses"),
			},
			{
				Name:     "status_qos_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.QOSClass"),
			},
			{
				Name:     "status_ephemeral_container_statuses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.EphemeralContainerStatuses"),
			},
		},
	}
}
