FROM golangci/golangci-lint 
COPY . $GOPATH/src/rust-discord-bot
WORKDIR $GOPATH/src/rust-discord-bot/
RUN go mod download && \
    golangci-lint run -v && \
    go test -v ./...

FROM golang:alpine as builder
RUN apk update && apk add ca-certificates --no-cache
RUN adduser -D -g '' rustbot
COPY . $GOPATH/src/rust-discord-bot/
WORKDIR $GOPATH/src/rust-discord-bot/
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64 
RUN go build -a -installsuffix cgo -o /go/bin/rust-discord-bot

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/rust-discord-bot /go/bin/rust-discord-bot
USER rustbot
ENTRYPOINT ["/go/bin/rust-discord-bot"]
