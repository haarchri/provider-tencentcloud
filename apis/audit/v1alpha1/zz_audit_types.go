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

type AuditInitParameters struct {

	// Indicate whether to turn on audit logging or not.
	// Indicate whether to turn on audit logging or not.
	AuditSwitch *bool `json:"auditSwitch,omitempty" tf:"audit_switch,omitempty"`

	// Name of the cos bucket to save audit log.
	// Name of the cos bucket to save audit log.
	CosBucket *string `json:"cosBucket,omitempty" tf:"cos_bucket,omitempty"`

	// Region of the cos bucket.
	// Region of the cos bucket.
	CosRegion *string `json:"cosRegion,omitempty" tf:"cos_region,omitempty"`

	// Indicate whether the log is encrypt with KMS algorithm or not.
	// Indicate whether the log is encrypt with KMS algorithm or not.
	EnableKMSEncry *bool `json:"enableKmsEncry,omitempty" tf:"enable_kms_encry,omitempty"`

	// Existing CMK unique key. This field can be get by data source tencentcloud_audit_key_alias. Caution: the region of the KMS must be as same as the cos_region.
	// Existing CMK unique key. This field can be get by data source `tencentcloud_audit_key_alias`. Caution: the region of the KMS must be as same as the `cos_region`.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The log file name prefix. The length ranges from 3 to 40. If not set, the account ID will be the log file prefix.
	// The log file name prefix. The length ranges from 3 to 40. If not set, the account ID will be the log file prefix.
	LogFilePrefix *string `json:"logFilePrefix,omitempty" tf:"log_file_prefix,omitempty"`

	// Name of audit. Valid length ranges from 3 to 128. Only alpha character or numbers or '_' supported.
	// Name of audit. Valid length ranges from 3 to 128. Only alpha character or numbers or '_' supported.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Event attribute filter. Valid values: 1, 2, 3. 1 for readonly, 2 for write-only, 3 for all.
	// Event attribute filter. Valid values: `1`, `2`, `3`. `1` for readonly, `2` for write-only, `3` for all.
	ReadWriteAttribute *float64 `json:"readWriteAttribute,omitempty" tf:"read_write_attribute,omitempty"`
}

