# go-todo

mkdir todo && cd todo

git init

go mod init github.com/hambalee/go-todo

go get -u github.com/gin-gonic/gin

PORT=8081 go run main.go

## ldflags

```
go build \
-ldflags "-X main.buildcommit=`git rev-parse --short HEAD` \
-X main.buildtime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
-o app
```

## Liveness Probe Readiness Probe
`cat /tmp/live`

`echo $?`

## Rate Limit
`echo "GET http://:8081/limitz" | vegeta attack -rate=10/s -duration=1s | vegeta report`

