// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/login": {
            "post": {
                "description": "用户登陆模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户登陆"
                ],
                "summary": "用户登陆接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "telephone",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登陆成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "登陆失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/auth/register": {
            "post": {
                "description": "用户注册模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户注册"
                ],
                "summary": "用户注册接口",
                "parameters": [
                    {
                        "type": "string",
                        "name": "createdAt",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "telephone",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "updatedAt",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "注册成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "注册失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/example/helloworld": {
            "get": {
                "description": "swagger测试用例",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "swagger的列子"
                ],
                "summary": "example",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "false",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/info": {
            "get": {
                "description": "中间件验证模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "中间件验证"
                ],
                "summary": "中间件验证接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登陆成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "登陆失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/categories": {
            "post": {
                "description": "创建类别模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "创建类别"
                ],
                "summary": "创建类别接口",
                "parameters": [
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "分类创建成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "数据验证错误，分类名称必填",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/categories/{id}": {
            "get": {
                "description": "查看类别模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "查看类别"
                ],
                "summary": "查看类别接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "类别ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "分类查看成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "分类不存在",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "更新类别模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "更新类别"
                ],
                "summary": "更新类别接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "类别ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "修改分类成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "分类不存在",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除类别模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "删除类别"
                ],
                "summary": "删除类别接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "类别ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "分类删除成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "删除失败，请重试",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "swagger API",
	Description: "简单的后端登陆注册和文章分类API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}