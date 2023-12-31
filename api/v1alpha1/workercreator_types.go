/*
Copyright 2023 clementreiffers.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WorkerCreatorSpec defines the desired state of WorkerCreator
type WorkerCreatorSpec struct {
	WorkerDeploymentId string `json:"worker-deployment-id"`
	WorkerDefinitionId string `json:"worker-definition-id"`
}

type Port struct {
	PortName   string `json:"portName"`
	PortNumber int32  `json:"portNumber"`
}

// WorkerCreatorStatus defines the observed state of WorkerCreator
type WorkerCreatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// WorkerCreator is the Schema for the workercreators API
type WorkerCreator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkerCreatorSpec   `json:"spec,omitempty"`
	Status WorkerCreatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WorkerCreatorList contains a list of WorkerCreator
type WorkerCreatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkerCreator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkerCreator{}, &WorkerCreatorList{})
}
