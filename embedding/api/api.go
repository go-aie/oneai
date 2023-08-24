package api

import (
	"context"
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
	Model string   `json:"model,omitempty"`
	Input []string `json:"input,omitempty"`
	User  string   `json:"user,omitempty"`
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
