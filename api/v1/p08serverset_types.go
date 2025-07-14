/*
Copyright 2025.

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
	v1 "k8s.io/api/core/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// P08ServerSetSpec defines the desired state of P08ServerSet.
type P08ServerSetSpec struct {
	// Replicas is the desired number of replicas of the given Template.
	// These are replicas in the sense that they are instantiations of the
	// same Template, but individual replicas also have a consistent identity.
	// If unspecified, defaults to 1.
	// TODO: Consider a rename of this field.
	// +optional
	Replicas int32 `json:"replicas,omitempty"`

	// Selector is a label query over pods that should match the replica count.
	// If empty, defaulted to labels on the pod template.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`

	// Template is the object that describes the pod that will be created if
	// insufficient replicas are detected. Each pod stamped out by the StatefulSet
	// will fulfill this Template, but have a unique identity from the rest
	// of the StatefulSet.
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Template v1.PodTemplateSpec `json:"template"`

	// VolumeClaimTemplates is a list of claims that pods are allowed to reference.
	// The StatefulSet controller is responsible for mapping network identities to
	// claims in a way that maintains the identity of a pod. Every claim in
	// this list must have at least one matching (by name) volumeMount in one
	// container in the template. A claim in this list takes precedence over
	// any volumes in the template, with the same name.
	// TODO: Define the behavior if a claim already exists with the same name.
	// +optional
	VolumeClaimTemplates []v1.PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty"`

	// ServiceName is the name of the service that governs this StatefulSet.
	// This service must exist before the StatefulSet, and is responsible for
	// the network identity of the set. Pods get DNS/hostnames that follow the
	// pattern: pod-specific-string.serviceName.default.svc.cluster.local
	// where "pod-specific-string" is managed by the StatefulSet controller.
	ServiceName string `json:"serviceName"`

	// PodManagementPolicy controls how pods are created during initial scale up,
	// when replacing pods on nodes, or when scaling down. The default policy is
	// `OrderedReady`, where pods are created in increasing order (pod-0, then
	// pod-1, etc) and the controller will wait until each pod is ready before
	// continuing. When scaling down, the pods are removed in the opposite order.
	// The alternative policy is `Parallel` which will create pods in parallel
	// to match the desired scale without waiting, and on scale down will delete
	// all pods at once.
	// +optional
	PodManagementPolicy PodManagementPolicyType `json:"podManagementPolicy,omitempty"`

	// updateStrategy indicates the StatefulSetUpdateStrategy that will be
	// employed to update Pods in the StatefulSet when a revision is made to
	// Template.
	UpdateStrategy StatefulSetUpdateStrategy `json:"updateStrategy"`

	// revisionHistoryLimit is the maximum number of revisions that will
	// be maintained in the StatefulSet's revision history. The revision history
	// consists of all revisions not represented by a currently applied
	// StatefulSetSpec version. The default value is 10.
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit"`

	// Minimum number of seconds for which a newly created pod should be ready
	// without any of its container crashing for it to be considered available.
	// Defaults to 0 (pod will be considered available as soon as it is ready)
	// This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate.
	// +optional
	MinReadySeconds int32 `json:"minReadySeconds,omitempty"`

	// PersistentVolumeClaimRetentionPolicy describes the policy used for PVCs created from
	// the StatefulSet VolumeClaimTemplates. This requires the
	// StatefulSetAutoDeletePVC feature gate to be enabled, which is alpha.
	// +optional
	PersistentVolumeClaimRetentionPolicy *StatefulSetPersistentVolumeClaimRetentionPolicy `json:"persistentVolumeClaimRetentionPolicy,omitempty"`
}


// P08ServerSetStatus defines the observed state of P08ServerSet.
type P08ServerSetStatus struct {
	// observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the
	// StatefulSet's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// replicas is the number of Pods created by the StatefulSet controller.
	Replicas int32 `json:"replicas"`

	// readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
	ReadyReplicas int32 `json:"readyReplicas"`

	// currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version
	// indicated by currentRevision.
	CurrentReplicas int32 `json:"currentReplicas"`

	// updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version
	// indicated by updateRevision.
	UpdatedReplicas int32 `json:"updatedReplicas"`

	// currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the
	// sequence [0,currentReplicas).
	CurrentRevision string `json:"currentRevision"`

	// updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence
	// [replicas-updatedReplicas,replicas)
	UpdateRevision string `json:"updateRevision"`

	// collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller
	// uses this field as a collision avoidance mechanism when it needs to create the name for the
	// newest ControllerRevision.
	// +optional
	CollisionCount *int32 `json:"collisionCount,omitempty"`

	// Represents the latest available observations of a statefulset's current state.
	Conditions []StatefulSetCondition `json:"conditions"`

	// Total number of available pods (ready for at least minReadySeconds) targeted by this statefulset.
	// This is a beta field and requires enabling StatefulSetMinReadySeconds feature gate.
	// +optional
	AvailableReplicas int32 `json:"avaliableReplicas,omitempty"`

	// compatible to crd scale
	LabelSelector string `json:"labelSelector,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=p08ss
// +kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas,selectorpath=.status.labelSelector

// P08ServerSet is the Schema for the p08serversets API.
type P08ServerSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   P08ServerSetSpec   `json:"spec,omitempty"`
	Status P08ServerSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// P08ServerSetList contains a list of P08ServerSet.
type P08ServerSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []P08ServerSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&P08ServerSet{}, &P08ServerSetList{})
}
