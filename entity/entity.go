package entity

import "encoding/json"

type Message struct {
	Action  string          `json:"action"`
	UserID  string          `json:"user_id"`
	Payload json.RawMessage `json:"payload"`
}

type MessagePayload struct {
	Text string `json:"text"`
}
