package googlesmarthome

import "encoding/json"

func UnmarshalQuery(data []byte) (Query, error) {
	var r Query
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Query) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Query struct {
	RequestID string       `json:"requestId"`
	Inputs    []QueryInput `json:"inputs"`
}

type QueryInput struct {
	Intent  string       `json:"intent"`
	Payload QueryPayload `json:"payload"`
}

type QueryPayload struct {
	Devices []QueryDevice `json:"devices"`
}

type QueryDevice struct {
	ID string `json:"id"`
}

func UnmarshalQueryResponse(data []byte) (QueryResponse, error) {
	var r QueryResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *QueryResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type QueryResponse struct {
	RequestID string               `json:"requestId"`
	Payload   QueryResponsePayload `json:"payload"`
}

type QueryResponsePayload struct {
	Devices map[string]map[string]interface{} `json:"devices"`
}
