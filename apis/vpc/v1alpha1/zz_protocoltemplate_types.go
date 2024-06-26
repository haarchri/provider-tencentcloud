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

type ProtocolTemplateInitParameters struct {

	// Name of the protocol template.
	// Name of the protocol template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Protocol list. Valid protocols are  tcp, udp, icmp, gre. Single port(tcp:80), multi-port(tcp:80,443), port range(tcp:3306-20000), all(tcp:all) format are support. Protocol icmp and gre cannot specify port.
	// Protocol list. Valid protocols are  `tcp`, `udp`, `icmp`, `gre`. Single port(tcp:80), multi-port(tcp:80,443), port range(tcp:3306-20000), all(tcp:all) format are support. Protocol `icmp` and `gre` cannot specify port.
	// +listType=set
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`
}

type ProtocolTemplateObservation struct {

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the protocol template.
	// Name of the protocol template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Protocol list. Valid protocols are  tcp, udp, icmp, gre. Single port(tcp:80), multi-port(tcp:80,443), port range(tcp:3306-20000), all(tcp:all) format are support. Protocol icmp and gre cannot specify port.
	// Protocol list. Valid protocols are  `tcp`, `udp`, `icmp`, `gre`. Single port(tcp:80), multi-port(tcp:80,443), port range(tcp:3306-20000), all(tcp:all) format are support. Protocol `icmp` and `gre` cannot specify port.
	// +listType=set
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`
}

type ProtocolTemplateParameters struct {

	// Name of the protocol template.
	// Name of the protocol template.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Protocol list. Valid protocols are  tcp, udp, icmp, gre. Single port(tcp:80), multi-port(tcp:80,443), port range(tcp:3306-20000), all(tcp:all) format are support. Protocol icmp and gre cannot specify port.
	// Protocol list. Valid protocols are  `tcp`, `udp`, `icmp`, `gre`. Single port(tcp:80), multi-port(tcp:80,443), port range(tcp:3306-20000), all(tcp:all) format are support. Protocol `icmp` and `gre` cannot specify port.
	// +kubebuilder:validation:Optional
	// +listType=set
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`
}

// ProtocolTemplateSpec defines the desired state of ProtocolTemplate
type ProtocolTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProtocolTemplateParameters `json:"forProvider"`
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
	InitProvider ProtocolTemplateInitParameters `json:"initProvider,omitempty"`
}

// ProtocolTemplateStatus defines the observed state of ProtocolTemplate.
type ProtocolTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProtocolTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProtocolTemplate is the Schema for the ProtocolTemplates API. Provides a resource to manage protocol template.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type ProtocolTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocols) || (has(self.initProvider) && has(self.initProvider.protocols))",message="spec.forProvider.protocols is a required parameter"
	Spec   ProtocolTemplateSpec   `json:"spec"`
	Status ProtocolTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtocolTemplateList contains a list of ProtocolTemplates
type ProtocolTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProtocolTemplate `json:"items"`
}

// Repository type metadata.
var (
	ProtocolTemplate_Kind             = "ProtocolTemplate"
	ProtocolTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProtocolTemplate_Kind}.String()
	ProtocolTemplate_KindAPIVersion   = ProtocolTemplate_Kind + "." + CRDGroupVersion.String()
	ProtocolTemplate_GroupVersionKind = CRDGroupVersion.WithKind(ProtocolTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&ProtocolTemplate{}, &ProtocolTemplateList{})
}
