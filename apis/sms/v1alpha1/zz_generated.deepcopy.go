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
func (in *Sign) DeepCopyInto(out *Sign) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sign.
func (in *Sign) DeepCopy() *Sign {
	if in == nil {
		return nil
	}
	out := new(Sign)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Sign) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignInitParameters) DeepCopyInto(out *SignInitParameters) {
	*out = *in
	if in.CommissionImage != nil {
		in, out := &in.CommissionImage, &out.CommissionImage
		*out = new(string)
		**out = **in
	}
	if in.DocumentType != nil {
		in, out := &in.DocumentType, &out.DocumentType
		*out = new(float64)
		**out = **in
	}
	if in.International != nil {
		in, out := &in.International, &out.International
		*out = new(float64)
		**out = **in
	}
	if in.ProofImage != nil {
		in, out := &in.ProofImage, &out.ProofImage
		*out = new(string)
		**out = **in
	}
	if in.Remark != nil {
		in, out := &in.Remark, &out.Remark
		*out = new(string)
		**out = **in
	}
	if in.SignName != nil {
		in, out := &in.SignName, &out.SignName
		*out = new(string)
		**out = **in
	}
	if in.SignPurpose != nil {
		in, out := &in.SignPurpose, &out.SignPurpose
		*out = new(float64)
		**out = **in
	}
	if in.SignType != nil {
		in, out := &in.SignType, &out.SignType
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignInitParameters.
func (in *SignInitParameters) DeepCopy() *SignInitParameters {
	if in == nil {
		return nil
	}
	out := new(SignInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignList) DeepCopyInto(out *SignList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sign, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignList.
func (in *SignList) DeepCopy() *SignList {
	if in == nil {
		return nil
	}
	out := new(SignList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SignList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignObservation) DeepCopyInto(out *SignObservation) {
	*out = *in
	if in.CommissionImage != nil {
		in, out := &in.CommissionImage, &out.CommissionImage
		*out = new(string)
		**out = **in
	}
	if in.DocumentType != nil {
		in, out := &in.DocumentType, &out.DocumentType
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.International != nil {
		in, out := &in.International, &out.International
		*out = new(float64)
		**out = **in
	}
	if in.ProofImage != nil {
		in, out := &in.ProofImage, &out.ProofImage
		*out = new(string)
		**out = **in
	}
	if in.Remark != nil {
		in, out := &in.Remark, &out.Remark
		*out = new(string)
		**out = **in
	}
	if in.SignName != nil {
		in, out := &in.SignName, &out.SignName
		*out = new(string)
		**out = **in
	}
	if in.SignPurpose != nil {
		in, out := &in.SignPurpose, &out.SignPurpose
		*out = new(float64)
		**out = **in
	}
	if in.SignType != nil {
		in, out := &in.SignType, &out.SignType
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignObservation.
func (in *SignObservation) DeepCopy() *SignObservation {
	if in == nil {
		return nil
	}
	out := new(SignObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignParameters) DeepCopyInto(out *SignParameters) {
	*out = *in
	if in.CommissionImage != nil {
		in, out := &in.CommissionImage, &out.CommissionImage
		*out = new(string)
		**out = **in
	}
	if in.DocumentType != nil {
		in, out := &in.DocumentType, &out.DocumentType
		*out = new(float64)
		**out = **in
	}
	if in.International != nil {
		in, out := &in.International, &out.International
		*out = new(float64)
		**out = **in
	}
	if in.ProofImage != nil {
		in, out := &in.ProofImage, &out.ProofImage
		*out = new(string)
		**out = **in
	}
	if in.Remark != nil {
		in, out := &in.Remark, &out.Remark
		*out = new(string)
		**out = **in
	}
	if in.SignName != nil {
		in, out := &in.SignName, &out.SignName
		*out = new(string)
		**out = **in
	}
	if in.SignPurpose != nil {
		in, out := &in.SignPurpose, &out.SignPurpose
		*out = new(float64)
		**out = **in
	}
	if in.SignType != nil {
		in, out := &in.SignType, &out.SignType
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignParameters.
func (in *SignParameters) DeepCopy() *SignParameters {
	if in == nil {
		return nil
	}
	out := new(SignParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignSpec) DeepCopyInto(out *SignSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignSpec.
func (in *SignSpec) DeepCopy() *SignSpec {
	if in == nil {
		return nil
	}
	out := new(SignSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignStatus) DeepCopyInto(out *SignStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignStatus.
func (in *SignStatus) DeepCopy() *SignStatus {
	if in == nil {
		return nil
	}
	out := new(SignStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Template) DeepCopyInto(out *Template) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Template.
func (in *Template) DeepCopy() *Template {
	if in == nil {
		return nil
	}
	out := new(Template)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Template) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateInitParameters) DeepCopyInto(out *TemplateInitParameters) {
	*out = *in
	if in.International != nil {
		in, out := &in.International, &out.International
		*out = new(float64)
		**out = **in
	}
	if in.Remark != nil {
		in, out := &in.Remark, &out.Remark
		*out = new(string)
		**out = **in
	}
	if in.SMSType != nil {
		in, out := &in.SMSType, &out.SMSType
		*out = new(float64)
		**out = **in
	}
	if in.TemplateContent != nil {
		in, out := &in.TemplateContent, &out.TemplateContent
		*out = new(string)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateInitParameters.
func (in *TemplateInitParameters) DeepCopy() *TemplateInitParameters {
	if in == nil {
		return nil
	}
	out := new(TemplateInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateList) DeepCopyInto(out *TemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Template, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateList.
func (in *TemplateList) DeepCopy() *TemplateList {
	if in == nil {
		return nil
	}
	out := new(TemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateObservation) DeepCopyInto(out *TemplateObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.International != nil {
		in, out := &in.International, &out.International
		*out = new(float64)
		**out = **in
	}
	if in.Remark != nil {
		in, out := &in.Remark, &out.Remark
		*out = new(string)
		**out = **in
	}
	if in.SMSType != nil {
		in, out := &in.SMSType, &out.SMSType
		*out = new(float64)
		**out = **in
	}
	if in.TemplateContent != nil {
		in, out := &in.TemplateContent, &out.TemplateContent
		*out = new(string)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateObservation.
func (in *TemplateObservation) DeepCopy() *TemplateObservation {
	if in == nil {
		return nil
	}
	out := new(TemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateParameters) DeepCopyInto(out *TemplateParameters) {
	*out = *in
	if in.International != nil {
		in, out := &in.International, &out.International
		*out = new(float64)
		**out = **in
	}
	if in.Remark != nil {
		in, out := &in.Remark, &out.Remark
		*out = new(string)
		**out = **in
	}
	if in.SMSType != nil {
		in, out := &in.SMSType, &out.SMSType
		*out = new(float64)
		**out = **in
	}
	if in.TemplateContent != nil {
		in, out := &in.TemplateContent, &out.TemplateContent
		*out = new(string)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateParameters.
func (in *TemplateParameters) DeepCopy() *TemplateParameters {
	if in == nil {
		return nil
	}
	out := new(TemplateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpec) DeepCopyInto(out *TemplateSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpec.
func (in *TemplateSpec) DeepCopy() *TemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateStatus) DeepCopyInto(out *TemplateStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateStatus.
func (in *TemplateStatus) DeepCopy() *TemplateStatus {
	if in == nil {
		return nil
	}
	out := new(TemplateStatus)
	in.DeepCopyInto(out)
	return out
}
