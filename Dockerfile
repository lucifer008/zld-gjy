FROM golang:1.18 as builder
# 启用go module

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

#拷贝包依赖
COPY go.mod go.sum ./
#下载依赖
RUN go mod download && go mod verify

#指定环境变量
RUN go env -w GO111MODULE=on

#构建程序
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

#拷贝编译程序到publish目录
RUN mkdir publish && cp main app.yaml publish

# 运行阶段指定scratch作为基础镜像
FROM alpine:3.10

WORKDIR /app

#拷贝发布文件到gopath
COPY --from=builder /app/publish .

# 指定运行时环境变量
ENV GIN_MODE=release

EXPOSE 5000

CMD ["/app/main"]
