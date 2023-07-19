FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
RUN go mod download
RUN go build main.go

FROM alpine as greeter-client
WORKDIR /
COPY --from=builder /app/main .
USER 65523:65523
ENTRYPOINT ["/main"]