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

type KeyPairInitParameters struct {

	// The key pair's name. It is the only in one TencentCloud account.
	// The key pair's name. It is the only in one TencentCloud account.
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// Specifys to which project the key pair belongs.
	// Specifys to which project the key pair belongs.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// You can import an existing public key and using TencentCloud key pair to manage it.
	// You can import an existing public key and using TencentCloud key pair to manage it.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// Tags of the key pair.
	// Tags of the key pair.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KeyPairObservation struct {

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The key pair's name. It is the only in one TencentCloud account.
	// The key pair's name. It is the only in one TencentCloud account.
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// Specifys to which project the key pair belongs.
	// Specifys to which project the key pair belongs.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// You can import an existing public key and using TencentCloud key pair to manage it.
	// You can import an existing public key and using TencentCloud key pair to manage it.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// Tags of the key pair.
	// Tags of the key pair.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KeyPairParameters struct {

	// The key pair's name. It is the only in one TencentCloud account.
	// The key pair's name. It is the only in one TencentCloud account.
	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// Specifys to which project the key pair belongs.
	// Specifys to which project the key pair belongs.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// You can import an existing public key and using TencentCloud key pair to manage it.
	// You can import an existing public key and using TencentCloud key pair to manage it.
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// Tags of the key pair.
	// Tags of the key pair.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// KeyPairSpec defines the desired state of KeyPair
type KeyPairSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyPairParameters `json:"forProvider"`
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
	InitProvider KeyPairInitParameters `json:"initProvider,omitempty"`
}

// KeyPairStatus defines the observed state of KeyPair.
type KeyPairStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyPairObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// KeyPair is the Schema for the KeyPairs API. Provides a key pair resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type KeyPair struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.keyName) || (has(self.initProvider) && has(self.initProvider.keyName))",message="spec.forProvider.keyName is a required parameter"
	Spec   KeyPairSpec   `json:"spec"`
	Status KeyPairStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyPairList contains a list of KeyPairs
type KeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyPair `json:"items"`
}

// Repository type metadata.
var (
	KeyPair_Kind             = "KeyPair"
	KeyPair_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeyPair_Kind}.String()
	KeyPair_KindAPIVersion   = KeyPair_Kind + "." + CRDGroupVersion.String()
	KeyPair_GroupVersionKind = CRDGroupVersion.WithKind(KeyPair_Kind)
)

func init() {
	SchemeBuilder.Register(&KeyPair{}, &KeyPairList{})
}
