// Package modifier exposes a header modifier for generating correlation id
package modifier

import (
	"encoding/json"
	"net/http"

	"github.com/google/martian/parse"
	"github.com/google/uuid"
)

type CorrelationHeaderModifier struct {
	HeaderName string `json:"header_name"`
}

func (m *CorrelationHeaderModifier) ModifyRequest(req *http.Request) error {
	if req.Header.Get(m.HeaderName) == "" {
		req.Header.Set(m.HeaderName, uuid.NewString())
	}

	return nil
}

func FromJSON(b []byte) (*parse.Result, error) {
	msg := &CorrelationHeaderModifier{}
	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	return parse.NewResult(msg, []parse.ModifierType{parse.Request})
}
