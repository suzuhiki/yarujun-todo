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
        "license": {
            "name": "suzuhiki"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/current_user": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "現在のユーザーidを返す",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.GetUserIdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/refresh_token": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "認証情報の更新",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/tasks": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Todo一覧を配列で返す",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/types.TaskEntity"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "summary": "タスクを作成する",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "body param",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTaskRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/create_account": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "アカウント作成",
                "parameters": [
                    {
                        "description": "body param",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CreateAccountResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "ログイン",
                "parameters": [
                    {
                        "description": "body param",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/test": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "hello worldを返す",
                "responses": {
                    "200": {
                        "description": "Hello, World!!!!!!!!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.CreateAccountRequest": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "taro"
                },
                "password": {
                    "type": "string",
                    "example": "tarodesu"
                }
            }
        },
        "types.CreateAccountResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "name": {
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "types.CreateTaskRequest": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "deadline": {
                    "type": "string",
                    "example": "2024-09-20T03:12:53+09:00"
                },
                "memo": {
                    "type": "string",
                    "example": "概要"
                },
                "title": {
                    "type": "string",
                    "example": "やること"
                },
                "waitlist_num": {
                    "type": "integer"
                }
            }
        },
        "types.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "types.GetUserIdResponse": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string",
                    "example": "1"
                }
            }
        },
        "types.LoginRequest": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "taro"
                },
                "password": {
                    "type": "string",
                    "example": "tarodesu"
                }
            }
        },
        "types.LoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "expier": {
                    "type": "string",
                    "example": "2024-09-20T03:12:53+09:00"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "types.SuccessResponse": {
            "type": "object",
            "properties": {
                "data": {}
            }
        },
        "types.TaskEntity": {
            "type": "object",
            "properties": {
                "deadline": {
                    "type": "string"
                },
                "memo": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "waitlist_num": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "gin-swagger todos",
	Description:      "このswaggerはyarujunのAPIを定義しています。 JWTトークンの前に\"Bearer\"を追加してください。",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
