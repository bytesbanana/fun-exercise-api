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
        "/api/v1/users/{user_id}/wallets": {
            "get": {
                "description": "Get all wallets by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all wallets by user id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/wallet.Wallet"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/user.Err"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/user.Err"
                        }
                    }
                }
            }
        },
        "/api/v1/wallets": {
            "get": {
                "description": "Get all wallets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get all wallets",
                "parameters": [
                    {
                        "enum": [
                            "Savings",
                            "CreditCard",
                            "CryptoWallet"
                        ],
                        "type": "string",
                        "description": "Wallet type",
                        "name": "wallet_type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/wallet.Wallet"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/wallet.Err"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wallet.Err"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Create new wallet",
                "parameters": [
                    {
                        "description": "Wallet",
                        "name": "wallet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/wallet.Wallet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/wallet.Wallet"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/wallet.Err"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wallet.Err"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "user.Err": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "wallet.Err": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "wallet.Wallet": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number",
                    "example": 100
                },
                "created_at": {
                    "type": "string",
                    "example": "2024-03-25T14:19:00.729237Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                },
                "user_name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "wallet_name": {
                    "type": "string",
                    "example": "John's Wallet"
                },
                "wallet_type": {
                    "type": "string",
                    "example": "Credit Card"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1323",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Wallet API",
	Description:      "Sophisticated Wallet API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
