FROM golang:1.22.4

RUN adduser --disabled-password --gecos '' gouser

RUN mkdir -p /home/gouser/project

WORKDIR /home/gouser/project

COPY . .

RUN chown -R gouser:gouser /home/gouser/project

USER gouser

ENV SERVER_PORT=8080

RUN go mod tidy
RUN go build -o build/server cmd/main.go

EXPOSE 8080

CMD ["./build/server"]
 