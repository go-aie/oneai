// Code generated by kun; DO NOT EDIT.
// github.com/RussellLuo/kun

package http

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/RussellLuo/kun/pkg/httpcodec"
	"github.com/go-aie/oneai/vectorstore/api"
	"github.com/go-aie/oneai/vectorstore/controller/endpoint"
)

type HTTPClient struct {
	codecs     httpcodec.Codecs
	httpClient *http.Client
	scheme     string
	host       string
	pathPrefix string
}

func NewHTTPClient(codecs httpcodec.Codecs, httpClient *http.Client, baseURL string) (*HTTPClient, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &HTTPClient{
		codecs:     codecs,
		httpClient: httpClient,
		scheme:     u.Scheme,
		host:       u.Host,
		pathPrefix: strings.TrimSuffix(u.Path, "/"),
	}, nil
}

func (c *HTTPClient) Delete(ctx context.Context, vendor string, documentIDs ...string) (err error) {
	codec := c.codecs.EncodeDecoder("Delete")

	path := "/delete"
	u := &url.URL{
		Scheme: c.scheme,
		Host:   c.host,
		Path:   c.pathPrefix + path,
	}

	reqBody := struct {
		Vendor      string   `json:"vendor"`
		DocumentIDs []string `json:"document_i_ds"`
	}{
		Vendor:      vendor,
		DocumentIDs: documentIDs,
	}
	reqBodyReader, headers, err := codec.EncodeRequestBody(&reqBody)
	if err != nil {
		return err
	}

	_req, err := http.NewRequestWithContext(ctx, "POST", u.String(), reqBodyReader)
	if err != nil {
		return err
	}

	for k, v := range headers {
		_req.Header.Set(k, v)
	}
	for _, v := range codec.EncodeRequestParam("__", nil) {
		_req.Header.Add("Authorization", v)
	}

	_resp, err := c.httpClient.Do(_req)
	if err != nil {
		return err
	}
	defer _resp.Body.Close()

	if _resp.StatusCode < http.StatusOK || _resp.StatusCode > http.StatusNoContent {
		var respErr error
		err := codec.DecodeFailureResponse(_resp.Body, &respErr)
		if err == nil {
			err = respErr
		}
		return err
	}

	return nil
}

func (c *HTTPClient) Query(ctx context.Context, vendor string, vector []float64, topK int) (similarities []*api.Similarity, err error) {
	codec := c.codecs.EncodeDecoder("Query")

	path := "/query"
	u := &url.URL{
		Scheme: c.scheme,
		Host:   c.host,
		Path:   c.pathPrefix + path,
	}

	reqBody := struct {
		Vendor string    `json:"vendor"`
		Vector []float64 `json:"vector"`
		TopK   int       `json:"top_k"`
	}{
		Vendor: vendor,
		Vector: vector,
		TopK:   topK,
	}
	reqBodyReader, headers, err := codec.EncodeRequestBody(&reqBody)
	if err != nil {
		return nil, err
	}

	_req, err := http.NewRequestWithContext(ctx, "POST", u.String(), reqBodyReader)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		_req.Header.Set(k, v)
	}
	for _, v := range codec.EncodeRequestParam("__", nil) {
		_req.Header.Add("Authorization", v)
	}

	_resp, err := c.httpClient.Do(_req)
	if err != nil {
		return nil, err
	}
	defer _resp.Body.Close()

	if _resp.StatusCode < http.StatusOK || _resp.StatusCode > http.StatusNoContent {
		var respErr error
		err := codec.DecodeFailureResponse(_resp.Body, &respErr)
		if err == nil {
			err = respErr
		}
		return nil, err
	}

	respBody := &endpoint.QueryResponse{}
	err = codec.DecodeSuccessResponse(_resp.Body, respBody.Body())
	if err != nil {
		return nil, err
	}
	return respBody.Similarities, nil
}

func (c *HTTPClient) Upsert(ctx context.Context, vendor string, chunks map[string][]*api.Chunk) (err error) {
	codec := c.codecs.EncodeDecoder("Upsert")

	path := "/upsert"
	u := &url.URL{
		Scheme: c.scheme,
		Host:   c.host,
		Path:   c.pathPrefix + path,
	}

	reqBody := struct {
		Vendor string                  `json:"vendor"`
		Chunks map[string][]*api.Chunk `json:"chunks"`
	}{
		Vendor: vendor,
		Chunks: chunks,
	}
	reqBodyReader, headers, err := codec.EncodeRequestBody(&reqBody)
	if err != nil {
		return err
	}

	_req, err := http.NewRequestWithContext(ctx, "POST", u.String(), reqBodyReader)
	if err != nil {
		return err
	}

	for k, v := range headers {
		_req.Header.Set(k, v)
	}
	for _, v := range codec.EncodeRequestParam("__", nil) {
		_req.Header.Add("Authorization", v)
	}

	_resp, err := c.httpClient.Do(_req)
	if err != nil {
		return err
	}
	defer _resp.Body.Close()

	if _resp.StatusCode < http.StatusOK || _resp.StatusCode > http.StatusNoContent {
		var respErr error
		err := codec.DecodeFailureResponse(_resp.Body, &respErr)
		if err == nil {
			err = respErr
		}
		return err
	}

	return nil
}