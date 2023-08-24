// Code generated by kun; DO NOT EDIT.
// github.com/RussellLuo/kun

package http

import (
	"context"
	"net/http"

	"github.com/RussellLuo/kun/pkg/httpcodec"
	"github.com/RussellLuo/kun/pkg/httpoption"
	"github.com/RussellLuo/kun/pkg/oas2"
	"github.com/go-aie/oneai/vectorstore/api"
	"github.com/go-aie/oneai/vectorstore/controller/endpoint"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func NewHTTPRouter(svc api.VectorStore, codecs httpcodec.Codecs, opts ...httpoption.Option) chi.Router {
	r := chi.NewRouter()
	options := httpoption.NewOptions(opts...)

	r.Method("GET", "/api", oas2.Handler(OASv2APIDoc, options.ResponseSchema()))

	var codec httpcodec.Codec
	var validator httpoption.Validator
	var kitOptions []kithttp.ServerOption

	codec = codecs.EncodeDecoder("Delete")
	validator = options.RequestValidator("Delete")
	r.Method(
		"POST", "/delete",
		kithttp.NewServer(
			endpoint.MakeEndpointOfDelete(svc),
			decodeDeleteRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("Query")
	validator = options.RequestValidator("Query")
	r.Method(
		"POST", "/query",
		kithttp.NewServer(
			endpoint.MakeEndpointOfQuery(svc),
			decodeQueryRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("Upsert")
	validator = options.RequestValidator("Upsert")
	r.Method(
		"POST", "/upsert",
		kithttp.NewServer(
			endpoint.MakeEndpointOfUpsert(svc),
			decodeUpsertRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	return r
}

func decodeDeleteRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req endpoint.DeleteRequest

		if err := codec.DecodeRequestBody(r, &_req); err != nil {
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

func decodeQueryRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req endpoint.QueryRequest

		if err := codec.DecodeRequestBody(r, &_req); err != nil {
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

func decodeUpsertRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req endpoint.UpsertRequest

		if err := codec.DecodeRequestBody(r, &_req); err != nil {
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