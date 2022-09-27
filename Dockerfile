FROM golang:1.18 as builder
# 启用go module

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /usr/local/go/src/zld-jy

COPY . .

#拷贝包依赖
COPY go.mod go.sum ./
#下载依赖
RUN go mod download && go mod verify

#指定环境变量
RUN go env -w GO111MODULE=on

RUN go build /usr/local/go/src/zld-jy/main.go

# 运行阶段指定scratch作为基础镜像
FROM alpine:3.10

COPY --from=builder /usr/local/go/src/zld-jy /usr/local/go/src/zld-jy

# 指定运行时环境变量
ENV GIN_MODE=release

EXPOSE 5000

CMD ["/usr/local/go/src/zld-jy/main"]
