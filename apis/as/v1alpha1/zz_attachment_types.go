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

type AttachmentInitParameters struct {

	// ID list of CVM instances to be attached to the scaling group.
	// ID list of CVM instances to be attached to the scaling group.
	// +listType=set
	InstanceIds []*string `json:"instanceIds,omitempty" tf:"instance_ids,omitempty"`

	// ID of a scaling group.
	// ID of a scaling group.
	// +crossplane:generate:reference:type=ScalingGroup
	ScalingGroupID *string `json:"scalingGroupId,omitempty" tf:"scaling_group_id,omitempty"`

	// Reference to a ScalingGroup to populate scalingGroupId.
	// +kubebuilder:validation:Optional
	ScalingGroupIDRef *v1.Reference `json:"scalingGroupIdRef,omitempty" tf:"-"`

	// Selector for a ScalingGroup to populate scalingGroupId.
	// +kubebuilder:validation:Optional
	ScalingGroupIDSelector *v1.Selector `json:"scalingGroupIdSelector,omitempty" tf:"-"`
}

type AttachmentObservation struct {

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID list of CVM instances to be attached to the scaling group.
	// ID list of CVM instances to be attached to the scaling group.
	// +listType=set
	InstanceIds []*string `json:"instanceIds,omitempty" tf:"instance_ids,omitempty"`

	// ID of a scaling group.
	// ID of a scaling group.
	ScalingGroupID *string `json:"scalingGroupId,omitempty" tf:"scaling_group_id,omitempty"`
}

type AttachmentParameters struct {

	// ID list of CVM instances to be attached to the scaling group.
	// ID list of CVM instances to be attached to the scaling group.
	// +kubebuilder:validation:Optional
	// +listType=set
	InstanceIds []*string `json:"instanceIds,omitempty" tf:"instance_ids,omitempty"`

	// ID of a scaling group.
	// ID of a scaling group.
	// +crossplane:generate:reference:type=ScalingGroup
	// +kubebuilder:validation:Optional
	ScalingGroupID *string `json:"scalingGroupId,omitempty" tf:"scaling_group_id,omitempty"`

	// Reference to a ScalingGroup to populate scalingGroupId.
	// +kubebuilder:validation:Optional
	ScalingGroupIDRef *v1.Reference `json:"scalingGroupIdRef,omitempty" tf:"-"`

	// Selector for a ScalingGroup to populate scalingGroupId.
	// +kubebuilder:validation:Optional
	ScalingGroupIDSelector *v1.Selector `json:"scalingGroupIdSelector,omitempty" tf:"-"`
}

// AttachmentSpec defines the desired state of Attachment
type AttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttachmentParameters `json:"forProvider"`
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
	InitProvider AttachmentInitParameters `json:"initProvider,omitempty"`
}

// AttachmentStatus defines the observed state of Attachment.
type AttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Attachment is the Schema for the Attachments API. Provides a resource to attach or detach CVM instances to a specified scaling group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Attachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceIds) || (has(self.initProvider) && has(self.initProvider.instanceIds))",message="spec.forProvider.instanceIds is a required parameter"
	Spec   AttachmentSpec   `json:"spec"`
	Status AttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttachmentList contains a list of Attachments
type AttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Attachment `json:"items"`
}

// Repository type metadata.
var (
	Attachment_Kind             = "Attachment"
	Attachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Attachment_Kind}.String()
	Attachment_KindAPIVersion   = Attachment_Kind + "." + CRDGroupVersion.String()
	Attachment_GroupVersionKind = CRDGroupVersion.WithKind(Attachment_Kind)
)

func init() {
	SchemeBuilder.Register(&Attachment{}, &AttachmentList{})
}
