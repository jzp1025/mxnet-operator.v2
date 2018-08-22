// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

// This file was automatically generated by informer-gen

package kubeflow

import (
	internalinterfaces "github.com/kubeflow/mxnet-operator.v2/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/kubeflow/mxnet-operator.v2/pkg/client/informers/externalversions/kubeflow/v1alpha2"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1alpha2 provides access to shared informers for resources in V1alpha2.
	V1alpha2() v1alpha2.Interface
}

type group struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &group{f}
}

// V1alpha2 returns a new v1alpha2.Interface.
func (g *group) V1alpha2() v1alpha2.Interface {
	return v1alpha2.New(g.SharedInformerFactory)
}
