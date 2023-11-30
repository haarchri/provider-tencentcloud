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

type ActionsCodeActionObservation struct {
}

type ActionsCodeActionParameters struct {

	// Action name.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Action parameters.
	// +kubebuilder:validation:Required
	Parameters []CodeActionParametersParameters `json:"parameters" tf:"parameters,omitempty"`
}

type ActionsNormalActionObservation struct {
}

type ActionsNormalActionParameters struct {

	// Action name.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Action parameters.
	// +kubebuilder:validation:Required
	Parameters []ActionsNormalActionParametersParameters `json:"parameters" tf:"parameters,omitempty"`
}

type ActionsNormalActionParametersObservation struct {
}

type ActionsNormalActionParametersParameters struct {

	// Parameter Name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Parameter Values.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ActionsObservation struct {
}

type ActionsParameters struct {

	// Define a code action.
	// +kubebuilder:validation:Optional
	CodeAction []CodeActionParameters `json:"codeAction,omitempty" tf:"code_action,omitempty"`

	// Define a normal action.
	// +kubebuilder:validation:Optional
	NormalAction []NormalActionParameters `json:"normalAction,omitempty" tf:"normal_action,omitempty"`

	// Define a rewrite action.
	// +kubebuilder:validation:Optional
	RewriteAction []RewriteActionParameters `json:"rewriteAction,omitempty" tf:"rewrite_action,omitempty"`
}

type ActionsRewriteActionObservation struct {
}

type ActionsRewriteActionParameters struct {

	// Action name.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Action parameters.
	// +kubebuilder:validation:Required
	Parameters []ActionsRewriteActionParametersParameters `json:"parameters" tf:"parameters,omitempty"`
}

type ActionsRewriteActionParametersObservation struct {
}

type ActionsRewriteActionParametersParameters struct {

	// Action to take on the HEADER. Valid values: `add`, `del`, `set`.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Target HEADER name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Parameter Value.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type AndObservation struct {
}

type AndParameters struct {

	// Whether to ignore the case of the parameter value, the default value is false.
	// +kubebuilder:validation:Optional
	IgnoreCase *bool `json:"ignoreCase,omitempty" tf:"ignore_case,omitempty"`

	// The parameter name corresponding to the matching type is valid when the Target value is the following, and the valid value cannot be empty: `query_string` (query string): The parameter name of the query string in the URL request under the current site, such as lang and version in lang=cn&version=1; `request_header` (HTTP request header): HTTP request header field name, such as Accept-Language in Accept-Language:zh-CN,zh;q=0.9.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Condition operator. Valid values are `equal`, `notequal`.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Condition target. Valid values:- `host`: Host of the URL.- `filename`: filename of the URL.- `extension`: file extension of the URL.- `full_url`: full url.- `url`: path of the URL.
	// +kubebuilder:validation:Required
	Target *string `json:"target" tf:"target,omitempty"`

	// Condition Value.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type CodeActionObservation struct {
}

type CodeActionParameters struct {

	// Action name.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Action parameters.
	// +kubebuilder:validation:Required
	Parameters []ParametersParameters `json:"parameters" tf:"parameters,omitempty"`
}

type CodeActionParametersObservation struct {
}

type CodeActionParametersParameters struct {

	// Parameter Name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// HTTP status code to use.
	// +kubebuilder:validation:Required
	StatusCode *float64 `json:"statusCode" tf:"status_code,omitempty"`

	// Parameter Values.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type NormalActionObservation struct {
}

type NormalActionParameters struct {

	// Action name.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Action parameters.
	// +kubebuilder:validation:Required
	Parameters []NormalActionParametersParameters `json:"parameters" tf:"parameters,omitempty"`
}

type NormalActionParametersObservation struct {
}

type NormalActionParametersParameters struct {

	// Parameter Name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Parameter Values.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type OrAndObservation struct {
}

type OrAndParameters struct {

	// Whether to ignore the case of the parameter value, the default value is false.
	// +kubebuilder:validation:Optional
	IgnoreCase *bool `json:"ignoreCase,omitempty" tf:"ignore_case,omitempty"`

	// The parameter name corresponding to the matching type is valid when the Target value is the following, and the valid value cannot be empty:- `query_string` (query string): The parameter name of the query string in the URL request under the current site, such as lang and version in lang=cn&version=1; `request_header` (HTTP request header): HTTP request header field name, such as Accept-Language in Accept-Language:zh-CN,zh;q=0.9.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Condition operator. Valid values are `equal`, `notequal`.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Condition target. Valid values:- `host`: Host of the URL.- `filename`: filename of the URL.- `extension`: file extension of the URL.- `full_url`: full url.- `url`: path of the URL.
	// +kubebuilder:validation:Required
	Target *string `json:"target" tf:"target,omitempty"`

	// Condition Value.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type OrObservation struct {
}

