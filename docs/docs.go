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
        "/publish": {
            "post": {
                "description": "Mock an MQTT weather publish, fields are formatted as Protobuf and sent via the MQTT broker",
                "tags": [
                    "mqtt"
                ],
                "summary": "Mock an MQTT weather publish",
                "parameters": [
                    {
                        "description": "json",
                        "name": "request-body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.publishModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.publishModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.publishModel": {
            "type": "object",
            "required": [
                "topic"
            ],
            "properties": {
                "humidity": {
                    "type": "number",
                    "example": 50
                },
                "pressure": {
                    "type": "number",
                    "example": 50
                },
                "temperature": {
                    "type": "number",
                    "example": 100
                },
                "topic": {
                    "type": "string",
                    "example": "topic/telemetry"
                },
                "water_amount": {
                    "type": "number",
                    "example": 10
                },
                "wind_direction": {
                    "type": "number",
                    "example": 60
                },
                "wind_speed": {
                    "type": "number",
                    "example": 25
                }
            }
        },
        "util.ErrorModel": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "something went wrong"
                },
                "status": {
                    "type": "integer",
                    "example": 400
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Weatherman Telemetry Server APIs",
	Description:      "The comprehensive list of all Weatherman Telemetry Server APIs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
