
package swaggerui
//This file is generated automatically. Do not try to edit it manually.

var resourceListingJson = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://localhost:1323/swagger-ui",
    "apis": [
        {
            "path": "/hello",
            "description": "hello"
        },
        {
            "path": "/auth",
            "description": "auth"
        }
    ],
    "info": {
        "title": "tas-blog-api",
        "description": "tas-blog-api"
    }
}`
var apiDescriptionsJson = map[string]string{"auth":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://localhost:1323",
    "resourcePath": "/auth",
    "apis": [
        {
            "path": "/auth/login",
            "description": "ログイン",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "ログインAPI",
                    "type": "github.com.team-a-hacks.tas-blog-api.auth.Token",
                    "items": {},
                    "summary": "ログイン",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "auth",
                            "description": "auth",
                            "dataType": "github.com.team-a-hacks.tas-blog-api.auth.Login",
                            "type": "github.com.team-a-hacks.tas-blog-api.auth.Login",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "トークン",
                            "responseType": "object",
                            "responseModel": "github.com.team-a-hacks.tas-blog-api.auth.Token"
                        },
                        {
                            "code": 400,
                            "message": "bad request",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.team-a-hacks.tas-blog-api.auth.Login": {
            "id": "github.com.team-a-hacks.tas-blog-api.auth.Login",
            "properties": {
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "password": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.team-a-hacks.tas-blog-api.auth.Token": {
            "id": "github.com.team-a-hacks.tas-blog-api.auth.Token",
            "properties": {
                "accessToken": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "refreshToken": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,"hello":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://localhost:1323",
    "resourcePath": "/hello",
    "apis": [
        {
            "path": "/hello",
            "description": "Hello 'your name'",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Hello",
                    "type": "string",
                    "items": {},
                    "summary": "Hello 'your name'",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Hello World",
                            "responseType": "object",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        }
    ]
}`,}
