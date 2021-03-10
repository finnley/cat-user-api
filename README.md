# cat-user-api

## 生成 grpc

```shell
protoc -I . user.proto --go_out=plugins=grpc:.
```

## zap

```shell
go get -u go.uber.org/zap
```

## gin

```shell
go get -u github.com/gin-gonic/gin
```