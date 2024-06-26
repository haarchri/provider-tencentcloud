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

type IpStrategyInitParameters struct {

	// The ID of the API gateway service.
	// The ID of the API gateway service.
	// +crossplane:generate:reference:type=Service
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Reference to a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// Selector for a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`

	// IP address data.
	// IP address data.
	StrategyData *string `json:"strategyData,omitempty" tf:"strategy_data,omitempty"`

	// User defined strategy name.
	// User defined strategy name.
	StrategyName *string `json:"strategyName,omitempty" tf:"strategy_name,omitempty"`

	// Blacklist or whitelist.
	// Blacklist or whitelist.
	StrategyType *string `json:"strategyType,omitempty" tf:"strategy_type,omitempty"`
}

type IpStrategyObservation struct {

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the API gateway service.
	// The ID of the API gateway service.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// IP address data.
	// IP address data.
	StrategyData *string `json:"strategyData,omitempty" tf:"strategy_data,omitempty"`

	// IP policy ID.
	// IP policy ID.
	StrategyID *string `json:"strategyId,omitempty" tf:"strategy_id,omitempty"`

	// User defined strategy name.
	// User defined strategy name.
	StrategyName *string `json:"strategyName,omitempty" tf:"strategy_name,omitempty"`

	// Blacklist or whitelist.
	// Blacklist or whitelist.
	StrategyType *string `json:"strategyType,omitempty" tf:"strategy_type,omitempty"`
}

type IpStrategyParameters struct {

	// The ID of the API gateway service.
	// The ID of the API gateway service.
	// +crossplane:generate:reference:type=Service
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Reference to a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// Selector for a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`

	// IP address data.
	// IP address data.
	// +kubebuilder:validation:Optional
	StrategyData *string `json:"strategyData,omitempty" tf:"strategy_data,omitempty"`

	// User defined strategy name.
	// User defined strategy name.
	// +kubebuilder:validation:Optional
	StrategyName *string `json:"strategyName,omitempty" tf:"strategy_name,omitempty"`

	// Blacklist or whitelist.
	// Blacklist or whitelist.
	// +kubebuilder:validation:Optional
	StrategyType *string `json:"strategyType,omitempty" tf:"strategy_type,omitempty"`
}

// IpStrategySpec defines the desired state of IpStrategy
type IpStrategySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IpStrategyParameters `json:"forProvider"`
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
	InitProvider IpStrategyInitParameters `json:"initProvider,omitempty"`
}

// IpStrategyStatus defines the observed state of IpStrategy.
type IpStrategyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IpStrategyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IpStrategy is the Schema for the IpStrategys API. Use this resource to create IP strategy of API gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type IpStrategy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.strategyData) || (has(self.initProvider) && has(self.initProvider.strategyData))",message="spec.forProvider.strategyData is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.strategyName) || (has(self.initProvider) && has(self.initProvider.strategyName))",message="spec.forProvider.strategyName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.strategyType) || (has(self.initProvider) && has(self.initProvider.strategyType))",message="spec.forProvider.strategyType is a required parameter"
	Spec   IpStrategySpec   `json:"spec"`
	Status IpStrategyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpStrategyList contains a list of IpStrategys
type IpStrategyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpStrategy `json:"items"`
}

// Repository type metadata.
var (
	IpStrategy_Kind             = "IpStrategy"
	IpStrategy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IpStrategy_Kind}.String()
	IpStrategy_KindAPIVersion   = IpStrategy_Kind + "." + CRDGroupVersion.String()
	IpStrategy_GroupVersionKind = CRDGroupVersion.WithKind(IpStrategy_Kind)
)

func init() {
	SchemeBuilder.Register(&IpStrategy{}, &IpStrategyList{})
}
