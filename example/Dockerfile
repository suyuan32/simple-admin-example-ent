FROM golang:1.19.1-alpine3.16 as builder

WORKDIR /home
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -ldflags="-s -w" -o /home/example_api example.go

FROM alpine:latest

WORKDIR /home

COPY --from=builder /home/example_api ./
COPY --from=builder /home/etc/example.yaml ./

EXPOSE ᾑ
ENTRYPOINT ./example_api -f example.yaml