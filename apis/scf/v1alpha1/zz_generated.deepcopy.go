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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CfsConfigObservation) DeepCopyInto(out *CfsConfigObservation) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.MountSubnetID != nil {
		in, out := &in.MountSubnetID, &out.MountSubnetID
		*out = new(string)
		**out = **in
	}
	if in.MountVPCID != nil {
		in, out := &in.MountVPCID, &out.MountVPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CfsConfigObservation.
func (in *CfsConfigObservation) DeepCopy() *CfsConfigObservation {
	if in == nil {
		return nil
	}
	out := new(CfsConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CfsConfigParameters) DeepCopyInto(out *CfsConfigParameters) {
	*out = *in
	if in.CfsID != nil {
		in, out := &in.CfsID, &out.CfsID
		*out = new(string)
		**out = **in
	}
	if in.LocalMountDir != nil {
		in, out := &in.LocalMountDir, &out.LocalMountDir
		*out = new(string)
		**out = **in
	}
	if in.MountInsID != nil {
		in, out := &in.MountInsID, &out.MountInsID
		*out = new(string)
		**out = **in
	}
	if in.RemoteMountDir != nil {
		in, out := &in.RemoteMountDir, &out.RemoteMountDir
		*out = new(string)
		**out = **in
	}
	if in.UserGroupID != nil {
		in, out := &in.UserGroupID, &out.UserGroupID
		*out = new(string)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CfsConfigParameters.
func (in *CfsConfigParameters) DeepCopy() *CfsConfigParameters {
	if in == nil {
		return nil
	}
	out := new(CfsConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentObservation) DeepCopyInto(out *ContentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentObservation.
func (in *ContentObservation) DeepCopy() *ContentObservation {
	if in == nil {
		return nil
	}
	out := new(ContentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentParameters) DeepCopyInto(out *ContentParameters) {
	*out = *in
	if in.CosBucketName != nil {
		in, out := &in.CosBucketName, &out.CosBucketName
		*out = new(string)
		**out = **in
	}
	if in.CosBucketRegion != nil {
		in, out := &in.CosBucketRegion, &out.CosBucketRegion
		*out = new(string)
		**out = **in
	}
	if in.CosObjectName != nil {
		in, out := &in.CosObjectName, &out.CosObjectName
		*out = new(string)
		**out = **in
	}
	if in.ZipFile != nil {
		in, out := &in.ZipFile, &out.ZipFile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentParameters.
func (in *ContentParameters) DeepCopy() *ContentParameters {
	if in == nil {
		return nil
	}
	out := new(ContentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionObservation) DeepCopyInto(out *FunctionObservation) {
	*out = *in
	if in.CfsConfig != nil {
		in, out := &in.CfsConfig, &out.CfsConfig
		*out = make([]CfsConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CodeError != nil {
		in, out := &in.CodeError, &out.CodeError
		*out = new(string)
		**out = **in
	}
	if in.CodeResult != nil {
		in, out := &in.CodeResult, &out.CodeResult
		*out = new(string)
		**out = **in
	}
	if in.CodeSize != nil {
		in, out := &in.CodeSize, &out.CodeSize
		*out = new(float64)
		**out = **in
	}
	if in.EIPFixed != nil {
		in, out := &in.EIPFixed, &out.EIPFixed
		*out = new(bool)
		**out = **in
	}
	if in.Eips != nil {
		in, out := &in.Eips, &out.Eips
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ErrNo != nil {
		in, out := &in.ErrNo, &out.ErrNo
		*out = new(float64)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InstallDependency != nil {
		in, out := &in.InstallDependency, &out.InstallDependency
		*out = new(bool)
		**out = **in
	}
	if in.ModifyTime != nil {
		in, out := &in.ModifyTime, &out.ModifyTime
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StatusDesc != nil {
		in, out := &in.StatusDesc, &out.StatusDesc
		*out = new(string)
		**out = **in
	}
	if in.TriggerInfo != nil {
		in, out := &in.TriggerInfo, &out.TriggerInfo
		*out = make([]TriggerInfoObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Vip != nil {
		in, out := &in.Vip, &out.Vip
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionObservation.
func (in *FunctionObservation) DeepCopy() *FunctionObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionParameters) DeepCopyInto(out *FunctionParameters) {
	*out = *in
	if in.CfsConfig != nil {
		in, out := &in.CfsConfig, &out.CfsConfig
		*out = make([]CfsConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClsLogsetID != nil {
		in, out := &in.ClsLogsetID, &out.ClsLogsetID
		*out = new(string)
		**out = **in
	}
	if in.ClsTopicID != nil {
		in, out := &in.ClsTopicID, &out.ClsTopicID
		*out = new(string)
		**out = **in
	}
	if in.CosBucketName != nil {
		in, out := &in.CosBucketName, &out.CosBucketName
		*out = new(string)
		**out = **in
	}
	if in.CosBucketRegion != nil {
		in, out := &in.CosBucketRegion, &out.CosBucketRegion
		*out = new(string)
		**out = **in
	}
	if in.CosObjectName != nil {
		in, out := &in.CosObjectName, &out.CosObjectName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnableEIPConfig != nil {
		in, out := &in.EnableEIPConfig, &out.EnableEIPConfig
		*out = new(bool)
		**out = **in
	}
	if in.EnablePublicNet != nil {
		in, out := &in.EnablePublicNet, &out.EnablePublicNet
		*out = new(bool)
		**out = **in
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
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
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = new(string)
		**out = **in
	}
	if in.ImageConfig != nil {
		in, out := &in.ImageConfig, &out.ImageConfig
		*out = make([]ImageConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.L5Enable != nil {
		in, out := &in.L5Enable, &out.L5Enable
		*out = new(bool)
		**out = **in
	}
	if in.Layers != nil {
		in, out := &in.Layers, &out.Layers
		*out = make([]LayersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemSize != nil {
		in, out := &in.MemSize, &out.MemSize
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
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
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(float64)
		**out = **in
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]TriggersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.ZipFile != nil {
		in, out := &in.ZipFile, &out.ZipFile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionParameters.
func (in *FunctionParameters) DeepCopy() *FunctionParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfigObservation) DeepCopyInto(out *ImageConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfigObservation.
func (in *ImageConfigObservation) DeepCopy() *ImageConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ImageConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfigParameters) DeepCopyInto(out *ImageConfigParameters) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = new(string)
		**out = **in
	}
	if in.EntryPoint != nil {
		in, out := &in.EntryPoint, &out.EntryPoint
		*out = new(string)
		**out = **in
	}
	if in.ImageType != nil {
		in, out := &in.ImageType, &out.ImageType
		*out = new(string)
		**out = **in
	}
	if in.ImageURI != nil {
		in, out := &in.ImageURI, &out.ImageURI
		*out = new(string)
		**out = **in
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfigParameters.
func (in *ImageConfigParameters) DeepCopy() *ImageConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ImageConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Layer) DeepCopyInto(out *Layer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Layer.
func (in *Layer) DeepCopy() *Layer {
	if in == nil {
		return nil
	}
	out := new(Layer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Layer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerList) DeepCopyInto(out *LayerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Layer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerList.
func (in *LayerList) DeepCopy() *LayerList {
	if in == nil {
		return nil
	}
	out := new(LayerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LayerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerObservation) DeepCopyInto(out *LayerObservation) {
	*out = *in
	if in.CodeSha256 != nil {
		in, out := &in.CodeSha256, &out.CodeSha256
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
	if in.LayerVersion != nil {
		in, out := &in.LayerVersion, &out.LayerVersion
		*out = new(float64)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerObservation.
func (in *LayerObservation) DeepCopy() *LayerObservation {
	if in == nil {
		return nil
	}
	out := new(LayerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerParameters) DeepCopyInto(out *LayerParameters) {
	*out = *in
	if in.CompatibleRuntimes != nil {
		in, out := &in.CompatibleRuntimes, &out.CompatibleRuntimes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = make([]ContentParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.LayerName != nil {
		in, out := &in.LayerName, &out.LayerName
		*out = new(string)
		**out = **in
	}
	if in.LicenseInfo != nil {
		in, out := &in.LicenseInfo, &out.LicenseInfo
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerParameters.
func (in *LayerParameters) DeepCopy() *LayerParameters {
	if in == nil {
		return nil
	}
	out := new(LayerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerSpec) DeepCopyInto(out *LayerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerSpec.
func (in *LayerSpec) DeepCopy() *LayerSpec {
	if in == nil {
		return nil
	}
	out := new(LayerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerStatus) DeepCopyInto(out *LayerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerStatus.
func (in *LayerStatus) DeepCopy() *LayerStatus {
	if in == nil {
		return nil
	}
	out := new(LayerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayersObservation) DeepCopyInto(out *LayersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayersObservation.
func (in *LayersObservation) DeepCopy() *LayersObservation {
	if in == nil {
		return nil
	}
	out := new(LayersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayersParameters) DeepCopyInto(out *LayersParameters) {
	*out = *in
	if in.LayerName != nil {
		in, out := &in.LayerName, &out.LayerName
		*out = new(string)
		**out = **in
	}
	if in.LayerVersion != nil {
		in, out := &in.LayerVersion, &out.LayerVersion
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayersParameters.
func (in *LayersParameters) DeepCopy() *LayersParameters {
	if in == nil {
		return nil
	}
	out := new(LayersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScfNamespace) DeepCopyInto(out *ScfNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScfNamespace.
func (in *ScfNamespace) DeepCopy() *ScfNamespace {
	if in == nil {
		return nil
	}
	out := new(ScfNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScfNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScfNamespaceList) DeepCopyInto(out *ScfNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScfNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScfNamespaceList.
func (in *ScfNamespaceList) DeepCopy() *ScfNamespaceList {
	if in == nil {
		return nil
	}
	out := new(ScfNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScfNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScfNamespaceObservation) DeepCopyInto(out *ScfNamespaceObservation) {
	*out = *in
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
	if in.ModifyTime != nil {
		in, out := &in.ModifyTime, &out.ModifyTime
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScfNamespaceObservation.
func (in *ScfNamespaceObservation) DeepCopy() *ScfNamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(ScfNamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScfNamespaceParameters) DeepCopyInto(out *ScfNamespaceParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScfNamespaceParameters.
func (in *ScfNamespaceParameters) DeepCopy() *ScfNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(ScfNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScfNamespaceSpec) DeepCopyInto(out *ScfNamespaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScfNamespaceSpec.
func (in *ScfNamespaceSpec) DeepCopy() *ScfNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(ScfNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScfNamespaceStatus) DeepCopyInto(out *ScfNamespaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScfNamespaceStatus.
func (in *ScfNamespaceStatus) DeepCopy() *ScfNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(ScfNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerInfoObservation) DeepCopyInto(out *TriggerInfoObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.CustomArgument != nil {
		in, out := &in.CustomArgument, &out.CustomArgument
		*out = new(string)
		**out = **in
	}
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.ModifyTime != nil {
		in, out := &in.ModifyTime, &out.ModifyTime
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TriggerDesc != nil {
		in, out := &in.TriggerDesc, &out.TriggerDesc
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerInfoObservation.
func (in *TriggerInfoObservation) DeepCopy() *TriggerInfoObservation {
	if in == nil {
		return nil
	}
	out := new(TriggerInfoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerInfoParameters) DeepCopyInto(out *TriggerInfoParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerInfoParameters.
func (in *TriggerInfoParameters) DeepCopy() *TriggerInfoParameters {
	if in == nil {
		return nil
	}
	out := new(TriggerInfoParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggersObservation) DeepCopyInto(out *TriggersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggersObservation.
func (in *TriggersObservation) DeepCopy() *TriggersObservation {
	if in == nil {
		return nil
	}
	out := new(TriggersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggersParameters) DeepCopyInto(out *TriggersParameters) {
	*out = *in
	if in.CosRegion != nil {
		in, out := &in.CosRegion, &out.CosRegion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TriggerDesc != nil {
		in, out := &in.TriggerDesc, &out.TriggerDesc
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggersParameters.
func (in *TriggersParameters) DeepCopy() *TriggersParameters {
	if in == nil {
		return nil
	}
	out := new(TriggersParameters)
	in.DeepCopyInto(out)
	return out
}
