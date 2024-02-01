FROM golang:1.21-alpine as build
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o service cmd/service/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build ./app/service /app

CMD ["/app/service"]
