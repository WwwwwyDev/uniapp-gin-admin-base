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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/c/decodeJwt": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysJwt"
                ],
                "summary": "jwt解码,反馈用户信息",
                "responses": {
                    "200": {
                        "description": "{\"code\": 20000,\"data\":{},\"msg\":\"解码成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/c/isValidJwt": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysJwt"
                ],
                "summary": "jwt验证",
                "responses": {
                    "200": {
                        "description": "{\"code\": 20000,\"data\":{},\"msg\":\"验证通过\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/c/jwt": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysJwt"
                ],
                "summary": "删除用户token",
                "responses": {
                    "200": {
                        "description": "{\"code\": 20000,\"data\":{},\"msg\":\"删除成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sysCaptcha": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysCaptcha"
                ],
                "summary": "生成验证码",
                "responses": {
                    "200": {
                        "description": "{\"code\": 20000,\"data\":{},\"msg\":\"验证码获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sysUser/changePassword": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysUser"
                ],
                "summary": "修改密码",
                "parameters": [
                    {
                        "description": "用户名, 旧密码, 新密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ChangePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 20000,\"data\": {},\"msg\": \"修改密码成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sysUser/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysUser"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名, 密码, 验证码，验证码ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 20000,\"data\":{},\"msg\":\"登陆成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sysUser/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysUser"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "必填：用户名, 密码, 昵称；非必填：头像链接，电子邮箱，联系电话",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Register"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 20000,\"data\": {},\"msg\": \"注册成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.ChangePassword": {
            "type": "object",
            "properties": {
                "newPassword": {
                    "description": "新密码",
                    "type": "string"
                },
                "oldPassword": {
                    "description": "旧密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "request.Login": {
            "type": "object",
            "properties": {
                "captcha": {
                    "description": "验证码",
                    "type": "string"
                },
                "captchaId": {
                    "description": "验证码ID",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "request.Register": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "电子邮箱",
                    "type": "string"
                },
                "headerImg": {
                    "description": "头像链接",
                    "type": "string"
                },
                "nickName": {
                    "description": "昵称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "联系电话",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
	swag.Register(swag.Name, &s{})
}