type OrParameters struct {

	// AND Conditions list of the rule. Rule would be triggered if all conditions are true.
	// +kubebuilder:validation:Required
	And []AndParameters `json:"and" tf:"and,omitempty"`
}

type ParametersObservation struct {
}

type ParametersParameters struct {

	// Parameter Name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// HTTP status code to use.
	// +kubebuilder:validation:Required
	StatusCode *float64 `json:"statusCode" tf:"status_code,omitempty"`

	// Parameter Values.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type RewriteActionObservation struct {
}

type RewriteActionParameters struct {

	// Action name.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Action parameters.
	// +kubebuilder:validation:Required
	Parameters []RewriteActionParametersParameters `json:"parameters" tf:"parameters,omitempty"`
}

type RewriteActionParametersObservation struct {
}

type RewriteActionParametersParameters struct {

	// Action to take on the HEADER. Valid values: `add`, `del`, `set`.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Target HEADER name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Parameter Value.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type RuleEngineObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`
}

type RuleEngineParameters struct {

	// The rule name (1 to 255 characters).
	// +kubebuilder:validation:Required
	RuleName *string `json:"ruleName" tf:"rule_name,omitempty"`

	// Rule items list.
	// +kubebuilder:validation:Required
	Rules []RulesParameters `json:"rules" tf:"rules,omitempty"`

	// Rule status. Values: `enable`: Enabled; `disable`: Disabled.
	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`

	// rule tag list.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of the site.
	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type RulesActionsObservation struct {
}

type RulesActionsParameters struct {

	// Define a code action.
	// +kubebuilder:validation:Optional
	CodeAction []ActionsCodeActionParameters `json:"codeAction,omitempty" tf:"code_action,omitempty"`

	// Define a normal action.
	// +kubebuilder:validation:Optional
	NormalAction []ActionsNormalActionParameters `json:"normalAction,omitempty" tf:"normal_action,omitempty"`

	// Define a rewrite action.
	// +kubebuilder:validation:Optional
	RewriteAction []ActionsRewriteActionParameters `json:"rewriteAction,omitempty" tf:"rewrite_action,omitempty"`
}

type RulesObservation struct {
}

type RulesOrObservation struct {
}

type RulesOrParameters struct {

	// AND Conditions list of the rule. Rule would be triggered if all conditions are true.
	// +kubebuilder:validation:Required
	And []OrAndParameters `json:"and" tf:"and,omitempty"`
}

type RulesParameters struct {

	// Actions list of the rule. See details in data source `rule_engine_setting`.
	// +kubebuilder:validation:Required
	Actions []ActionsParameters `json:"actions" tf:"actions,omitempty"`

	// OR Conditions list of the rule. Rule would be triggered if any of the condition is true.
	// +kubebuilder:validation:Required
	Or []OrParameters `json:"or" tf:"or,omitempty"`

	// Actions list of the rule. See details in data source `rule_engine_setting`.
	// +kubebuilder:validation:Optional
	SubRules []SubRulesParameters `json:"subRules,omitempty" tf:"sub_rules,omitempty"`
}

type SubRulesObservation struct {
}

type SubRulesParameters struct {

	// Rule items list.
	// +kubebuilder:validation:Required
	Rules []SubRulesRulesParameters `json:"rules" tf:"rules,omitempty"`

	// rule tag list.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SubRulesRulesObservation struct {
}

type SubRulesRulesParameters struct {

	// Actions list of the rule. See details in data source `rule_engine_setting`.
	// +kubebuilder:validation:Required
	Actions []RulesActionsParameters `json:"actions" tf:"actions,omitempty"`

	// OR Conditions list of the rule. Rule would be triggered if any of the condition is true.
	// +kubebuilder:validation:Required
	Or []RulesOrParameters `json:"or" tf:"or,omitempty"`
}

// RuleEngineSpec defines the desired state of RuleEngine
type RuleEngineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleEngineParameters `json:"forProvider"`
}

// RuleEngineStatus defines the observed state of RuleEngine.
type RuleEngineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleEngineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RuleEngine is the Schema for the RuleEngines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type RuleEngine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleEngineSpec   `json:"spec"`
	Status            RuleEngineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleEngineList contains a list of RuleEngines
type RuleEngineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuleEngine `json:"items"`
}

// Repository type metadata.
var (
	RuleEngine_Kind             = "RuleEngine"
	RuleEngine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RuleEngine_Kind}.String()
	RuleEngine_KindAPIVersion   = RuleEngine_Kind + "." + CRDGroupVersion.String()
	RuleEngine_GroupVersionKind = CRDGroupVersion.WithKind(RuleEngine_Kind)
)

func init() {
	SchemeBuilder.Register(&RuleEngine{}, &RuleEngineList{})
}
