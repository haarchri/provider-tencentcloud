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

type SecurityGroupLiteRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecurityGroupLiteRuleParameters struct {

	// Egress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP`, `ALL` and `ppm(g?)-xxxxxxxx`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
	// +kubebuilder:validation:Optional
	Egress []*string `json:"egress,omitempty" tf:"egress,omitempty"`

	// Ingress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP`, `ALL` and `ppm(g?)-xxxxxxxx`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
	// +kubebuilder:validation:Optional
	Ingress []*string `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// ID of the security group.
	// +crossplane:generate:reference:type=SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`
}

// SecurityGroupLiteRuleSpec defines the desired state of SecurityGroupLiteRule
type SecurityGroupLiteRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupLiteRuleParameters `json:"forProvider"`
}

// SecurityGroupLiteRuleStatus defines the observed state of SecurityGroupLiteRule.
type SecurityGroupLiteRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupLiteRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupLiteRule is the Schema for the SecurityGroupLiteRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type SecurityGroupLiteRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupLiteRuleSpec   `json:"spec"`
	Status            SecurityGroupLiteRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupLiteRuleList contains a list of SecurityGroupLiteRules
type SecurityGroupLiteRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroupLiteRule `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroupLiteRule_Kind             = "SecurityGroupLiteRule"
	SecurityGroupLiteRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroupLiteRule_Kind}.String()
	SecurityGroupLiteRule_KindAPIVersion   = SecurityGroupLiteRule_Kind + "." + CRDGroupVersion.String()
	SecurityGroupLiteRule_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroupLiteRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroupLiteRule{}, &SecurityGroupLiteRuleList{})
}