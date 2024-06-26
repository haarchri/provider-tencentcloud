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

type HttpRuleInitParameters struct {

	// Timeout of the health check response, default value is 2s.
	// Timeout of the health check response, default value is 2s.
	ConnectTimeout *float64 `json:"connectTimeout,omitempty" tf:"connect_timeout,omitempty"`

	// Forward domain of the forward rule.
	// Forward domain of the forward rule.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The default value of requested host which is forwarded to the realserver by the listener is default.
	// The default value of requested host which is forwarded to the realserver by the listener is `default`.
	ForwardHost *string `json:"forwardHost,omitempty" tf:"forward_host,omitempty"`

	// Indicates whether health check is enable.
	// Indicates whether health check is enable.
	HealthCheck *bool `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// Method of the health check. Valid value: GET and HEAD.
	// Method of the health check. Valid value: `GET` and `HEAD`.
	HealthCheckMethod *string `json:"healthCheckMethod,omitempty" tf:"health_check_method,omitempty"`

	// Path of health check. Maximum length is 80.
	// Path of health check. Maximum length is 80.
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// Return code of confirmed normal. Valid value: 100, 200, 300, 400 and 500.
	// Return code of confirmed normal. Valid value: `100`, `200`, `300`, `400` and `500`.
	// +listType=set
	HealthCheckStatusCodes []*float64 `json:"healthCheckStatusCodes,omitempty" tf:"health_check_status_codes,omitempty"`

	// Interval of the health check, default value is 5s.
	// Interval of the health check, default value is 5s.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// ID of the layer7 listener.
	// ID of the layer7 listener.
	// +crossplane:generate:reference:type=Layer7Listener
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// Reference to a Layer7Listener to populate listenerId.
	// +kubebuilder:validation:Optional
	ListenerIDRef *v1.Reference `json:"listenerIdRef,omitempty" tf:"-"`

	// Selector for a Layer7Listener to populate listenerId.
	// +kubebuilder:validation:Optional
	ListenerIDSelector *v1.Selector `json:"listenerIdSelector,omitempty" tf:"-"`

	// Path of the forward rule. Maximum length is 80.
	// Path of the forward rule. Maximum length is 80.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Type of the realserver. Valid value: IP and DOMAIN.
	// Type of the realserver. Valid value: `IP` and `DOMAIN`.
	RealserverType *string `json:"realserverType,omitempty" tf:"realserver_type,omitempty"`

	// An information list of GAAP realserver.
	// An information list of GAAP realserver.
	Realservers []RealserversInitParameters `json:"realservers,omitempty" tf:"realservers,omitempty"`

	// Scheduling policy of the forward rule, default value is rr. Valid value: rr, wrr and lc.
	// Scheduling policy of the forward rule, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler *string `json:"scheduler,omitempty" tf:"scheduler,omitempty"`

	// ServerNameIndication (SNI) is required when the SNI switch is turned on.
	// ServerNameIndication (SNI) is required when the SNI switch is turned on.
	Sni *string `json:"sni,omitempty" tf:"sni,omitempty"`

	// ServerNameIndication (SNI) switch. ON means on and OFF means off.
	// ServerNameIndication (SNI) switch. ON means on and OFF means off.
	SniSwitch *string `json:"sniSwitch,omitempty" tf:"sni_switch,omitempty"`
}

type HttpRuleObservation struct {

	// Timeout of the health check response, default value is 2s.
	// Timeout of the health check response, default value is 2s.
	ConnectTimeout *float64 `json:"connectTimeout,omitempty" tf:"connect_timeout,omitempty"`

	// Forward domain of the forward rule.
	// Forward domain of the forward rule.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The default value of requested host which is forwarded to the realserver by the listener is default.
	// The default value of requested host which is forwarded to the realserver by the listener is `default`.
	ForwardHost *string `json:"forwardHost,omitempty" tf:"forward_host,omitempty"`

	// Indicates whether health check is enable.
	// Indicates whether health check is enable.
	HealthCheck *bool `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// Method of the health check. Valid value: GET and HEAD.
	// Method of the health check. Valid value: `GET` and `HEAD`.
	HealthCheckMethod *string `json:"healthCheckMethod,omitempty" tf:"health_check_method,omitempty"`

	// Path of health check. Maximum length is 80.
	// Path of health check. Maximum length is 80.
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// Return code of confirmed normal. Valid value: 100, 200, 300, 400 and 500.
	// Return code of confirmed normal. Valid value: `100`, `200`, `300`, `400` and `500`.
	// +listType=set
	HealthCheckStatusCodes []*float64 `json:"healthCheckStatusCodes,omitempty" tf:"health_check_status_codes,omitempty"`

	// ID of the GAAP realserver.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Interval of the health check, default value is 5s.
	// Interval of the health check, default value is 5s.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// ID of the layer7 listener.
	// ID of the layer7 listener.
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// Path of the forward rule. Maximum length is 80.
	// Path of the forward rule. Maximum length is 80.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Type of the realserver. Valid value: IP and DOMAIN.
	// Type of the realserver. Valid value: `IP` and `DOMAIN`.
	RealserverType *string `json:"realserverType,omitempty" tf:"realserver_type,omitempty"`

	// An information list of GAAP realserver.
	// An information list of GAAP realserver.
	Realservers []RealserversObservation `json:"realservers,omitempty" tf:"realservers,omitempty"`

	// Scheduling policy of the forward rule, default value is rr. Valid value: rr, wrr and lc.
	// Scheduling policy of the forward rule, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler *string `json:"scheduler,omitempty" tf:"scheduler,omitempty"`

	// ServerNameIndication (SNI) is required when the SNI switch is turned on.
	// ServerNameIndication (SNI) is required when the SNI switch is turned on.
	Sni *string `json:"sni,omitempty" tf:"sni,omitempty"`

	// ServerNameIndication (SNI) switch. ON means on and OFF means off.
	// ServerNameIndication (SNI) switch. ON means on and OFF means off.
	SniSwitch *string `json:"sniSwitch,omitempty" tf:"sni_switch,omitempty"`
}

