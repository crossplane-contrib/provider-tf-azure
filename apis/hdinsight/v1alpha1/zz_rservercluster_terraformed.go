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

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this RServerCluster
func (mg *RServerCluster) GetTerraformResourceType() string {
	return "azurerm_hdinsight_rserver_cluster"
}

// GetConnectionDetailsMapping for this RServerCluster
func (tr *RServerCluster) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"gateway[*].password": "spec.forProvider.gateway[*].passwordSecretRef", "roles[*].edge_node[*].password": "spec.forProvider.roles[*].edgeNode[*].passwordSecretRef", "roles[*].head_node[*].password": "spec.forProvider.roles[*].headNode[*].passwordSecretRef", "roles[*].worker_node[*].password": "spec.forProvider.roles[*].workerNode[*].passwordSecretRef", "roles[*].zookeeper_node[*].password": "spec.forProvider.roles[*].zookeeperNode[*].passwordSecretRef", "storage_account[*].storage_account_key": "spec.forProvider.storageAccount[*].storageAccountKeySecretRef"}
}

// GetObservation of this RServerCluster
func (tr *RServerCluster) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RServerCluster
func (tr *RServerCluster) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RServerCluster
func (tr *RServerCluster) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RServerCluster
func (tr *RServerCluster) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RServerCluster
func (tr *RServerCluster) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this RServerCluster using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RServerCluster) LateInitialize(attrs []byte) (bool, error) {
	params := &RServerClusterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RServerCluster) GetTerraformSchemaVersion() int {
	return 0
}
