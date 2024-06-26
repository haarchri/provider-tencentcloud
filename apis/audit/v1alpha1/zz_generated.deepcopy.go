//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Audit) DeepCopyInto(out *Audit) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Audit.
func (in *Audit) DeepCopy() *Audit {
	if in == nil {
		return nil
	}
	out := new(Audit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Audit) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditInitParameters) DeepCopyInto(out *AuditInitParameters) {
	*out = *in
	if in.AuditSwitch != nil {
		in, out := &in.AuditSwitch, &out.AuditSwitch
		*out = new(bool)
		**out = **in
	}
	if in.CosBucket != nil {
		in, out := &in.CosBucket, &out.CosBucket
		*out = new(string)
		**out = **in
	}
	if in.CosRegion != nil {
		in, out := &in.CosRegion, &out.CosRegion
		*out = new(string)
		**out = **in
	}
	if in.EnableKMSEncry != nil {
		in, out := &in.EnableKMSEncry, &out.EnableKMSEncry
		*out = new(bool)
		**out = **in
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.LogFilePrefix != nil {
		in, out := &in.LogFilePrefix, &out.LogFilePrefix
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ReadWriteAttribute != nil {
		in, out := &in.ReadWriteAttribute, &out.ReadWriteAttribute
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditInitParameters.
func (in *AuditInitParameters) DeepCopy() *AuditInitParameters {
	if in == nil {
		return nil
	}
	out := new(AuditInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditList) DeepCopyInto(out *AuditList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Audit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditList.
func (in *AuditList) DeepCopy() *AuditList {
	if in == nil {
		return nil
	}
	out := new(AuditList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuditList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditObservation) DeepCopyInto(out *AuditObservation) {
	*out = *in
	if in.AuditSwitch != nil {
		in, out := &in.AuditSwitch, &out.AuditSwitch
		*out = new(bool)
		**out = **in
	}
	if in.CosBucket != nil {
		in, out := &in.CosBucket, &out.CosBucket
		*out = new(string)
		**out = **in
	}
	if in.CosRegion != nil {
		in, out := &in.CosRegion, &out.CosRegion
		*out = new(string)
		**out = **in
	}
	if in.EnableKMSEncry != nil {
		in, out := &in.EnableKMSEncry, &out.EnableKMSEncry
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.LogFilePrefix != nil {
		in, out := &in.LogFilePrefix, &out.LogFilePrefix
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ReadWriteAttribute != nil {
		in, out := &in.ReadWriteAttribute, &out.ReadWriteAttribute
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditObservation.
func (in *AuditObservation) DeepCopy() *AuditObservation {
	if in == nil {
		return nil
	}
	out := new(AuditObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditParameters) DeepCopyInto(out *AuditParameters) {
	*out = *in
	if in.AuditSwitch != nil {
		in, out := &in.AuditSwitch, &out.AuditSwitch
		*out = new(bool)
		**out = **in
	}
	if in.CosBucket != nil {
		in, out := &in.CosBucket, &out.CosBucket
		*out = new(string)
		**out = **in
	}
	if in.CosRegion != nil {
		in, out := &in.CosRegion, &out.CosRegion
		*out = new(string)
		**out = **in
	}
	if in.EnableKMSEncry != nil {
		in, out := &in.EnableKMSEncry, &out.EnableKMSEncry
		*out = new(bool)
		**out = **in
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.LogFilePrefix != nil {
		in, out := &in.LogFilePrefix, &out.LogFilePrefix
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ReadWriteAttribute != nil {
		in, out := &in.ReadWriteAttribute, &out.ReadWriteAttribute
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditParameters.
func (in *AuditParameters) DeepCopy() *AuditParameters {
	if in == nil {
		return nil
	}
	out := new(AuditParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditSpec) DeepCopyInto(out *AuditSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditSpec.
func (in *AuditSpec) DeepCopy() *AuditSpec {
	if in == nil {
		return nil
	}
	out := new(AuditSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditStatus) DeepCopyInto(out *AuditStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditStatus.
func (in *AuditStatus) DeepCopy() *AuditStatus {
	if in == nil {
		return nil
	}
	out := new(AuditStatus)
	in.DeepCopyInto(out)
	return out
}
