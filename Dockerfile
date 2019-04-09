FROM golang:latest
MAINTAINER kirio

# 设置工作目录
RUN mkdir -p /go/src
WORKDIR /go/src

# 添加应用
ADD . /go/src

# 构建服务
RUN make

# 运行服务
ENTRYPOINT ./admin.sh start
