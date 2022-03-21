FROM golang:1.17

WORKDIR /go/src/go-tech-dojo
COPY . .

RUN go mod download
RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]

