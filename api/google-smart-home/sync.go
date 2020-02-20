package googlesmarthome

import "encoding/json"

func UnmarshalSyncResponse(data []byte) (SyncResponse, error) {
	var r SyncResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SyncResponse struct {
	RequestID string      `json:"requestId"`
	Payload   SyncPayload `json:"payload"`
}

type SyncPayload struct {
	AgentUserID string       `json:"agentUserId"`
	Devices     []SyncDevice `json:"devices"`
}

type SyncDevice struct {
	ID              string   `json:"id"`
	Type            string   `json:"type"`
	Traits          []string `json:"traits"`
	Name            Name     `json:"name"`
	WillReportState bool     `json:"willReportState"`
	RoomHint        string   `json:"roomHint"`
}

type Name struct {
	Name string `json:"name,omitempty"`
}
