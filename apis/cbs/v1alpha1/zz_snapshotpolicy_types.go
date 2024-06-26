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

type SnapshotPolicyInitParameters struct {

	// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
	// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
	RepeatHours []*float64 `json:"repeatHours,omitempty" tf:"repeat_hours,omitempty"`

	// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
	// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
	RepeatWeekdays []*float64 `json:"repeatWeekdays,omitempty" tf:"repeat_weekdays,omitempty"`

	// Retention days of the snapshot, and the default value is 7.
	// Retention days of the snapshot, and the default value is 7.
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// Name of snapshot policy. The maximum length can not exceed 60 bytes.
	// Name of snapshot policy. The maximum length can not exceed 60 bytes.
	SnapshotPolicyName *string `json:"snapshotPolicyName,omitempty" tf:"snapshot_policy_name,omitempty"`
}

type SnapshotPolicyObservation struct {

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
	// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
	RepeatHours []*float64 `json:"repeatHours,omitempty" tf:"repeat_hours,omitempty"`

	// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
	// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
	RepeatWeekdays []*float64 `json:"repeatWeekdays,omitempty" tf:"repeat_weekdays,omitempty"`

	// Retention days of the snapshot, and the default value is 7.
	// Retention days of the snapshot, and the default value is 7.
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// Name of snapshot policy. The maximum length can not exceed 60 bytes.
	// Name of snapshot policy. The maximum length can not exceed 60 bytes.
	SnapshotPolicyName *string `json:"snapshotPolicyName,omitempty" tf:"snapshot_policy_name,omitempty"`
}

type SnapshotPolicyParameters struct {

	// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
	// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
	// +kubebuilder:validation:Optional
	RepeatHours []*float64 `json:"repeatHours,omitempty" tf:"repeat_hours,omitempty"`

	// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
	// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
	// +kubebuilder:validation:Optional
	RepeatWeekdays []*float64 `json:"repeatWeekdays,omitempty" tf:"repeat_weekdays,omitempty"`

	// Retention days of the snapshot, and the default value is 7.
	// Retention days of the snapshot, and the default value is 7.
	// +kubebuilder:validation:Optional
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// Name of snapshot policy. The maximum length can not exceed 60 bytes.
	// Name of snapshot policy. The maximum length can not exceed 60 bytes.
	// +kubebuilder:validation:Optional
	SnapshotPolicyName *string `json:"snapshotPolicyName,omitempty" tf:"snapshot_policy_name,omitempty"`
}

// SnapshotPolicySpec defines the desired state of SnapshotPolicy
type SnapshotPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotPolicyParameters `json:"forProvider"`
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
	InitProvider SnapshotPolicyInitParameters `json:"initProvider,omitempty"`
}

// SnapshotPolicyStatus defines the observed state of SnapshotPolicy.
type SnapshotPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SnapshotPolicy is the Schema for the SnapshotPolicys API. Provides a snapshot policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type SnapshotPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.repeatHours) || (has(self.initProvider) && has(self.initProvider.repeatHours))",message="spec.forProvider.repeatHours is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.repeatWeekdays) || (has(self.initProvider) && has(self.initProvider.repeatWeekdays))",message="spec.forProvider.repeatWeekdays is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.snapshotPolicyName) || (has(self.initProvider) && has(self.initProvider.snapshotPolicyName))",message="spec.forProvider.snapshotPolicyName is a required parameter"
	Spec   SnapshotPolicySpec   `json:"spec"`
	Status SnapshotPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotPolicyList contains a list of SnapshotPolicys
type SnapshotPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotPolicy `json:"items"`
}

// Repository type metadata.
var (
	SnapshotPolicy_Kind             = "SnapshotPolicy"
	SnapshotPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnapshotPolicy_Kind}.String()
	SnapshotPolicy_KindAPIVersion   = SnapshotPolicy_Kind + "." + CRDGroupVersion.String()
	SnapshotPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SnapshotPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SnapshotPolicy{}, &SnapshotPolicyList{})
}