type AuditObservation struct {

	// Indicate whether to turn on audit logging or not.
	// Indicate whether to turn on audit logging or not.
	AuditSwitch *bool `json:"auditSwitch,omitempty" tf:"audit_switch,omitempty"`

	// Name of the cos bucket to save audit log.
	// Name of the cos bucket to save audit log.
	CosBucket *string `json:"cosBucket,omitempty" tf:"cos_bucket,omitempty"`

	// Region of the cos bucket.
	// Region of the cos bucket.
	CosRegion *string `json:"cosRegion,omitempty" tf:"cos_region,omitempty"`

	// Indicate whether the log is encrypt with KMS algorithm or not.
	// Indicate whether the log is encrypt with KMS algorithm or not.
	EnableKMSEncry *bool `json:"enableKmsEncry,omitempty" tf:"enable_kms_encry,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Existing CMK unique key. This field can be get by data source tencentcloud_audit_key_alias. Caution: the region of the KMS must be as same as the cos_region.
	// Existing CMK unique key. This field can be get by data source `tencentcloud_audit_key_alias`. Caution: the region of the KMS must be as same as the `cos_region`.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The log file name prefix. The length ranges from 3 to 40. If not set, the account ID will be the log file prefix.
	// The log file name prefix. The length ranges from 3 to 40. If not set, the account ID will be the log file prefix.
	LogFilePrefix *string `json:"logFilePrefix,omitempty" tf:"log_file_prefix,omitempty"`

	// Name of audit. Valid length ranges from 3 to 128. Only alpha character or numbers or '_' supported.
	// Name of audit. Valid length ranges from 3 to 128. Only alpha character or numbers or '_' supported.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Event attribute filter. Valid values: 1, 2, 3. 1 for readonly, 2 for write-only, 3 for all.
	// Event attribute filter. Valid values: `1`, `2`, `3`. `1` for readonly, `2` for write-only, `3` for all.
	ReadWriteAttribute *float64 `json:"readWriteAttribute,omitempty" tf:"read_write_attribute,omitempty"`
}

type AuditParameters struct {

	// Indicate whether to turn on audit logging or not.
	// Indicate whether to turn on audit logging or not.
	// +kubebuilder:validation:Optional
	AuditSwitch *bool `json:"auditSwitch,omitempty" tf:"audit_switch,omitempty"`

	// Name of the cos bucket to save audit log.
	// Name of the cos bucket to save audit log.
	// +kubebuilder:validation:Optional
	CosBucket *string `json:"cosBucket,omitempty" tf:"cos_bucket,omitempty"`

	// Region of the cos bucket.
	// Region of the cos bucket.
	// +kubebuilder:validation:Optional
	CosRegion *string `json:"cosRegion,omitempty" tf:"cos_region,omitempty"`

	// Indicate whether the log is encrypt with KMS algorithm or not.
	// Indicate whether the log is encrypt with KMS algorithm or not.
	// +kubebuilder:validation:Optional
	EnableKMSEncry *bool `json:"enableKmsEncry,omitempty" tf:"enable_kms_encry,omitempty"`

	// Existing CMK unique key. This field can be get by data source tencentcloud_audit_key_alias. Caution: the region of the KMS must be as same as the cos_region.
	// Existing CMK unique key. This field can be get by data source `tencentcloud_audit_key_alias`. Caution: the region of the KMS must be as same as the `cos_region`.
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The log file name prefix. The length ranges from 3 to 40. If not set, the account ID will be the log file prefix.
	// The log file name prefix. The length ranges from 3 to 40. If not set, the account ID will be the log file prefix.
	// +kubebuilder:validation:Optional
	LogFilePrefix *string `json:"logFilePrefix,omitempty" tf:"log_file_prefix,omitempty"`

	// Name of audit. Valid length ranges from 3 to 128. Only alpha character or numbers or '_' supported.
	// Name of audit. Valid length ranges from 3 to 128. Only alpha character or numbers or '_' supported.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Event attribute filter. Valid values: 1, 2, 3. 1 for readonly, 2 for write-only, 3 for all.
	// Event attribute filter. Valid values: `1`, `2`, `3`. `1` for readonly, `2` for write-only, `3` for all.
	// +kubebuilder:validation:Optional
	ReadWriteAttribute *float64 `json:"readWriteAttribute,omitempty" tf:"read_write_attribute,omitempty"`
}

// AuditSpec defines the desired state of Audit
type AuditSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuditParameters `json:"forProvider"`
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
	InitProvider AuditInitParameters `json:"initProvider,omitempty"`
}

// AuditStatus defines the observed state of Audit.
type AuditStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuditObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Audit is the Schema for the Audits API. Provides a resource to create an audit.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Audit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.auditSwitch) || (has(self.initProvider) && has(self.initProvider.auditSwitch))",message="spec.forProvider.auditSwitch is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cosBucket) || (has(self.initProvider) && has(self.initProvider.cosBucket))",message="spec.forProvider.cosBucket is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cosRegion) || (has(self.initProvider) && has(self.initProvider.cosRegion))",message="spec.forProvider.cosRegion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.readWriteAttribute) || (has(self.initProvider) && has(self.initProvider.readWriteAttribute))",message="spec.forProvider.readWriteAttribute is a required parameter"
	Spec   AuditSpec   `json:"spec"`
	Status AuditStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuditList contains a list of Audits
type AuditList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Audit `json:"items"`
}

// Repository type metadata.
var (
	Audit_Kind             = "Audit"
	Audit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Audit_Kind}.String()
	Audit_KindAPIVersion   = Audit_Kind + "." + CRDGroupVersion.String()
	Audit_GroupVersionKind = CRDGroupVersion.WithKind(Audit_Kind)
)

func init() {
	SchemeBuilder.Register(&Audit{}, &AuditList{})
}
