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

type GrafanaInstanceInitParameters struct {

	// Whether to automatically use vouchers.
	// Whether to automatically use vouchers.
	AutoVoucher *bool `json:"autoVoucher,omitempty" tf:"auto_voucher,omitempty"`

	// Control whether grafana could be accessed by internet.
	// Control whether grafana could be accessed by internet.
	EnableInternet *bool `json:"enableInternet,omitempty" tf:"enable_internet,omitempty"`

	// Grafana server admin password.
	// Grafana server admin password.
	GrafanaInitPassword *string `json:"grafanaInitPassword,omitempty" tf:"grafana_init_password,omitempty"`

	// Instance name.
	// Instance name.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Whether to clean up completely, the default is false.
	// Whether to clean up completely, the default is false.
	IsDestroy *bool `json:"isDestroy,omitempty" tf:"is_destroy,omitempty"`

	// It has been deprecated from version 1.81.16. Whether to clean up completely, the default is false.
	// Whether to clean up completely, the default is false.
	IsDistroy *bool `json:"isDistroy,omitempty" tf:"is_distroy,omitempty"`

	// Subnet Id array.
	// Subnet Id array.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Tag description list.
	// Tag description list.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Vpc Id.
	// Vpc Id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type GrafanaInstanceObservation struct {

	// Whether to automatically use vouchers.
	// Whether to automatically use vouchers.
	AutoVoucher *bool `json:"autoVoucher,omitempty" tf:"auto_voucher,omitempty"`

	// Control whether grafana could be accessed by internet.
	// Control whether grafana could be accessed by internet.
	EnableInternet *bool `json:"enableInternet,omitempty" tf:"enable_internet,omitempty"`

	// Grafana server admin password.
	// Grafana server admin password.
	GrafanaInitPassword *string `json:"grafanaInitPassword,omitempty" tf:"grafana_init_password,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Grafana instance id.
	// Grafana instance id.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Instance name.
	// Instance name.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Grafana instance status, 1: Creating, 2: Running, 6: Stopped.
	// Grafana instance status, 1: Creating, 2: Running, 6: Stopped.
	InstanceStatus *float64 `json:"instanceStatus,omitempty" tf:"instance_status,omitempty"`

	// Grafana public address.
	// Grafana public address.
	InternalURL *string `json:"internalUrl,omitempty" tf:"internal_url,omitempty"`

	// Grafana intranet address.
	// Grafana intranet address.
	InternetURL *string `json:"internetUrl,omitempty" tf:"internet_url,omitempty"`

	// Whether to clean up completely, the default is false.
	// Whether to clean up completely, the default is false.
	IsDestroy *bool `json:"isDestroy,omitempty" tf:"is_destroy,omitempty"`

	// It has been deprecated from version 1.81.16. Whether to clean up completely, the default is false.
	// Whether to clean up completely, the default is false.
	IsDistroy *bool `json:"isDistroy,omitempty" tf:"is_distroy,omitempty"`

	// Grafana external url which could be accessed by user.
	// Grafana external url which could be accessed by user.
	RootURL *string `json:"rootUrl,omitempty" tf:"root_url,omitempty"`

	// Subnet Id array.
	// Subnet Id array.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Tag description list.
	// Tag description list.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Vpc Id.
	// Vpc Id.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type GrafanaInstanceParameters struct {

	// Whether to automatically use vouchers.
	// Whether to automatically use vouchers.
	// +kubebuilder:validation:Optional
	AutoVoucher *bool `json:"autoVoucher,omitempty" tf:"auto_voucher,omitempty"`

	// Control whether grafana could be accessed by internet.
	// Control whether grafana could be accessed by internet.
	// +kubebuilder:validation:Optional
	EnableInternet *bool `json:"enableInternet,omitempty" tf:"enable_internet,omitempty"`

	// Grafana server admin password.
	// Grafana server admin password.
	// +kubebuilder:validation:Optional
	GrafanaInitPassword *string `json:"grafanaInitPassword,omitempty" tf:"grafana_init_password,omitempty"`

	// Instance name.
	// Instance name.
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Whether to clean up completely, the default is false.
	// Whether to clean up completely, the default is false.
	// +kubebuilder:validation:Optional
	IsDestroy *bool `json:"isDestroy,omitempty" tf:"is_destroy,omitempty"`

	// It has been deprecated from version 1.81.16. Whether to clean up completely, the default is false.
	// Whether to clean up completely, the default is false.
	// +kubebuilder:validation:Optional
	IsDistroy *bool `json:"isDistroy,omitempty" tf:"is_distroy,omitempty"`

	// Subnet Id array.
	// Subnet Id array.
	// +kubebuilder:validation:Optional
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Tag description list.
	// Tag description list.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Vpc Id.
	// Vpc Id.
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

// GrafanaInstanceSpec defines the desired state of GrafanaInstance
type GrafanaInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GrafanaInstanceParameters `json:"forProvider"`
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
	InitProvider GrafanaInstanceInitParameters `json:"initProvider,omitempty"`
}

// GrafanaInstanceStatus defines the observed state of GrafanaInstance.
type GrafanaInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GrafanaInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GrafanaInstance is the Schema for the GrafanaInstances API. Provides a resource to create a monitor grafanaInstance
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type GrafanaInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceName) || (has(self.initProvider) && has(self.initProvider.instanceName))",message="spec.forProvider.instanceName is a required parameter"
	Spec   GrafanaInstanceSpec   `json:"spec"`
	Status GrafanaInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GrafanaInstanceList contains a list of GrafanaInstances
type GrafanaInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GrafanaInstance `json:"items"`
}

// Repository type metadata.
var (
	GrafanaInstance_Kind             = "GrafanaInstance"
	GrafanaInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GrafanaInstance_Kind}.String()
	GrafanaInstance_KindAPIVersion   = GrafanaInstance_Kind + "." + CRDGroupVersion.String()
	GrafanaInstance_GroupVersionKind = CRDGroupVersion.WithKind(GrafanaInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&GrafanaInstance{}, &GrafanaInstanceList{})
}
