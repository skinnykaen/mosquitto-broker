FROM golang:1.16-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o mqtt_broker ./cmd/main.go

CMD ["./mqtt_brokert"]