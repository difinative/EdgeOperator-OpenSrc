/*
Copyright 2022.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Vitals struct {
	UpOrDown   string `json:"upordown,omitempty"`
	FreeMemory string `json:"freememory,omitempty"`
	Temperatur int    `json:"temperature,omitempty"`
	// InferenceServerLastUpdate int    `json:"inferenceserverlastupdate,omitempty"`
}

type Camera struct {
	Resolution string `json:"resolution,omitempty"`
	UpOrDown   string `json:"upordown,omitempty"`
	// IP            string `json:"ip,omitempty"`
}

// ScEdgeSpec defines the desired state of ScEdge
type ScEdgeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ScEdge. Edit scedge_types.go to remove/update
	// Foo string `json:"foo,omitempty"`
	Edgename string            `json:"edgename,omitempty"`
	Type     string            `json:"type,omitempty"`
	Vitals   Vitals            `json:"vitals,omitempty"`
	Cameras  map[string]Camera `json:"cameras,omitempty"`
}

// ScEdgeStatus defines the observed state of ScEdge
type ScEdgeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Vitals  Vitals            `json:"vitals,omitempty"`
	Cameras map[string]Camera `json:"cameras,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="Up/Down?",type="string",JSONPath=`.status.vitals.upordown`
// +kubebuilder:printcolumn:name="Free Memory",type="string",JSONPath=`.status.vitals.freememory`
// +kubebuilder:printcolumn:name="Temperatur",type="string",JSONPath=`.status.vitals.temperature`
// ScEdge is the Schema for the scedges API
type ScEdge struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScEdgeSpec   `json:"spec,omitempty"`
	Status ScEdgeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ScEdgeList contains a list of ScEdge
type ScEdgeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScEdge `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ScEdge{}, &ScEdgeList{})
}
