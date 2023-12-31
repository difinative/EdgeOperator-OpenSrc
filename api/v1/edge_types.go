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

// EdgeSpec defines the desired state of Edge
type EdgeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Edge. Edit edge_types.go to remove/update
	// Name    string `json:"name,omitempty"`
	MacId            string   `json:"macId,omitempty"`
	Usecase          string   `json:"usecase,omitempty"`
	HealthPercentage int      `json:"healthPercentage,omitempty"`
	Processes        []string `json:"processes,omitempty"`
}

// EdgeStatus defines the observed state of Edge
type EdgeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	LUT                   string             `json:"LUT,omitempty"`
	HealthPercentage      string             `json:"healthPercentage,omitempty"`
	VitalsStatsPercentage string             `json:"vitalsPercentage,omitempty"`
	Uptime                int                `json:"uptime,omitempty"`
	UpOrDown              string             `json:"up/down,omitempty"`
	SqNet                 string             `json:"sqnet,omitempty"`
	HealthVitals          HealthVitalsStatus `json:"healthVitals,omitempty"`
	Usecase_Vitals        map[string]string  `json:"usecaseVitals,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="Usecase",type="string",JSONPath=`.spec.usecase`
// +kubebuilder:printcolumn:name="Up/Down?",type="string",JSONPath=`.status.up/down`
// +kubebuilder:printcolumn:name="HealthPercentage",type="string",JSONPath=`.status.healthPercentage`
// +kubebuilder:printcolumn:name="Uptime",type="integer",JSONPath=`.status.uptime`
// +kubebuilder:printcolumn:name="VitalsPercentage",type="string",JSONPath=`.status.vitalsPercentage`
// Edge is the Schema for the edges API
type Edge struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeSpec   `json:"spec,omitempty"`
	Status EdgeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EdgeList contains a list of Edge
type EdgeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Edge `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Edge{}, &EdgeList{})
}
