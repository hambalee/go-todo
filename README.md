# go-todo

mkdir todo && cd todo

git init

go mod init github.com/hambalee/go-todo

go get -u github.com/gin-gonic/gin

PORT=8081 go run main.go

```
go build \
-ldflags "-X main.buildcommit=`git rev-parse --short HEAD` \
-X main.buildtime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
-o app
```