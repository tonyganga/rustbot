FROM golang:alpine as builder
# install git, ca-certificates 
RUN apk update && apk add ca-certificates --no-cache
# create rustbot to run the binary
RUN adduser -D -g '' rustbot
# copy source
COPY . $GOPATH/src/rust-discord-bot/
WORKDIR $GOPATH/src/rust-discord-bot/
# set ENV to build binary for scratch
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64 
# test and build
#RUN go test -v ./...
RUN go build -a -installsuffix cgo -o /go/bin/rust-discord-bot

# STEP 2 build a small image
FROM scratch
# copy certs, user and binary
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/rust-discord-bot /go/bin/rust-discord-bot
# define user
USER rustbot
# run binary
ENTRYPOINT ["/go/bin/rust-discord-bot"]
