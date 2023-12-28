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
        "/_hc": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Health Check for the service",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Health Check",
                "operationId": "HealthCheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/callback": {
            "get": {
                "description": "After successfully logging in with a @chula account, you'll receive a token. If you attempt to log in using a different domain, Google will not allow the login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "receive a token after successfully login with Google",
                "operationId": "GoogleCallback",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.CallbackResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/auth.CallbackInvalidResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/auth.CallbackErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "get": {
                "description": "Redirect to Google login page",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Redirect to Google login page",
                "operationId": "GoogleLogin",
                "responses": {}
            }
        },
        "/auth/me": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get Profile of current user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Get Profile of current user",
                "operationId": "GetProfile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.GetProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/auth.GetProfileUnauthorized"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/auth.GetProfileUserNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/auth.GetProfileErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Register new account with email",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "operationId": "Register",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterUnauthorized"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterInvalidResponse"
                        }
                    },
                    "498": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterInvalidToken"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterErrorResponse"
                        }
                    }
                }
            }
        },
        "/events": {
            "get": {
                "description": "Get all events as array of events",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "summary": "Get all events",
                "operationId": "GetAllEvents",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/event.GetAllEventResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/event.EventInvalidResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/event.EventAllErrorResponse"
                        }
                    }
                }
            }
        },
        "/events/{eventId}": {
            "get": {
                "description": "Get event by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "summary": "get event by id",
                "operationId": "GetEventById",
                "parameters": [
                    {
                        "type": "string",
                        "description": "event id",
                        "name": "eventId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/event.EventDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/event.EventInvalidResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/event.EventErrorResponse"
                        }
                    }
                }
            }
        },
        "/events/{eventId}/register": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Register event",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "summary": "Register event",
                "operationId": "RegisterEvent",
                "responses": {}
            }
        },
        "/live": {
            "get": {
                "description": "Get livestream flag",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FeatureFlag"
                ],
                "summary": "Get livestream flag",
                "operationId": "GetLivestreamInfo",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/featureflag.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/featureflag.invalidResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/featureflag.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.BillingualName": {
            "type": "object",
            "properties": {
                "en": {
                    "type": "string"
                },
                "th": {
                    "type": "string"
                }
            }
        },
        "auth.CallbackErrorResponse": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/auth/callback"
                },
                "title": {
                    "type": "string",
                    "example": "internal-server-error"
                }
            }
        },
        "auth.CallbackInvalidResponse": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/auth/callback"
                },
                "title": {
                    "type": "string",
                    "example": "bad-request"
                }
            }
        },
        "auth.CallbackResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "gbxnZjiHVzb_4mDQTQNiJdrZFOCactWXkZvZOxS2_qZsy7vAQY7uA2RFIHe2JABoEjhT0Y3KlOJuOEvE2YJMLrJDagwhpAITGex"
                }
            }
        },
        "auth.DesiredRound": {
            "type": "object",
            "properties": {
                "order": {
                    "type": "integer",
                    "example": 1
                },
                "round": {
                    "type": "string",
                    "example": "1"
                }
            }
        },
        "auth.FacultyInfo": {
            "type": "object",
            "properties": {
                "department": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "string"
                        },
                        "name": {
                            "$ref": "#/definitions/auth.BillingualName"
                        }
                    }
                },
                "faculty": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "string"
                        },
                        "name": {
                            "$ref": "#/definitions/auth.BillingualName"
                        }
                    }
                },
                "section": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "string"
                        },
                        "name": {
                            "$ref": "#/definitions/auth.BillingualName"
                        }
                    }
                }
            }
        },
        "auth.FacultyInfoId": {
            "type": "object",
            "properties": {
                "department_code": {
                    "type": "string",
                    "example": "10"
                },
                "faculty_code": {
                    "type": "string",
                    "example": "21"
                },
                "order": {
                    "type": "integer",
                    "example": 1
                },
                "section_code": {
                    "type": "string",
                    "example": "-"
                }
            }
        },
        "auth.GetProfileErrorResponse": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/auth/me"
                },
                "title": {
                    "type": "string",
                    "example": "internal-server-error"
                }
            }
        },
        "auth.GetProfileResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/auth.User"
                }
            }
        },
        "auth.GetProfileUnauthorized": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/auth/me"
                },
                "title": {
                    "type": "string",
                    "example": "unauthorized"
                }
            }
        },
        "auth.GetProfileUserNotFound": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/auth/me"
                },
                "title": {
                    "type": "string",
                    "example": "user-not-found"
                }
            }
        },
        "auth.RegisterErrorResponse": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/auth/register"
                },
                "title": {
                    "type": "string",
                    "example": "internal-server-error"
                }
            }
        },
        "auth.RegisterInvalidResponse": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/auth/register"
                },
                "title": {
                    "type": "string",
                    "example": "bad-request"
                }
            }
        },
        "auth.RegisterInvalidToken": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/auth/register"
                },
                "title": {
                    "type": "string",
                    "example": "invalid-token"
                }
            }
        },
        "auth.RegisterRequestDTO": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "Bangkok"
                },
                "allergy": {
                    "type": "string",
                    "example": "None"
                },
                "birth_date": {
                    "type": "string",
                    "example": "1990-01-01"
                },
                "desired_rounds": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/auth.DesiredRound"
                    }
                },
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "from_abroad": {
                    "type": "string",
                    "example": "no"
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "grade": {
                    "type": "string",
                    "example": "undergraduate"
                },
                "interested_faculties": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/auth.FacultyInfoId"
                    }
                },
                "join_cu_reason": {
                    "type": "string",
                    "example": "Interested in the programs offered"
                },
                "last_name": {
                    "type": "string",
                    "example": "Doe"
                },
                "medical_condition": {
                    "type": "string",
                    "example": "None"
                },
                "news_source": {
                    "type": "string",
                    "example": "Facebook"
                },
                "school": {
                    "type": "string",
                    "example": "CU"
                },
                "status": {
                    "type": "string",
                    "example": "student"
                }
            }
        },
        "auth.RegisterResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/auth.User"
                }
            }
        },
        "auth.RegisterUnauthorized": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/auth/register"
                },
                "title": {
                    "type": "string",
                    "example": "unauthorized"
                }
            }
        },
        "auth.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "Bangkok"
                },
                "allergy": {
                    "type": "string",
                    "example": "None"
                },
                "birth_date": {
                    "type": "string",
                    "example": "1990-01-01"
                },
                "desired_rounds": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/auth.DesiredRound"
                    }
                },
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "from_abroad": {
                    "type": "string",
                    "example": "no"
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "grade": {
                    "type": "string",
                    "example": "undergraduate"
                },
                "interested_faculties": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/auth.FacultyInfo"
                    }
                },
                "join_cu_reason": {
                    "type": "string",
                    "example": "Interested in the programs offered"
                },
                "last_name": {
                    "type": "string",
                    "example": "Doe"
                },
                "medical_condition": {
                    "type": "string",
                    "example": "None"
                },
                "news_source": {
                    "type": "string",
                    "example": "Facebook"
                },
                "school": {
                    "type": "string",
                    "example": "CU"
                },
                "status": {
                    "type": "string",
                    "example": "student"
                }
            }
        },
        "event.DepartmentBilingual": {
            "type": "object",
            "properties": {
                "en": {
                    "type": "string",
                    "example": "Computer Engineering"
                },
                "th": {
                    "type": "string",
                    "example": "ภาควิชาคอมพิวเตอร์"
                }
            }
        },
        "event.DescriptionBilingual": {
            "type": "object",
            "properties": {
                "en": {
                    "type": "string",
                    "example": "This is the first event."
                },
                "th": {
                    "type": "string",
                    "example": "รายละเอียดอีเวนท์แรก"
                }
            }
        },
        "event.EventAllErrorResponse": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/events"
                },
                "title": {
                    "type": "string",
                    "example": "internal-server-error"
                }
            }
        },
        "event.EventDTO": {
            "type": "object",
            "properties": {
                "department": {
                    "$ref": "#/definitions/event.DepartmentBilingual"
                },
                "description": {
                    "$ref": "#/definitions/event.DescriptionBilingual"
                },
                "faculty": {
                    "$ref": "#/definitions/event.Faculty"
                },
                "id": {
                    "type": "string",
                    "example": "first-event"
                },
                "location": {
                    "$ref": "#/definitions/event.LocationBilingual"
                },
                "max_capacity": {
                    "type": "integer",
                    "example": 100
                },
                "name": {
                    "$ref": "#/definitions/event.NameEventBilingual"
                },
                "require_registration": {
                    "type": "boolean",
                    "example": true
                },
                "schedules": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/event.Schedule"
                    }
                }
            }
        },
        "event.EventErrorResponse": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/events/:eventId"
                },
                "title": {
                    "type": "string",
                    "example": "internal-server-error"
                }
            }
        },
        "event.EventInvalidResponse": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/events/:eventId"
                },
                "title": {
                    "type": "string",
                    "example": "invalid-event-id"
                }
            }
        },
        "event.Faculty": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "21"
                },
                "name": {
                    "$ref": "#/definitions/event.NameFacultyBilingual"
                }
            }
        },
        "event.GetAllEventResponse": {
            "type": "object"
        },
        "event.LocationBilingual": {
            "type": "object",
            "properties": {
                "en": {
                    "type": "string",
                    "example": "SIT Building"
                },
                "th": {
                    "type": "string",
                    "example": "อาคาร SIT"
                }
            }
        },
        "event.NameEventBilingual": {
            "type": "object",
            "properties": {
                "en": {
                    "type": "string",
                    "example": "First Event"
                },
                "th": {
                    "type": "string",
                    "example": "อีเวนท์แรก"
                }
            }
        },
        "event.NameFacultyBilingual": {
            "type": "object",
            "properties": {
                "en": {
                    "type": "string",
                    "example": "Faculty of Engineering"
                },
                "th": {
                    "type": "string",
                    "example": "คณะวิศวกรรมศาสตร์"
                }
            }
        },
        "event.Schedule": {
            "type": "object",
            "properties": {
                "ends_at": {
                    "type": "string",
                    "example": "2021-08-01T00:00:00+07:00"
                },
                "starts_at": {
                    "type": "string",
                    "example": "2021-08-01T00:00:00+07:00"
                }
            }
        },
        "featureflag.errorResponse": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/featureflag/live"
                },
                "title": {
                    "type": "string",
                    "example": "internal-server-error"
                }
            }
        },
        "featureflag.invalidResponse": {
            "type": "object",
            "properties": {
                "instance": {
                    "type": "string",
                    "example": "/featureflag/live"
                },
                "title": {
                    "type": "string",
                    "example": "invalid-feature-flag-key"
                }
            }
        },
        "featureflag.response": {
            "type": "object",
            "properties": {
                "enabled": {
                    "type": "boolean",
                    "example": true
                },
                "extra_info": {
                    "type": "string",
                    "example": "https://www.youtube.com/watch?v=6n3pFFPSlW4"
                },
                "key": {
                    "type": "string",
                    "example": "livestream"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
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
	BasePath:         "",
	Schemes:          []string{"http", "https"},
	Title:            "OPH-66 Backend API",
	Description:      "Documentation outlines the specifications and endpoints for the OPH-66 Backend API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
