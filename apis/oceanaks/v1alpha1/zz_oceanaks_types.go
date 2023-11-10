/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AutomaticObservation struct {

	// Optionally set a number between 0-100 to control the percentage of total cluster resources dedicated to headroom.
	Percentage *float64 `json:"percentage,omitempty" tf:"percentage,omitempty"`
}

type AutomaticParameters struct {

	// Optionally set a number between 0-100 to control the percentage of total cluster resources dedicated to headroom.
	// +kubebuilder:validation:Optional
	Percentage *float64 `json:"percentage,omitempty" tf:"percentage,omitempty"`
}

type AutoscaleDownObservation struct {

	// The maximum percentage allowed to scale down in a single scaling action.
	MaxScaleDownPercentage *float64 `json:"maxScaleDownPercentage,omitempty" tf:"max_scale_down_percentage,omitempty"`
}

type AutoscaleDownParameters struct {

	// The maximum percentage allowed to scale down in a single scaling action.
	// +kubebuilder:validation:Optional
	MaxScaleDownPercentage *float64 `json:"maxScaleDownPercentage,omitempty" tf:"max_scale_down_percentage,omitempty"`
}

type AutoscaleHeadroomObservation struct {

	// Automatic headroom configuration.
	Automatic []AutomaticObservation `json:"automatic,omitempty" tf:"automatic,omitempty"`
}

type AutoscaleHeadroomParameters struct {

	// Automatic headroom configuration.
	// +kubebuilder:validation:Optional
	Automatic []AutomaticParameters `json:"automatic,omitempty" tf:"automatic,omitempty"`
}

type AutoscalerObservation struct {

	// Auto Scaling scale down operations.
	AutoscaleDown []AutoscaleDownObservation `json:"autoscaleDown,omitempty" tf:"autoscale_down,omitempty"`

	// Spare resource capacity management enabling fast assignment of pods without waiting for new resources to launch.
	AutoscaleHeadroom []AutoscaleHeadroomObservation `json:"autoscaleHeadroom,omitempty" tf:"autoscale_headroom,omitempty"`

	// Enable the Ocean Kubernetes Autoscaler.
	AutoscaleIsEnabled *bool `json:"autoscaleIsEnabled,omitempty" tf:"autoscale_is_enabled,omitempty"`

	// Optionally set upper and lower bounds on the resource usage of the cluster.
	ResourceLimits []ResourceLimitsObservation `json:"resourceLimits,omitempty" tf:"resource_limits,omitempty"`
}

type AutoscalerParameters struct {

	// Auto Scaling scale down operations.
	// +kubebuilder:validation:Optional
	AutoscaleDown []AutoscaleDownParameters `json:"autoscaleDown,omitempty" tf:"autoscale_down,omitempty"`

	// Spare resource capacity management enabling fast assignment of pods without waiting for new resources to launch.
	// +kubebuilder:validation:Optional
	AutoscaleHeadroom []AutoscaleHeadroomParameters `json:"autoscaleHeadroom,omitempty" tf:"autoscale_headroom,omitempty"`

	// Enable the Ocean Kubernetes Autoscaler.
	// +kubebuilder:validation:Optional
	AutoscaleIsEnabled *bool `json:"autoscaleIsEnabled,omitempty" tf:"autoscale_is_enabled,omitempty"`

	// Optionally set upper and lower bounds on the resource usage of the cluster.
	// +kubebuilder:validation:Optional
	ResourceLimits []ResourceLimitsParameters `json:"resourceLimits,omitempty" tf:"resource_limits,omitempty"`
}

