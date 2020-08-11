# Rustbot
![ci](https://github.com/tonyganga/rustbot/workflows/ci/badge.svg)

Rustbot is a [discord](https://discordapp.com/) bot written using the [discordgo](https://github.com/bwmarrin/discordgo) framework. I created it one night when a friend asked me if there was a queue and used it as an excuse to play with some go.

![server](docs/server.png)

### Usage

In any authorized Discord channel type `!rustbot help` for a list of commands.

```
Rustbot is a discord bot to help you pick a rust server to play on.

!rustbot top - the top command will return the top 25 ranked rust servers on battlemetrics.
!rustbot server [id] - the server command will lookup the id you provided and return you detailed information about the rust server. you can get the id from the !rustbot top command.
!rustbot search [query] - the search command will perform a search based off your query. it will return you a list of servers matching the search sorted by rank.
!rustbot commits - the commits command will link you to the facepunch official site to view the latest commits for the game.
!rustbot roadmap - the roadmap command will link you to the rust roadmap.
```

### Building the Application


```
$ make build
```

Yep. That's it. You should see the following output:

```
docker build -t rustbot:latest .
Sending build context to Docker daemon  26.02MB
Step 1/15 : FROM golang:alpine as builder
 ---> ccda0e5ccbfc
Step 2/15 : RUN apk update && apk add ca-certificates --no-cache
 ---> Using cache
 ---> 38d67435d23a
Step 3/15 : RUN adduser -D -g '' rustbot
 ---> Using cache
 ---> 94656f13e2d7
Step 4/15 : COPY . $GOPATH/src/rust-discord-bot/
 ---> 1037cf1c3782
Step 5/15 : WORKDIR $GOPATH/src/rust-discord-bot/
 ---> Running in 5221381450a7
Removing intermediate container 5221381450a7
 ---> 84799e1fc50f
Step 6/15 : ENV CGO_ENABLED 0
 ---> Running in 6b2f756cee1b
Removing intermediate container 6b2f756cee1b
 ---> 9a9cc2058878
Step 7/15 : ENV GOOS linux
 ---> Running in 143fe7e39175
Removing intermediate container 143fe7e39175
 ---> 3dd045f68b9a
Step 8/15 : ENV GOARCH amd64
 ---> Running in bc6e650b0620
Removing intermediate container bc6e650b0620
 ---> bdf3a14a0aa3
Step 9/15 : RUN go build -a -installsuffix cgo -o /go/bin/rust-discord-bot
 ---> Running in f8a33198eb48
go: downloading github.com/bwmarrin/discordgo v0.19.0
go: downloading golang.org/x/crypto v0.0.0-20181112202954-3d3f9f413869
go: downloading github.com/gorilla/websocket v1.4.0
Removing intermediate container f8a33198eb48
 ---> af4014629fd2
Step 10/15 : FROM scratch
 --->
Step 11/15 : COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
 ---> Using cache
 ---> 97433d6c14b4
Step 12/15 : COPY --from=builder /etc/passwd /etc/passwd
 ---> Using cache
 ---> e08a169157b1
Step 13/15 : COPY --from=builder /go/bin/rust-discord-bot /go/bin/rust-discord-bot
 ---> 455c27e2d03f
Step 14/15 : USER rustbot
 ---> Running in f234d01a5dde
Removing intermediate container f234d01a5dde
 ---> 0412dc3cea69
Step 15/15 : ENTRYPOINT ["/go/bin/rust-discord-bot"]
 ---> Running in 30c2f793b14f
Removing intermediate container 30c2f793b14f
 ---> b494c16dbadb
Successfully built b494c16dbadb
Successfully tagged rustbot:latest
```

### Deploying

```
$ make build
$ docker run -e "DISCORD_TOKEN=XXXXXXXXXXXXXXXXXXXXXX" rustbot:latest
Bot is now running.  Press CTRL-C to exit.
```
