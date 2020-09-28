FROM golang:1.14-alpine3.11 AS builder

ENV GOPROXY https://goproxy.io

WORKDIR /build
COPY go.mod .
COPY go.sum .

# 新增用户
RUN adduser -u 10001 -D app-runner
# 编译
COPY . .
COPY conf ./conf
COPY controllers ./controllers
COPY models ./models
COPY routers ./routers
COPY static ./static
COPY tests ./tests
COPY views ./views

RUN go mod download; \
    go get github.com/beego/bee; \
    bee pack -be GOOS=linux -a run
    # CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o run .;

FROM alpine:3.11 AS final

WORKDIR /app
COPY --from=builder /build/run.tar.gz /app/
RUN tar -xf ./run.tar.gz && rm -rf ./run.tar.gz
# 获取最新数据
ADD http://data.haifengat.com/instrument.csv .

#USER app-runner
ENTRYPOINT ["./run"]
