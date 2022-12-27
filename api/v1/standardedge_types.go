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

// StandardEdgeSpec defines the desired state of StandardEdge
type StandardEdgeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of StandardEdge. Edit standardedge_types.go to remove/update
	// Foo string `json:"foo,omitempty"`

	Edgename string `json:"edgename,omitempty"`
	Type     string `json:"type,omitempty"`
	Vitals   Vitals `json:"vitals,omitempty"`
	// Cameras  map[string]Camera `json:"cameras,omitempty"`
}

// StandardEdgeStatus defines the observed state of StandardEdge
type StandardEdgeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Vitals Vitals `json:"vitals,omitempty"`
	// Cameras map[string]Camera `json:"cameras,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="Up/Down?",type="string",JSONPath=`.status.vitals.upordown`
// +kubebuilder:printcolumn:name="Free Memory",type="string",JSONPath=`.status.vitals.freememory`
// +kubebuilder:printcolumn:name="Temperature",type="string",JSONPath=`.status.vitals.temperature`

// StandardEdge is the Schema for the standardedges API
type StandardEdge struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StandardEdgeSpec   `json:"spec,omitempty"`
	Status StandardEdgeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StandardEdgeList contains a list of StandardEdge
type StandardEdgeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StandardEdge `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StandardEdge{}, &StandardEdgeList{})
}
