// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package message

import "fmt"
import "encoding/json"

type PlaybookRunResponseMessageYamlEventsElem struct {
	// Counter corresponds to the JSON schema field "counter".
	Counter int `json:"counter"`

	// EndLine corresponds to the JSON schema field "end_line".
	EndLine int `json:"end_line"`

	// Event corresponds to the JSON schema field "event".
	Event string `json:"event"`

	// EventData corresponds to the JSON schema field "event_data".
	EventData *PlaybookRunResponseMessageYamlEventsElemEventData `json:"event_data,omitempty"`

	// StartLine corresponds to the JSON schema field "start_line".
	StartLine int `json:"start_line"`

	// Stdout corresponds to the JSON schema field "stdout".
	Stdout *string `json:"stdout,omitempty"`

	// Uuid corresponds to the JSON schema field "uuid".
	Uuid string `json:"uuid"`
}

type PlaybookRunResponseMessageYamlEventsElemEventData struct {
	// CrcDispatcherCorrelationId corresponds to the JSON schema field
	// "crc_dispatcher_correlation_id".
	CrcDispatcherCorrelationId *string `json:"crc_dispatcher_correlation_id,omitempty"`

	// Host corresponds to the JSON schema field "host".
	Host *string `json:"host,omitempty"`

	// Playbook corresponds to the JSON schema field "playbook".
	Playbook *string `json:"playbook,omitempty"`

	// PlaybookUuid corresponds to the JSON schema field "playbook_uuid".
	PlaybookUuid *string `json:"playbook_uuid,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PlaybookRunResponseMessageYamlEventsElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["counter"]; !ok || v == nil {
		return fmt.Errorf("field counter: required")
	}
	if v, ok := raw["end_line"]; !ok || v == nil {
		return fmt.Errorf("field end_line: required")
	}
	if v, ok := raw["event"]; !ok || v == nil {
		return fmt.Errorf("field event: required")
	}
	if v, ok := raw["start_line"]; !ok || v == nil {
		return fmt.Errorf("field start_line: required")
	}
	if v, ok := raw["uuid"]; !ok || v == nil {
		return fmt.Errorf("field uuid: required")
	}
	type Plain PlaybookRunResponseMessageYamlEventsElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PlaybookRunResponseMessageYamlEventsElem(plain)
	return nil
}

type PlaybookRunResponseMessageYaml struct {
	// Account corresponds to the JSON schema field "account".
	Account string `json:"account"`

	// B64Identity corresponds to the JSON schema field "b64_identity".
	B64Identity string `json:"b64_identity"`

	// Events corresponds to the JSON schema field "events".
	Events []PlaybookRunResponseMessageYamlEventsElem `json:"events"`

	// RequestId corresponds to the JSON schema field "request_id".
	RequestId string `json:"request_id"`

	// UploadTimestamp corresponds to the JSON schema field "upload_timestamp".
	UploadTimestamp string `json:"upload_timestamp"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PlaybookRunResponseMessageYaml) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["account"]; !ok || v == nil {
		return fmt.Errorf("field account: required")
	}
	if v, ok := raw["b64_identity"]; !ok || v == nil {
		return fmt.Errorf("field b64_identity: required")
	}
	if v, ok := raw["events"]; !ok || v == nil {
		return fmt.Errorf("field events: required")
	}
	if v, ok := raw["request_id"]; !ok || v == nil {
		return fmt.Errorf("field request_id: required")
	}
	if v, ok := raw["upload_timestamp"]; !ok || v == nil {
		return fmt.Errorf("field upload_timestamp: required")
	}
	type Plain PlaybookRunResponseMessageYaml
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PlaybookRunResponseMessageYaml(plain)
	return nil
}