type FiltersObservation struct {

	// In case acceleratedNetworking is set to Enabled, accelerated networking applies only to the VM that enables it.
	AcceleratedNetworking *string `json:"acceleratedNetworking,omitempty" tf:"accelerated_networking,omitempty"`

	// The filtered vm sizes will support at least one of the architectures from this list. x86_64 includes both intel64 and amd64.
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// The filtered vm sizes will support at least one of the classes from this list.
	DiskPerformance *string `json:"diskPerformance,omitempty" tf:"disk_performance,omitempty"`

	// Vm sizes belonging to a series from the list will not be available for scaling
	ExcludeSeries []*string `json:"excludeSeries,omitempty" tf:"exclude_series,omitempty"`

	// Maximum number of GPUs available.
	MaxGpu *float64 `json:"maxGpu,omitempty" tf:"max_gpu,omitempty"`

	// The maximum memory in GiB units that can be allocated to the cluster.
	MaxMemoryGib *float64 `json:"maxMemoryGib,omitempty" tf:"max_memory_gib,omitempty"`

	// The maximum cpu in vCpu units that can be allocated to the cluster.
	MaxVcpu *float64 `json:"maxVcpu,omitempty" tf:"max_vcpu,omitempty"`

	// Minimum number of data disks available.
	MinData *float64 `json:"minData,omitempty" tf:"min_data,omitempty"`

	// Minimum number of GPUs available.
	MinGpu *float64 `json:"minGpu,omitempty" tf:"min_gpu,omitempty"`

	// Minimum amount of Memory (GiB).
	MinMemoryGib *float64 `json:"minMemoryGib,omitempty" tf:"min_memory_gib,omitempty"`

	// Minimum number of network interfaces.
	MinNics *float64 `json:"minNics,omitempty" tf:"min_nics,omitempty"`

	// Minimum number of vcpus available.
	MinVcpu *float64 `json:"minVcpu,omitempty" tf:"min_vcpu,omitempty"`

	// Vm sizes belonging to a series from the list will be available for scaling. We can specify include list and series can be specified with capital or small letters, with space, without space or with underscore '_' .  For example all of these "DSv2", "Ds v2", "ds_v2" refer to same DS_v2 series.
	Series []*string `json:"series,omitempty" tf:"series,omitempty"`

	// The filtered vm types will belong to one of the vm types from this list.
	VMTypes []*string `json:"vmTypes,omitempty" tf:"vm_types,omitempty"`
}

type FiltersParameters struct {

	// In case acceleratedNetworking is set to Enabled, accelerated networking applies only to the VM that enables it.
	// +kubebuilder:validation:Optional
	AcceleratedNetworking *string `json:"acceleratedNetworking,omitempty" tf:"accelerated_networking,omitempty"`

	// The filtered vm sizes will support at least one of the architectures from this list. x86_64 includes both intel64 and amd64.
	// +kubebuilder:validation:Optional
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// The filtered vm sizes will support at least one of the classes from this list.
	// +kubebuilder:validation:Optional
	DiskPerformance *string `json:"diskPerformance,omitempty" tf:"disk_performance,omitempty"`

	// Vm sizes belonging to a series from the list will not be available for scaling
	// +kubebuilder:validation:Optional
	ExcludeSeries []*string `json:"excludeSeries,omitempty" tf:"exclude_series,omitempty"`

	// Maximum number of GPUs available.
	// +kubebuilder:validation:Optional
	MaxGpu *float64 `json:"maxGpu,omitempty" tf:"max_gpu,omitempty"`

	// The maximum memory in GiB units that can be allocated to the cluster.
	// +kubebuilder:validation:Optional
	MaxMemoryGib *float64 `json:"maxMemoryGib,omitempty" tf:"max_memory_gib,omitempty"`

	// The maximum cpu in vCpu units that can be allocated to the cluster.
	// +kubebuilder:validation:Optional
	MaxVcpu *float64 `json:"maxVcpu,omitempty" tf:"max_vcpu,omitempty"`

	// Minimum number of data disks available.
	// +kubebuilder:validation:Optional
	MinData *float64 `json:"minData,omitempty" tf:"min_data,omitempty"`

	// Minimum number of GPUs available.
	// +kubebuilder:validation:Optional
	MinGpu *float64 `json:"minGpu,omitempty" tf:"min_gpu,omitempty"`

	// Minimum amount of Memory (GiB).
	// +kubebuilder:validation:Optional
	MinMemoryGib *float64 `json:"minMemoryGib,omitempty" tf:"min_memory_gib,omitempty"`

	// Minimum number of network interfaces.
	// +kubebuilder:validation:Optional
	MinNics *float64 `json:"minNics,omitempty" tf:"min_nics,omitempty"`

	// Minimum number of vcpus available.
	// +kubebuilder:validation:Optional
	MinVcpu *float64 `json:"minVcpu,omitempty" tf:"min_vcpu,omitempty"`

	// Vm sizes belonging to a series from the list will be available for scaling. We can specify include list and series can be specified with capital or small letters, with space, without space or with underscore '_' .  For example all of these "DSv2", "Ds v2", "ds_v2" refer to same DS_v2 series.
	// +kubebuilder:validation:Optional
	Series []*string `json:"series,omitempty" tf:"series,omitempty"`

	// The filtered vm types will belong to one of the vm types from this list.
	// +kubebuilder:validation:Optional
	VMTypes []*string `json:"vmTypes,omitempty" tf:"vm_types,omitempty"`
}

