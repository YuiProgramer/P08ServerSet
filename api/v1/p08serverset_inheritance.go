package v1


import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/apimachinery/pkg/util/intstr"
)


// PodManagementPolicyType defines the policy for creating pods under a stateful set.
type PodManagementPolicyType string

const (
	// OrderedReadyPodManagement will create pods in strictly increasing order on
	// scale up and strictly decreasing order on scale down, progressing only when
	// the previous pod is ready or terminated. At most one pod will be changed
	// at any time.
	OrderedReadyPodManagement PodManagementPolicyType = "OrderedReady"
	// ParallelPodManagement will create and delete pods as soon as the stateful set
	// replica count is changed, and will not wait for pods to be ready or complete
	// termination.
	ParallelPodManagement PodManagementPolicyType = "Parallel"
)

// StatefulSetUpdateStrategy indicates the strategy that the StatefulSet
// controller will use to perform updates. It includes any additional parameters
// necessary to perform the update for the indicated strategy.
type StatefulSetUpdateStrategy struct {
	// Type indicates the type of the StatefulSetUpdateStrategy.
	Type StatefulSetUpdateStrategyType `json:"type"`
	// RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.
	RollingUpdate *RollingUpdateStatefulSetStrategy `json:"rollingUpdate,,omitempty"`
}

// StatefulSetUpdateStrategyType is a string enumeration type that enumerates
// all possible update strategies for the StatefulSet controller.
type StatefulSetUpdateStrategyType string

const (
	// RollingUpdateStatefulSetStrategyType indicates that update will be
	// applied to all Pods in the StatefulSet with respect to the StatefulSet
	// ordering constraints. When a scale operation is performed with this
	// strategy, new Pods will be created from the specification version indicated
	// by the StatefulSet's updateRevision.
	RollingUpdateStatefulSetStrategyType StatefulSetUpdateStrategyType = "RollingUpdate"
	// OnDeleteStatefulSetStrategyType triggers the legacy behavior. Version
	// tracking and ordered rolling restarts are disabled. Pods are recreated
	// from the StatefulSetSpec when they are manually deleted. When a scale
	// operation is performed with this strategy,specification version indicated
	// by the StatefulSet's currentRevision.
	OnDeleteStatefulSetStrategyType StatefulSetUpdateStrategyType = "OnDelete"
)

// RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.
type RollingUpdateStatefulSetStrategy struct {
	// Partition indicates the ordinal at which the StatefulSet should be partitioned
	// for updates. During a rolling update, all pods from ordinal Replicas-1 to
	// Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched.
	// This is helpful in being able to do a canary based deployment. The default value is 0.
	Partition int32 `json:"partition"`
	// The maximum number of pods that can be unavailable during the update.
	// Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
	// Absolute number is calculated from percentage by rounding up. This can not be 0.
	// Defaults to 1. This field is alpha-level and is only honored by servers that enable the
	// MaxUnavailableStatefulSet feature. The field applies to all pods in the range 0 to
	// Replicas-1. That means if there is any unavailable pod in the range 0 to Replicas-1, it
	// will be counted towards MaxUnavailable.
	// +optional
	MaxUnavailable *intstr.IntOrString `json:"maxUnvailable"`
}

// PersistentVolumeClaimRetentionPolicyType is a string enumeration of the policies that will determine
// when volumes from the VolumeClaimTemplates will be deleted when the controlling StatefulSet is
// deleted or scaled down.
type PersistentVolumeClaimRetentionPolicyType string

const (
	// RetainPersistentVolumeClaimRetentionPolicyType is the default
	// PersistentVolumeClaimRetentionPolicy and specifies that
	// PersistentVolumeClaims associated with StatefulSet VolumeClaimTemplates
	// will not be deleted.
	RetainPersistentVolumeClaimRetentionPolicyType PersistentVolumeClaimRetentionPolicyType = "Retain"
	// DeletePersistentVolumeClaimRetentionPolicyType specifies that
	// PersistentVolumeClaims associated with StatefulSet VolumeClaimTemplates
	// will be deleted in the scenario specified in
	// StatefulSetPersistentVolumeClaimPolicy.
	DeletePersistentVolumeClaimRetentionPolicyType PersistentVolumeClaimRetentionPolicyType = "Delete"
)

// StatefulSetPersistentVolumeClaimRetentionPolicy describes the policy used for PVCs
// created from the StatefulSet VolumeClaimTemplates.
type StatefulSetPersistentVolumeClaimRetentionPolicy struct {
	// WhenDeleted specifies what happens to PVCs created from StatefulSet
	// VolumeClaimTemplates when the StatefulSet is deleted. The default policy
	// of `Retain` causes PVCs to not be affected by StatefulSet deletion. The
	// `Delete` policy causes those PVCs to be deleted.
	WhenDeleted PersistentVolumeClaimRetentionPolicyType `json:"whenDeleted,omitempty"`
	// WhenScaled specifies what happens to PVCs created from StatefulSet
	// VolumeClaimTemplates when the StatefulSet is scaled down. The default
	// policy of `Retain` causes PVCs to not be affected by a scaledown. The
	// `Delete` policy causes the associated PVCs for any excess pods above
	// the replica count to be deleted.
	WhenScaled PersistentVolumeClaimRetentionPolicyType `json:"whenScaled,omitempty"`
}


// StatefulSetConditionType describes the condition types of StatefulSets.
type StatefulSetConditionType string

// TODO: Add valid condition types for Statefulsets.

// StatefulSetCondition describes the state of a statefulset at a certain point.
type StatefulSetCondition struct {
	// Type of statefulset condition.
	Type StatefulSetConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status api.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// The reason for the condition's last transition.
	Reason string `json:"reason"`
	// A human readable message indicating details about the transition.
	Message string `json:"message"`
}