type HttpRuleParameters struct {

	// Timeout of the health check response, default value is 2s.
	// Timeout of the health check response, default value is 2s.
	// +kubebuilder:validation:Optional
	ConnectTimeout *float64 `json:"connectTimeout,omitempty" tf:"connect_timeout,omitempty"`

	// Forward domain of the forward rule.
	// Forward domain of the forward rule.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The default value of requested host which is forwarded to the realserver by the listener is default.
	// The default value of requested host which is forwarded to the realserver by the listener is `default`.
	// +kubebuilder:validation:Optional
	ForwardHost *string `json:"forwardHost,omitempty" tf:"forward_host,omitempty"`

	// Indicates whether health check is enable.
	// Indicates whether health check is enable.
	// +kubebuilder:validation:Optional
	HealthCheck *bool `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// Method of the health check. Valid value: GET and HEAD.
	// Method of the health check. Valid value: `GET` and `HEAD`.
	// +kubebuilder:validation:Optional
	HealthCheckMethod *string `json:"healthCheckMethod,omitempty" tf:"health_check_method,omitempty"`

	// Path of health check. Maximum length is 80.
	// Path of health check. Maximum length is 80.
	// +kubebuilder:validation:Optional
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// Return code of confirmed normal. Valid value: 100, 200, 300, 400 and 500.
	// Return code of confirmed normal. Valid value: `100`, `200`, `300`, `400` and `500`.
	// +kubebuilder:validation:Optional
	// +listType=set
	HealthCheckStatusCodes []*float64 `json:"healthCheckStatusCodes,omitempty" tf:"health_check_status_codes,omitempty"`

	// Interval of the health check, default value is 5s.
	// Interval of the health check, default value is 5s.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// ID of the layer7 listener.
	// ID of the layer7 listener.
	// +crossplane:generate:reference:type=Layer7Listener
	// +kubebuilder:validation:Optional
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// Reference to a Layer7Listener to populate listenerId.
	// +kubebuilder:validation:Optional
	ListenerIDRef *v1.Reference `json:"listenerIdRef,omitempty" tf:"-"`

	// Selector for a Layer7Listener to populate listenerId.
	// +kubebuilder:validation:Optional
	ListenerIDSelector *v1.Selector `json:"listenerIdSelector,omitempty" tf:"-"`

	// Path of the forward rule. Maximum length is 80.
	// Path of the forward rule. Maximum length is 80.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Type of the realserver. Valid value: IP and DOMAIN.
	// Type of the realserver. Valid value: `IP` and `DOMAIN`.
	// +kubebuilder:validation:Optional
	RealserverType *string `json:"realserverType,omitempty" tf:"realserver_type,omitempty"`

	// An information list of GAAP realserver.
	// An information list of GAAP realserver.
	// +kubebuilder:validation:Optional
	Realservers []RealserversParameters `json:"realservers,omitempty" tf:"realservers,omitempty"`

	// Scheduling policy of the forward rule, default value is rr. Valid value: rr, wrr and lc.
	// Scheduling policy of the forward rule, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	// +kubebuilder:validation:Optional
	Scheduler *string `json:"scheduler,omitempty" tf:"scheduler,omitempty"`

	// ServerNameIndication (SNI) is required when the SNI switch is turned on.
	// ServerNameIndication (SNI) is required when the SNI switch is turned on.
	// +kubebuilder:validation:Optional
	Sni *string `json:"sni,omitempty" tf:"sni,omitempty"`

	// ServerNameIndication (SNI) switch. ON means on and OFF means off.
	// ServerNameIndication (SNI) switch. ON means on and OFF means off.
	// +kubebuilder:validation:Optional
	SniSwitch *string `json:"sniSwitch,omitempty" tf:"sni_switch,omitempty"`
}

type RealserversInitParameters struct {

	// ID of the GAAP realserver.
	// ID of the GAAP realserver.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP of the GAAP realserver.
	// IP of the GAAP realserver.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Port of the GAAP realserver.
	// Port of the GAAP realserver.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Scheduling weight, default value is 1. Valid value ranges: (1~100).
	// Scheduling weight, default value is `1`. Valid value ranges: (1~100).
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type RealserversObservation struct {

	// ID of the GAAP realserver.
	// ID of the GAAP realserver.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP of the GAAP realserver.
	// IP of the GAAP realserver.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Port of the GAAP realserver.
	// Port of the GAAP realserver.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Scheduling weight, default value is 1. Valid value ranges: (1~100).
	// Scheduling weight, default value is `1`. Valid value ranges: (1~100).
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type RealserversParameters struct {

	// ID of the GAAP realserver.
	// ID of the GAAP realserver.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// IP of the GAAP realserver.
	// IP of the GAAP realserver.
	// +kubebuilder:validation:Optional
	IP *string `json:"ip" tf:"ip,omitempty"`

	// Port of the GAAP realserver.
	// Port of the GAAP realserver.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// Scheduling weight, default value is 1. Valid value ranges: (1~100).
	// Scheduling weight, default value is `1`. Valid value ranges: (1~100).
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

// HttpRuleSpec defines the desired state of HttpRule
type HttpRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HttpRuleParameters `json:"forProvider"`
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
	InitProvider HttpRuleInitParameters `json:"initProvider,omitempty"`
}

// HttpRuleStatus defines the observed state of HttpRule.
type HttpRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HttpRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HttpRule is the Schema for the HttpRules API. Provides a resource to create a forward rule of layer7 listener.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type HttpRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domain) || (has(self.initProvider) && has(self.initProvider.domain))",message="spec.forProvider.domain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.healthCheck) || (has(self.initProvider) && has(self.initProvider.healthCheck))",message="spec.forProvider.healthCheck is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.path) || (has(self.initProvider) && has(self.initProvider.path))",message="spec.forProvider.path is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.realserverType) || (has(self.initProvider) && has(self.initProvider.realserverType))",message="spec.forProvider.realserverType is a required parameter"
	Spec   HttpRuleSpec   `json:"spec"`
	Status HttpRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HttpRuleList contains a list of HttpRules
type HttpRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HttpRule `json:"items"`
}

// Repository type metadata.
var (
	HttpRule_Kind             = "HttpRule"
	HttpRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HttpRule_Kind}.String()
	HttpRule_KindAPIVersion   = HttpRule_Kind + "." + CRDGroupVersion.String()
	HttpRule_GroupVersionKind = CRDGroupVersion.WithKind(HttpRule_Kind)
)

func init() {
	SchemeBuilder.Register(&HttpRule{}, &HttpRuleList{})
}