type HeadroomsObservation struct {

	// Configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
	CPUPerUnit *float64 `json:"cpuPerUnit,omitempty" tf:"cpu_per_unit,omitempty"`

	// Amount of GPU to allocate for headroom unit.
	GpuPerUnit *float64 `json:"gpuPerUnit,omitempty" tf:"gpu_per_unit,omitempty"`

	// Configure the amount of memory (MiB) to allocate the headroom.
	MemoryPerUnit *float64 `json:"memoryPerUnit,omitempty" tf:"memory_per_unit,omitempty"`

	// The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
	NumOfUnits *float64 `json:"numOfUnits,omitempty" tf:"num_of_units,omitempty"`
}

type HeadroomsParameters struct {

	// Configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
	// +kubebuilder:validation:Optional
	CPUPerUnit *float64 `json:"cpuPerUnit,omitempty" tf:"cpu_per_unit,omitempty"`

	// Amount of GPU to allocate for headroom unit.
	// +kubebuilder:validation:Optional
	GpuPerUnit *float64 `json:"gpuPerUnit,omitempty" tf:"gpu_per_unit,omitempty"`

	// Configure the amount of memory (MiB) to allocate the headroom.
	// +kubebuilder:validation:Optional
	MemoryPerUnit *float64 `json:"memoryPerUnit,omitempty" tf:"memory_per_unit,omitempty"`

	// The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
	// +kubebuilder:validation:Optional
	NumOfUnits *float64 `json:"numOfUnits,omitempty" tf:"num_of_units,omitempty"`
}

type HealthObservation struct {

	// The amount of time to wait, in seconds, from the moment the instance has launched until monitoring of its health checks begins.
	GracePeriod *float64 `json:"gracePeriod,omitempty" tf:"grace_period,omitempty"`
}

type HealthParameters struct {

	// The amount of time to wait, in seconds, from the moment the instance has launched until monitoring of its health checks begins.
	// +kubebuilder:validation:Optional
	GracePeriod *float64 `json:"gracePeriod,omitempty" tf:"grace_period,omitempty"`
}

