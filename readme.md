# Golang API Calculate Number

## Todo
- Create api from net/http
- Create mock some function for expect response when error case such as BadRequest, InternalServerError, etc
- Make logger middleware

## Start Server
```sh
go run cmd/main.go
# default port: 8888
```

## Manual Test with cURL
```sh
curl 127.0.0.1:8888/calculate -X POST -I
# HTTP/1.1 400 Bad Request

curl 127.0.0.1:8888/calculate -I
# HTTP/1.1 404 Not Found

curl 127.0.0.1:8888/calculate -X POST -d '{"number1": 3, "number2": 3}'
# {"result":6}
```
