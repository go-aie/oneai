package api

import (
	"context"
	"encoding/json"
	"fmt"
)

//go:generate kungen -out=../controller -flat=false ./api.go Embedding

type Embedding interface {
	//kun:op POST /
	//kun:param __ in=header name=Authorization required=true
	//kun:body req
	//kun:success body=resp
	Encode(ctx context.Context, req *Request) (resp *Response, err error)
}

type Request struct {
	Model string                  `json:"model,omitempty"`
	Input OneOf[string, []string] `json:"input,omitempty"`
	User  string                  `json:"user,omitempty"`
}

type Response struct {
	Model  string  `json:"model,omitempty"`
	Object string  `json:"object,omitempty"`
	Data   []*Data `json:"data,omitempty"`
}

type Data struct {
	Object    string    `json:"object,omitempty"`
	Embedding []float64 `json:"embedding,omitempty"`
	Index     int       `json:"index,omitempty"`
}

type OneOf[A, B any] struct{ Value any }

func (o *OneOf[A, B]) UnmarshalJSON(data []byte) error {
	var a A
	if err := json.Unmarshal(data, &a); err == nil {
		o.Value = a
		return nil
	}

	var b B
	if err := json.Unmarshal(data, &b); err == nil {
		o.Value = b
		return nil
	}

	return fmt.Errorf("json: cannot unmarshal object into Go value of type %T or %T", a, b)
}
