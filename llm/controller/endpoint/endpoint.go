// Code generated by kun; DO NOT EDIT.
// github.com/RussellLuo/kun

package endpoint

import (
	"context"

	"github.com/RussellLuo/kun/pkg/httpoption"
	"github.com/RussellLuo/validating/v3"
	"github.com/go-aie/oneai/llm/api"
	"github.com/go-kit/kit/endpoint"
)

type ChatRequest struct {
	Req *api.Request `json:"req"`
}

// ValidateChatRequest creates a validator for ChatRequest.
func ValidateChatRequest(newSchema func(*ChatRequest) validating.Schema) httpoption.Validator {
	return httpoption.FuncValidator(func(value interface{}) error {
		req := value.(*ChatRequest)
		return httpoption.Validate(newSchema(req))
	})
}

type ChatResponse struct {
	Resp *api.Response `json:"resp"`
	Err  error         `json:"-"`
}

func (r *ChatResponse) Body() interface{} { return &r.Resp }

// Failed implements endpoint.Failer.
func (r *ChatResponse) Failed() error { return r.Err }

// MakeEndpointOfChat creates the endpoint for s.Chat.
func MakeEndpointOfChat(s api.LLM) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*ChatRequest)
		resp, err := s.Chat(
			ctx,
			req.Req,
		)
		return &ChatResponse{
			Resp: resp,
			Err:  err,
		}, nil
	}
}