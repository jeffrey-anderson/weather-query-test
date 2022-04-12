FROM golang:1.18 AS builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/query-test main.go

# FROM scratch
# COPY --from=builder /usr/local/bin/query-test /query-test

CMD ["query-test","-sampleSize", "100"]
