FROM alpine:latest

MAINTAINER pdf "运行在alpha环境go build的时候要设置CGO_ENABLED=0"

COPY metrics /app/

WORKDIR /app/

CMD ["./metrics"]