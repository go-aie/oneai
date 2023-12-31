// Code generated by kun; DO NOT EDIT.
// github.com/RussellLuo/kun

package http

import (
	"reflect"

	"github.com/RussellLuo/kun/pkg/oas2"
	"github.com/go-aie/oneai/embedding/controller/endpoint"
)

var (
	base = `swagger: "2.0"
info:
  title: "No Title"
  version: "0.0.0"
  description: ""
  license:
    name: "MIT"
host: "example.com"
basePath: "/"
schemes:
  - "https"
consumes:
  - "application/json"
produces:
  - "application/json"
`

	paths = `
paths:
  /:
    post:
      description: ""
      summary: ""
      operationId: "Encode"
      parameters:
        - name: Authorization
          in: header
          required: true
          type: string
          description: ""
        - name: body
          in: body
          schema:
            $ref: "#/definitions/EncodeRequestBody"
      %s
`
)

func getResponses(schema oas2.Schema) []oas2.OASResponses {
	return []oas2.OASResponses{
		oas2.GetOASResponses(schema, "Encode", 200, &endpoint.EncodeResponse{}),
	}
}

func getDefinitions(schema oas2.Schema) map[string]oas2.Definition {
	defs := make(map[string]oas2.Definition)

	oas2.AddDefinition(defs, "EncodeRequestBody", reflect.ValueOf((&endpoint.EncodeRequest{}).Req))
	oas2.AddResponseDefinitions(defs, schema, "Encode", 200, (&endpoint.EncodeResponse{}).Body())

	return defs
}

func OASv2APIDoc(schema oas2.Schema) string {
	resps := getResponses(schema)
	paths := oas2.GenPaths(resps, paths)

	defs := getDefinitions(schema)
	definitions := oas2.GenDefinitions(defs)

	return base + paths + definitions
}
