FROM reg.pdf.cn/library/golang:1.13 as builder

WORKDIR /go/src/ci-test

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM reg.pdf.cn/library/alpine as prod

WORKDIR /

COPY --from=0 /go/src/ci-test/app .

CMD [ "./app" ]