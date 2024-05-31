package sdk

import (
	"encoding/json"
	"io"
)

type APIResponse[A any] struct {
	Payload A `json:"payload"`
}

type JSON[A any] struct {
	body io.Reader
}

func (t *JSON[A]) Unmarshal() (*A, error) {
	body, err := io.ReadAll(t.body)
	if err != nil {
		return nil, err
	}
	var response APIResponse[A]
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response.Payload, nil
}

func NewTwiplaJSON[A any](r io.Reader) *JSON[A] {
	return &JSON[A]{
		body: r,
	}
}
