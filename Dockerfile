FROM golang:1.18 as builder
RUN apt update && apt install -y upx
WORKDIR /work
COPY hello-httpd.go /work
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" hello-httpd.go
RUN upx hello-httpd

FROM scratch
COPY --from=builder /work/hello-httpd /
CMD ["/hello-httpd"]
