/*
Copyright 2023.

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

// BarSpec defines the desired state of Bar
type BarSpec struct {

	// Quantity defines the number associated with the bar resource
	// Json:"foo" defines how this object will be refered to elsewhere
	Quantity int `json:"quantity"`
	// Foo references the Foo owner which is then passed as a string
	Foo string `json:"foo"`
}

// BarStatus defines the observed state of Bar
type BarStatus struct {
	Conditions []metav1.Condition `json:"conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Bar is the Schema for the bars API
type Bar struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BarSpec   `json:"spec,omitempty"`
	Status BarStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BarList contains a list of Bar
type BarList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bar `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bar{}, &BarList{})
}
