// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.swagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.swagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Simple To Do List API",
    "version": "0.1.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "todos"
        ],
        "operationId": "find",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "name": "limit",
            "in": "formData",
            "required": true,
            "allowEmptyValue": true
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "X-Rate-Limit",
            "in": "header",
            "required": true
          },
          {
            "type": "array",
            "items": {
              "type": "integer",
              "format": "int32"
            },
            "collectionFormat": "multi",
            "name": "tags",
            "in": "formData",
            "required": true,
            "allowEmptyValue": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/item"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "todos"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "put": {
        "tags": [
          "todos"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "todos"
        ],
        "operationId": "destroyOne",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "item": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-todolist-token",
      "in": "header"
    }
  },
  "security": [
    {
      "key": []
    }
  ],
  "x-schemes": [
    "unix"
  ]
}`))
}
