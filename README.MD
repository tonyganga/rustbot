# rustbot
 
rustbot is a [discord](https://discordapp.com/) bot written using the [discordgo](https://github.com/bwmarrin/discordgo) framework. I created it one night when a friend asked me if there was a queue and used it as an excuse to play with some go.

### Usage

In any authorized Discord channel type `!lowpop` and rustbot will grab you info from https://rust-servers.info/api.html and return it in an embeded message.

### Building the Application

Now that you have not done anything it's time to build your application:

```
$ make build
```

Yep. That's it. You should see the following output:

```
docker build -t rustbot:latest .
Sending build context to Docker daemon  25.94MB
Step 1/16 : FROM golang:alpine as builder
 ---> ccda0e5ccbfc
Step 2/16 : RUN apk update && apk add git && apk add ca-certificates
 ---> Using cache
 ---> 4cbc84ad9250
Step 3/16 : RUN adduser -D -g '' rustbot
 ---> Using cache
 ---> 6fbf010edcca
Step 4/16 : COPY . $GOPATH/src/rust-discord-bot/
 ---> c3f5e109a13a
Step 5/16 : WORKDIR $GOPATH/src/rust-discord-bot/
 ---> Running in 2a8925b49638
Removing intermediate container 2a8925b49638
 ---> a4f08757641d
Step 6/16 : ENV CGO_ENABLED 0
 ---> Running in a7fbacfca472
Removing intermediate container a7fbacfca472
 ---> fe7d5ef560f0
Step 7/16 : ENV GOOS linux
 ---> Running in 2f3af266c8db
Removing intermediate container 2f3af266c8db
 ---> db5a9da06e7c
Step 8/16 : ENV GOARCH amd64
 ---> Running in 232e826104b2
Removing intermediate container 232e826104b2
 ---> 4f38e4899d19
Step 9/16 : RUN go test -v
 ---> Running in 26e0de5cf9d5
go: downloading github.com/bwmarrin/discordgo v0.19.0
go: downloading github.com/gorilla/websocket v1.4.0
go: downloading golang.org/x/crypto v0.0.0-20181112202954-3d3f9f413869
?       github.com/tonyganga/rustbot    [no test files]
Removing intermediate container 26e0de5cf9d5
 ---> 1a0b11087a70
Step 10/16 : RUN go build -a -installsuffix cgo -o /go/bin/rust-discord-bot
 ---> Running in 94fcf9cf492a
Removing intermediate container 94fcf9cf492a
 ---> 075955d5686d
Step 11/16 : FROM scratch
 ---> 
Step 12/16 : COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
 ---> Using cache
 ---> 97433d6c14b4
Step 13/16 : COPY --from=builder /etc/passwd /etc/passwd
 ---> Using cache
 ---> e08a169157b1
Step 14/16 : COPY --from=builder /go/bin/rust-discord-bot /go/bin/rust-discord-bot
 ---> Using cache
 ---> a58268dd4fb2
Step 15/16 : USER rustbot
 ---> Using cache
 ---> c7726cfd5502
Step 16/16 : ENTRYPOINT ["/go/bin/rust-discord-bot"]
 ---> Using cache
 ---> 985c2251c740
Successfully built 985c2251c740
Successfully tagged rustbot:latest
```

### Deploying

While you still have not done anything it's time to deploy your application. By running the following command you can deploy your application absolutely.

```
# docker pull rustbot:latest
$ docker run -e "DISCORD_TOKEN=XXXXXXXXXXXXXXXXXXXXXX" rustbot
Bot is now running.  Press CTRL-C to exit.
```

