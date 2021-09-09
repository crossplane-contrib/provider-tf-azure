/*
Copyright 2020 The Crossplane Authors.

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
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/terrajet/pkg/conversion"
	"github.com/crossplane-contrib/terrajet/pkg/conversion/lateinit"
)

// GetTerraformResourceType returns Terraform resource type for this HdinsightStormCluster
func (mg *HdinsightStormCluster) GetTerraformResourceType() string {
	return "azurerm_hdinsight_storm_cluster"
}

// GetTerraformResourceIdField returns Terraform identifier field for this HdinsightStormCluster
func (tr *HdinsightStormCluster) GetTerraformResourceIdField() string {
	return "id"
}

// GetObservation of this HdinsightStormCluster
func (tr *HdinsightStormCluster) GetObservation() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Status.AtProvider)
}

// SetObservation for this HdinsightStormCluster
func (tr *HdinsightStormCluster) SetObservation(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Status.AtProvider)
}

// GetParameters of this HdinsightStormCluster
func (tr *HdinsightStormCluster) GetParameters() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Spec.ForProvider)
}

// SetParameters for this HdinsightStormCluster
func (tr *HdinsightStormCluster) SetParameters(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Spec.ForProvider)
}

// LateInitialize this HdinsightStormCluster using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HdinsightStormCluster) LateInitialize(tfState []byte) (bool, error) {
	stateObject := &HdinsightStormClusterParameters{}
	if err := conversion.TFParser.Unmarshal(tfState, stateObject); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state for late initialization")
	}

	return lateinit.LateInitializeFromResponse("", &tr.Spec.ForProvider, stateObject,
		lateinit.ZeroValueJSONOmitEmptyFilter(lateinit.CNameWildcard), lateinit.ZeroElemPtrFilter(lateinit.CNameWildcard))
}
