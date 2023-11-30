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

type AddonAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ResponseBody *string `json:"responseBody,omitempty" tf:"response_body,omitempty"`

	Status map[string]*string `json:"status,omitempty" tf:"status,omitempty"`
}

type AddonAttachmentParameters struct {

	// ID of cluster.
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Name of addon.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Raw Values. Conflict with `request_body`. Required with `raw_values_type`.
	// +kubebuilder:validation:Optional
	RawValues *string `json:"rawValues,omitempty" tf:"raw_values,omitempty"`

	// The type of raw Values. Required with `raw_values`.
	// +kubebuilder:validation:Optional
	RawValuesType *string `json:"rawValuesType,omitempty" tf:"raw_values_type,omitempty"`

	// Serialized json string as request body of addon spec. If set, will ignore `version` and `values`.
	// +kubebuilder:validation:Optional
	RequestBody *string `json:"requestBody,omitempty" tf:"request_body,omitempty"`

	// Values the addon passthroughs. Conflict with `request_body`.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`

	// Addon version, default latest version. Conflict with `request_body`.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// AddonAttachmentSpec defines the desired state of AddonAttachment
type AddonAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddonAttachmentParameters `json:"forProvider"`
}

// AddonAttachmentStatus defines the observed state of AddonAttachment.
type AddonAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddonAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AddonAttachment is the Schema for the AddonAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type AddonAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddonAttachmentSpec   `json:"spec"`
	Status            AddonAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddonAttachmentList contains a list of AddonAttachments
type AddonAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddonAttachment `json:"items"`
}

// Repository type metadata.
var (
	AddonAttachment_Kind             = "AddonAttachment"
	AddonAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AddonAttachment_Kind}.String()
	AddonAttachment_KindAPIVersion   = AddonAttachment_Kind + "." + CRDGroupVersion.String()
	AddonAttachment_GroupVersionKind = CRDGroupVersion.WithKind(AddonAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&AddonAttachment{}, &AddonAttachmentList{})
}
