FROM golang:1.23.7-alpine3.21

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "cmd/server/main.go"]
EXPOSE 4000