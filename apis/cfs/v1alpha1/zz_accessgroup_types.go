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

type AccessGroupInitParameters struct {

	// Description of the access group, and max length is 255.
	// Description of the access group, and max length is 255.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the access group, and max length is 64.
	// Name of the access group, and max length is 64.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AccessGroupObservation struct {

	// Create time of the access group.
	// Create time of the access group.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Description of the access group, and max length is 255.
	// Description of the access group, and max length is 255.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the access group, and max length is 64.
	// Name of the access group, and max length is 64.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AccessGroupParameters struct {

	// Description of the access group, and max length is 255.
	// Description of the access group, and max length is 255.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the access group, and max length is 64.
	// Name of the access group, and max length is 64.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// AccessGroupSpec defines the desired state of AccessGroup
type AccessGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessGroupParameters `json:"forProvider"`
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
	InitProvider AccessGroupInitParameters `json:"initProvider,omitempty"`
}

// AccessGroupStatus defines the observed state of AccessGroup.
type AccessGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccessGroup is the Schema for the AccessGroups API. Provides a resource to create a CFS access group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type AccessGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AccessGroupSpec   `json:"spec"`
	Status AccessGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessGroupList contains a list of AccessGroups
type AccessGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessGroup `json:"items"`
}

// Repository type metadata.
var (
	AccessGroup_Kind             = "AccessGroup"
	AccessGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessGroup_Kind}.String()
	AccessGroup_KindAPIVersion   = AccessGroup_Kind + "." + CRDGroupVersion.String()
	AccessGroup_GroupVersionKind = CRDGroupVersion.WithKind(AccessGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessGroup{}, &AccessGroupList{})
}
