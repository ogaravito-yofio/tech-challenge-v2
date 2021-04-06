// Code generated by go-swagger; DO NOT EDIT.

package tech_challenge_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "YoFio Recruitment - Tech Challenge",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "ogaravito@yofio.co"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "tech-challenge-v2.herokuapp.com",
  "paths": {
    "/": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "general"
        ],
        "summary": "Add a new registration to the database",
        "operationId": "healthcheck",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/registration": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "form"
        ],
        "summary": "Add a new registration to the database",
        "operationId": "addPeople",
        "parameters": [
          {
            "description": "People object that needs to be added to the database",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/People"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "People register was successfully created in database",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "400": {
            "description": "Invalid input."
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid"
        }
      }
    },
    "Location": {
      "type": "object",
      "properties": {
        "latitude": {
          "type": "number",
          "format": "float64"
        },
        "longitude": {
          "type": "number",
          "format": "float64"
        }
      },
      "xml": {
        "name": "Tag"
      }
    },
    "People": {
      "type": "object",
      "required": [
        "birthdate",
        "lastName",
        "name"
      ],
      "properties": {
        "birthdate": {
          "type": "string",
          "format": "date"
        },
        "lastName": {
          "type": "string",
          "example": "Parker"
        },
        "location": {
          "$ref": "#/definitions/Location"
        },
        "name": {
          "type": "string",
          "example": "Peter"
        },
        "photo": {
          "type": "string",
          "format": "byte",
          "pattern": "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$"
        }
      },
      "xml": {
        "name": "Pet"
      }
    }
  },
  "tags": [
    {
      "description": "Basic test for collecting a form",
      "name": "form"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "YoFio Recruitment - Tech Challenge",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "ogaravito@yofio.co"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "tech-challenge-v2.herokuapp.com",
  "paths": {
    "/": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "general"
        ],
        "summary": "Add a new registration to the database",
        "operationId": "healthcheck",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/registration": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "form"
        ],
        "summary": "Add a new registration to the database",
        "operationId": "addPeople",
        "parameters": [
          {
            "description": "People object that needs to be added to the database",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/People"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "People register was successfully created in database",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "400": {
            "description": "Invalid input."
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid"
        }
      }
    },
    "Location": {
      "type": "object",
      "properties": {
        "latitude": {
          "type": "number",
          "format": "float64"
        },
        "longitude": {
          "type": "number",
          "format": "float64"
        }
      },
      "xml": {
        "name": "Tag"
      }
    },
    "People": {
      "type": "object",
      "required": [
        "birthdate",
        "lastName",
        "name"
      ],
      "properties": {
        "birthdate": {
          "type": "string",
          "format": "date"
        },
        "lastName": {
          "type": "string",
          "example": "Parker"
        },
        "location": {
          "$ref": "#/definitions/Location"
        },
        "name": {
          "type": "string",
          "example": "Peter"
        },
        "photo": {
          "type": "string",
          "format": "byte",
          "pattern": "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$"
        }
      },
      "xml": {
        "name": "Pet"
      }
    }
  },
  "tags": [
    {
      "description": "Basic test for collecting a form",
      "name": "form"
    }
  ]
}`))
}
