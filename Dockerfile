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
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o run .;

FROM alpine:3.11 AS final

WORKDIR /app
COPY --from=builder /build/run /app/

#USER app-runner
ENTRYPOINT ["./run"]
