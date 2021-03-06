/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package fake

import (
	api "k8s.io/kubernetes/pkg/api"
	core "k8s.io/kubernetes/pkg/client/testing/core"
	labels "k8s.io/kubernetes/pkg/labels"
	watch "k8s.io/kubernetes/pkg/watch"
)

// FakeSecrets implements SecretInterface
type FakeSecrets struct {
	Fake *FakeCore
	ns   string
}

func (c *FakeSecrets) Create(secret *api.Secret) (result *api.Secret, err error) {
	obj, err := c.Fake.
		Invokes(core.NewCreateAction("secrets", c.ns, secret), &api.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Secret), err
}

func (c *FakeSecrets) Update(secret *api.Secret) (result *api.Secret, err error) {
	obj, err := c.Fake.
		Invokes(core.NewUpdateAction("secrets", c.ns, secret), &api.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Secret), err
}

func (c *FakeSecrets) Delete(name string, options *api.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewDeleteAction("secrets", c.ns, name), &api.Secret{})

	return err
}

func (c *FakeSecrets) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	action := core.NewDeleteCollectionAction("secrets", c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &api.SecretList{})
	return err
}

func (c *FakeSecrets) Get(name string) (result *api.Secret, err error) {
	obj, err := c.Fake.
		Invokes(core.NewGetAction("secrets", c.ns, name), &api.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Secret), err
}

func (c *FakeSecrets) List(opts api.ListOptions) (result *api.SecretList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewListAction("secrets", c.ns, opts), &api.SecretList{})

	if obj == nil {
		return nil, err
	}

	label := opts.LabelSelector
	if label == nil {
		label = labels.Everything()
	}
	list := &api.SecretList{}
	for _, item := range obj.(*api.SecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secrets.
func (c *FakeSecrets) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewWatchAction("secrets", c.ns, opts))

}
