// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "email": "1239989762@qq.com"
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
        "/auth": {
            "post": {
                "description": "对原图片进行盲水印提取",
                "consumes": [
                    "form-data/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "水印相关接口"
                ],
                "summary": "对图片盲水印进行提取",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamAuth"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "url",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/blockNumber": {
            "get": {
                "description": "获取fisco-bcos最新块高",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "区块链信息相关接口"
                ],
                "summary": "获取最新块高",
                "responses": {
                    "200": {
                        "description": "height",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/evidence": {
            "get": {
                "description": "查询上链的版权详情信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "水印相关接口"
                ],
                "summary": "查看版权详情",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamEvidenceDeatil"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "$ref": "#/definitions/models.EvidenceDetailResp"
                        }
                    }
                }
            }
        },
        "/getLatestBlock": {
            "get": {
                "description": "获取fisco-bcos区块链获取最新十个区块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "区块链信息相关接口"
                ],
                "summary": "获取最新十个区块",
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.BlockDetailResp"
                            }
                        }
                    }
                }
            }
        },
        "/getLatestTx": {
            "get": {
                "description": "获取fisco-bcos区块链获取最新十笔区块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "区块链信息相关接口"
                ],
                "summary": "获取最新十个十笔交易",
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.TxDetailResp"
                            }
                        }
                    }
                }
            }
        },
        "/getPeers": {
            "get": {
                "description": "获取fisco-bcos区块链连接的节点ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "区块链信息相关接口"
                ],
                "summary": "最新区块链连接的节点ID",
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/getTransSums": {
            "get": {
                "description": "获取印系统的交易数量最新用户交易数量",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "区块链信息相关接口"
                ],
                "summary": "获取通过水印系统的交易数量",
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/mysql.TxSum"
                            }
                        }
                    }
                }
            }
        },
        "/getTransactionTotal": {
            "get": {
                "description": "获取fisco-bcos区块链交易总数",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "区块链信息相关接口"
                ],
                "summary": "最新区块链交易总数",
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "$ref": "#/definitions/models.TxTotalResp"
                        }
                    }
                }
            }
        },
        "/getUserCount": {
            "get": {
                "description": "用户数量",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "水印相关接口"
                ],
                "summary": "用户数量",
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/getUserInfo": {
            "post": {
                "description": "通过Header中Authorization查询用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "获取用户详情信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/img": {
            "post": {
                "description": "对原图片进行盲水印提取",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "水印相关接口"
                ],
                "summary": "图片水印提取",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamImgEvidence"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "url",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "description": "查询上链的版权信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "水印相关接口"
                ],
                "summary": "查询版权列表",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "description": "每页数据量",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "",
                        "description": "用户名",
                        "name": "user_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.EvidenceResp"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "用户登陆接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户登陆",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "查询参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/refresh": {
            "get": {
                "description": "用户获取新的AccessToken接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "获取新AccessToken",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Access_Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer Refresh_Token",
                        "name": "RefreshToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "access_token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "用户注册接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamSignUp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/upwatermark": {
            "post": {
                "description": "上传水印图片，为后续图片水印进行处理",
                "consumes": [
                    "form-data/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "水印相关接口"
                ],
                "summary": "上传图片水印",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamWaterMark"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "url",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/watermark": {
            "post": {
                "description": "通过用户私钥，为原图添加盲水印",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "水印相关接口"
                ],
                "summary": "添加文字水印",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamEvidence"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "url",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ResCode": {
            "type": "integer",
            "enum": [
                1000,
                1001,
                1002,
                1003,
                1004,
                1005,
                1006,
                1007,
                1008
            ],
            "x-enum-varnames": [
                "CodeSuccess",
                "CodeInvalidParam",
                "CodeUserExist",
                "CodeUserNotExist",
                "CodeInvalidPassword",
                "CodeServerBusy",
                "CodeNeedLogin",
                "CodeInvalidToken",
                "CodeInvalidUploadFile"
            ]
        },
        "controller.ResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/controller.ResCode"
                },
                "data": {},
                "msg": {}
            }
        },
        "models.BlockDetailResp": {
            "type": "object",
            "properties": {
                "height": {
                    "description": "当前块高",
                    "type": "string"
                },
                "sealer": {
                    "description": "出块节点",
                    "type": "string"
                },
                "timestamp": {
                    "description": "时间戳",
                    "type": "string"
                },
                "tx_sum": {
                    "description": "交易总数",
                    "type": "string"
                }
            }
        },
        "models.EvidenceDetailResp": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string"
                },
                "created_at": {
                    "description": "时间",
                    "type": "string"
                },
                "evidence": {
                    "type": "string"
                },
                "img_url": {
                    "description": "图片",
                    "type": "string"
                },
                "total": {
                    "description": "总数量",
                    "type": "integer"
                },
                "tx_id": {
                    "description": "版权编号/交易id",
                    "type": "string"
                },
                "user_name": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "models.EvidenceResp": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "时间",
                    "type": "string"
                },
                "img_url": {
                    "description": "图片",
                    "type": "string"
                },
                "total": {
                    "description": "总数量",
                    "type": "integer"
                },
                "tx_id": {
                    "description": "版权编号/交易id",
                    "type": "string"
                },
                "user_name": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "models.ParamAuth": {
            "type": "object"
        },
        "models.ParamEvidence": {
            "type": "object"
        },
        "models.ParamEvidenceDeatil": {
            "type": "object",
            "required": [
                "tx_id"
            ],
            "properties": {
                "tx_id": {
                    "type": "string"
                }
            }
        },
        "models.ParamImgEvidence": {
            "type": "object",
            "required": [
                "hexPrivateKey",
                "img",
                "watermark"
            ],
            "properties": {
                "hexPrivateKey": {
                    "description": "私钥",
                    "type": "string"
                },
                "img": {
                    "description": "原图",
                    "type": "string"
                },
                "watermark": {
                    "description": "水印图片",
                    "type": "string"
                }
            }
        },
        "models.ParamLogin": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "用户密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "models.ParamSignUp": {
            "type": "object",
            "required": [
                "password",
                "re_password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "用户密码",
                    "type": "string"
                },
                "re_password": {
                    "description": "用户再次输入密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "models.ParamWaterMark": {
            "type": "object"
        },
        "models.TxDetailResp": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                },
                "tx_id": {
                    "type": "string"
                }
            }
        },
        "models.TxTotalResp": {
            "type": "object",
            "properties": {
                "failed_tx_sum": {
                    "description": "失败交易总数",
                    "type": "string"
                },
                "tx_sum": {
                    "description": "交易总数",
                    "type": "string"
                }
            }
        },
        "mysql.TxSum": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "trans_count": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8084",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "fisco-cert-app 社区",
	Description:      "关于fisco-cert-app社区接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
