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

type ClusterInitParameters struct {

	// Name of the TcaplusDB cluster. Name length should be between 1 and 30.
	// Name of the TcaplusDB cluster. Name length should be between 1 and 30.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// IDL type of the TcaplusDB cluster. Valid values: PROTO and TDR.
	// IDL type of the TcaplusDB cluster. Valid values: `PROTO` and `TDR`.
	IdlType *string `json:"idlType,omitempty" tf:"idl_type,omitempty"`

	// Expiration time of old password after password update, unit: second.
	// Expiration time of old password after password update, unit: second.
	OldPasswordExpireLast *float64 `json:"oldPasswordExpireLast,omitempty" tf:"old_password_expire_last,omitempty"`

	// Subnet id of the TcaplusDB cluster.
	// Subnet id of the TcaplusDB cluster.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// VPC id of the TcaplusDB cluster.
	// VPC id of the TcaplusDB cluster.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type ClusterObservation struct {

	// Access ID of the TcaplusDB cluster.For TcaplusDB SDK connect.
	// Access ID of the TcaplusDB cluster.For TcaplusDB SDK connect.
	APIAccessID *string `json:"apiAccessId,omitempty" tf:"api_access_id,omitempty"`

	// Access IP of the TcaplusDB cluster.For TcaplusDB SDK connect.
	// Access IP of the TcaplusDB cluster.For TcaplusDB SDK connect.
	APIAccessIP *string `json:"apiAccessIp,omitempty" tf:"api_access_ip,omitempty"`

	// Access port of the TcaplusDB cluster.For TcaplusDB SDK connect.
	// Access port of the TcaplusDB cluster.For TcaplusDB SDK connect.
	APIAccessPort *float64 `json:"apiAccessPort,omitempty" tf:"api_access_port,omitempty"`

	// Name of the TcaplusDB cluster. Name length should be between 1 and 30.
	// Name of the TcaplusDB cluster. Name length should be between 1 and 30.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Create time of the TcaplusDB cluster.
	// Create time of the TcaplusDB cluster.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IDL type of the TcaplusDB cluster. Valid values: PROTO and TDR.
	// IDL type of the TcaplusDB cluster. Valid values: `PROTO` and `TDR`.
	IdlType *string `json:"idlType,omitempty" tf:"idl_type,omitempty"`

	// Network type of the TcaplusDB cluster.
	// Network type of the TcaplusDB cluster.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// Expiration time of old password after password update, unit: second.
	// Expiration time of old password after password update, unit: second.
	OldPasswordExpireLast *float64 `json:"oldPasswordExpireLast,omitempty" tf:"old_password_expire_last,omitempty"`

	// Expiration time of the old password. If password_status is unmodifiable, it means the old password has not yet expired.
	// Expiration time of the old password. If `password_status` is `unmodifiable`, it means the old password has not yet expired.
	OldPasswordExpireTime *string `json:"oldPasswordExpireTime,omitempty" tf:"old_password_expire_time,omitempty"`

	// Password status of the TcaplusDB cluster. Valid values: unmodifiable, modifiable. unmodifiable. which means the password can not be changed in this moment; modifiable, which means the password can be changed in this moment.
	// Password status of the TcaplusDB cluster. Valid values: `unmodifiable`, `modifiable`. `unmodifiable`. which means the password can not be changed in this moment; `modifiable`, which means the password can be changed in this moment.
	PasswordStatus *string `json:"passwordStatus,omitempty" tf:"password_status,omitempty"`

	// Subnet id of the TcaplusDB cluster.
	// Subnet id of the TcaplusDB cluster.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// VPC id of the TcaplusDB cluster.
	// VPC id of the TcaplusDB cluster.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type ClusterParameters struct {

	// Name of the TcaplusDB cluster. Name length should be between 1 and 30.
	// Name of the TcaplusDB cluster. Name length should be between 1 and 30.
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// IDL type of the TcaplusDB cluster. Valid values: PROTO and TDR.
	// IDL type of the TcaplusDB cluster. Valid values: `PROTO` and `TDR`.
	// +kubebuilder:validation:Optional
	IdlType *string `json:"idlType,omitempty" tf:"idl_type,omitempty"`

	// Expiration time of old password after password update, unit: second.
	// Expiration time of old password after password update, unit: second.
	// +kubebuilder:validation:Optional
	OldPasswordExpireLast *float64 `json:"oldPasswordExpireLast,omitempty" tf:"old_password_expire_last,omitempty"`

	// Password of the TcaplusDB cluster. Password length should be between 12 and 16. The password must be a mix of uppercase letters (A-Z), lowercase letters (a-z) and numbers (0-9).
	// Password of the TcaplusDB cluster. Password length should be between 12 and 16. The password must be a *mix* of uppercase letters (A-Z), lowercase *letters* (a-z) and *numbers* (0-9).
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Subnet id of the TcaplusDB cluster.
	// Subnet id of the TcaplusDB cluster.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// VPC id of the TcaplusDB cluster.
	// VPC id of the TcaplusDB cluster.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
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
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Cluster is the Schema for the Clusters API. Use this resource to create TcaplusDB cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clusterName) || (has(self.initProvider) && has(self.initProvider.clusterName))",message="spec.forProvider.clusterName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.idlType) || (has(self.initProvider) && has(self.initProvider.idlType))",message="spec.forProvider.idlType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	Spec   ClusterSpec   `json:"spec"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
