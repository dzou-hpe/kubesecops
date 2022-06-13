/*
Copyright 2017 The Kubernetes Authors.

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Zap is a specification for a Zap resource
type Zap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZapSpec   `json:"spec"`
	Status ZapStatus `json:"status"`
}

// ZapSpec is the spec for a Zap resource
type ZapSpec struct {
	ScanName   string `json:"scanName"`
	AppUrl     string `json:"appUrl"`
	OpenApiUrl string `json:"openApiUrl"`
}

// ZapStatus is the status for a Zap resource
type ZapStatus struct {
	AvailableReports int32 `json:"availableReports"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ZapList is a list of Zap resources
type ZapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Zap `json:"items"`
}
