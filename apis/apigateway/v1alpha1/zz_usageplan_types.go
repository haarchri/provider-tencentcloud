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

type AttachListInitParameters struct {
}

type AttachListObservation struct {

	// The API ID, this value is empty if attach service.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// The API name, this value is empty if attach service.
	APIName *string `json:"apiName,omitempty" tf:"api_name,omitempty"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The environment name.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The API method, this value is empty if attach service.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// The API path, this value is empty if attach service.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The service ID.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// The service name.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type AttachListParameters struct {
}

type UsagePlanInitParameters struct {

	// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.
	// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.
	MaxRequestNum *float64 `json:"maxRequestNum,omitempty" tf:"max_request_num,omitempty"`

	// Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.
	// Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.
	MaxRequestNumPreSec *float64 `json:"maxRequestNumPreSec,omitempty" tf:"max_request_num_pre_sec,omitempty"`

	// Custom usage plan description.
	// Custom usage plan description.
	UsagePlanDesc *string `json:"usagePlanDesc,omitempty" tf:"usage_plan_desc,omitempty"`

	// Custom usage plan name.
	// Custom usage plan name.
	UsagePlanName *string `json:"usagePlanName,omitempty" tf:"usage_plan_name,omitempty"`
}

type UsagePlanObservation struct {

	// Attach API keys list.
	// Attach API keys list.
	AttachAPIKeys []*string `json:"attachApiKeys,omitempty" tf:"attach_api_keys,omitempty"`

	// Attach service and API list.
	// Attach service and API list.
	AttachList []AttachListObservation `json:"attachList,omitempty" tf:"attach_list,omitempty"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.
	// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.
	MaxRequestNum *float64 `json:"maxRequestNum,omitempty" tf:"max_request_num,omitempty"`

	// Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.
	// Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.
	MaxRequestNumPreSec *float64 `json:"maxRequestNumPreSec,omitempty" tf:"max_request_num_pre_sec,omitempty"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// Custom usage plan description.
	// Custom usage plan description.
	UsagePlanDesc *string `json:"usagePlanDesc,omitempty" tf:"usage_plan_desc,omitempty"`

	// Custom usage plan name.
	// Custom usage plan name.
	UsagePlanName *string `json:"usagePlanName,omitempty" tf:"usage_plan_name,omitempty"`
}

type UsagePlanParameters struct {

	// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.
	// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.
	// +kubebuilder:validation:Optional
	MaxRequestNum *float64 `json:"maxRequestNum,omitempty" tf:"max_request_num,omitempty"`

	// Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.
	// Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.
	// +kubebuilder:validation:Optional
	MaxRequestNumPreSec *float64 `json:"maxRequestNumPreSec,omitempty" tf:"max_request_num_pre_sec,omitempty"`

	// Custom usage plan description.
	// Custom usage plan description.
	// +kubebuilder:validation:Optional
	UsagePlanDesc *string `json:"usagePlanDesc,omitempty" tf:"usage_plan_desc,omitempty"`

	// Custom usage plan name.
	// Custom usage plan name.
	// +kubebuilder:validation:Optional
	UsagePlanName *string `json:"usagePlanName,omitempty" tf:"usage_plan_name,omitempty"`
}

// UsagePlanSpec defines the desired state of UsagePlan
type UsagePlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UsagePlanParameters `json:"forProvider"`
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
	InitProvider UsagePlanInitParameters `json:"initProvider,omitempty"`
}

// UsagePlanStatus defines the observed state of UsagePlan.
type UsagePlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UsagePlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UsagePlan is the Schema for the UsagePlans API. Use this resource to create API gateway usage plan.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type UsagePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.usagePlanName) || (has(self.initProvider) && has(self.initProvider.usagePlanName))",message="spec.forProvider.usagePlanName is a required parameter"
	Spec   UsagePlanSpec   `json:"spec"`
	Status UsagePlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UsagePlanList contains a list of UsagePlans
type UsagePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UsagePlan `json:"items"`
}

// Repository type metadata.
var (
	UsagePlan_Kind             = "UsagePlan"
	UsagePlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UsagePlan_Kind}.String()
	UsagePlan_KindAPIVersion   = UsagePlan_Kind + "." + CRDGroupVersion.String()
	UsagePlan_GroupVersionKind = CRDGroupVersion.WithKind(UsagePlan_Kind)
)

func init() {
	SchemeBuilder.Register(&UsagePlan{}, &UsagePlanList{})
}
