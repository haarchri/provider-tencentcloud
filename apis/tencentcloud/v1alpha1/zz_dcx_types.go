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

type DcxInitParameters struct {

	// BGP ASN of the user. A required field within BGP.
	// BGP ASN of the user. A required field within BGP.
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// BGP key of the user.
	// BGP key of the user.
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// Bandwidth of the DC.
	// Bandwidth of the DC.
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// Interconnect IP of the DC within client.
	// Interconnect IP of the DC within client.
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// ID of the DC to be queried, application deployment offline.
	// ID of the DC to be queried, application deployment offline.
	DcID *string `json:"dcId,omitempty" tf:"dc_id,omitempty"`

	// Connection owner, who is the current customer by default. The developer account ID should be entered for shared connections.
	// Connection owner, who is the current customer by default. The developer account ID should be entered for shared connections.
	DcOwnerAccount *string `json:"dcOwnerAccount,omitempty" tf:"dc_owner_account,omitempty"`

	// ID of the DC Gateway. Currently only new in the console.
	// ID of the DC Gateway. Currently only new in the console.
	DcgID *string `json:"dcgId,omitempty" tf:"dcg_id,omitempty"`

	// Name of the dedicated tunnel.
	// Name of the dedicated tunnel.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network region.
	// Network region.
	NetworkRegion *string `json:"networkRegion,omitempty" tf:"network_region,omitempty"`

	// Type of the network. Valid value: VPC, BMVPC and CCN. The default value is VPC.
	// Type of the network. Valid value: `VPC`, `BMVPC` and `CCN`. The default value is `VPC`.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// Static route, the network address of the user IDC. It can be modified after setting but cannot be deleted. AN unable field within BGP.
	// Static route, the network address of the user IDC. It can be modified after setting but cannot be deleted. AN unable field within BGP.
	// +listType=set
	RouteFilterPrefixes []*string `json:"routeFilterPrefixes,omitempty" tf:"route_filter_prefixes,omitempty"`

	// Type of the route, and available values include BGP and STATIC. The default value is BGP.
	// Type of the route, and available values include BGP and STATIC. The default value is `BGP`.
	RouteType *string `json:"routeType,omitempty" tf:"route_type,omitempty"`

	// Interconnect IP of the DC within Tencent.
	// Interconnect IP of the DC within Tencent.
	TencentAddress *string `json:"tencentAddress,omitempty" tf:"tencent_address,omitempty"`

	// ID of the VPC or BMVPC.
	// ID of the VPC or BMVPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Vlan of the dedicated tunnels. Valid value ranges: (0~3000). 0 means that only one tunnel can be created for the physical connect.
	// Vlan of the dedicated tunnels. Valid value ranges: (0~3000). `0` means that only one tunnel can be created for the physical connect.
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

type DcxObservation struct {

	// BGP ASN of the user. A required field within BGP.
	// BGP ASN of the user. A required field within BGP.
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// BGP key of the user.
	// BGP key of the user.
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// Bandwidth of the DC.
	// Bandwidth of the DC.
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// Creation time of resource.
	// Creation time of resource.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Interconnect IP of the DC within client.
	// Interconnect IP of the DC within client.
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// ID of the DC to be queried, application deployment offline.
	// ID of the DC to be queried, application deployment offline.
	DcID *string `json:"dcId,omitempty" tf:"dc_id,omitempty"`

	// Connection owner, who is the current customer by default. The developer account ID should be entered for shared connections.
	// Connection owner, who is the current customer by default. The developer account ID should be entered for shared connections.
	DcOwnerAccount *string `json:"dcOwnerAccount,omitempty" tf:"dc_owner_account,omitempty"`

	// ID of the DC Gateway. Currently only new in the console.
	// ID of the DC Gateway. Currently only new in the console.
	DcgID *string `json:"dcgId,omitempty" tf:"dcg_id,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the dedicated tunnel.
	// Name of the dedicated tunnel.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network region.
	// Network region.
	NetworkRegion *string `json:"networkRegion,omitempty" tf:"network_region,omitempty"`

	// Type of the network. Valid value: VPC, BMVPC and CCN. The default value is VPC.
	// Type of the network. Valid value: `VPC`, `BMVPC` and `CCN`. The default value is `VPC`.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// Static route, the network address of the user IDC. It can be modified after setting but cannot be deleted. AN unable field within BGP.
	// Static route, the network address of the user IDC. It can be modified after setting but cannot be deleted. AN unable field within BGP.
	// +listType=set
	RouteFilterPrefixes []*string `json:"routeFilterPrefixes,omitempty" tf:"route_filter_prefixes,omitempty"`

	// Type of the route, and available values include BGP and STATIC. The default value is BGP.
	// Type of the route, and available values include BGP and STATIC. The default value is `BGP`.
	RouteType *string `json:"routeType,omitempty" tf:"route_type,omitempty"`

	// State of the dedicated tunnels. Valid value: PENDING, ALLOCATING, ALLOCATED, ALTERING, DELETING, DELETED, COMFIRMING and REJECTED.
	// State of the dedicated tunnels. Valid value: `PENDING`, `ALLOCATING`, `ALLOCATED`, `ALTERING`, `DELETING`, `DELETED`, `COMFIRMING` and `REJECTED`.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Interconnect IP of the DC within Tencent.
	// Interconnect IP of the DC within Tencent.
	TencentAddress *string `json:"tencentAddress,omitempty" tf:"tencent_address,omitempty"`

	// ID of the VPC or BMVPC.
	// ID of the VPC or BMVPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Vlan of the dedicated tunnels. Valid value ranges: (0~3000). 0 means that only one tunnel can be created for the physical connect.
	// Vlan of the dedicated tunnels. Valid value ranges: (0~3000). `0` means that only one tunnel can be created for the physical connect.
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

type DcxParameters struct {

	// BGP ASN of the user. A required field within BGP.
	// BGP ASN of the user. A required field within BGP.
	// +kubebuilder:validation:Optional
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// BGP key of the user.
	// BGP key of the user.
	// +kubebuilder:validation:Optional
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// Bandwidth of the DC.
	// Bandwidth of the DC.
	// +kubebuilder:validation:Optional
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// Interconnect IP of the DC within client.
	// Interconnect IP of the DC within client.
	// +kubebuilder:validation:Optional
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// ID of the DC to be queried, application deployment offline.
	// ID of the DC to be queried, application deployment offline.
	// +kubebuilder:validation:Optional
	DcID *string `json:"dcId,omitempty" tf:"dc_id,omitempty"`

	// Connection owner, who is the current customer by default. The developer account ID should be entered for shared connections.
	// Connection owner, who is the current customer by default. The developer account ID should be entered for shared connections.
	// +kubebuilder:validation:Optional
	DcOwnerAccount *string `json:"dcOwnerAccount,omitempty" tf:"dc_owner_account,omitempty"`

	// ID of the DC Gateway. Currently only new in the console.
	// ID of the DC Gateway. Currently only new in the console.
	// +kubebuilder:validation:Optional
	DcgID *string `json:"dcgId,omitempty" tf:"dcg_id,omitempty"`

	// Name of the dedicated tunnel.
	// Name of the dedicated tunnel.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network region.
	// Network region.
	// +kubebuilder:validation:Optional
	NetworkRegion *string `json:"networkRegion,omitempty" tf:"network_region,omitempty"`

	// Type of the network. Valid value: VPC, BMVPC and CCN. The default value is VPC.
	// Type of the network. Valid value: `VPC`, `BMVPC` and `CCN`. The default value is `VPC`.
	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// Static route, the network address of the user IDC. It can be modified after setting but cannot be deleted. AN unable field within BGP.
	// Static route, the network address of the user IDC. It can be modified after setting but cannot be deleted. AN unable field within BGP.
	// +kubebuilder:validation:Optional
	// +listType=set
	RouteFilterPrefixes []*string `json:"routeFilterPrefixes,omitempty" tf:"route_filter_prefixes,omitempty"`

	// Type of the route, and available values include BGP and STATIC. The default value is BGP.
	// Type of the route, and available values include BGP and STATIC. The default value is `BGP`.
	// +kubebuilder:validation:Optional
	RouteType *string `json:"routeType,omitempty" tf:"route_type,omitempty"`

	// Interconnect IP of the DC within Tencent.
	// Interconnect IP of the DC within Tencent.
	// +kubebuilder:validation:Optional
	TencentAddress *string `json:"tencentAddress,omitempty" tf:"tencent_address,omitempty"`

	// ID of the VPC or BMVPC.
	// ID of the VPC or BMVPC.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Vlan of the dedicated tunnels. Valid value ranges: (0~3000). 0 means that only one tunnel can be created for the physical connect.
	// Vlan of the dedicated tunnels. Valid value ranges: (0~3000). `0` means that only one tunnel can be created for the physical connect.
	// +kubebuilder:validation:Optional
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

// DcxSpec defines the desired state of Dcx
type DcxSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DcxParameters `json:"forProvider"`
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
	InitProvider DcxInitParameters `json:"initProvider,omitempty"`
}

// DcxStatus defines the observed state of Dcx.
type DcxStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DcxObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Dcx is the Schema for the Dcxs API. Provides a resource to creating dedicated tunnels instances.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Dcx struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dcId) || (has(self.initProvider) && has(self.initProvider.dcId))",message="spec.forProvider.dcId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dcgId) || (has(self.initProvider) && has(self.initProvider.dcgId))",message="spec.forProvider.dcgId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   DcxSpec   `json:"spec"`
	Status DcxStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DcxList contains a list of Dcxs
type DcxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dcx `json:"items"`
}

// Repository type metadata.
var (
	Dcx_Kind             = "Dcx"
	Dcx_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dcx_Kind}.String()
	Dcx_KindAPIVersion   = Dcx_Kind + "." + CRDGroupVersion.String()
	Dcx_GroupVersionKind = CRDGroupVersion.WithKind(Dcx_Kind)
)

func init() {
	SchemeBuilder.Register(&Dcx{}, &DcxList{})
}
