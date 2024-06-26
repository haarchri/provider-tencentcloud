/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-tencentcloud/apis/cvm/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Snapshot.
func (mg *Snapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageIDRef,
		Selector:     mg.Spec.ForProvider.StorageIDSelector,
		To: reference.To{
			List:    &StorageList{},
			Managed: &Storage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageID")
	}
	mg.Spec.ForProvider.StorageID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.StorageIDRef,
		Selector:     mg.Spec.InitProvider.StorageIDSelector,
		To: reference.To{
			List:    &StorageList{},
			Managed: &Storage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StorageID")
	}
	mg.Spec.InitProvider.StorageID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StorageIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SnapshotPolicyAttachment.
func (mg *SnapshotPolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnapshotPolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SnapshotPolicyIDRef,
		Selector:     mg.Spec.ForProvider.SnapshotPolicyIDSelector,
		To: reference.To{
			List:    &SnapshotPolicyList{},
			Managed: &SnapshotPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnapshotPolicyID")
	}
	mg.Spec.ForProvider.SnapshotPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnapshotPolicyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageIDRef,
		Selector:     mg.Spec.ForProvider.StorageIDSelector,
		To: reference.To{
			List:    &StorageList{},
			Managed: &Storage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageID")
	}
	mg.Spec.ForProvider.StorageID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SnapshotPolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SnapshotPolicyIDRef,
		Selector:     mg.Spec.InitProvider.SnapshotPolicyIDSelector,
		To: reference.To{
			List:    &SnapshotPolicyList{},
			Managed: &SnapshotPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SnapshotPolicyID")
	}
	mg.Spec.InitProvider.SnapshotPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SnapshotPolicyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.StorageIDRef,
		Selector:     mg.Spec.InitProvider.StorageIDSelector,
		To: reference.To{
			List:    &StorageList{},
			Managed: &Storage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StorageID")
	}
	mg.Spec.InitProvider.StorageID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StorageIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StorageAttachment.
func (mg *StorageAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &v1alpha1.InstanceList{},
			Managed: &v1alpha1.Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageIDRef,
		Selector:     mg.Spec.ForProvider.StorageIDSelector,
		To: reference.To{
			List:    &StorageList{},
			Managed: &Storage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageID")
	}
	mg.Spec.ForProvider.StorageID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &v1alpha1.InstanceList{},
			Managed: &v1alpha1.Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.StorageIDRef,
		Selector:     mg.Spec.InitProvider.StorageIDSelector,
		To: reference.To{
			List:    &StorageList{},
			Managed: &Storage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StorageID")
	}
	mg.Spec.InitProvider.StorageID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StorageIDRef = rsp.ResolvedReference

	return nil
}
