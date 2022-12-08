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

type ClusterObservation struct {
	Charset *string `json:"charset,omitempty" tf:"charset,omitempty"`

	ClusterStatus *string `json:"clusterStatus,omitempty" tf:"cluster_status,omitempty"`

	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	InstanceStatus *string `json:"instanceStatus,omitempty" tf:"instance_status,omitempty"`

	InstanceStorageSize *float64 `json:"instanceStorageSize,omitempty" tf:"instance_storage_size,omitempty"`

	RoGroupAddr []RoGroupAddrObservation `json:"roGroupAddr,omitempty" tf:"ro_group_addr,omitempty"`

	RoGroupID *string `json:"roGroupId,omitempty" tf:"ro_group_id,omitempty"`

	RoGroupInstances []RoGroupInstancesObservation `json:"roGroupInstances,omitempty" tf:"ro_group_instances,omitempty"`

	RwGroupAddr []RwGroupAddrObservation `json:"rwGroupAddr,omitempty" tf:"rw_group_addr,omitempty"`

	RwGroupID *string `json:"rwGroupId,omitempty" tf:"rw_group_id,omitempty"`

	RwGroupInstances []RwGroupInstancesObservation `json:"rwGroupInstances,omitempty" tf:"rw_group_instances,omitempty"`

	ServerlessStatus *string `json:"serverlessStatus,omitempty" tf:"serverless_status,omitempty"`

	StorageUsed *float64 `json:"storageUsed,omitempty" tf:"storage_used,omitempty"`
}

