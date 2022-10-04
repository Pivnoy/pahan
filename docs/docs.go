// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/v1/accept_subsidy": {
            "post": {
                "description": "Accept and link subsidy",
                "tags": [
                    "Posts"
                ],
                "summary": "accept subsidy",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.createAcceptSubsidyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/create_order": {
            "post": {
                "description": "Create and link new order with",
                "tags": [
                    "Posts"
                ],
                "summary": "create new order",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.createOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/create_shipment": {
            "post": {
                "description": "Create and link new shipment",
                "tags": [
                    "Posts"
                ],
                "summary": "create new shipment",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.createShipmentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/create_subsidy": {
            "post": {
                "description": "Create and link subsidy with dependent values",
                "tags": [
                    "Posts"
                ],
                "summary": "create subsidy",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.createSubsidyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/get_components_by_vendor_and_type": {
            "get": {
                "description": "Get all components depend on vendorID and typeID",
                "tags": [
                    "Gets"
                ],
                "summary": "get component",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of a vendor",
                        "name": "vendor-id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id of a type",
                        "name": "type-id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Component"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/get_engineers_by_vendor": {
            "get": {
                "description": "Get all engineers with current vendorID",
                "tags": [
                    "Gets"
                ],
                "summary": "list of engineers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of a vendor",
                        "name": "vendor-id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Engineer"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/get_factories_by_vendor": {
            "get": {
                "description": "Get all factories with current vendorID",
                "tags": [
                    "Gets"
                ],
                "summary": "list of factories",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of a vendor",
                        "name": "vendor-id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Factory"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/get_models": {
            "get": {
                "description": "Get all models",
                "tags": [
                    "Gets"
                ],
                "summary": "list of models",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Model"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/get_subsidies": {
            "get": {
                "description": "Get all subsidies",
                "tags": [
                    "Gets"
                ],
                "summary": "list of subsidies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Subsidy"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/get_types": {
            "get": {
                "description": "Get all types",
                "tags": [
                    "Gets"
                ],
                "summary": "list of types",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Type"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pahan_internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Component": {
            "type": "object",
            "properties": {
                "additional_info": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "type_id": {
                    "type": "integer"
                },
                "vendor_id": {
                    "type": "integer"
                }
            }
        },
        "entity.Engineer": {
            "type": "object",
            "properties": {
                "experience": {
                    "type": "integer"
                },
                "factory_id": {
                    "type": "integer"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                },
                "vendor_id": {
                    "type": "integer"
                }
            }
        },
        "entity.Factory": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "max_workers": {
                    "type": "integer"
                },
                "productivity": {
                    "type": "integer"
                },
                "vendor_id": {
                    "type": "integer"
                }
            }
        },
        "entity.Model": {
            "type": "object",
            "properties": {
                "engineer_id": {
                    "type": "integer"
                },
                "factory_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "prod_cost": {
                    "type": "integer"
                },
                "sales": {
                    "type": "integer"
                },
                "significance": {
                    "type": "integer"
                },
                "vendor_id": {
                    "type": "integer"
                },
                "wheeldrive": {
                    "type": "string"
                }
            }
        },
        "entity.Subsidy": {
            "type": "object",
            "properties": {
                "country_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "require_price": {
                    "type": "integer"
                },
                "required_wd": {
                    "type": "string"
                }
            }
        },
        "entity.Type": {
            "type": "object",
            "properties": {
                "additional_info": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "internal_controller_http_v1.createAcceptSubsidyRequest": {
            "type": "object",
            "properties": {
                "component-bumper-id": {
                    "type": "integer"
                },
                "component-door-id": {
                    "type": "integer"
                },
                "component-engine-id": {
                    "type": "integer"
                },
                "component-transmission-id": {
                    "type": "integer"
                },
                "engineer-id": {
                    "type": "integer"
                },
                "factory-id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "significance": {
                    "type": "integer"
                },
                "subsidy-id": {
                    "type": "integer"
                },
                "vendor-id": {
                    "type": "integer"
                }
            }
        },
        "internal_controller_http_v1.createOrderRequest": {
            "type": "object",
            "properties": {
                "model_id": {
                    "type": "integer"
                },
                "order_type": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "internal_controller_http_v1.createShipmentRequest": {
            "type": "object",
            "properties": {
                "country_id": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "order_id": {
                    "type": "integer"
                }
            }
        },
        "internal_controller_http_v1.createSubsidyRequest": {
            "type": "object",
            "properties": {
                "country-id-by": {
                    "type": "integer"
                },
                "require-price-by": {
                    "type": "number"
                },
                "required-wd-by": {
                    "type": "string"
                }
            }
        },
        "internal_controller_http_v1.errResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "pahan_internal_controller_http_v1.createAcceptSubsidyRequest": {
            "type": "object",
            "properties": {
                "component-bumper-id": {
                    "type": "integer"
                },
                "component-door-id": {
                    "type": "integer"
                },
                "component-engine-id": {
                    "type": "integer"
                },
                "component-transmission-id": {
                    "type": "integer"
                },
                "engineer-id": {
                    "type": "integer"
                },
                "factory-id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "significance": {
                    "type": "integer"
                },
                "subsidy-id": {
                    "type": "integer"
                },
                "vendor-id": {
                    "type": "integer"
                }
            }
        },
        "pahan_internal_controller_http_v1.createOrderRequest": {
            "type": "object",
            "properties": {
                "model_id": {
                    "type": "integer"
                },
                "order_type": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "pahan_internal_controller_http_v1.createShipmentRequest": {
            "type": "object",
            "properties": {
                "country_id": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "order_id": {
                    "type": "integer"
                }
            }
        },
        "pahan_internal_controller_http_v1.createSubsidyRequest": {
            "type": "object",
            "properties": {
                "country-id-by": {
                    "type": "integer"
                },
                "require-price-by": {
                    "type": "number"
                },
                "required-wd-by": {
                    "type": "string"
                }
            }
        },
        "pahan_internal_controller_http_v1.errResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Автомобилестроение в экономике",
	Description:      "Курсовая работа по предмету \"Информационные системы и базы данных\" студента группы P34312, Соловьева Павла",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
