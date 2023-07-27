/*
Copyright 2023.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/project-codeflare/codeflare-operator/api/codeflare/v1alpha1"
	codeflarev1alpha1 "github.com/project-codeflare/codeflare-operator/client/applyconfiguration/codeflare/v1alpha1"
)

// FakeMCADs implements MCADInterface
type FakeMCADs struct {
	Fake *FakeCodeflareV1alpha1
	ns   string
}

var mcadsResource = v1alpha1.SchemeGroupVersion.WithResource("mcads")

var mcadsKind = v1alpha1.SchemeGroupVersion.WithKind("MCAD")

// Get takes name of the mCAD, and returns the corresponding mCAD object, and an error if there is any.
func (c *FakeMCADs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MCAD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mcadsResource, c.ns, name), &v1alpha1.MCAD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MCAD), err
}

// List takes label and field selectors, and returns the list of MCADs that match those selectors.
func (c *FakeMCADs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MCADList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mcadsResource, mcadsKind, c.ns, opts), &v1alpha1.MCADList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MCADList{ListMeta: obj.(*v1alpha1.MCADList).ListMeta}
	for _, item := range obj.(*v1alpha1.MCADList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mCADs.
func (c *FakeMCADs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mcadsResource, c.ns, opts))

}

// Create takes the representation of a mCAD and creates it.  Returns the server's representation of the mCAD, and an error, if there is any.
func (c *FakeMCADs) Create(ctx context.Context, mCAD *v1alpha1.MCAD, opts v1.CreateOptions) (result *v1alpha1.MCAD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mcadsResource, c.ns, mCAD), &v1alpha1.MCAD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MCAD), err
}

// Update takes the representation of a mCAD and updates it. Returns the server's representation of the mCAD, and an error, if there is any.
func (c *FakeMCADs) Update(ctx context.Context, mCAD *v1alpha1.MCAD, opts v1.UpdateOptions) (result *v1alpha1.MCAD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mcadsResource, c.ns, mCAD), &v1alpha1.MCAD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MCAD), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMCADs) UpdateStatus(ctx context.Context, mCAD *v1alpha1.MCAD, opts v1.UpdateOptions) (*v1alpha1.MCAD, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mcadsResource, "status", c.ns, mCAD), &v1alpha1.MCAD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MCAD), err
}

// Delete takes name of the mCAD and deletes it. Returns an error if one occurs.
func (c *FakeMCADs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(mcadsResource, c.ns, name, opts), &v1alpha1.MCAD{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMCADs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mcadsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MCADList{})
	return err
}

// Patch applies the patch and returns the patched mCAD.
func (c *FakeMCADs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MCAD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mcadsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MCAD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MCAD), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied mCAD.
func (c *FakeMCADs) Apply(ctx context.Context, mCAD *codeflarev1alpha1.MCADApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.MCAD, err error) {
	if mCAD == nil {
		return nil, fmt.Errorf("mCAD provided to Apply must not be nil")
	}
	data, err := json.Marshal(mCAD)
	if err != nil {
		return nil, err
	}
	name := mCAD.Name
	if name == nil {
		return nil, fmt.Errorf("mCAD.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mcadsResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.MCAD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MCAD), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeMCADs) ApplyStatus(ctx context.Context, mCAD *codeflarev1alpha1.MCADApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.MCAD, err error) {
	if mCAD == nil {
		return nil, fmt.Errorf("mCAD provided to Apply must not be nil")
	}
	data, err := json.Marshal(mCAD)
	if err != nil {
		return nil, err
	}
	name := mCAD.Name
	if name == nil {
		return nil, fmt.Errorf("mCAD.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mcadsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.MCAD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MCAD), err
}
