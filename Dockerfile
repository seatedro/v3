ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app ./cmd/api/...


FROM debian:bookworm

WORKDIR /usr/src/app
COPY --from=builder /run-app /usr/src/app/server
COPY . .
CMD ["./server"]
