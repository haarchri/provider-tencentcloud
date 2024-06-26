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

type DrmStreamingInfoInitParameters struct {

	// ID of the adaptive dynamic streaming template whose protection type is SimpleAES.
	// ID of the adaptive dynamic streaming template whose protection type is `SimpleAES`.
	SimpleAesDefinition *string `json:"simpleAesDefinition,omitempty" tf:"simple_aes_definition,omitempty"`
}

type DrmStreamingInfoObservation struct {

	// ID of the adaptive dynamic streaming template whose protection type is SimpleAES.
	// ID of the adaptive dynamic streaming template whose protection type is `SimpleAES`.
	SimpleAesDefinition *string `json:"simpleAesDefinition,omitempty" tf:"simple_aes_definition,omitempty"`
}

type DrmStreamingInfoParameters struct {

	// ID of the adaptive dynamic streaming template whose protection type is SimpleAES.
	// ID of the adaptive dynamic streaming template whose protection type is `SimpleAES`.
	// +kubebuilder:validation:Optional
	SimpleAesDefinition *string `json:"simpleAesDefinition,omitempty" tf:"simple_aes_definition,omitempty"`
}

type ResolutionNamesInitParameters struct {

	// Length of video short side in px.
	// Length of video short side in px.
	MinEdgeLength *float64 `json:"minEdgeLength,omitempty" tf:"min_edge_length,omitempty"`

	// Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	// Display name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ResolutionNamesObservation struct {

	// Length of video short side in px.
	// Length of video short side in px.
	MinEdgeLength *float64 `json:"minEdgeLength,omitempty" tf:"min_edge_length,omitempty"`

	// Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	// Display name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ResolutionNamesParameters struct {

	// Length of video short side in px.
	// Length of video short side in px.
	// +kubebuilder:validation:Optional
	MinEdgeLength *float64 `json:"minEdgeLength" tf:"min_edge_length,omitempty"`

	// Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	// Display name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type SuperPlayerConfigInitParameters struct {

	// ID of the unencrypted adaptive bitrate streaming template that allows output, which is required if drm_switch is false.
	// ID of the unencrypted adaptive bitrate streaming template that allows output, which is required if `drm_switch` is `false`.
	AdaptiveDynamicStreamingDefinition *string `json:"adaptiveDynamicStreamingDefinition,omitempty" tf:"adaptive_dynamic_streaming_definition,omitempty"`

	// Template description. Length limit: 256 characters.
	// Template description. Length limit: 256 characters.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Domain name used for playback. If it is left empty or set to Default, the domain name configured in Default Distribution Configuration will be used. Default by default.
	// Domain name used for playback. If it is left empty or set to `Default`, the domain name configured in [Default Distribution Configuration](https://cloud.tencent.com/document/product/266/33373) will be used. `Default` by default.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Content of the DRM-protected adaptive bitrate streaming template that allows output, which is required if drm_switch is true.
	// Content of the DRM-protected adaptive bitrate streaming template that allows output, which is required if `drm_switch` is `true`.
	DrmStreamingInfo []DrmStreamingInfoInitParameters `json:"drmStreamingInfo,omitempty" tf:"drm_streaming_info,omitempty"`

	// Switch of DRM-protected adaptive bitstream playback: true: enabled, indicating to play back only output adaptive bitstreams protected by DRM; false: disabled, indicating to play back unencrypted output adaptive bitstreams. Default value: false.
	// Switch of DRM-protected adaptive bitstream playback: `true`: enabled, indicating to play back only output adaptive bitstreams protected by DRM; `false`: disabled, indicating to play back unencrypted output adaptive bitstreams. Default value: `false`.
	DrmSwitch *bool `json:"drmSwitch,omitempty" tf:"drm_switch,omitempty"`

	// ID of the image sprite template that allows output.
	// ID of the image sprite template that allows output.
	ImageSpriteDefinition *string `json:"imageSpriteDefinition,omitempty" tf:"image_sprite_definition,omitempty"`

	// Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	// Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Display name of player for substreams with different resolutions. If this parameter is left empty or an empty array, the default configuration will be used: min_edge_length: 240, name: LD; min_edge_length: 480, name: SD; min_edge_length: 720, name: HD; min_edge_length: 1080, name: FHD; min_edge_length: 1440, name: 2K; min_edge_length: 2160, name: 4K; min_edge_length: 4320, name: 8K.
	// Display name of player for substreams with different resolutions. If this parameter is left empty or an empty array, the default configuration will be used: `min_edge_length: 240, name: LD`; `min_edge_length: 480, name: SD`; `min_edge_length: 720, name: HD`; `min_edge_length: 1080, name: FHD`; `min_edge_length: 1440, name: 2K`; `min_edge_length: 2160, name: 4K`; `min_edge_length: 4320, name: 8K`.
	ResolutionNames []ResolutionNamesInitParameters `json:"resolutionNames,omitempty" tf:"resolution_names,omitempty"`

	// Scheme used for playback. If it is left empty or set to Default, the scheme configured in Default Distribution Configuration will be used. Other valid values: HTTP; HTTPS.
	// Scheme used for playback. If it is left empty or set to `Default`, the scheme configured in [Default Distribution Configuration](https://cloud.tencent.com/document/product/266/33373) will be used. Other valid values: `HTTP`; `HTTPS`.
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppID *float64 `json:"subAppId,omitempty" tf:"sub_app_id,omitempty"`
}

