// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "semenovem@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Авторизация пользователя",
                "parameters": [
                    {
                        "description": "Логин/пароль",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_rest_controller_auth.loginForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_rest_controller_auth.loginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_semenovem_portal_pkg_failing.Response"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Выход из авторизованной сессии пользователя",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_rest_controller_auth.loginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_semenovem_portal_pkg_failing.Response"
                        }
                    }
                }
            }
        },
        "/auth/onetime": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Создание ссылки для одноразовой авторизации",
                "parameters": [
                    {
                        "description": "данные для создания сессии",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_rest_controller_auth.onetimeAuthForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_rest_controller_auth.onetimeAuthResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_semenovem_portal_pkg_failing.Response"
                        }
                    }
                }
            }
        },
        "/auth/onetime/:entry_id": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Логин по одноразовой ссылке",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id сессии авторизации",
                        "name": "session_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_rest_controller_auth.loginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_semenovem_portal_pkg_failing.Response"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Обновление токена авторизации",
                "parameters": [
                    {
                        "type": "string",
                        "description": "asdfasdf",
                        "name": "refresh-token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_rest_controller_auth.loginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_semenovem_portal_pkg_failing.Response"
                        }
                    }
                }
            }
        },
        "/vehicles": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicles"
                ],
                "summary": "Ищет автомобили по фильтру",
                "parameters": [
                    {
                        "type": "string",
                        "name": "end_time",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "name": "slug[]",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "start_time",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "name": "user_id[]",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_rest_controller_vehicle.ListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_semenovem_portal_pkg_failing.Response"
                        }
                    }
                }
            }
        },
        "/vehicles/:vehicle_id": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicles"
                ],
                "summary": "Получает данные автомобиля по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID автомобиля",
                        "name": "vehicle_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_semenovem_portal_internal_rest_view.Vehicle"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_semenovem_portal_pkg_failing.Response"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicles"
                ],
                "summary": "Обновляет данные автомобиля",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID автомобиля",
                        "name": "vehicle_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_semenovem_portal_internal_rest_view.Vehicle"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_semenovem_portal_pkg_failing.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicles"
                ],
                "summary": "Удаляет автомобиль",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID автомобиля",
                        "name": "vehicle_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "no content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_semenovem_portal_pkg_failing.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_semenovem_portal_internal_rest_view.Vehicle": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "integer"
                }
            }
        },
        "github_com_semenovem_portal_internal_rest_view.VehicleShort": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "integer"
                }
            }
        },
        "github_com_semenovem_portal_pkg_failing.Response": {
            "type": "object",
            "properties": {
                "additional_fields": {
                    "type": "object",
                    "additionalProperties": true
                },
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "validation_errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_semenovem_portal_pkg_failing.ValidationError"
                    }
                }
            }
        },
        "github_com_semenovem_portal_pkg_failing.ValidationError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "internal_rest_controller_auth.loginForm": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "device_id": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "internal_rest_controller_auth.loginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "description": "TODO для разработки",
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "internal_rest_controller_auth.onetimeAuthForm": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "internal_rest_controller_auth.onetimeAuthResponse": {
            "type": "object",
            "properties": {
                "uri": {
                    "type": "string"
                }
            }
        },
        "internal_rest_controller_vehicle.ListResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_semenovem_portal_internal_rest_view.VehicleShort"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/[v1]/",
	Schemes:          []string{},
	Title:            "portal API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
