
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
        },
        {
            "path": "/accounts",
            "description": "accounts"
        },
        {
            "path": "/articles",
            "description": "articles"
        },
        {
            "path": "/comments",
            "description": "comments"
        }
    ],
    "info": {
        "title": "tas-blog-api",
        "description": "tas-blog-api"
    }
}`
var apiDescriptionsJson = map[string]string{"accounts":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://localhost:1323",
    "resourcePath": "/accounts",
    "apis": [
        {
            "path": "/accounts/{id}",
            "description": "",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "アカウント取得API",
                    "type": "github.com.team-a-hacks.tas-blog-api.account.AccountData",
                    "items": {},
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "ID",
                            "dataType": "github.com.satori.go.uuid.UUID",
                            "type": "github.com.satori.go.uuid.UUID",
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
                            "message": "アカウント情報",
                            "responseType": "object",
                            "responseModel": "github.com.team-a-hacks.tas-blog-api.account.AccountData"
                        },
                        {
                            "code": 404,
                            "message": "not found error",
                            "responseType": "object",
                            "responseModel": "error"
                        },
                        {
                            "code": 500,
                            "message": "internal server error",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.satori.go.uuid.UUID": {
            "id": "github.com.satori.go.uuid.UUID",
            "properties": null
        },
        "github.com.team-a-hacks.tas-blog-api.account.AccountData": {
            "id": "github.com.team-a-hacks.tas-blog-api.account.AccountData",
            "properties": {
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,"articles":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://localhost:1323",
    "resourcePath": "/articles",
    "apis": [
        {
            "path": "/articles",
            "description": "記事一覧取得API",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "記事一覧取得API",
                    "type": "array",
                    "items": {
                        "$ref": "github.com.team-a-hacks.tas-blog-api.article.ArticleData"
                    },
                    "summary": "記事一覧取得API",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "記事",
                            "responseType": "array",
                            "responseModel": "github.com.team-a-hacks.tas-blog-api.article.ArticleData"
                        },
                        {
                            "code": 404,
                            "message": "not found error",
                            "responseType": "object",
                            "responseModel": "error"
                        },
                        {
                            "code": 500,
                            "message": "internal server error",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.satori.go.uuid.UUID": {
            "id": "github.com.satori.go.uuid.UUID",
            "properties": null
        },
        "github.com.team-a-hacks.tas-blog-api.article.ArticleData": {
            "id": "github.com.team-a-hacks.tas-blog-api.article.ArticleData",
            "properties": {
                "content": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,"auth":`{
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
        },
        {
            "path": "/auth/refresh",
            "description": "リフレッシュトークンの再発行",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "レフレッシュAPI",
                    "type": "github.com.team-a-hacks.tas-blog-api.auth.Token",
                    "items": {},
                    "summary": "リフレッシュトークンの再発行",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "token",
                            "description": "リフレッシュトークン",
                            "dataType": "github.com.team-a-hacks.tas-blog-api.auth.Refresh",
                            "type": "github.com.team-a-hacks.tas-blog-api.auth.Refresh",
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
        },
        {
            "path": "/auth/logout",
            "description": "店舗従業員ログアウト",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "ログアウトAPI",
                    "type": "",
                    "items": {},
                    "summary": "店舗従業員ログアウト",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "logout",
                            "description": "アカウントID",
                            "dataType": "github.com.team-a-hacks.tas-blog-api.auth.Logout",
                            "type": "github.com.team-a-hacks.tas-blog-api.auth.Logout",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "no content",
                            "responseType": "object",
                            "responseModel": "error"
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
        "github.com.satori.go.uuid.UUID": {
            "id": "github.com.satori.go.uuid.UUID",
            "properties": null
        },
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
        "github.com.team-a-hacks.tas-blog-api.auth.Logout": {
            "id": "github.com.team-a-hacks.tas-blog-api.auth.Logout",
            "properties": {
                "id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.team-a-hacks.tas-blog-api.auth.Refresh": {
            "id": "github.com.team-a-hacks.tas-blog-api.auth.Refresh",
            "properties": {
                "accountID": {
                    "type": "github.com.satori.go.uuid.UUID",
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
}`,"comments":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://localhost:1323",
    "resourcePath": "/comments",
    "apis": [
        {
            "path": "/comments",
            "description": "コメント作成",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "コメント作成API",
                    "type": "github.com.team-a-hacks.tas-blog-api.comment.ResponseData",
                    "items": {},
                    "summary": "コメント作成",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "comment",
                            "description": "コメント",
                            "dataType": "github.com.team-a-hacks.tas-blog-api.comment.Payload",
                            "type": "github.com.team-a-hacks.tas-blog-api.comment.Payload",
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
                            "message": "コメント",
                            "responseType": "object",
                            "responseModel": "github.com.team-a-hacks.tas-blog-api.comment.ResponseData"
                        },
                        {
                            "code": 404,
                            "message": "not found error",
                            "responseType": "object",
                            "responseModel": "error"
                        },
                        {
                            "code": 500,
                            "message": "internal server error",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.satori.go.uuid.UUID": {
            "id": "github.com.satori.go.uuid.UUID",
            "properties": null
        },
        "github.com.team-a-hacks.tas-blog-api.comment.Payload": {
            "id": "github.com.team-a-hacks.tas-blog-api.comment.Payload",
            "properties": {
                "articleId": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "author": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "authorId": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "content": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdBy": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.team-a-hacks.tas-blog-api.comment.ResponseData": {
            "id": "github.com.team-a-hacks.tas-blog-api.comment.ResponseData",
            "properties": {
                "articleId": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "author": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "content": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "postDate": {
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
