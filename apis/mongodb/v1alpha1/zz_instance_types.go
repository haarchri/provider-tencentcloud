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

type InstanceObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	StandbyInstanceList []StandbyInstanceListObservation `json:"standbyInstanceList,omitempty" tf:"standby_instance_list,omitempty"`

	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`

	Vip *string `json:"vip,omitempty" tf:"vip,omitempty"`

	Vport *float64 `json:"vport,omitempty" tf:"vport,omitempty"`
}

type InstanceParameters struct {

	// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
	// +kubebuilder:validation:Optional
	AutoRenewFlag *float64 `json:"autoRenewFlag,omitempty" tf:"auto_renew_flag,omitempty"`

	// A list of nodes deployed in multiple availability zones. For more information, please use the API DescribeSpecInfo.
	// - Multi-availability zone deployment nodes can only be deployed in 3 different availability zones. It is not supported to deploy most nodes of the cluster in the same availability zone. For example, a 3-node cluster does not support the deployment of 2 nodes in the same zone.
	// - Version 4.2 and above are not supported.
	// - Read-only disaster recovery instances are not supported.
	// - Basic network cannot be selected.
	// +kubebuilder:validation:Optional
	AvailabilityZoneList []*string `json:"availabilityZoneList,omitempty" tf:"availability_zone_list,omitempty"`

	// The available zone of the Mongodb.
	// +kubebuilder:validation:Required
	AvailableZone *string `json:"availableZone" tf:"available_zone,omitempty"`

	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
	// +kubebuilder:validation:Optional
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
	// +kubebuilder:validation:Required
	EngineVersion *string `json:"engineVersion" tf:"engine_version,omitempty"`

	// The availability zone to which the Hidden node belongs. This parameter must be configured to deploy instances across availability zones.
	// +kubebuilder:validation:Optional
	HiddenZone *string `json:"hiddenZone,omitempty" tf:"hidden_zone,omitempty"`

	// Name of the Mongodb instance.
	// +kubebuilder:validation:Required
	InstanceName *string `json:"instanceName" tf:"instance_name,omitempty"`

	// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
	// +kubebuilder:validation:Required
	MachineType *string `json:"machineType" tf:"machine_type,omitempty"`

	// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	// +kubebuilder:validation:Required
	Memory *float64 `json:"memory" tf:"memory,omitempty"`

	// The number of nodes in each replica set. Default value: 3.
	// +kubebuilder:validation:Optional
	NodeNum *float64 `json:"nodeNum,omitempty" tf:"node_num,omitempty"`

	// Password of this Mongodb account.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when charge_type is set to `PREPAID`.
	// +kubebuilder:validation:Optional
	PrepaidPeriod *float64 `json:"prepaidPeriod,omitempty" tf:"prepaid_period,omitempty"`

	// ID of the project which the instance belongs.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// ID of the security group.
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// ID of the subnet within this VPC. The value is required if `vpc_id` is set.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of the VPC.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcidSelector,omitempty" tf:"-"`

	// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	// +kubebuilder:validation:Required
	Volume *float64 `json:"volume" tf:"volume,omitempty"`
}

type StandbyInstanceListObservation struct {
	StandbyInstanceID *string `json:"standbyInstanceId,omitempty" tf:"standby_instance_id,omitempty"`

	StandbyInstanceRegion *string `json:"standbyInstanceRegion,omitempty" tf:"standby_instance_region,omitempty"`
}

type StandbyInstanceListParameters struct {
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
