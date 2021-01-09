FROM golang:latest
WORKDIR /mSystem/svc/gateway
ADD . /mSystem/svc/gateway
RUN go env -w GOPROXY=https://goproxy.cn,direct

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
#RUN  go build -o user
# RUN apk add --no-cache gcc musl-dev git && go build -I -v
RUN CGO_ENABLED=0 GOOS=linux go build -o gateway -a -tags netgo -ldflags '-s -w' .
EXPOSE 9099
ENTRYPOINT ["./gateway"]