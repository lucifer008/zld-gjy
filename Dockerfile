FROM golang:1.18
# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

# CGO_ENABLED禁用cgo 然后指定OS等，并go build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

RUN mkdir publish && cp main app.yaml publish


# 运行阶段指定scratch作为基础镜像
FROM alpine:3.10
RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \

WORKDIR /app


# 将上一个阶段publish文件夹下的所有文件复制进来
#COPY --from=builder /app/publish .


# 指定运行时环境变量
ENV GIN_MODE=release

EXPOSE 5000

CMD ["/app/main", "-c", "/app.yaml"]
