/*
Copyright 2021.

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

package v1

import (
	commonv1 "github.com/kubeflow/common/pkg/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KubeflowJobSpec defines the desired state of KubeflowJob
type KubeflowJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	RunPolicy    *commonv1.RunPolicy                            `json:"run_policy,omitempty"`
	ReplicaSpecs map[commonv1.ReplicaType]*commonv1.ReplicaSpec `json:"replica_specs"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KubeflowJob is the Schema for the kubeflowjobs API
type KubeflowJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeflowJobSpec    `json:"spec,omitempty"`
	Status commonv1.JobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KubeflowJobList contains a list of KubeflowJob
type KubeflowJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeflowJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeflowJob{}, &KubeflowJobList{})
}
