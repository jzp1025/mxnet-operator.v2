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

package validation

import (
	//"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	mxv2 "github.com/kubeflow/mx-operator.v2/pkg/apis/mxnet/v1alpha2"
	//"github.com/kubeflow/tf-operator/pkg/util"
)

// ValidateAlphaTwoMXJobSpec checks that the v1alpha2.MXJobSpec is valid.
func ValidateAlphaTwoMXJobSpec(c *mxv2.MXJobSpec) error {
	if c.MXReplicaSpecs == nil {
		return fmt.Errorf("MXJobSpec is NULL")
	}
	for rType, value := range c.MXReplicaSpecs {
		if value == nil || len(value.Template.Spec.Containers) == 0 {
			return fmt.Errorf("MXJobSpec is not valid")
		}
		//Make sure the image is defined in the container
		numNamedMXNet := 0
		for _, container := range value.Template.Spec.Containers {
			if container.Image == "" {
				log.Warn("Image is undefined in the container")
				return fmt.Errorf("MXJobSpec is not valid")
			}
			if container.Name == mxv2.DefaultContainerName {
				numNamedMXNet++
			}
		}
		//Make sure there has at least one container named "mxnet"
		if numNamedMXNet == 0 {
			log.Warnf("There is no container named mxnet in %v", rType)
			return fmt.Errorf("MXJobSpec is not valid")
		}
	}
	return nil
}

