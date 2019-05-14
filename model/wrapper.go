package model

import "encoding/json"

type Wrapper struct {
	Kind string          `json:"kind"`
	Raw  json.RawMessage `json:"raw"`
}