type SuperPlayerConfigObservation struct {

	// ID of the unencrypted adaptive bitrate streaming template that allows output, which is required if drm_switch is false.
	// ID of the unencrypted adaptive bitrate streaming template that allows output, which is required if `drm_switch` is `false`.
	AdaptiveDynamicStreamingDefinition *string `json:"adaptiveDynamicStreamingDefinition,omitempty" tf:"adaptive_dynamic_streaming_definition,omitempty"`

	// Template description. Length limit: 256 characters.
	// Template description. Length limit: 256 characters.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Creation time of template in ISO date format.
	// Creation time of template in ISO date format.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Domain name used for playback. If it is left empty or set to Default, the domain name configured in Default Distribution Configuration will be used. Default by default.
	// Domain name used for playback. If it is left empty or set to `Default`, the domain name configured in [Default Distribution Configuration](https://cloud.tencent.com/document/product/266/33373) will be used. `Default` by default.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Content of the DRM-protected adaptive bitrate streaming template that allows output, which is required if drm_switch is true.
	// Content of the DRM-protected adaptive bitrate streaming template that allows output, which is required if `drm_switch` is `true`.
	DrmStreamingInfo []DrmStreamingInfoObservation `json:"drmStreamingInfo,omitempty" tf:"drm_streaming_info,omitempty"`

	// Switch of DRM-protected adaptive bitstream playback: true: enabled, indicating to play back only output adaptive bitstreams protected by DRM; false: disabled, indicating to play back unencrypted output adaptive bitstreams. Default value: false.
	// Switch of DRM-protected adaptive bitstream playback: `true`: enabled, indicating to play back only output adaptive bitstreams protected by DRM; `false`: disabled, indicating to play back unencrypted output adaptive bitstreams. Default value: `false`.
	DrmSwitch *bool `json:"drmSwitch,omitempty" tf:"drm_switch,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the image sprite template that allows output.
	// ID of the image sprite template that allows output.
	ImageSpriteDefinition *string `json:"imageSpriteDefinition,omitempty" tf:"image_sprite_definition,omitempty"`

	// Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	// Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Display name of player for substreams with different resolutions. If this parameter is left empty or an empty array, the default configuration will be used: min_edge_length: 240, name: LD; min_edge_length: 480, name: SD; min_edge_length: 720, name: HD; min_edge_length: 1080, name: FHD; min_edge_length: 1440, name: 2K; min_edge_length: 2160, name: 4K; min_edge_length: 4320, name: 8K.
	// Display name of player for substreams with different resolutions. If this parameter is left empty or an empty array, the default configuration will be used: `min_edge_length: 240, name: LD`; `min_edge_length: 480, name: SD`; `min_edge_length: 720, name: HD`; `min_edge_length: 1080, name: FHD`; `min_edge_length: 1440, name: 2K`; `min_edge_length: 2160, name: 4K`; `min_edge_length: 4320, name: 8K`.
	ResolutionNames []ResolutionNamesObservation `json:"resolutionNames,omitempty" tf:"resolution_names,omitempty"`

	// Scheme used for playback. If it is left empty or set to Default, the scheme configured in Default Distribution Configuration will be used. Other valid values: HTTP; HTTPS.
	// Scheme used for playback. If it is left empty or set to `Default`, the scheme configured in [Default Distribution Configuration](https://cloud.tencent.com/document/product/266/33373) will be used. Other valid values: `HTTP`; `HTTPS`.
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppID *float64 `json:"subAppId,omitempty" tf:"sub_app_id,omitempty"`

	// Last modified time of template in ISO date format.
	// Last modified time of template in ISO date format.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type SuperPlayerConfigParameters struct {

	// ID of the unencrypted adaptive bitrate streaming template that allows output, which is required if drm_switch is false.
	// ID of the unencrypted adaptive bitrate streaming template that allows output, which is required if `drm_switch` is `false`.
	// +kubebuilder:validation:Optional
	AdaptiveDynamicStreamingDefinition *string `json:"adaptiveDynamicStreamingDefinition,omitempty" tf:"adaptive_dynamic_streaming_definition,omitempty"`

	// Template description. Length limit: 256 characters.
	// Template description. Length limit: 256 characters.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Domain name used for playback. If it is left empty or set to Default, the domain name configured in Default Distribution Configuration will be used. Default by default.
	// Domain name used for playback. If it is left empty or set to `Default`, the domain name configured in [Default Distribution Configuration](https://cloud.tencent.com/document/product/266/33373) will be used. `Default` by default.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Content of the DRM-protected adaptive bitrate streaming template that allows output, which is required if drm_switch is true.
	// Content of the DRM-protected adaptive bitrate streaming template that allows output, which is required if `drm_switch` is `true`.
	// +kubebuilder:validation:Optional
	DrmStreamingInfo []DrmStreamingInfoParameters `json:"drmStreamingInfo,omitempty" tf:"drm_streaming_info,omitempty"`

	// Switch of DRM-protected adaptive bitstream playback: true: enabled, indicating to play back only output adaptive bitstreams protected by DRM; false: disabled, indicating to play back unencrypted output adaptive bitstreams. Default value: false.
	// Switch of DRM-protected adaptive bitstream playback: `true`: enabled, indicating to play back only output adaptive bitstreams protected by DRM; `false`: disabled, indicating to play back unencrypted output adaptive bitstreams. Default value: `false`.
	// +kubebuilder:validation:Optional
	DrmSwitch *bool `json:"drmSwitch,omitempty" tf:"drm_switch,omitempty"`

	// ID of the image sprite template that allows output.
	// ID of the image sprite template that allows output.
	// +kubebuilder:validation:Optional
	ImageSpriteDefinition *string `json:"imageSpriteDefinition,omitempty" tf:"image_sprite_definition,omitempty"`

	// Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	// Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Display name of player for substreams with different resolutions. If this parameter is left empty or an empty array, the default configuration will be used: min_edge_length: 240, name: LD; min_edge_length: 480, name: SD; min_edge_length: 720, name: HD; min_edge_length: 1080, name: FHD; min_edge_length: 1440, name: 2K; min_edge_length: 2160, name: 4K; min_edge_length: 4320, name: 8K.
	// Display name of player for substreams with different resolutions. If this parameter is left empty or an empty array, the default configuration will be used: `min_edge_length: 240, name: LD`; `min_edge_length: 480, name: SD`; `min_edge_length: 720, name: HD`; `min_edge_length: 1080, name: FHD`; `min_edge_length: 1440, name: 2K`; `min_edge_length: 2160, name: 4K`; `min_edge_length: 4320, name: 8K`.
	// +kubebuilder:validation:Optional
	ResolutionNames []ResolutionNamesParameters `json:"resolutionNames,omitempty" tf:"resolution_names,omitempty"`

	// Scheme used for playback. If it is left empty or set to Default, the scheme configured in Default Distribution Configuration will be used. Other valid values: HTTP; HTTPS.
	// Scheme used for playback. If it is left empty or set to `Default`, the scheme configured in [Default Distribution Configuration](https://cloud.tencent.com/document/product/266/33373) will be used. Other valid values: `HTTP`; `HTTPS`.
	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	// +kubebuilder:validation:Optional
	SubAppID *float64 `json:"subAppId,omitempty" tf:"sub_app_id,omitempty"`
}

// SuperPlayerConfigSpec defines the desired state of SuperPlayerConfig
type SuperPlayerConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SuperPlayerConfigParameters `json:"forProvider"`
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
	InitProvider SuperPlayerConfigInitParameters `json:"initProvider,omitempty"`
}

// SuperPlayerConfigStatus defines the observed state of SuperPlayerConfig.
type SuperPlayerConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SuperPlayerConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SuperPlayerConfig is the Schema for the SuperPlayerConfigs API. Provide a resource to create a VOD super player config.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type SuperPlayerConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SuperPlayerConfigSpec   `json:"spec"`
	Status SuperPlayerConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SuperPlayerConfigList contains a list of SuperPlayerConfigs
type SuperPlayerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SuperPlayerConfig `json:"items"`
}

// Repository type metadata.
var (
	SuperPlayerConfig_Kind             = "SuperPlayerConfig"
	SuperPlayerConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SuperPlayerConfig_Kind}.String()
	SuperPlayerConfig_KindAPIVersion   = SuperPlayerConfig_Kind + "." + CRDGroupVersion.String()
	SuperPlayerConfig_GroupVersionKind = CRDGroupVersion.WithKind(SuperPlayerConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&SuperPlayerConfig{}, &SuperPlayerConfigList{})
}