type OceanAksObservation struct {

	// The name of the AKS Cluster.
	AksClusterName *string `json:"aksClusterName,omitempty" tf:"aks_cluster_name,omitempty"`

	// The name of the cluster's infrastructure resource group.
	AksInfrastructureResourceGroupName *string `json:"aksInfrastructureResourceGroupName,omitempty" tf:"aks_infrastructure_resource_group_name,omitempty"`

	// The cluster's region.
	AksRegion *string `json:"aksRegion,omitempty" tf:"aks_region,omitempty"`

	// The name of the cluster's resource group.
	AksResourceGroupName *string `json:"aksResourceGroupName,omitempty" tf:"aks_resource_group_name,omitempty"`

	// The Ocean Kubernetes Autoscaler object.
	Autoscaler []AutoscalerObservation `json:"autoscaler,omitempty" tf:"autoscaler,omitempty"`

	// An Array holding Availability Zones, this configures the availability zones the Ocean may launch instances in per VNG.
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Enter a unique Ocean cluster identifier. Cannot be updated. This needs to match with string that was used to install the controller in the cluster, typically clusterName + 8 digit string.
	ControllerClusterID *string `json:"controllerClusterId,omitempty" tf:"controller_cluster_id,omitempty"`

	// Enable node public IP.
	EnableNodePublicIP *bool `json:"enableNodePublicIp,omitempty" tf:"enable_node_public_ip,omitempty"`

	// If no spot VM markets are available, enable Ocean to launch regular (pay-as-you-go) nodes instead.
	FallbackToOndemand *bool `json:"fallbackToOndemand,omitempty" tf:"fallback_to_ondemand,omitempty"`

	// Filters for the VM sizes that can be launched from the virtual node group.
	Filters []FiltersObservation `json:"filters,omitempty" tf:"filters,omitempty"`

	// Specify the custom headroom per VNG. Provide a list of headroom objects.
	Headrooms []HeadroomsObservation `json:"headrooms,omitempty" tf:"headrooms,omitempty"`

	// The Ocean AKS Health object.
	Health []HealthObservation `json:"health,omitempty" tf:"health,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The desired Kubernetes version of the launched nodes. In case the value is null, the Kubernetes version of the control plane is used.
	KubernetesVersion *string `json:"kubernetesVersion,omitempty" tf:"kubernetes_version,omitempty"`

	// An array of labels to add to the virtual node group. Only custom user labels are allowed, and not Kubernetes well-known labels or  Azure AKS labels or Spot labels.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maximum node count limit.
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// The maximum number of pods per node in the node pools.
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// Minimum node count limit.
	MinCount *float64 `json:"minCount,omitempty" tf:"min_count,omitempty"`

	// Add a name for the Ocean cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The size of the OS disk in GB.
	OsDiskSizeGb *float64 `json:"osDiskSizeGb,omitempty" tf:"os_disk_size_gb,omitempty"`

	// The type of the OS disk.
	OsDiskType *string `json:"osDiskType,omitempty" tf:"os_disk_type,omitempty"`

	// The OS SKU of the OS type. Must correlate with the os type.
	OsSku *string `json:"osSku,omitempty" tf:"os_sku,omitempty"`

	// The OS type of the OS disk. Can't be modified once set.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The IDs of subnets in an existing VNet into which to assign pods in the cluster (requires azure network-plugin).
	PodSubnetIds []*string `json:"podSubnetIds,omitempty" tf:"pod_subnet_ids,omitempty"`

	// An object used to specify times when the cluster will turn off. Once the shutdown time will be over, the cluster will return to its previous state.
	Scheduling []SchedulingObservation `json:"scheduling,omitempty" tf:"scheduling,omitempty"`

	// Percentage of spot VMs to maintain.
	SpotPercentage *float64 `json:"spotPercentage,omitempty" tf:"spot_percentage,omitempty"`

	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Add taints to a virtual node group. Only custom user taints are allowed, and not Kubernetes well-known taints or Azure AKS ScaleSetPrioirty (Spot VM) taint. For all Spot VMs, AKS injects a taint kubernetes.azure.com/scalesetpriority=spot:NoSchedule, to ensure that only workloads that can handle interruptions are scheduled on Spot nodes. To schedule a pod to run on Spot node, add a toleration but dont include the nodeAffinity (not supported for Spot Ocean), this will prevent the pod from being scheduled using Spot Ocean.
	Taints []TaintsObservation `json:"taints,omitempty" tf:"taints,omitempty"`

	// The IDs of subnets in an existing VNet into which to assign nodes in the cluster (requires azure network-plugin).
	VnetSubnetIds []*string `json:"vnetSubnetIds,omitempty" tf:"vnet_subnet_ids,omitempty"`
}

type OceanAksParameters struct {

	// The name of the AKS Cluster.
	// +kubebuilder:validation:Optional
	AksClusterName *string `json:"aksClusterName,omitempty" tf:"aks_cluster_name,omitempty"`

	// The name of the cluster's infrastructure resource group.
	// +kubebuilder:validation:Optional
	AksInfrastructureResourceGroupName *string `json:"aksInfrastructureResourceGroupName,omitempty" tf:"aks_infrastructure_resource_group_name,omitempty"`

	// The cluster's region.
	// +kubebuilder:validation:Optional
	AksRegion *string `json:"aksRegion,omitempty" tf:"aks_region,omitempty"`

	// The name of the cluster's resource group.
	// +kubebuilder:validation:Optional
	AksResourceGroupName *string `json:"aksResourceGroupName,omitempty" tf:"aks_resource_group_name,omitempty"`

	// The Ocean Kubernetes Autoscaler object.
	// +kubebuilder:validation:Optional
	Autoscaler []AutoscalerParameters `json:"autoscaler,omitempty" tf:"autoscaler,omitempty"`

	// An Array holding Availability Zones, this configures the availability zones the Ocean may launch instances in per VNG.
	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Enter a unique Ocean cluster identifier. Cannot be updated. This needs to match with string that was used to install the controller in the cluster, typically clusterName + 8 digit string.
	// +kubebuilder:validation:Optional
	ControllerClusterID *string `json:"controllerClusterId,omitempty" tf:"controller_cluster_id,omitempty"`

	// Enable node public IP.
	// +kubebuilder:validation:Optional
	EnableNodePublicIP *bool `json:"enableNodePublicIp,omitempty" tf:"enable_node_public_ip,omitempty"`

	// If no spot VM markets are available, enable Ocean to launch regular (pay-as-you-go) nodes instead.
	// +kubebuilder:validation:Optional
	FallbackToOndemand *bool `json:"fallbackToOndemand,omitempty" tf:"fallback_to_ondemand,omitempty"`

	// Filters for the VM sizes that can be launched from the virtual node group.
	// +kubebuilder:validation:Optional
	Filters []FiltersParameters `json:"filters,omitempty" tf:"filters,omitempty"`

	// Specify the custom headroom per VNG. Provide a list of headroom objects.
	// +kubebuilder:validation:Optional
	Headrooms []HeadroomsParameters `json:"headrooms,omitempty" tf:"headrooms,omitempty"`

	// The Ocean AKS Health object.
	// +kubebuilder:validation:Optional
	Health []HealthParameters `json:"health,omitempty" tf:"health,omitempty"`

	// The desired Kubernetes version of the launched nodes. In case the value is null, the Kubernetes version of the control plane is used.
	// +kubebuilder:validation:Optional
	KubernetesVersion *string `json:"kubernetesVersion,omitempty" tf:"kubernetes_version,omitempty"`

	// An array of labels to add to the virtual node group. Only custom user labels are allowed, and not Kubernetes well-known labels or  Azure AKS labels or Spot labels.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maximum node count limit.
	// +kubebuilder:validation:Optional
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// The maximum number of pods per node in the node pools.
	// +kubebuilder:validation:Optional
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// Minimum node count limit.
	// +kubebuilder:validation:Optional
	MinCount *float64 `json:"minCount,omitempty" tf:"min_count,omitempty"`

	// Add a name for the Ocean cluster.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The size of the OS disk in GB.
	// +kubebuilder:validation:Optional
	OsDiskSizeGb *float64 `json:"osDiskSizeGb,omitempty" tf:"os_disk_size_gb,omitempty"`

	// The type of the OS disk.
	// +kubebuilder:validation:Optional
	OsDiskType *string `json:"osDiskType,omitempty" tf:"os_disk_type,omitempty"`

	// The OS SKU of the OS type. Must correlate with the os type.
	// +kubebuilder:validation:Optional
	OsSku *string `json:"osSku,omitempty" tf:"os_sku,omitempty"`

	// The OS type of the OS disk. Can't be modified once set.
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The IDs of subnets in an existing VNet into which to assign pods in the cluster (requires azure network-plugin).
	// +kubebuilder:validation:Optional
	PodSubnetIds []*string `json:"podSubnetIds,omitempty" tf:"pod_subnet_ids,omitempty"`

	// An object used to specify times when the cluster will turn off. Once the shutdown time will be over, the cluster will return to its previous state.
	// +kubebuilder:validation:Optional
	Scheduling []SchedulingParameters `json:"scheduling,omitempty" tf:"scheduling,omitempty"`

	// Percentage of spot VMs to maintain.
	// +kubebuilder:validation:Optional
	SpotPercentage *float64 `json:"spotPercentage,omitempty" tf:"spot_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Add taints to a virtual node group. Only custom user taints are allowed, and not Kubernetes well-known taints or Azure AKS ScaleSetPrioirty (Spot VM) taint. For all Spot VMs, AKS injects a taint kubernetes.azure.com/scalesetpriority=spot:NoSchedule, to ensure that only workloads that can handle interruptions are scheduled on Spot nodes. To schedule a pod to run on Spot node, add a toleration but dont include the nodeAffinity (not supported for Spot Ocean), this will prevent the pod from being scheduled using Spot Ocean.
	// +kubebuilder:validation:Optional
	Taints []TaintsParameters `json:"taints,omitempty" tf:"taints,omitempty"`

	// The IDs of subnets in an existing VNet into which to assign nodes in the cluster (requires azure network-plugin).
	// +kubebuilder:validation:Optional
	VnetSubnetIds []*string `json:"vnetSubnetIds,omitempty" tf:"vnet_subnet_ids,omitempty"`
}

