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

type TargetGroupInitParameters struct {

	// The default port of target group, add server after can use it.
	// The default port of target group, add server after can use it.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// It has been deprecated from version 1.77.3. please use tencentcloud_clb_target_group_instance_attachment instead. The backend server of target group bind.
	// The backend server of target group bind.
	TargetGroupInstances []TargetGroupInstancesInitParameters `json:"targetGroupInstances,omitempty" tf:"target_group_instances,omitempty"`

	// Target group name.
	// Target group name.
	TargetGroupName *string `json:"targetGroupName,omitempty" tf:"target_group_name,omitempty"`

	// VPC ID, default is based on the network.
	// VPC ID, default is based on the network.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type TargetGroupInstancesInitParameters struct {

	// The internal ip of target group instance.
	// The internal ip of target group instance.
	BindIP *string `json:"bindIp,omitempty" tf:"bind_ip,omitempty"`

	// The new port of target group instance.
	// The new port of target group instance.
	NewPort *float64 `json:"newPort,omitempty" tf:"new_port,omitempty"`

	// The default port of target group, add server after can use it.
	// The port of target group instance.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The weight of target group instance.
	// The weight of target group instance.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TargetGroupInstancesObservation struct {

	// The internal ip of target group instance.
	// The internal ip of target group instance.
	BindIP *string `json:"bindIp,omitempty" tf:"bind_ip,omitempty"`

	// The new port of target group instance.
	// The new port of target group instance.
	NewPort *float64 `json:"newPort,omitempty" tf:"new_port,omitempty"`

	// The default port of target group, add server after can use it.
	// The port of target group instance.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The weight of target group instance.
	// The weight of target group instance.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TargetGroupInstancesParameters struct {

	// The internal ip of target group instance.
	// The internal ip of target group instance.
	// +kubebuilder:validation:Optional
	BindIP *string `json:"bindIp" tf:"bind_ip,omitempty"`

	// The new port of target group instance.
	// The new port of target group instance.
	// +kubebuilder:validation:Optional
	NewPort *float64 `json:"newPort,omitempty" tf:"new_port,omitempty"`

	// The default port of target group, add server after can use it.
	// The port of target group instance.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// The weight of target group instance.
	// The weight of target group instance.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TargetGroupObservation struct {

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The default port of target group, add server after can use it.
	// The default port of target group, add server after can use it.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// It has been deprecated from version 1.77.3. please use tencentcloud_clb_target_group_instance_attachment instead. The backend server of target group bind.
	// The backend server of target group bind.
	TargetGroupInstances []TargetGroupInstancesObservation `json:"targetGroupInstances,omitempty" tf:"target_group_instances,omitempty"`

	// Target group name.
	// Target group name.
	TargetGroupName *string `json:"targetGroupName,omitempty" tf:"target_group_name,omitempty"`

	// VPC ID, default is based on the network.
	// VPC ID, default is based on the network.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type TargetGroupParameters struct {

	// The default port of target group, add server after can use it.
	// The default port of target group, add server after can use it.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// It has been deprecated from version 1.77.3. please use tencentcloud_clb_target_group_instance_attachment instead. The backend server of target group bind.
	// The backend server of target group bind.
	// +kubebuilder:validation:Optional
	TargetGroupInstances []TargetGroupInstancesParameters `json:"targetGroupInstances,omitempty" tf:"target_group_instances,omitempty"`

	// Target group name.
	// Target group name.
	// +kubebuilder:validation:Optional
	TargetGroupName *string `json:"targetGroupName,omitempty" tf:"target_group_name,omitempty"`

	// VPC ID, default is based on the network.
	// VPC ID, default is based on the network.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

// TargetGroupSpec defines the desired state of TargetGroup
type TargetGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetGroupParameters `json:"forProvider"`
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
	InitProvider TargetGroupInitParameters `json:"initProvider,omitempty"`
}

// TargetGroupStatus defines the observed state of TargetGroup.
type TargetGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TargetGroup is the Schema for the TargetGroups API. Provides a resource to create a CLB target group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type TargetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetGroupSpec   `json:"spec"`
	Status            TargetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetGroupList contains a list of TargetGroups
type TargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetGroup `json:"items"`
}

// Repository type metadata.
var (
	TargetGroup_Kind             = "TargetGroup"
	TargetGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TargetGroup_Kind}.String()
	TargetGroup_KindAPIVersion   = TargetGroup_Kind + "." + CRDGroupVersion.String()
	TargetGroup_GroupVersionKind = CRDGroupVersion.WithKind(TargetGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&TargetGroup{}, &TargetGroupList{})
}
