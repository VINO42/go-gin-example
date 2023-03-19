使用scratch镜像编译需要添加参数
``CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin-example .
``