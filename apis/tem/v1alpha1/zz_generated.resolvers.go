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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AppConfig.
func (mg *AppConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnvironmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EnvironmentIDRef,
		Selector:     mg.Spec.ForProvider.EnvironmentIDSelector,
		To: reference.To{
			List:    &EnvironmentList{},
			Managed: &Environment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EnvironmentID")
	}
	mg.Spec.ForProvider.EnvironmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnvironmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LogConfig.
func (mg *LogConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnvironmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EnvironmentIDRef,
		Selector:     mg.Spec.ForProvider.EnvironmentIDSelector,
		To: reference.To{
			List:    &EnvironmentList{},
			Managed: &Environment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EnvironmentID")
	}
	mg.Spec.ForProvider.EnvironmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnvironmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkloadID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.WorkloadIDRef,
		Selector:     mg.Spec.ForProvider.WorkloadIDSelector,
		To: reference.To{
			List:    &WorkloadList{},
			Managed: &Workload{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkloadID")
	}
	mg.Spec.ForProvider.WorkloadID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkloadIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ScaleRule.
func (mg *ScaleRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnvironmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EnvironmentIDRef,
		Selector:     mg.Spec.ForProvider.EnvironmentIDSelector,
		To: reference.To{
			List:    &EnvironmentList{},
			Managed: &Environment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EnvironmentID")
	}
	mg.Spec.ForProvider.EnvironmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnvironmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkloadID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.WorkloadIDRef,
		Selector:     mg.Spec.ForProvider.WorkloadIDSelector,
		To: reference.To{
			List:    &WorkloadList{},
			Managed: &Workload{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkloadID")
	}
	mg.Spec.ForProvider.WorkloadID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkloadIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Workload.
func (mg *Workload) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnvironmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EnvironmentIDRef,
		Selector:     mg.Spec.ForProvider.EnvironmentIDSelector,
		To: reference.To{
			List:    &EnvironmentList{},
			Managed: &Environment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EnvironmentID")
	}
	mg.Spec.ForProvider.EnvironmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnvironmentIDRef = rsp.ResolvedReference

	return nil
}