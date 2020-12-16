FROM golang:latest

WORKDIR /root/go/src/mSystem/svc/user

COPY / /root/go/src/mSystem/svc/user

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go build -o user

ENTRYPOINT ["./user"]