type ClusterParameters struct {

	// Specify whether the cluster can auto-pause while `db_mode` is `SERVERLESS`. Values: `yes` (default), `no`.
	// +kubebuilder:validation:Optional
	AutoPause *string `json:"autoPause,omitempty" tf:"auto_pause,omitempty"`

	// Specify auto-pause delay in second while `db_mode` is `SERVERLESS`. Value range: `[600, 691200]`. Default: `600`.
	// +kubebuilder:validation:Optional
	AutoPauseDelay *float64 `json:"autoPauseDelay,omitempty" tf:"auto_pause_delay,omitempty"`

	// Auto renew flag. Valid values are `0`(MANUAL_RENEW), `1`(AUTO_RENEW). Default value is `0`. Only works for PREPAID cluster.
	// +kubebuilder:validation:Optional
	AutoRenewFlag *float64 `json:"autoRenewFlag,omitempty" tf:"auto_renew_flag,omitempty"`

	// The available zone of the CynosDB Cluster.
	// +kubebuilder:validation:Required
	AvailableZone *string `json:"availableZone" tf:"available_zone,omitempty"`

	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`.
	// +kubebuilder:validation:Optional
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Name of CynosDB cluster.
	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// Specify DB mode, only available when `db_type` is `MYSQL`. Values: `NORMAL` (Default), `SERVERLESS`.
	// +kubebuilder:validation:Optional
	DBMode *string `json:"dbMode,omitempty" tf:"db_mode,omitempty"`

	// Type of CynosDB, and available values include `MYSQL`.
	// +kubebuilder:validation:Required
	DBType *string `json:"dbType" tf:"db_type,omitempty"`

	// Version of CynosDB, which is related to `db_type`. For `MYSQL`, available value is `5.7`.
	// +kubebuilder:validation:Required
	DBVersion *string `json:"dbVersion" tf:"db_version,omitempty"`

	// Indicate whether to delete cluster instance directly or not. Default is false. If set true, the cluster and its `All RELATED INSTANCES` will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
	// +kubebuilder:validation:Optional
	InstanceCPUCore *float64 `json:"instanceCpuCore,omitempty" tf:"instance_cpu_core,omitempty"`

	// Duration time for maintenance, unit in second. `3600` by default.
	// +kubebuilder:validation:Optional
	InstanceMaintainDuration *float64 `json:"instanceMaintainDuration,omitempty" tf:"instance_maintain_duration,omitempty"`

	// Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
	// +kubebuilder:validation:Optional
	InstanceMaintainStartTime *float64 `json:"instanceMaintainStartTime,omitempty" tf:"instance_maintain_start_time,omitempty"`

	// Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
	// +kubebuilder:validation:Optional
	InstanceMaintainWeekdays []*string `json:"instanceMaintainWeekdays,omitempty" tf:"instance_maintain_weekdays,omitempty"`

	// Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
	// +kubebuilder:validation:Optional
	InstanceMemorySize *float64 `json:"instanceMemorySize,omitempty" tf:"instance_memory_size,omitempty"`

	// Maximum CPU core count, required while `db_mode` is `SERVERLESS`, request DescribeServerlessInstanceSpecs for more reference.
	// +kubebuilder:validation:Optional
	MaxCPU *float64 `json:"maxCpu,omitempty" tf:"max_cpu,omitempty"`

	// Minimum CPU core count, required while `db_mode` is `SERVERLESS`, request DescribeServerlessInstanceSpecs for more reference.
	// +kubebuilder:validation:Optional
	MinCPU *float64 `json:"minCpu,omitempty" tf:"min_cpu,omitempty"`

	// Specify parameter list of database. It is valid when prarm_template_id is set in create cluster. Use `data.tencentcloud_mysql_default_params` to query available parameter details.
	// +kubebuilder:validation:Optional
	ParamItems []ParamItemsParameters `json:"paramItems,omitempty" tf:"param_items,omitempty"`

	// Password of `root` account.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Port of CynosDB cluster.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the parameter template.
	// +kubebuilder:validation:Optional
	PrarmTemplateID *float64 `json:"prarmTemplateId,omitempty" tf:"prarm_template_id,omitempty"`

	// The tenancy (time unit is month) of the prepaid instance. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. NOTE: it only works when charge_type is set to `PREPAID`.
	// +kubebuilder:validation:Optional
	PrepaidPeriod *float64 `json:"prepaidPeriod,omitempty" tf:"prepaid_period,omitempty"`

	// ID of the project. `0` by default.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// IDs of security group for `ro_group`.
	// +kubebuilder:validation:Optional
	RoGroupSg []*string `json:"roGroupSg,omitempty" tf:"ro_group_sg,omitempty"`

	// IDs of security group for `rw_group`.
	// +kubebuilder:validation:Optional
	RwGroupSg []*string `json:"rwGroupSg,omitempty" tf:"rw_group_sg,omitempty"`

	// Specify whether to pause or resume serverless cluster. values: `resume`, `pause`.
	// +kubebuilder:validation:Optional
	ServerlessStatusFlag *string `json:"serverlessStatusFlag,omitempty" tf:"serverless_status_flag,omitempty"`

	// Storage limit of CynosDB cluster instance, unit in GB. The maximum storage of a non-serverless instance in GB. NOTE: If db_type is `MYSQL` and charge_type is `PREPAID`, the value cannot exceed the maximum storage corresponding to the CPU and memory specifications, when charge_type is `POSTPAID_BY_HOUR`, this argument is unnecessary.
	// +kubebuilder:validation:Optional
	StorageLimit *float64 `json:"storageLimit,omitempty" tf:"storage_limit,omitempty"`

	// ID of the subnet within this VPC.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The tags of the CynosDB cluster.
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
}

type ParamItemsObservation struct {
}

type ParamItemsParameters struct {

	// Param expected value to set.
	// +kubebuilder:validation:Required
	CurrentValue *string `json:"currentValue" tf:"current_value,omitempty"`

	// Name of param, e.g. `character_set_server`.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Param old value, indicates the value which already set, this value is required when modifying current_value.
	// +kubebuilder:validation:Optional
	OldValue *string `json:"oldValue,omitempty" tf:"old_value,omitempty"`
}

type RoGroupAddrObservation struct {
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type RoGroupAddrParameters struct {
}

type RoGroupInstancesObservation struct {
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`
}

type RoGroupInstancesParameters struct {
}

type RwGroupAddrObservation struct {
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type RwGroupAddrParameters struct {
}

type RwGroupInstancesObservation struct {
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`
}

type RwGroupInstancesParameters struct {
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
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
