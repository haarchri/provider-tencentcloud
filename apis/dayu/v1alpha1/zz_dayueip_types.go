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

type DayuEipInitParameters struct {

	// Resource id to bind.
	// Resource id to bind.
	BindResourceID *string `json:"bindResourceId,omitempty" tf:"bind_resource_id,omitempty"`

	// Resource region to bind.
	// Resource region to bind.
	BindResourceRegion *string `json:"bindResourceRegion,omitempty" tf:"bind_resource_region,omitempty"`

	// Resource type to bind, value range [clb, cvm].
	// Resource type to bind, value range [`clb`, `cvm`].
	BindResourceType *string `json:"bindResourceType,omitempty" tf:"bind_resource_type,omitempty"`

	// Eip of the resource.
	// Eip of the resource.
	EIP *string `json:"eip,omitempty" tf:"eip,omitempty"`

	// ID of the resource.
	// ID of the resource.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type DayuEipObservation struct {

	// Resource id to bind.
	// Resource id to bind.
	BindResourceID *string `json:"bindResourceId,omitempty" tf:"bind_resource_id,omitempty"`

	// Resource region to bind.
	// Resource region to bind.
	BindResourceRegion *string `json:"bindResourceRegion,omitempty" tf:"bind_resource_region,omitempty"`

	// Resource type to bind, value range [clb, cvm].
	// Resource type to bind, value range [`clb`, `cvm`].
	BindResourceType *string `json:"bindResourceType,omitempty" tf:"bind_resource_type,omitempty"`

	// Created time of the resource instance.
	// Created time of the resource instance.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// Eip of the resource.
	// Eip of the resource.
	EIP *string `json:"eip,omitempty" tf:"eip,omitempty"`

	// Eip address status of the resource instance.
	// Eip address status of the resource instance.
	EIPAddressStatus *string `json:"eipAddressStatus,omitempty" tf:"eip_address_status,omitempty"`

	// Eip bound rsc eni of the resource instance.
	// Eip bound rsc eni of the resource instance.
	EIPBoundRscEni *string `json:"eipBoundRscEni,omitempty" tf:"eip_bound_rsc_eni,omitempty"`

	// Eip bound rsc ins of the resource instance.
	// Eip bound rsc ins of the resource instance.
	EIPBoundRscIns *string `json:"eipBoundRscIns,omitempty" tf:"eip_bound_rsc_ins,omitempty"`

	// Eip bound rsc vip of the resource instance.
	// Eip bound rsc vip of the resource instance.
	EIPBoundRscVip *string `json:"eipBoundRscVip,omitempty" tf:"eip_bound_rsc_vip,omitempty"`

	// Expired time of the resource instance.
	// Expired time of the resource instance.
	ExpiredTime *string `json:"expiredTime,omitempty" tf:"expired_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Modify time of the resource instance.
	// Modify time of the resource instance.
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// Protection status of the resource instance.
	// Protection status of the resource instance.
	ProtectionStatus *string `json:"protectionStatus,omitempty" tf:"protection_status,omitempty"`

	// ID of the resource.
	// ID of the resource.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Region of the resource instance.
	// Region of the resource instance.
	ResourceRegion *string `json:"resourceRegion,omitempty" tf:"resource_region,omitempty"`
}

type DayuEipParameters struct {

	// Resource id to bind.
	// Resource id to bind.
	// +kubebuilder:validation:Optional
	BindResourceID *string `json:"bindResourceId,omitempty" tf:"bind_resource_id,omitempty"`

	// Resource region to bind.
	// Resource region to bind.
	// +kubebuilder:validation:Optional
	BindResourceRegion *string `json:"bindResourceRegion,omitempty" tf:"bind_resource_region,omitempty"`

	// Resource type to bind, value range [clb, cvm].
	// Resource type to bind, value range [`clb`, `cvm`].
	// +kubebuilder:validation:Optional
	BindResourceType *string `json:"bindResourceType,omitempty" tf:"bind_resource_type,omitempty"`

	// Eip of the resource.
	// Eip of the resource.
	// +kubebuilder:validation:Optional
	EIP *string `json:"eip,omitempty" tf:"eip,omitempty"`

	// ID of the resource.
	// ID of the resource.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

// DayuEipSpec defines the desired state of DayuEip
type DayuEipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DayuEipParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DayuEipInitParameters `json:"initProvider,omitempty"`
}

// DayuEipStatus defines the observed state of DayuEip.
type DayuEipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DayuEipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DayuEip is the Schema for the DayuEips API. Use this resource to create dayu eip rule
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type DayuEip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bindResourceId) || (has(self.initProvider) && has(self.initProvider.bindResourceId))",message="spec.forProvider.bindResourceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bindResourceRegion) || (has(self.initProvider) && has(self.initProvider.bindResourceRegion))",message="spec.forProvider.bindResourceRegion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bindResourceType) || (has(self.initProvider) && has(self.initProvider.bindResourceType))",message="spec.forProvider.bindResourceType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.eip) || (has(self.initProvider) && has(self.initProvider.eip))",message="spec.forProvider.eip is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceId) || (has(self.initProvider) && has(self.initProvider.resourceId))",message="spec.forProvider.resourceId is a required parameter"
	Spec   DayuEipSpec   `json:"spec"`
	Status DayuEipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DayuEipList contains a list of DayuEips
type DayuEipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DayuEip `json:"items"`
}

// Repository type metadata.
var (
	DayuEip_Kind             = "DayuEip"
	DayuEip_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DayuEip_Kind}.String()
	DayuEip_KindAPIVersion   = DayuEip_Kind + "." + CRDGroupVersion.String()
	DayuEip_GroupVersionKind = CRDGroupVersion.WithKind(DayuEip_Kind)
)

func init() {
	SchemeBuilder.Register(&DayuEip{}, &DayuEipList{})
}
