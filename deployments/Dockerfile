FROM golang:latest as builder
RUN mkdir /app
ADD / /app/
WORKDIR /app

RUN GOOS=linux GOARCH=amd64 go build -o /app/postbox /app/cmd/main.go

#поднимаем бинарик
FROM scratch

COPY --from=builder /app/postbox /usr/bin/postbox

ENTRYPOINT [ "/usr/bin/postbox" ]
