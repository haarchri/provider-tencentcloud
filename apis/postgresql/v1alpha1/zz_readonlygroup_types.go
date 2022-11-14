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

type ReadonlyGroupObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ReadonlyGroupParameters struct {

	// Primary instance ID.
	// +kubebuilder:validation:Required
	MasterDBInstanceID *string `json:"masterDbInstanceId" tf:"master_db_instance_id,omitempty"`

	// Delay threshold in ms.
	// +kubebuilder:validation:Required
	MaxReplayLag *float64 `json:"maxReplayLag" tf:"max_replay_lag,omitempty"`

	// Delayed log size threshold in MB.
	// +kubebuilder:validation:Required
	MaxReplayLatency *float64 `json:"maxReplayLatency" tf:"max_replay_latency,omitempty"`

	// The minimum number of read-only replicas that must be retained in an RO group.
	// +kubebuilder:validation:Required
	MinDelayEliminateReserve *float64 `json:"minDelayEliminateReserve" tf:"min_delay_eliminate_reserve,omitempty"`

	// RO group name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Project ID.
	// +kubebuilder:validation:Required
	ProjectID *float64 `json:"projectId" tf:"project_id,omitempty"`

	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	// +kubebuilder:validation:Required
	ReplayLagEliminate *float64 `json:"replayLagEliminate" tf:"replay_lag_eliminate,omitempty"`

	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	// +kubebuilder:validation:Required
	ReplayLatencyEliminate *float64 `json:"replayLatencyEliminate" tf:"replay_latency_eliminate,omitempty"`

	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// +kubebuilder:validation:Optional
	SecurityGroupsIds []*string `json:"securityGroupsIds,omitempty" tf:"security_groups_ids,omitempty"`

	// VPC subnet ID.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// VPC ID.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcidSelector,omitempty" tf:"-"`
}

// ReadonlyGroupSpec defines the desired state of ReadonlyGroup
type ReadonlyGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReadonlyGroupParameters `json:"forProvider"`
}

// ReadonlyGroupStatus defines the observed state of ReadonlyGroup.
type ReadonlyGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReadonlyGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReadonlyGroup is the Schema for the ReadonlyGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type ReadonlyGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReadonlyGroupSpec   `json:"spec"`
	Status            ReadonlyGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReadonlyGroupList contains a list of ReadonlyGroups
type ReadonlyGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReadonlyGroup `json:"items"`
}

// Repository type metadata.
var (
	ReadonlyGroup_Kind             = "ReadonlyGroup"
	ReadonlyGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReadonlyGroup_Kind}.String()
	ReadonlyGroup_KindAPIVersion   = ReadonlyGroup_Kind + "." + CRDGroupVersion.String()
	ReadonlyGroup_GroupVersionKind = CRDGroupVersion.WithKind(ReadonlyGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ReadonlyGroup{}, &ReadonlyGroupList{})
}