type ResourceLimitsObservation struct {

	// The maximum memory in GiB units that can be allocated to the cluster.
	MaxMemoryGib *float64 `json:"maxMemoryGib,omitempty" tf:"max_memory_gib,omitempty"`

	// The maximum cpu in vCpu units that can be allocated to the cluster.
	MaxVcpu *float64 `json:"maxVcpu,omitempty" tf:"max_vcpu,omitempty"`
}

type ResourceLimitsParameters struct {

	// The maximum memory in GiB units that can be allocated to the cluster.
	// +kubebuilder:validation:Optional
	MaxMemoryGib *float64 `json:"maxMemoryGib,omitempty" tf:"max_memory_gib,omitempty"`

	// The maximum cpu in vCpu units that can be allocated to the cluster.
	// +kubebuilder:validation:Optional
	MaxVcpu *float64 `json:"maxVcpu,omitempty" tf:"max_vcpu,omitempty"`
}

type SchedulingObservation struct {

	// Shutdown HoursAn object used to specify times that the nodes in the cluster will be taken down.
	ShutdownHours []ShutdownHoursObservation `json:"shutdownHours,omitempty" tf:"shutdown_hours,omitempty"`
}

type SchedulingParameters struct {

	// Shutdown HoursAn object used to specify times that the nodes in the cluster will be taken down.
	// +kubebuilder:validation:Optional
	ShutdownHours []ShutdownHoursParameters `json:"shutdownHours,omitempty" tf:"shutdown_hours,omitempty"`
}

