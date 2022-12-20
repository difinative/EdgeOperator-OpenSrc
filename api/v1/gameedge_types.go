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

// GameEdgeSpec defines the desired state of GameEdge
type GameEdgeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of GameEdge. Edit gameedge_types.go to remove/update
	// Foo string `json:"foo,omitempty"`
	Edgename string            `json:"edgename,omitempty"`
	Type     string            `json:"type,omitempty"`
	Vitals   Vitals            `json:"vitals,omitempty"`
	Cameras  map[string]Camera `json:"cameras,omitempty"`
}

// GameEdgeStatus defines the observed state of GameEdge
type GameEdgeStatus struct {
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
// GameEdge is the Schema for the gameedges API
type GameEdge struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GameEdgeSpec   `json:"spec,omitempty"`
	Status GameEdgeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GameEdgeList contains a list of GameEdge
type GameEdgeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GameEdge `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GameEdge{}, &GameEdgeList{})
}
