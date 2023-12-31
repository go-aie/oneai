// Code generated by kun; DO NOT EDIT.
// github.com/RussellLuo/kun

package http

import (
	"context"
	"net/http"

	"github.com/RussellLuo/kun/pkg/httpcodec"
	"github.com/RussellLuo/kun/pkg/httpoption"
	"github.com/RussellLuo/kun/pkg/oas2"
	"github.com/go-aie/oneai/llm/api"
	"github.com/go-aie/oneai/llm/controller/endpoint"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func NewHTTPRouter(svc api.LLM, codecs httpcodec.Codecs, opts ...httpoption.Option) chi.Router {
	r := chi.NewRouter()
	options := httpoption.NewOptions(opts...)

	r.Method("GET", "/api", oas2.Handler(OASv2APIDoc, options.ResponseSchema()))

	var codec httpcodec.Codec
	var validator httpoption.Validator
	var kitOptions []kithttp.ServerOption

	codec = codecs.EncodeDecoder("Chat")
	validator = options.RequestValidator("Chat")
	r.Method(
		"POST", "/completions",
		kithttp.NewServer(
			endpoint.MakeEndpointOfChat(svc),
			decodeChatRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	return r
}

func decodeChatRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req endpoint.ChatRequest

		if err := codec.DecodeRequestBody(r, &_req.Req); err != nil {
			return nil, err
		}

		__ := r.Header.Values("Authorization")
		if err := codec.DecodeRequestParam("__", __, nil); err != nil {
			return nil, err
		}

		if err := validator.Validate(&_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}
