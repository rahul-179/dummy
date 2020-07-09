/*
Copyright 2019 The Tekton Authors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterTriggerBindingLister helps list ClusterTriggerBindings.
type ClusterTriggerBindingLister interface {
	// List lists all ClusterTriggerBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterTriggerBinding, err error)
	// Get retrieves the ClusterTriggerBinding from the index for a given name.
	Get(name string) (*v1alpha1.ClusterTriggerBinding, error)
	ClusterTriggerBindingListerExpansion
}

// clusterTriggerBindingLister implements the ClusterTriggerBindingLister interface.
type clusterTriggerBindingLister struct {
	indexer cache.Indexer
}

// NewClusterTriggerBindingLister returns a new ClusterTriggerBindingLister.
func NewClusterTriggerBindingLister(indexer cache.Indexer) ClusterTriggerBindingLister {
	return &clusterTriggerBindingLister{indexer: indexer}
}

// List lists all ClusterTriggerBindings in the indexer.
func (s *clusterTriggerBindingLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterTriggerBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterTriggerBinding))
	})
	return ret, err
}

// Get retrieves the ClusterTriggerBinding from the index for a given name.
func (s *clusterTriggerBindingLister) Get(name string) (*v1alpha1.ClusterTriggerBinding, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustertriggerbinding"), name)
	}
	return obj.(*v1alpha1.ClusterTriggerBinding), nil
}
