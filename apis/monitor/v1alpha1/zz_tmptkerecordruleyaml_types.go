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

type TmpTkeRecordRuleYamlInitParameters struct {

	// Contents of record rules in yaml format.
	// Contents of record rules in yaml format.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Instance Id.
	// Instance Id.
	// +crossplane:generate:reference:type=TmpInstance
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a TmpInstance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a TmpInstance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`
}

type TmpTkeRecordRuleYamlObservation struct {

	// An ID identify the cluster, like cls-xxxxxx.
	// An ID identify the cluster, like cls-xxxxxx.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Contents of record rules in yaml format.
	// Contents of record rules in yaml format.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Instance Id.
	// Instance Id.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Name of the instance.
	// Name of the instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`

	// Last modified time of record rule.
	// Last modified time of record rule.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type TmpTkeRecordRuleYamlParameters struct {

	// Contents of record rules in yaml format.
	// Contents of record rules in yaml format.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Instance Id.
	// Instance Id.
	// +crossplane:generate:reference:type=TmpInstance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a TmpInstance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a TmpInstance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`
}

// TmpTkeRecordRuleYamlSpec defines the desired state of TmpTkeRecordRuleYaml
type TmpTkeRecordRuleYamlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TmpTkeRecordRuleYamlParameters `json:"forProvider"`
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
	InitProvider TmpTkeRecordRuleYamlInitParameters `json:"initProvider,omitempty"`
}

// TmpTkeRecordRuleYamlStatus defines the observed state of TmpTkeRecordRuleYaml.
type TmpTkeRecordRuleYamlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TmpTkeRecordRuleYamlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TmpTkeRecordRuleYaml is the Schema for the TmpTkeRecordRuleYamls API. Provides a resource to create a tke tmpRecordRule
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type TmpTkeRecordRuleYaml struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.content) || (has(self.initProvider) && has(self.initProvider.content))",message="spec.forProvider.content is a required parameter"
	Spec   TmpTkeRecordRuleYamlSpec   `json:"spec"`
	Status TmpTkeRecordRuleYamlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TmpTkeRecordRuleYamlList contains a list of TmpTkeRecordRuleYamls
type TmpTkeRecordRuleYamlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TmpTkeRecordRuleYaml `json:"items"`
}

// Repository type metadata.
var (
	TmpTkeRecordRuleYaml_Kind             = "TmpTkeRecordRuleYaml"
	TmpTkeRecordRuleYaml_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TmpTkeRecordRuleYaml_Kind}.String()
	TmpTkeRecordRuleYaml_KindAPIVersion   = TmpTkeRecordRuleYaml_Kind + "." + CRDGroupVersion.String()
	TmpTkeRecordRuleYaml_GroupVersionKind = CRDGroupVersion.WithKind(TmpTkeRecordRuleYaml_Kind)
)

func init() {
	SchemeBuilder.Register(&TmpTkeRecordRuleYaml{}, &TmpTkeRecordRuleYamlList{})
}
