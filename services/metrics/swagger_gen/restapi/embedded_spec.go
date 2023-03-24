// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Metrics API",
    "title": "Metrics",
    "version": "1.0.0"
  },
  "paths": {
    "/v1/metrics": {
      "get": {
        "security": [],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "metrics"
        ],
        "summary": "get all metrics",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/LoginResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "LoginResponse": {
      "type": "object"
    }
  },
  "security": [
    {
      "hasRole": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Metrics API",
    "title": "Metrics",
    "version": "1.0.0"
  },
  "paths": {
    "/v1/metrics": {
      "get": {
        "security": [],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "metrics"
        ],
        "summary": "get all metrics",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/LoginResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "LoginResponse": {
      "type": "object"
    }
  },
  "security": [
    {
      "hasRole": []
    }
  ]
}`))
}
