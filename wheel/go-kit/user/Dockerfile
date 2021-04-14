## docker build -t user .
FROM golang:latest
WORKDIR /root/microService/user
COPY / /root/microService/user
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
go build -o user
EXPOSE 10086
ENTRYPOINT ["./user"]
