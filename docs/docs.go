// Code generated by swaggo/swag. DO NOT EDIT.

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
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization code",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CallbackResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.CallbackInvalidResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.CallbackErrorResponse"
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
                            "$ref": "#/definitions/dto.GetProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.GetProfileUnauthorized"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.GetProfileUserNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.GetProfileErrorResponse"
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
                            "$ref": "#/definitions/dto.RegisterRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterUnauthorized"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterInvalidResponse"
                        }
                    },
                    "498": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterInvalidToken"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterErrorResponse"
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
                            "$ref": "#/definitions/dto.GetAllEventResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.EventInvalidResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.EventAllErrorResponse"
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
                            "$ref": "#/definitions/dto.GetEventByIdResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.EventInvalidResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.EventErrorResponse"
                        }
                    }
                }
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
                            "$ref": "#/definitions/dto.FeatureFlagResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.FeatureFlagInvalidKeyResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.FeatureFlagInternalErrorResponse"
                        }
                    }
                }
            }
        },
        "/schedules/{scheduleId}/register": {
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
                "parameters": [
                    {
                        "type": "integer",
                        "description": "schedule id",
                        "name": "scheduleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Event register body",
                        "name": "registerEventDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.EventRegistrationDTO"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/staff/checkin/{userId}": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Checkin attendee which perform by staff",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "staff"
                ],
                "summary": "checkin attendee",
                "operationId": "AttendeeStaffCheckin",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AttendeeStaffCheckinResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.EventInvalidResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.EventInvalidResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.EventAllErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AttendeeStaffCheckinResponse": {
            "type": "object",
            "properties": {
                "already_checkin": {
                    "type": "boolean"
                },
                "user": {
                    "$ref": "#/definitions/dto.AttendeeStaffCheckinUser"
                }
            }
        },
        "dto.AttendeeStaffCheckinUser": {
            "type": "object",
            "properties": {
                "allergies": {
                    "type": "string",
                    "example": "Romantic"
                },
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "last_name": {
                    "type": "string",
                    "example": "Doe"
                },
                "medical_condition": {
                    "type": "string",
                    "example": "Unlovable"
                }
            }
        },
        "dto.BilingualField": {
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
        "dto.CallbackErrorResponse": {
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
        "dto.CallbackInvalidResponse": {
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
        "dto.CallbackResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "gbxnZjiHVzb_4mDQTQNiJdrZFOCactWXkZvZOxS2_qZsy7vAQY7uA2RFIHe2JABoEjhT0Y3KlOJuOEvE2YJMLrJDagwhpAITGex"
                }
            }
        },
        "dto.Department": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "21"
                },
                "name": {
                    "$ref": "#/definitions/dto.BilingualField"
                }
            }
        },
        "dto.Event": {
            "type": "object",
            "properties": {
                "department": {
                    "$ref": "#/definitions/dto.Department"
                },
                "description": {
                    "$ref": "#/definitions/dto.BilingualField"
                },
                "faculty": {
                    "$ref": "#/definitions/dto.Faculty"
                },
                "id": {
                    "type": "string",
                    "example": "first-event"
                },
                "location": {
                    "$ref": "#/definitions/dto.BilingualField"
                },
                "max_capacity": {
                    "type": "integer",
                    "example": 100
                },
                "name": {
                    "$ref": "#/definitions/dto.BilingualField"
                },
                "require_registration": {
                    "type": "boolean",
                    "example": true
                },
                "schedules": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Schedule"
                    }
                }
            }
        },
        "dto.EventAllErrorResponse": {
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
        "dto.EventErrorResponse": {
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
        "dto.EventInvalidResponse": {
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
        "dto.EventRegistrationDTO": {
            "type": "object",
            "properties": {
                "news_sources": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "facebook",
                        "instagram",
                        "faculty",
                        "chula-student",
                        "friend",
                        "parent",
                        "school",
                        "other"
                    ]
                }
            }
        },
        "dto.Faculty": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "21"
                },
                "name": {
                    "$ref": "#/definitions/dto.BilingualField"
                }
            }
        },
        "dto.FacultyInfo": {
            "type": "object",
            "properties": {
                "department": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "string"
                        },
                        "name": {
                            "$ref": "#/definitions/dto.BilingualField"
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
                            "$ref": "#/definitions/dto.BilingualField"
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
                            "$ref": "#/definitions/dto.BilingualField"
                        }
                    }
                }
            }
        },
        "dto.FacultyInfoId": {
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
        "dto.FeatureFlagInternalErrorResponse": {
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
        "dto.FeatureFlagInvalidKeyResponse": {
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
        "dto.FeatureFlagResponse": {
            "type": "object",
            "properties": {
                "enabled": {
                    "type": "boolean",
                    "example": true
                },
                "extra_info": {
                    "type": "string",
                    "example": "\u003cjsonobject\u003e"
                },
                "key": {
                    "type": "string",
                    "example": "livestream"
                }
            }
        },
        "dto.GetAllEventResponse": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Event"
                    }
                }
            }
        },
        "dto.GetEventByIdResponse": {
            "type": "object",
            "properties": {
                "event": {
                    "$ref": "#/definitions/dto.Event"
                }
            }
        },
        "dto.GetProfileErrorResponse": {
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
        "dto.GetProfileResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/dto.User"
                }
            }
        },
        "dto.GetProfileUnauthorized": {
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
        "dto.GetProfileUserNotFound": {
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
        "dto.RegisterErrorResponse": {
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
        "dto.RegisterInvalidResponse": {
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
        "dto.RegisterInvalidToken": {
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
        "dto.RegisterRequestDTO": {
            "type": "object",
            "properties": {
                "allergies": {
                    "type": "string",
                    "example": "Dog"
                },
                "birth_date": {
                    "type": "string",
                    "example": "1990-01-01"
                },
                "country": {
                    "type": "string",
                    "example": "Japan"
                },
                "desired_round": {
                    "type": "string",
                    "example": "3"
                },
                "educational_level": {
                    "type": "string",
                    "example": "Ph.D."
                },
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "interested_faculties": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.FacultyInfoId"
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
                    "example": "Dog"
                },
                "news_sources": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "facebook",
                        "instagram",
                        "faculty",
                        "chula-student",
                        "friend",
                        "parent",
                        "school",
                        "other"
                    ]
                },
                "province": {
                    "type": "string",
                    "example": "Tokyo"
                },
                "status": {
                    "type": "string",
                    "example": "student"
                },
                "visiting_faculties": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.FacultyInfoId"
                    }
                }
            }
        },
        "dto.RegisterResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/dto.User"
                }
            }
        },
        "dto.RegisterUnauthorized": {
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
        "dto.Schedule": {
            "type": "object",
            "properties": {
                "current_attendee": {
                    "type": "integer",
                    "example": 83
                },
                "ends_at": {
                    "type": "string",
                    "example": "2021-08-01T00:00:00+07:00"
                },
                "id": {
                    "type": "integer",
                    "example": 5
                },
                "period": {
                    "type": "string",
                    "example": "20-morning"
                },
                "starts_at": {
                    "type": "string",
                    "example": "2021-08-01T00:00:00+07:00"
                }
            }
        },
        "dto.User": {
            "type": "object",
            "properties": {
                "allergies": {
                    "type": "string",
                    "example": "Dog"
                },
                "birth_date": {
                    "type": "string",
                    "example": "1990-01-01"
                },
                "country": {
                    "type": "string",
                    "example": "Japan"
                },
                "desired_round": {
                    "type": "string"
                },
                "educational_level": {
                    "type": "string",
                    "example": "Ph.D."
                },
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "id": {
                    "type": "integer",
                    "example": 10000
                },
                "interested_faculties": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.FacultyInfo"
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
                    "example": "Dog"
                },
                "news_sources": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "facebook",
                        "instagram"
                    ]
                },
                "province": {
                    "type": "string",
                    "example": "Austin"
                },
                "registered_events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Schedule"
                    }
                },
                "status": {
                    "type": "string",
                    "example": "student"
                },
                "visiting_faculties": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.FacultyInfo"
                    }
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
