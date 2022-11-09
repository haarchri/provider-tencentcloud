/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VPNCustomerGatewayObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VPNCustomerGatewayParameters struct {

	// Name of the customer gateway. The length of character is limited to 1-60.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Public IP of the customer gateway.
	// +kubebuilder:validation:Required
	PublicIPAddress *string `json:"publicIpAddress" tf:"public_ip_address,omitempty"`

	// A list of tags used to associate different resources.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPNCustomerGatewaySpec defines the desired state of VPNCustomerGateway
type VPNCustomerGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNCustomerGatewayParameters `json:"forProvider"`
}

// VPNCustomerGatewayStatus defines the observed state of VPNCustomerGateway.
type VPNCustomerGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNCustomerGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPNCustomerGateway is the Schema for the VPNCustomerGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type VPNCustomerGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPNCustomerGatewaySpec   `json:"spec"`
	Status            VPNCustomerGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNCustomerGatewayList contains a list of VPNCustomerGateways
type VPNCustomerGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNCustomerGateway `json:"items"`
}

// Repository type metadata.
var (
	VPNCustomerGateway_Kind             = "VPNCustomerGateway"
	VPNCustomerGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNCustomerGateway_Kind}.String()
	VPNCustomerGateway_KindAPIVersion   = VPNCustomerGateway_Kind + "." + CRDGroupVersion.String()
	VPNCustomerGateway_GroupVersionKind = CRDGroupVersion.WithKind(VPNCustomerGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNCustomerGateway{}, &VPNCustomerGatewayList{})
}
