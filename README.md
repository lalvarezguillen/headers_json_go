![Build Status](https://travis-ci.org/lalvarezguillen/headers_json_go.svg?branch=master) [![codecov](https://codecov.io/gh/lalvarezguillen/headers_json_go/branch/master/graph/badge.svg)](https://codecov.io/gh/lalvarezguillen/headers_json_go)

Toy project to test Golang's Echo.

It's a simple API that takes a URL, makes a HEAD request to it, and returns the response's headers as JSON.

Get the service running:
```bash
$ go run headers_json.go

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v3.2.3
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
```

Example using [httpie](https://httpie.org/):
```bash
$ http get localhost:8080?url=www.ip-api.com/json

HTTP/1.1 200 OK
Content-Length: 163
Content-Type: application/json; charset=UTF-8
Date: Sat, 23 Dec 2017 11:22:57 GMT

{
    "Access-Control-Allow-Origin": "*",
    "Content-Length": "289",
    "Content-Type": "application/json; charset=utf-8",
    "Date": "Sat, 23 Dec 2017 11:22:57 GMT"
}
```

TODO:
  * Extend tests