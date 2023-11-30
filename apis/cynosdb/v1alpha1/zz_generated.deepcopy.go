//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
	*out = *in
	if in.Charset != nil {
		in, out := &in.Charset, &out.Charset
		*out = new(string)
		**out = **in
	}
	if in.ClusterStatus != nil {
		in, out := &in.ClusterStatus, &out.ClusterStatus
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceName != nil {
		in, out := &in.InstanceName, &out.InstanceName
		*out = new(string)
		**out = **in
	}
	if in.InstanceStatus != nil {
		in, out := &in.InstanceStatus, &out.InstanceStatus
		*out = new(string)
		**out = **in
	}
	if in.InstanceStorageSize != nil {
		in, out := &in.InstanceStorageSize, &out.InstanceStorageSize
		*out = new(float64)
		**out = **in
	}
	if in.RoGroupAddr != nil {
		in, out := &in.RoGroupAddr, &out.RoGroupAddr
		*out = make([]RoGroupAddrObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RoGroupID != nil {
		in, out := &in.RoGroupID, &out.RoGroupID
		*out = new(string)
		**out = **in
	}
	if in.RoGroupInstances != nil {
		in, out := &in.RoGroupInstances, &out.RoGroupInstances
		*out = make([]RoGroupInstancesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RwGroupAddr != nil {
		in, out := &in.RwGroupAddr, &out.RwGroupAddr
		*out = make([]RwGroupAddrObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RwGroupID != nil {
		in, out := &in.RwGroupID, &out.RwGroupID
		*out = new(string)
		**out = **in
	}
	if in.RwGroupInstances != nil {
		in, out := &in.RwGroupInstances, &out.RwGroupInstances
		*out = make([]RwGroupInstancesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServerlessStatus != nil {
		in, out := &in.ServerlessStatus, &out.ServerlessStatus
		*out = new(string)
		**out = **in
	}
	if in.StorageUsed != nil {
		in, out := &in.StorageUsed, &out.StorageUsed
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
	*out = *in
	if in.AutoPause != nil {
		in, out := &in.AutoPause, &out.AutoPause
		*out = new(string)
		**out = **in
	}
	if in.AutoPauseDelay != nil {
		in, out := &in.AutoPauseDelay, &out.AutoPauseDelay
		*out = new(float64)
		**out = **in
	}
	if in.AutoRenewFlag != nil {
		in, out := &in.AutoRenewFlag, &out.AutoRenewFlag
		*out = new(float64)
		**out = **in
	}
	if in.AvailableZone != nil {
		in, out := &in.AvailableZone, &out.AvailableZone
		*out = new(string)
		**out = **in
	}
	if in.ChargeType != nil {
		in, out := &in.ChargeType, &out.ChargeType
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.DBMode != nil {
		in, out := &in.DBMode, &out.DBMode
		*out = new(string)
		**out = **in
	}
	if in.DBType != nil {
		in, out := &in.DBType, &out.DBType
		*out = new(string)
		**out = **in
	}
	if in.DBVersion != nil {
		in, out := &in.DBVersion, &out.DBVersion
		*out = new(string)
		**out = **in
	}
	if in.ForceDelete != nil {
		in, out := &in.ForceDelete, &out.ForceDelete
		*out = new(bool)
		**out = **in
	}
	if in.InstanceCPUCore != nil {
		in, out := &in.InstanceCPUCore, &out.InstanceCPUCore
		*out = new(float64)
		**out = **in
	}
	if in.InstanceMaintainDuration != nil {
		in, out := &in.InstanceMaintainDuration, &out.InstanceMaintainDuration
		*out = new(float64)
		**out = **in
	}
	if in.InstanceMaintainStartTime != nil {
		in, out := &in.InstanceMaintainStartTime, &out.InstanceMaintainStartTime
		*out = new(float64)
		**out = **in
	}
	if in.InstanceMaintainWeekdays != nil {
		in, out := &in.InstanceMaintainWeekdays, &out.InstanceMaintainWeekdays
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.InstanceMemorySize != nil {
		in, out := &in.InstanceMemorySize, &out.InstanceMemorySize
		*out = new(float64)
		**out = **in
	}
	if in.MaxCPU != nil {
		in, out := &in.MaxCPU, &out.MaxCPU
		*out = new(float64)
		**out = **in
	}
	if in.MinCPU != nil {
		in, out := &in.MinCPU, &out.MinCPU
		*out = new(float64)
		**out = **in
	}
	if in.OldIPReserveHours != nil {
		in, out := &in.OldIPReserveHours, &out.OldIPReserveHours
		*out = new(float64)
		**out = **in
	}
	if in.ParamItems != nil {
		in, out := &in.ParamItems, &out.ParamItems
		*out = make([]ParamItemsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.PrarmTemplateID != nil {
		in, out := &in.PrarmTemplateID, &out.PrarmTemplateID
		*out = new(float64)
		**out = **in
	}
	if in.PrepaidPeriod != nil {
		in, out := &in.PrepaidPeriod, &out.PrepaidPeriod
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(float64)
		**out = **in
	}
	if in.RoGroupSg != nil {
		in, out := &in.RoGroupSg, &out.RoGroupSg
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RwGroupSg != nil {
		in, out := &in.RwGroupSg, &out.RwGroupSg
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ServerlessStatusFlag != nil {
		in, out := &in.ServerlessStatusFlag, &out.ServerlessStatusFlag
		*out = new(string)
		**out = **in
	}
	if in.StorageLimit != nil {
		in, out := &in.StorageLimit, &out.StorageLimit
		*out = new(float64)
		**out = **in
	}
	if in.StoragePayMode != nil {
		in, out := &in.StoragePayMode, &out.StoragePayMode
		*out = new(float64)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
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
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParamItemsObservation) DeepCopyInto(out *ParamItemsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParamItemsObservation.
func (in *ParamItemsObservation) DeepCopy() *ParamItemsObservation {
	if in == nil {
		return nil
	}
	out := new(ParamItemsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParamItemsParameters) DeepCopyInto(out *ParamItemsParameters) {
	*out = *in
	if in.CurrentValue != nil {
		in, out := &in.CurrentValue, &out.CurrentValue
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OldValue != nil {
		in, out := &in.OldValue, &out.OldValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParamItemsParameters.
func (in *ParamItemsParameters) DeepCopy() *ParamItemsParameters {
	if in == nil {
		return nil
	}
	out := new(ParamItemsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadonlyInstance) DeepCopyInto(out *ReadonlyInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadonlyInstance.
func (in *ReadonlyInstance) DeepCopy() *ReadonlyInstance {
	if in == nil {
		return nil
	}
	out := new(ReadonlyInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReadonlyInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadonlyInstanceList) DeepCopyInto(out *ReadonlyInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReadonlyInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadonlyInstanceList.
func (in *ReadonlyInstanceList) DeepCopy() *ReadonlyInstanceList {
	if in == nil {
		return nil
	}
	out := new(ReadonlyInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReadonlyInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadonlyInstanceObservation) DeepCopyInto(out *ReadonlyInstanceObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InstanceStatus != nil {
		in, out := &in.InstanceStatus, &out.InstanceStatus
		*out = new(string)
		**out = **in
	}
	if in.InstanceStorageSize != nil {
		in, out := &in.InstanceStorageSize, &out.InstanceStorageSize
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadonlyInstanceObservation.
func (in *ReadonlyInstanceObservation) DeepCopy() *ReadonlyInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(ReadonlyInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadonlyInstanceParameters) DeepCopyInto(out *ReadonlyInstanceParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterIDRef != nil {
		in, out := &in.ClusterIDRef, &out.ClusterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterIDSelector != nil {
		in, out := &in.ClusterIDSelector, &out.ClusterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ForceDelete != nil {
		in, out := &in.ForceDelete, &out.ForceDelete
		*out = new(bool)
		**out = **in
	}
	if in.InstanceCPUCore != nil {
		in, out := &in.InstanceCPUCore, &out.InstanceCPUCore
		*out = new(float64)
		**out = **in
	}
	if in.InstanceMaintainDuration != nil {
		in, out := &in.InstanceMaintainDuration, &out.InstanceMaintainDuration
		*out = new(float64)
		**out = **in
	}
	if in.InstanceMaintainStartTime != nil {
		in, out := &in.InstanceMaintainStartTime, &out.InstanceMaintainStartTime
		*out = new(float64)
		**out = **in
	}
	if in.InstanceMaintainWeekdays != nil {
		in, out := &in.InstanceMaintainWeekdays, &out.InstanceMaintainWeekdays
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.InstanceMemorySize != nil {
		in, out := &in.InstanceMemorySize, &out.InstanceMemorySize
		*out = new(float64)
		**out = **in
	}
	if in.InstanceName != nil {
		in, out := &in.InstanceName, &out.InstanceName
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadonlyInstanceParameters.
func (in *ReadonlyInstanceParameters) DeepCopy() *ReadonlyInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(ReadonlyInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadonlyInstanceSpec) DeepCopyInto(out *ReadonlyInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadonlyInstanceSpec.
func (in *ReadonlyInstanceSpec) DeepCopy() *ReadonlyInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(ReadonlyInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadonlyInstanceStatus) DeepCopyInto(out *ReadonlyInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadonlyInstanceStatus.
func (in *ReadonlyInstanceStatus) DeepCopy() *ReadonlyInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(ReadonlyInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoGroupAddrObservation) DeepCopyInto(out *RoGroupAddrObservation) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoGroupAddrObservation.
func (in *RoGroupAddrObservation) DeepCopy() *RoGroupAddrObservation {
	if in == nil {
		return nil
	}
	out := new(RoGroupAddrObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoGroupAddrParameters) DeepCopyInto(out *RoGroupAddrParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoGroupAddrParameters.
func (in *RoGroupAddrParameters) DeepCopy() *RoGroupAddrParameters {
	if in == nil {
		return nil
	}
	out := new(RoGroupAddrParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoGroupInstancesObservation) DeepCopyInto(out *RoGroupInstancesObservation) {
	*out = *in
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceName != nil {
		in, out := &in.InstanceName, &out.InstanceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoGroupInstancesObservation.
func (in *RoGroupInstancesObservation) DeepCopy() *RoGroupInstancesObservation {
	if in == nil {
		return nil
	}
	out := new(RoGroupInstancesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoGroupInstancesParameters) DeepCopyInto(out *RoGroupInstancesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoGroupInstancesParameters.
func (in *RoGroupInstancesParameters) DeepCopy() *RoGroupInstancesParameters {
	if in == nil {
		return nil
	}
	out := new(RoGroupInstancesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RwGroupAddrObservation) DeepCopyInto(out *RwGroupAddrObservation) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RwGroupAddrObservation.
func (in *RwGroupAddrObservation) DeepCopy() *RwGroupAddrObservation {
	if in == nil {
		return nil
	}
	out := new(RwGroupAddrObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RwGroupAddrParameters) DeepCopyInto(out *RwGroupAddrParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RwGroupAddrParameters.
func (in *RwGroupAddrParameters) DeepCopy() *RwGroupAddrParameters {
	if in == nil {
		return nil
	}
	out := new(RwGroupAddrParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RwGroupInstancesObservation) DeepCopyInto(out *RwGroupInstancesObservation) {
	*out = *in
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceName != nil {
		in, out := &in.InstanceName, &out.InstanceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RwGroupInstancesObservation.
func (in *RwGroupInstancesObservation) DeepCopy() *RwGroupInstancesObservation {
	if in == nil {
		return nil
	}
	out := new(RwGroupInstancesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RwGroupInstancesParameters) DeepCopyInto(out *RwGroupInstancesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RwGroupInstancesParameters.
func (in *RwGroupInstancesParameters) DeepCopy() *RwGroupInstancesParameters {
	if in == nil {
		return nil
	}
	out := new(RwGroupInstancesParameters)
	in.DeepCopyInto(out)
	return out
}
