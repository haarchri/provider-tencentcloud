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

type RoleSSOInitParameters struct {

	// Client ids.
	// Client ids.
	// +listType=set
	ClientIds []*string `json:"clientIds,omitempty" tf:"client_ids,omitempty"`

	// The description of resource.
	// The description of resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Sign the public key.
	// Sign the public key.
	IdentityKey *string `json:"identityKey,omitempty" tf:"identity_key,omitempty"`

	// Identity provider URL.
	// Identity provider URL.
	IdentityURL *string `json:"identityUrl,omitempty" tf:"identity_url,omitempty"`

	// The name of resource.
	// The name of resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type RoleSSOObservation struct {

	// Client ids.
	// Client ids.
	// +listType=set
	ClientIds []*string `json:"clientIds,omitempty" tf:"client_ids,omitempty"`

	// The description of resource.
	// The description of resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Sign the public key.
	// Sign the public key.
	IdentityKey *string `json:"identityKey,omitempty" tf:"identity_key,omitempty"`

	// Identity provider URL.
	// Identity provider URL.
	IdentityURL *string `json:"identityUrl,omitempty" tf:"identity_url,omitempty"`

	// The name of resource.
	// The name of resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type RoleSSOParameters struct {

	// Client ids.
	// Client ids.
	// +kubebuilder:validation:Optional
	// +listType=set
	ClientIds []*string `json:"clientIds,omitempty" tf:"client_ids,omitempty"`

	// The description of resource.
	// The description of resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Sign the public key.
	// Sign the public key.
	// +kubebuilder:validation:Optional
	IdentityKey *string `json:"identityKey,omitempty" tf:"identity_key,omitempty"`

	// Identity provider URL.
	// Identity provider URL.
	// +kubebuilder:validation:Optional
	IdentityURL *string `json:"identityUrl,omitempty" tf:"identity_url,omitempty"`

	// The name of resource.
	// The name of resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// RoleSSOSpec defines the desired state of RoleSSO
type RoleSSOSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleSSOParameters `json:"forProvider"`
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
	InitProvider RoleSSOInitParameters `json:"initProvider,omitempty"`
}

// RoleSSOStatus defines the observed state of RoleSSO.
type RoleSSOStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleSSOObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RoleSSO is the Schema for the RoleSSOs API. Provides a resource to create a CAM-ROLE-SSO (Only support OIDC).
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type RoleSSO struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientIds) || (has(self.initProvider) && has(self.initProvider.clientIds))",message="spec.forProvider.clientIds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identityKey) || (has(self.initProvider) && has(self.initProvider.identityKey))",message="spec.forProvider.identityKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identityUrl) || (has(self.initProvider) && has(self.initProvider.identityUrl))",message="spec.forProvider.identityUrl is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RoleSSOSpec   `json:"spec"`
	Status RoleSSOStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleSSOList contains a list of RoleSSOs
type RoleSSOList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleSSO `json:"items"`
}

// Repository type metadata.
var (
	RoleSSO_Kind             = "RoleSSO"
	RoleSSO_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleSSO_Kind}.String()
	RoleSSO_KindAPIVersion   = RoleSSO_Kind + "." + CRDGroupVersion.String()
	RoleSSO_GroupVersionKind = CRDGroupVersion.WithKind(RoleSSO_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleSSO{}, &RoleSSOList{})
}
