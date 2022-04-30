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
        "termsOfService": "https://github.com/BakhodiribnYashinibnMansur/Medium_uz",
        "contact": {
            "name": "Bakhodir Yashin Mansur",
            "email": "phapp0224mb@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/account/resend": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "resend code to email for  verification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Resend cod for  Verification Email",
                "operationId": "resend-code-email",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/account/upload-image": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Upload Account Image",
                "consumes": [
                    "application/json",
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json",
                    "application/octet-stream",
                    "image/png",
                    "image/jpeg",
                    "image/jpg"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Upload Account Image",
                "operationId": "upload-image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/account/verify": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "verification email with code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Verification Email",
                "operationId": "verify-email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "code",
                        "name": "code",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "login account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "SignIn",
                "operationId": "login-account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseSign"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "SignUp",
                "operationId": "create-account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseSign"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/error.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "error.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "error.errorResponseData": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "model.ResponseSign": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "model.ResponseSuccess": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "model.SignInInput": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "default": "0224mb"
                },
                "username": {
                    "type": "string",
                    "default": "MRB"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string",
                    "default": "I am Golang dev"
                },
                "city": {
                    "type": "string",
                    "default": "Navoi"
                },
                "email": {
                    "type": "string",
                    "default": "phapp0224mb@gmail.com"
                },
                "first_name": {
                    "type": "string",
                    "default": "MRB"
                },
                "image": {
                    "type": "string"
                },
                "interesting": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "nickname": {
                    "type": "string",
                    "default": "mrb"
                },
                "password": {
                    "type": "string",
                    "default": "0224mb"
                },
                "phone": {
                    "type": "string",
                    "default": "+9989 93 753 65 71"
                },
                "second_name": {
                    "type": "string",
                    "default": "HERO"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
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
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MediumuZ API",
	Description:      "API Server for MediumuZ Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
