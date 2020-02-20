package googlesmarthome

import "encoding/json"

const (
	IntentQuery = "action.devices.QUERY"

	IntentSync = "action.devices.SYNC"

	IntentExecute = "action.devices.EXECUTE"
)

func UnmarshalRequest(data []byte) (Request, error) {
	var r Request
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Request) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Request struct {
	RequestID string  `json:"requestId"`
	Inputs    []Input `json:"inputs"`
}

type Input struct {
	Intent string `json:"intent"`
}
