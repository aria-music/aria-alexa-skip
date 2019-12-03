package alexa

import (
	"encoding/json"
	"time"
)

type RootRequest struct {
	Request json.RawMessage
	Context `json:"context"`
}

type Context struct {
	System `json:"system"`
}

type System struct {
	Application `json:"application"`
}

type Application struct {
	ApplicationID string
}

type Request struct {
	Type      string
	Timestamp time.Time
}

type IntentRequest struct {
	Request
	Intent `json:"intent"`
}

type Intent struct {
	Name string
}
