//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AksNpVirtualNodeGroup) DeepCopyInto(out *AksNpVirtualNodeGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AksNpVirtualNodeGroup.
func (in *AksNpVirtualNodeGroup) DeepCopy() *AksNpVirtualNodeGroup {
	if in == nil {
		return nil
	}
	out := new(AksNpVirtualNodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AksNpVirtualNodeGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AksNpVirtualNodeGroupList) DeepCopyInto(out *AksNpVirtualNodeGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AksNpVirtualNodeGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AksNpVirtualNodeGroupList.
func (in *AksNpVirtualNodeGroupList) DeepCopy() *AksNpVirtualNodeGroupList {
	if in == nil {
		return nil
	}
	out := new(AksNpVirtualNodeGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AksNpVirtualNodeGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AksNpVirtualNodeGroupObservation) DeepCopyInto(out *AksNpVirtualNodeGroupObservation) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EnableNodePublicIP != nil {
		in, out := &in.EnableNodePublicIP, &out.EnableNodePublicIP
		*out = new(bool)
		**out = **in
	}
	if in.FallbackToOndemand != nil {
		in, out := &in.FallbackToOndemand, &out.FallbackToOndemand
		*out = new(bool)
		**out = **in
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]FiltersObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Headrooms != nil {
		in, out := &in.Headrooms, &out.Headrooms
		*out = make([]HeadroomsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.MaxCount != nil {
		in, out := &in.MaxCount, &out.MaxCount
		*out = new(float64)
		**out = **in
	}
	if in.MaxPodsPerNode != nil {
		in, out := &in.MaxPodsPerNode, &out.MaxPodsPerNode
		*out = new(float64)
		**out = **in
	}
	if in.MinCount != nil {
		in, out := &in.MinCount, &out.MinCount
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OceanID != nil {
		in, out := &in.OceanID, &out.OceanID
		*out = new(string)
		**out = **in
	}
	if in.OsDiskSizeGb != nil {
		in, out := &in.OsDiskSizeGb, &out.OsDiskSizeGb
		*out = new(float64)
		**out = **in
	}
	if in.OsDiskType != nil {
		in, out := &in.OsDiskType, &out.OsDiskType
		*out = new(string)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(string)
		**out = **in
	}
	if in.SpotPercentage != nil {
		in, out := &in.SpotPercentage, &out.SpotPercentage
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]TaintsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AksNpVirtualNodeGroupObservation.
func (in *AksNpVirtualNodeGroupObservation) DeepCopy() *AksNpVirtualNodeGroupObservation {
	if in == nil {
		return nil
	}
	out := new(AksNpVirtualNodeGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AksNpVirtualNodeGroupParameters) DeepCopyInto(out *AksNpVirtualNodeGroupParameters) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EnableNodePublicIP != nil {
		in, out := &in.EnableNodePublicIP, &out.EnableNodePublicIP
		*out = new(bool)
		**out = **in
	}
	if in.FallbackToOndemand != nil {
		in, out := &in.FallbackToOndemand, &out.FallbackToOndemand
		*out = new(bool)
		**out = **in
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]FiltersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Headrooms != nil {
		in, out := &in.Headrooms, &out.Headrooms
		*out = make([]HeadroomsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.MaxCount != nil {
		in, out := &in.MaxCount, &out.MaxCount
		*out = new(float64)
		**out = **in
	}
	if in.MaxPodsPerNode != nil {
		in, out := &in.MaxPodsPerNode, &out.MaxPodsPerNode
		*out = new(float64)
		**out = **in
	}
	if in.MinCount != nil {
		in, out := &in.MinCount, &out.MinCount
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OceanID != nil {
		in, out := &in.OceanID, &out.OceanID
		*out = new(string)
		**out = **in
	}
	if in.OsDiskSizeGb != nil {
		in, out := &in.OsDiskSizeGb, &out.OsDiskSizeGb
		*out = new(float64)
		**out = **in
	}
	if in.OsDiskType != nil {
		in, out := &in.OsDiskType, &out.OsDiskType
		*out = new(string)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(string)
		**out = **in
	}
	if in.SpotPercentage != nil {
		in, out := &in.SpotPercentage, &out.SpotPercentage
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]TaintsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AksNpVirtualNodeGroupParameters.
func (in *AksNpVirtualNodeGroupParameters) DeepCopy() *AksNpVirtualNodeGroupParameters {
	if in == nil {
		return nil
	}
	out := new(AksNpVirtualNodeGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AksNpVirtualNodeGroupSpec) DeepCopyInto(out *AksNpVirtualNodeGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AksNpVirtualNodeGroupSpec.
func (in *AksNpVirtualNodeGroupSpec) DeepCopy() *AksNpVirtualNodeGroupSpec {
	if in == nil {
		return nil
	}
	out := new(AksNpVirtualNodeGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AksNpVirtualNodeGroupStatus) DeepCopyInto(out *AksNpVirtualNodeGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AksNpVirtualNodeGroupStatus.
func (in *AksNpVirtualNodeGroupStatus) DeepCopy() *AksNpVirtualNodeGroupStatus {
	if in == nil {
		return nil
	}
	out := new(AksNpVirtualNodeGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FiltersObservation) DeepCopyInto(out *FiltersObservation) {
	*out = *in
	if in.Architectures != nil {
		in, out := &in.Architectures, &out.Architectures
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MaxMemoryGib != nil {
		in, out := &in.MaxMemoryGib, &out.MaxMemoryGib
		*out = new(float64)
		**out = **in
	}
	if in.MaxVcpu != nil {
		in, out := &in.MaxVcpu, &out.MaxVcpu
		*out = new(float64)
		**out = **in
	}
	if in.MinMemoryGib != nil {
		in, out := &in.MinMemoryGib, &out.MinMemoryGib
		*out = new(float64)
		**out = **in
	}
	if in.MinVcpu != nil {
		in, out := &in.MinVcpu, &out.MinVcpu
		*out = new(float64)
		**out = **in
	}
	if in.Series != nil {
		in, out := &in.Series, &out.Series
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FiltersObservation.
func (in *FiltersObservation) DeepCopy() *FiltersObservation {
	if in == nil {
		return nil
	}
	out := new(FiltersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FiltersParameters) DeepCopyInto(out *FiltersParameters) {
	*out = *in
	if in.Architectures != nil {
		in, out := &in.Architectures, &out.Architectures
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MaxMemoryGib != nil {
		in, out := &in.MaxMemoryGib, &out.MaxMemoryGib
		*out = new(float64)
		**out = **in
	}
	if in.MaxVcpu != nil {
		in, out := &in.MaxVcpu, &out.MaxVcpu
		*out = new(float64)
		**out = **in
	}
	if in.MinMemoryGib != nil {
		in, out := &in.MinMemoryGib, &out.MinMemoryGib
		*out = new(float64)
		**out = **in
	}
	if in.MinVcpu != nil {
		in, out := &in.MinVcpu, &out.MinVcpu
		*out = new(float64)
		**out = **in
	}
	if in.Series != nil {
		in, out := &in.Series, &out.Series
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FiltersParameters.
func (in *FiltersParameters) DeepCopy() *FiltersParameters {
	if in == nil {
		return nil
	}
	out := new(FiltersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeadroomsObservation) DeepCopyInto(out *HeadroomsObservation) {
	*out = *in
	if in.CPUPerUnit != nil {
		in, out := &in.CPUPerUnit, &out.CPUPerUnit
		*out = new(float64)
		**out = **in
	}
	if in.GpuPerUnit != nil {
		in, out := &in.GpuPerUnit, &out.GpuPerUnit
		*out = new(float64)
		**out = **in
	}
	if in.MemoryPerUnit != nil {
		in, out := &in.MemoryPerUnit, &out.MemoryPerUnit
		*out = new(float64)
		**out = **in
	}
	if in.NumOfUnits != nil {
		in, out := &in.NumOfUnits, &out.NumOfUnits
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeadroomsObservation.
func (in *HeadroomsObservation) DeepCopy() *HeadroomsObservation {
	if in == nil {
		return nil
	}
	out := new(HeadroomsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeadroomsParameters) DeepCopyInto(out *HeadroomsParameters) {
	*out = *in
	if in.CPUPerUnit != nil {
		in, out := &in.CPUPerUnit, &out.CPUPerUnit
		*out = new(float64)
		**out = **in
	}
	if in.GpuPerUnit != nil {
		in, out := &in.GpuPerUnit, &out.GpuPerUnit
		*out = new(float64)
		**out = **in
	}
	if in.MemoryPerUnit != nil {
		in, out := &in.MemoryPerUnit, &out.MemoryPerUnit
		*out = new(float64)
		**out = **in
	}
	if in.NumOfUnits != nil {
		in, out := &in.NumOfUnits, &out.NumOfUnits
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeadroomsParameters.
func (in *HeadroomsParameters) DeepCopy() *HeadroomsParameters {
	if in == nil {
		return nil
	}
	out := new(HeadroomsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaintsObservation) DeepCopyInto(out *TaintsObservation) {
	*out = *in
	if in.Effect != nil {
		in, out := &in.Effect, &out.Effect
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaintsObservation.
func (in *TaintsObservation) DeepCopy() *TaintsObservation {
	if in == nil {
		return nil
	}
	out := new(TaintsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaintsParameters) DeepCopyInto(out *TaintsParameters) {
	*out = *in
	if in.Effect != nil {
		in, out := &in.Effect, &out.Effect
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaintsParameters.
func (in *TaintsParameters) DeepCopy() *TaintsParameters {
	if in == nil {
		return nil
	}
	out := new(TaintsParameters)
	in.DeepCopyInto(out)
	return out
}