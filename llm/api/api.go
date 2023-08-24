package api

import (
	"context"
)

//go:generate kungen -out=../controller -flat=false ./api.go LLM

type LLM interface {
	//kun:op POST /completions
	//kun:param __ in=header name=Authorization required=true
	//kun:body req
	//kun:success body=resp
	Chat(ctx context.Context, req *Request) (resp *Response, err error)
}

type Request struct {
	Model    string     `json:"model,omitempty"`
	Messages []*Message `json:"messages,omitempty"`

	Temperature float64 `json:"temperature,omitempty"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	N           int     `json:"n,omitempty"`

	PresencePenalty  float64 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	User string `json:"user,omitempty"`
}

type Response struct {
	ID      string    `json:"id,omitempty"`
	Object  string    `json:"object,omitempty"`
	Created int       `json:"created,omitempty"`
	Model   string    `json:"model,omitempty"`
	Choices []*Choice `json:"choices,omitempty"`
}

type Choice struct {
	Index        int      `json:"index,omitempty"`
	Message      *Message `json:"message,omitempty"`
	FinishReason string   `json:"finish_reason,omitempty"`
}

type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}
