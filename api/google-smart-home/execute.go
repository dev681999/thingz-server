package googlesmarthome

import "encoding/json"

func UnmarshalExecute(data []byte) (Execute, error) {
	var r Execute
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Execute) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Execute struct {
	RequestID string         `json:"requestId,omitempty"`
	Inputs    []ExecuteInput `json:"inputs,omitempty"`
}

type ExecuteInput struct {
	Intent  string         `json:"intent,omitempty"`
	Payload ExecutePayload `json:"payload,omitempty"`
}

type ExecutePayload struct {
	Commands []Command `json:"commands,omitempty"`
}

type Command struct {
	Devices   []ExecuteDevice `json:"devices,omitempty"`
	Execution []Execution     `json:"execution,omitempty"`
}

type ExecuteDevice struct {
	ID string `json:"id,omitempty"`
}

type Execution struct {
	Command string                 `json:"command,omitempty"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

type Params struct {
	On bool `json:"on,omitempty"`
}

func UnmarshalExecuteResponse(data []byte) (ExecuteResponse, error) {
	var r ExecuteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExecuteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ExecuteResponse struct {
	RequestID string                 `json:"requestId,omitempty"`
	Payload   ExecuteResponsePayload `json:"payload,omitempty"`
}

type ExecuteResponsePayload struct {
	Commands []ExecuteResponseCommand `json:"commands,omitempty"`
}

type ExecuteResponseCommand struct {
	IDS       []string               `json:"ids,omitempty"`
	Status    string                 `json:"status,omitempty"`
	States    map[string]interface{} `json:"states,omitempty"`
	ErrorCode string                 `json:"errorCode,omitempty"`
}
