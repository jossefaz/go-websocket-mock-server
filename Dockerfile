
############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/package/app/
COPY ./ ./
# Fetch dependencies.
# Using go get.
RUN go get -d -v ./... && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main
############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
WORKDIR /app


ENV MESSAGE_INTERVAL=5
ENV MOCK_MESSAGE="Default Message from Go Websocket Mock"

EXPOSE 8085

COPY --from=builder /go/bin/main .
# Run the hello binary.

CMD ["./main"]