// Copyright 2022 Expedia Group
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	hcv2beta1 "github.com/fluxcd/helm-controller/api/v2beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Chart defines the Helm Chart that should be applied
	// +required
	Chart hcv2beta1.HelmChartTemplate `json:"chart,omitempty"`

	// Interval at which to reconcile the application
	// +required
	Interval metav1.Duration `json:"interval"`

	// HelmReleaseName is the name of the Helm Manifest, also referenced as Release.Name in the Helm Chart.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=53
	// +required
	HelmReleaseName string `json:"helmReleaseName,omitempty"`

	// Timeout is the time required to wait for the individual Kubernetes resources to be running, including passing health checks
	// Defaults to '5m0s'.
	// +optional
	Timeout *metav1.Duration `json:"timeout,omitempty"`

	// Install holds configurations to be applied when the application is created
	// +optional
	Install *hcv2beta1.Install `json:"create,omitempty"`

	// Upgrade holds configurations to be applied when the application is updated
	// +optional
	Upgrade *hcv2beta1.Upgrade `json:"update,omitempty"`

	// Uninstall holds configurations to be applied when the application is uninstalled
	// +optional
	Uninstall *hcv2beta1.Uninstall `json:"delete,omitempty"`

	// Rollback holds configurations to be applied when the current revision of the application fails to run
	// +optional
	Rollback *hcv2beta1.Rollback `json:"rollback,omitempty"`

	// Test holds configurations to be applied for the Helm Tests
	// +optional
	Test *hcv2beta1.Test `json:"test,omitempty"`

	// Data to be consolidated for the Helm Chart's values.yaml file
	// +optional
	Data map[string]string `json:"data,omitempty"`

	// PreRenderer holds custom templating delimiters and a flag to enable helm templating. If helm templating is enabled, custom delimiters must be specified.
	// +optional
	PreRenderer PreRenderer `json:"preRenderer,omitempty"`

	// PostRenderers hold customizations on the Kubernetes resources
	// +optional
	PostRenderers []hcv2beta1.PostRenderer `json:"postRenderers,omitempty"`

	// TTL is a time to live string. The Application resource and all Helm chart applied resources
	// will be deleted after this period.
	// The TTL string is a positive decimal number with a unit suffix,
	// such as "300m" or "1.5h" or "2h45m" or "2d4h45m" or "2d1.5h45m".
	// Valid time units are "m", "h", "d" for mins, hours or days respectively.
	// +kubebuilder:validation:Pattern=`(^(([1-9]|([0-9]*(.[0-9]+)))(m|h|d){1})+$)`
	TTL string `json:"ttl,omitempty"`
}

// ApplicationStatus defines the observed state of Application
type ApplicationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ObservedGeneration is the last observed generation.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions holds the conditions for the Application.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// LastAppliedRevision is the revision of the last successfully applied source.
	// +optional
	LastAppliedRevision string `json:"lastAppliedRevision,omitempty"`

	// LastAttemptedRevision is the revision of the last reconciliation attempt.
	// +optional
	LastAttemptedRevision string `json:"lastAttemptedRevision,omitempty"`

	// LastAttemptedValuesChecksum is the SHA1 checksum of the values of the last
	// reconciliation attempt.
	// +optional
	LastAttemptedValuesChecksum string `json:"lastAttemptedValuesChecksum,omitempty"`

	// LastApplicationRevision is the revision of the last successful Application.
	// +optional
	LastReleaseRevision int `json:"lastReleaseRevision,omitempty"`

	// HelmChart is the namespaced name of the HelmChart resource created by
	// the operator.
	// +optional
	HelmChart string `json:"helmChart,omitempty"`

	// Failures is the reconciliation failure count against the latest desired
	// state. It is reset after a successful reconciliation.
	// +optional
	Failures int64 `json:"failures,omitempty"`

	// InstallFailures is the install failure count against the latest desired
	// state. It is reset after a successful reconciliation.
	// +optional
	InstallFailures int64 `json:"installFailures,omitempty"`

	// UpgradeFailures is the upgrade failure count against the latest desired
	// state. It is reset after a successful reconciliation.
	// +optional
	UpgradeFailures int64 `json:"upgradeFailures,omitempty"`
}

// +genclient
// +genclient:Namespaced
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=app
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status",description=""
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].message",description=""
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description=""

// Application is the Schema for the applications API
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ApplicationList contains a list of Application
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
