// Code generated by kun; DO NOT EDIT.
// github.com/RussellLuo/kun

package endpoint

import (
	"context"

	"github.com/RussellLuo/kun/pkg/httpoption"
	"github.com/RussellLuo/validating/v3"
	"github.com/go-aie/oneai/vectorstore/api"
	"github.com/go-kit/kit/endpoint"
)

type DeleteRequest struct {
	Vendor    string   `json:"vendor"`
	SourceIDs []string `json:"source_i_ds"`
}

// ValidateDeleteRequest creates a validator for DeleteRequest.
func ValidateDeleteRequest(newSchema func(*DeleteRequest) validating.Schema) httpoption.Validator {
	return httpoption.FuncValidator(func(value interface{}) error {
		req := value.(*DeleteRequest)
		return httpoption.Validate(newSchema(req))
	})
}

type DeleteResponse struct {
	Err error `json:"-"`
}

func (r *DeleteResponse) Body() interface{} { return r }

// Failed implements endpoint.Failer.
func (r *DeleteResponse) Failed() error { return r.Err }

// MakeEndpointOfDelete creates the endpoint for s.Delete.
func MakeEndpointOfDelete(s api.VectorStore) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*DeleteRequest)
		err := s.Delete(
			ctx,
			req.Vendor,
			req.SourceIDs...,
		)
		return &DeleteResponse{
			Err: err,
		}, nil
	}
}

type QueryRequest struct {
	Vendor string    `json:"vendor"`
	Vector []float64 `json:"vector"`
	TopK   int       `json:"top_k"`
}

// ValidateQueryRequest creates a validator for QueryRequest.
func ValidateQueryRequest(newSchema func(*QueryRequest) validating.Schema) httpoption.Validator {
	return httpoption.FuncValidator(func(value interface{}) error {
		req := value.(*QueryRequest)
		return httpoption.Validate(newSchema(req))
	})
}

type QueryResponse struct {
	Similarities []*api.Similarity `json:"similarities"`
	Err          error             `json:"-"`
}

func (r *QueryResponse) Body() interface{} { return r }

// Failed implements endpoint.Failer.
func (r *QueryResponse) Failed() error { return r.Err }

// MakeEndpointOfQuery creates the endpoint for s.Query.
func MakeEndpointOfQuery(s api.VectorStore) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*QueryRequest)
		similarities, err := s.Query(
			ctx,
			req.Vendor,
			req.Vector,
			req.TopK,
		)
		return &QueryResponse{
			Similarities: similarities,
			Err:          err,
		}, nil
	}
}

type UpsertRequest struct {
	Vendor    string          `json:"vendor"`
	Documents []*api.Document `json:"documents"`
}

// ValidateUpsertRequest creates a validator for UpsertRequest.
func ValidateUpsertRequest(newSchema func(*UpsertRequest) validating.Schema) httpoption.Validator {
	return httpoption.FuncValidator(func(value interface{}) error {
		req := value.(*UpsertRequest)
		return httpoption.Validate(newSchema(req))
	})
}

type UpsertResponse struct {
	Err error `json:"-"`
}

func (r *UpsertResponse) Body() interface{} { return r }

// Failed implements endpoint.Failer.
func (r *UpsertResponse) Failed() error { return r.Err }

// MakeEndpointOfUpsert creates the endpoint for s.Upsert.
func MakeEndpointOfUpsert(s api.VectorStore) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*UpsertRequest)
		err := s.Upsert(
			ctx,
			req.Vendor,
			req.Documents,
		)
		return &UpsertResponse{
			Err: err,
		}, nil
	}
}
