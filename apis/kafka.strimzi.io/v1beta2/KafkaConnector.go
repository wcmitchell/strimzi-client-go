// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package v1beta2

import (
	"encoding/json"
	"fmt"
	"reflect"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

		metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// KafkaConnector
type KafkaConnector struct {
	metav1.TypeMeta   `json:",inline"`

	// APIVersion defines the versioned schema of this representation of an object.
	// Servers should convert recognized schemas to the latest internal value, and may
	// reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty" mapstructure:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to. Cannot
	// be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty" yaml:"kind,omitempty" mapstructure:"kind,omitempty"`

	// Metadata corresponds to the JSON schema field "metadata".
	Metadata *apiextensions.JSON `json:"metadata,omitempty" yaml:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// The specification of the Kafka Connector.
	Spec *KafkaConnectorSpec `json:"spec,omitempty" yaml:"spec,omitempty" mapstructure:"spec,omitempty"`

	// The status of the Kafka Connector.
	Status *KafkaConnectorStatus `json:"status,omitempty" yaml:"status,omitempty" mapstructure:"status,omitempty"`
}

// +kubebuilder:object:root=true
// KafkaConnectorList contains a list of instances.
type KafkaConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// A list of KafkaConnector objects.
	Items []KafkaConnector `json:"items,omitempty"`
}

//type KafkaConnectorMetadata map[string]interface{}

// The specification of the Kafka Connector.
type KafkaConnectorSpec struct {
	// Configuration for altering offsets.
	AlterOffsets *KafkaConnectorSpecAlterOffsets `json:"alterOffsets,omitempty" yaml:"alterOffsets,omitempty" mapstructure:"alterOffsets,omitempty"`

	// Automatic restart of connector and tasks configuration.
	AutoRestart *KafkaConnectorSpecAutoRestart `json:"autoRestart,omitempty" yaml:"autoRestart,omitempty" mapstructure:"autoRestart,omitempty"`

	// The Class for the Kafka Connector.
	Class *string `json:"class,omitempty" yaml:"class,omitempty" mapstructure:"class,omitempty"`

	// The Kafka Connector configuration. The following properties cannot be set:
	// name, connector.class, tasks.max.
	Config *apiextensions.JSON `json:"config,omitempty" yaml:"config,omitempty" mapstructure:"config,omitempty"`

	// Configuration for listing offsets.
	ListOffsets *KafkaConnectorSpecListOffsets `json:"listOffsets,omitempty" yaml:"listOffsets,omitempty" mapstructure:"listOffsets,omitempty"`

	// Whether the connector should be paused. Defaults to false.
	Pause *bool `json:"pause,omitempty" yaml:"pause,omitempty" mapstructure:"pause,omitempty"`

	// The state the connector should be in. Defaults to running.
	State *KafkaConnectorSpecState `json:"state,omitempty" yaml:"state,omitempty" mapstructure:"state,omitempty"`

	// The maximum number of tasks for the Kafka Connector.
	TasksMax *int32 `json:"tasksMax,omitempty" yaml:"tasksMax,omitempty" mapstructure:"tasksMax,omitempty"`
}

// Configuration for altering offsets.
type KafkaConnectorSpecAlterOffsets struct {
	// Reference to the ConfigMap where the new offsets are stored.
	FromConfigMap KafkaConnectorSpecAlterOffsetsFromConfigMap `json:"fromConfigMap" yaml:"fromConfigMap" mapstructure:"fromConfigMap"`
}

// Reference to the ConfigMap where the new offsets are stored.
type KafkaConnectorSpecAlterOffsetsFromConfigMap struct {
	// Name corresponds to the JSON schema field "name".
	Name *string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`
}

// Automatic restart of connector and tasks configuration.
type KafkaConnectorSpecAutoRestart struct {
	// Whether automatic restart for failed connectors and tasks should be enabled or
	// disabled.
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty" mapstructure:"enabled,omitempty"`

	// The maximum number of connector restarts that the operator will try. If the
	// connector remains in a failed state after reaching this limit, it must be
	// restarted manually by the user. Defaults to an unlimited number of restarts.
	MaxRestarts *int32 `json:"maxRestarts,omitempty" yaml:"maxRestarts,omitempty" mapstructure:"maxRestarts,omitempty"`
}

// The Kafka Connector configuration. The following properties cannot be set: name,
// connector.class, tasks.max.
//type KafkaConnectorSpecConfig map[string]interface{}

// Configuration for listing offsets.
type KafkaConnectorSpecListOffsets struct {
	// Reference to the ConfigMap where the list of offsets will be written to.
	ToConfigMap KafkaConnectorSpecListOffsetsToConfigMap `json:"toConfigMap" yaml:"toConfigMap" mapstructure:"toConfigMap"`
}

// Reference to the ConfigMap where the list of offsets will be written to.
type KafkaConnectorSpecListOffsetsToConfigMap struct {
	// Name corresponds to the JSON schema field "name".
	Name *string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`
}

type KafkaConnectorSpecState string

const KafkaConnectorSpecStatePaused KafkaConnectorSpecState = "paused"
const KafkaConnectorSpecStateRunning KafkaConnectorSpecState = "running"
const KafkaConnectorSpecStateStopped KafkaConnectorSpecState = "stopped"

// The status of the Kafka Connector.
type KafkaConnectorStatus struct {
	// The auto restart status.
	AutoRestart *KafkaConnectorStatusAutoRestart `json:"autoRestart,omitempty" yaml:"autoRestart,omitempty" mapstructure:"autoRestart,omitempty"`

	// List of status conditions.
	Conditions []KafkaConnectorStatusConditionsElem `json:"conditions,omitempty" yaml:"conditions,omitempty" mapstructure:"conditions,omitempty"`

	// The connector status, as reported by the Kafka Connect REST API.
	ConnectorStatus *apiextensions.JSON `json:"connectorStatus,omitempty" yaml:"connectorStatus,omitempty" mapstructure:"connectorStatus,omitempty"`

	// The generation of the CRD that was last reconciled by the operator.
	ObservedGeneration *int32 `json:"observedGeneration,omitempty" yaml:"observedGeneration,omitempty" mapstructure:"observedGeneration,omitempty"`

	// The maximum number of tasks for the Kafka Connector.
	TasksMax *int32 `json:"tasksMax,omitempty" yaml:"tasksMax,omitempty" mapstructure:"tasksMax,omitempty"`

	// The list of topics used by the Kafka Connector.
	Topics []string `json:"topics,omitempty" yaml:"topics,omitempty" mapstructure:"topics,omitempty"`
}

// The auto restart status.
type KafkaConnectorStatusAutoRestart struct {
	// The name of the connector being restarted.
	ConnectorName *string `json:"connectorName,omitempty" yaml:"connectorName,omitempty" mapstructure:"connectorName,omitempty"`

	// The number of times the connector or task is restarted.
	Count *int32 `json:"count,omitempty" yaml:"count,omitempty" mapstructure:"count,omitempty"`

	// The last time the automatic restart was attempted. The required format is
	// 'yyyy-MM-ddTHH:mm:ssZ' in the UTC time zone.
	LastRestartTimestamp *string `json:"lastRestartTimestamp,omitempty" yaml:"lastRestartTimestamp,omitempty" mapstructure:"lastRestartTimestamp,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaConnectorSpecState) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_KafkaConnectorSpecState {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_KafkaConnectorSpecState, v)
	}
	*j = KafkaConnectorSpecState(v)
	return nil
}

