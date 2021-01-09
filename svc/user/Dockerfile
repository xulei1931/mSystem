FROM golang:latest
WORKDIR /mSystem/svc/user
ADD . /mSystem/svc/user
RUN go env -w GOPROXY=https://goproxy.cn,direct

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
#RUN  go build -o user
# RUN apk add --no-cache gcc musl-dev git && go build -I -v
RUN CGO_ENABLED=0 GOOS=linux go build -o user -a -tags netgo -ldflags '-s -w' .
EXPOSE 9001
ENTRYPOINT ["./user"]