type ShutdownHoursObservation struct {

	// Flag to enable or disable the shutdown hours mechanism. When False, the mechanism is deactivated, and the cluster remains in its current state.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// The times that the shutdown hours will apply.
	TimeWindows []*string `json:"timeWindows,omitempty" tf:"time_windows,omitempty"`
}

type ShutdownHoursParameters struct {

	// Flag to enable or disable the shutdown hours mechanism. When False, the mechanism is deactivated, and the cluster remains in its current state.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// The times that the shutdown hours will apply.
	// +kubebuilder:validation:Required
	TimeWindows []*string `json:"timeWindows" tf:"time_windows,omitempty"`
}

type TaintsObservation struct {

	// Set taint effect.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// Set label key spot labels and Azure labels. The following are not allowed: ["kubernetes.azure.com/agentpool","kubernetes.io/arch","kubernetes.io/os","node.kubernetes.io/instance-type", "topology.kubernetes.io/region", "topology.kubernetes.io/zone", "kubernetes.azure.com/cluster", "kubernetes.azure.com/mode", "kubernetes.azure.com/role", "kubernetes.azure.com/scalesetpriority", "kubernetes.io/hostname", "kubernetes.azure.com/storageprofile", "kubernetes.azure.com/storagetier", "kubernetes.azure.com/instance-sku", "kubernetes.azure.com/node-image-version", "kubernetes.azure.com/subnet", "kubernetes.azure.com/vnet", "kubernetes.azure.com/ppg", "kubernetes.azure.com/encrypted-set", "kubernetes.azure.com/accelerator", "kubernetes.azure.com/fips_enabled", "kubernetes.azure.com/os-sku"]
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Set label value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TaintsParameters struct {

	// Set taint effect.
	// +kubebuilder:validation:Required
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// Set label key spot labels and Azure labels. The following are not allowed: ["kubernetes.azure.com/agentpool","kubernetes.io/arch","kubernetes.io/os","node.kubernetes.io/instance-type", "topology.kubernetes.io/region", "topology.kubernetes.io/zone", "kubernetes.azure.com/cluster", "kubernetes.azure.com/mode", "kubernetes.azure.com/role", "kubernetes.azure.com/scalesetpriority", "kubernetes.io/hostname", "kubernetes.azure.com/storageprofile", "kubernetes.azure.com/storagetier", "kubernetes.azure.com/instance-sku", "kubernetes.azure.com/node-image-version", "kubernetes.azure.com/subnet", "kubernetes.azure.com/vnet", "kubernetes.azure.com/ppg", "kubernetes.azure.com/encrypted-set", "kubernetes.azure.com/accelerator", "kubernetes.azure.com/fips_enabled", "kubernetes.azure.com/os-sku"]
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Set label value.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// OceanAksSpec defines the desired state of OceanAks
type OceanAksSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OceanAksParameters `json:"forProvider"`
}

// OceanAksStatus defines the observed state of OceanAks.
type OceanAksStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OceanAksObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OceanAks is the Schema for the OceanAkss API. Provides a Spotinst Ocean resource using AKS.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,spot}
type OceanAks struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.aksClusterName)",message="aksClusterName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.aksInfrastructureResourceGroupName)",message="aksInfrastructureResourceGroupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.aksRegion)",message="aksRegion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.aksResourceGroupName)",message="aksResourceGroupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.availabilityZones)",message="availabilityZones is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   OceanAksSpec   `json:"spec"`
	Status OceanAksStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OceanAksList contains a list of OceanAkss
type OceanAksList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OceanAks `json:"items"`
}

// Repository type metadata.
var (
	OceanAks_Kind             = "OceanAks"
	OceanAks_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OceanAks_Kind}.String()
	OceanAks_KindAPIVersion   = OceanAks_Kind + "." + CRDGroupVersion.String()
	OceanAks_GroupVersionKind = CRDGroupVersion.WithKind(OceanAks_Kind)
)

func init() {
	SchemeBuilder.Register(&OceanAks{}, &OceanAksList{})
}