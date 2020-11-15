/*


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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VMProbeSpec contains specification parameters for a Probe.
// +k8s:openapi-gen=true
type VMProbeSpec struct {
	// The job name assigned to scraped metrics by default.
	JobName string `json:"jobName,omitempty"`
	// Specification for the prober to use for probing targets.
	// The prober.URL parameter is required. Targets cannot be probed if left empty.
	VMProberSpec VMProberSpec `json:"vmProberSpec"`
	// The module to use for probing specifying how to probe the target.
	// Example module configuring in the blackbox exporter:
	// https://github.com/prometheus/blackbox_exporter/blob/master/example.yml
	Module string `json:"module,omitempty"`
	// Targets defines a set of static and/or dynamically discovered targets to be probed using the prober.
	Targets VMProbeTargets `json:"targets,omitempty"`
	// Interval at which targets are probed using the configured prober.
	// If not specified Prometheus' global scrape interval is used.
	Interval string `json:"interval,omitempty"`
	// Timeout for scraping metrics from the Prometheus exporter.
	ScrapeTimeout string `json:"scrapeTimeout,omitempty"`
}

// VMProbeTargets defines a set of static and dynamically discovered targets for the prober.
// +k8s:openapi-gen=true
type VMProbeTargets struct {
	// StaticConfig defines static targets which are considers for probing.
	// More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#static_config.
	StaticConfig *VMProbeTargetStaticConfig `json:"staticConfig,omitempty"`
	// Ingress defines the set of dynamically discovered ingress objects which hosts are considered for probing.
	Ingress *ProbeTargetIngress `json:"ingress,omitempty"`
}

// VMProbeTargetStaticConfig defines the set of static targets considered for probing.
// +k8s:openapi-gen=true
type VMProbeTargetStaticConfig struct {
	// Targets is a list of URLs to probe using the configured prober.
	Targets []string `json:"targets"`
	// Labels assigned to all metrics scraped from the targets.
	Labels map[string]string `json:"labels,omitempty"`
}

// ProbeTargetIngress defines the set of Ingress objects considered for probing.
// +k8s:openapi-gen=true
type ProbeTargetIngress struct {
	// Select Ingress objects by labels.
	Selector metav1.LabelSelector `json:"selector,omitempty"`
	// Select Ingress objects by namespace.
	NamespaceSelector NamespaceSelector `json:"namespaceSelector,omitempty"`
	// RelabelConfigs to apply to samples before ingestion.
	// More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config
	RelabelConfigs []*RelabelConfig `json:"relabelingConfigs,omitempty"`
}

// VMProberSpec contains specification parameters for the Prober used for probing.
// +k8s:openapi-gen=true
type VMProberSpec struct {
	// Mandatory URL of the prober.
	URL string `json:"url"`
	// HTTP scheme to use for scraping.
	// Defaults to `http`.
	Scheme string `json:"scheme,omitempty"`
	// Path to collect metrics from.
	// Defaults to `/probe`.
	Path string `json:"path,omitempty"`
}

// VMProbeStatus defines the observed state of VMProbe
type VMProbeStatus struct {
}

//  VMProbe defines a probe for targets, that will be executed with prober,
//  like blackbox exporter.
// It helps to monitor reachability of target with various checks.
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type VMProbe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VMProbeSpec   `json:"spec"`
	Status VMProbeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// VMProbeList contains a list of VMProbe
type VMProbeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMProbe `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMProbe{}, &VMProbeList{})
}