type KafkaConnectorStatusConditionsElem struct {
	// Last time the condition of a type changed from one status to another. The
	// required format is 'yyyy-MM-ddTHH:mm:ssZ', in the UTC time zone.
	LastTransitionTime *string `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty" mapstructure:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about the condition's last
	// transition.
	Message *string `json:"message,omitempty" yaml:"message,omitempty" mapstructure:"message,omitempty"`

	// The reason for the condition's last transition (a single word in CamelCase).
	Reason *string `json:"reason,omitempty" yaml:"reason,omitempty" mapstructure:"reason,omitempty"`

	// The status of the condition, either True, False or Unknown.
	Status *string `json:"status,omitempty" yaml:"status,omitempty" mapstructure:"status,omitempty"`

	// The unique identifier of a condition, used to distinguish between other
	// conditions in the resource.
	Type *string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
}

// The connector status, as reported by the Kafka Connect REST API.
//type KafkaConnectorStatusConnectorStatus map[string]interface{}

var enumValues_KafkaConnectorSpecState = []interface{}{
	"paused",
	"stopped",
	"running",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaConnectorSpecListOffsets) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["toConfigMap"]; !ok || v == nil {
		return fmt.Errorf("field toConfigMap in KafkaConnectorSpecListOffsets: required")
	}
	type Plain KafkaConnectorSpecListOffsets
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = KafkaConnectorSpecListOffsets(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaConnectorSpecAlterOffsets) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["fromConfigMap"]; !ok || v == nil {
		return fmt.Errorf("field fromConfigMap in KafkaConnectorSpecAlterOffsets: required")
	}
	type Plain KafkaConnectorSpecAlterOffsets
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = KafkaConnectorSpecAlterOffsets(plain)
	return nil
}
