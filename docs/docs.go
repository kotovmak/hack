// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/compaign": {
            "get": {
                "description": "Получить список Компаний",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "compaign"
                ],
                "summary": "Получить список Компаний",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Compaign"
                            }
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/file": {
            "get": {
                "description": "Если файл есть в расчете то получим его состояние",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Получить состояние текущего файла с датасетом",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.File"
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "Получаем файл с фронта с новым датасетом для расчета",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Сохранить файл с датасетом",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "form data with file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.File"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/lead": {
            "get": {
                "description": "Получить список Лидов",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lead"
                ],
                "summary": "Получить список Лидов",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Lead"
                            }
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/region/predict": {
            "get": {
                "description": "Результат работы модели загружаем в этом методе",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "region"
                ],
                "summary": "Получить список городов с придиктивной аналитикой",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.RegionPredict"
                            }
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/telegram": {
            "get": {
                "description": "Получить список телеграм каналов",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "telegram"
                ],
                "summary": "Получить список телеграм каналов",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Telegram"
                            }
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            }
        },
        "/ws/": {
            "get": {
                "description": "Обрабатываем WebSocket",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ws"
                ],
                "summary": "Обрабатываем WebSocket",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Compaign": {
            "type": "object",
            "properties": {
                "age_from": {
                    "type": "integer"
                },
                "age_to": {
                    "type": "integer"
                },
                "city": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "theme": {
                    "type": "string"
                },
                "utm_campaign": {
                    "type": "string"
                }
            }
        },
        "model.File": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "createAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "receivedAt": {
                    "type": "string"
                },
                "sendAt": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "status_message": {
                    "type": "string"
                }
            }
        },
        "model.Lead": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "cpc": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "product_category_name": {
                    "type": "string"
                },
                "utm_campaign": {
                    "type": "string"
                },
                "utm_content": {
                    "type": "string"
                },
                "utm_source": {
                    "type": "string"
                }
            }
        },
        "model.RegionPredict": {
            "type": "object",
            "required": [
                "city",
                "position"
            ],
            "properties": {
                "city": {
                    "type": "string"
                },
                "currentClientIndex": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "position": {
                    "type": "integer"
                },
                "predictArpu": {
                    "type": "number"
                },
                "predictClientIndex": {
                    "type": "number"
                },
                "predictScore": {
                    "type": "number"
                },
                "product": {
                    "type": "string"
                }
            }
        },
        "model.ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.Telegram": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "n_subscribers": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "name_id": {
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
	Version:     "1.0",
	Host:        "51.250.44.134",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "hack API",
	Description: "API for Moscow City Hack 2022.",
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
