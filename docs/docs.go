// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/onedep": {
            "post": {
                "description": "Create a new deposition by uploading experiments, files, and metadata to OneDep API.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deposition"
                ],
                "summary": "Create a new deposition to OneDep",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User's email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Experiment type (e.g., single-particle analysis)",
                        "name": "experiments",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "collectionFormat": "multi",
                        "description": "File(s) to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Scientific metadata as a JSON string",
                        "name": "metadata",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "File metadata as a JSON string",
                        "name": "fileMetadata",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deposition ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error response",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "Create a new deposition by uploading experiments, files, and metadata to OneDep API.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "version"
                ],
                "summary": "Return current version",
                "responses": {
                    "200": {
                        "description": "Depositior version",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error response",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "api/v1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "OpenEm Depositor API",
	Description:      "Rest API for communication between SciCat frontend and depositor backend. Backend service enables deposition of datasets to OneDep API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
