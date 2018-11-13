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
Sending build context to Docker daemon  586.8kB
Step 1/19 : FROM golang:alpine as builder
 ---> 95ec94706ff6
Step 2/19 : RUN apk update && apk add git && apk add ca-certificates
 ---> Using cache
 ---> 54d6a21f447e
Step 3/19 : ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
Downloading [==================================================>]  8.499MB/8.499MB
 ---> Using cache
 ---> 45228515a4fa
Step 4/19 : RUN chmod +x /usr/bin/dep
 ---> Using cache
 ---> 651fe51ce825
Step 5/19 : RUN adduser -D -g '' rustbot
 ---> Using cache
 ---> 0aac8e7dd146
Step 6/19 : COPY . $GOPATH/src/rust-discord-bot/
 ---> 29c8e35b23f1
Step 7/19 : WORKDIR $GOPATH/src/rust-discord-bot/
 ---> Running in 3d5df20a76ed
Removing intermediate container 3d5df20a76ed
 ---> dd6322beb620
Step 8/19 : ENV CGO_ENABLED 0
 ---> Running in 4d6a29d41215
Removing intermediate container 4d6a29d41215
 ---> e1d053ec8107
Step 9/19 : ENV GOOS linux
 ---> Running in 0c076e6bc611
Removing intermediate container 0c076e6bc611
 ---> 3ca3851c5f60
Step 10/19 : ENV GOARCH amd64
 ---> Running in 29e7b34a83f2
Removing intermediate container 29e7b34a83f2
 ---> 0f423e95c71a
Step 11/19 : RUN dep ensure --vendor-only
 ---> Running in 8418e9861d00
Removing intermediate container 8418e9861d00
 ---> b28936807e04
Step 12/19 : RUN go test -v
 ---> Running in 574bf474d0b9
?   	rust-discord-bot	[no test files]
Removing intermediate container 574bf474d0b9
 ---> 6d1f56603c21
Step 13/19 : RUN go build -a -installsuffix cgo -o /go/bin/rust-discord-bot
 ---> Running in f0123d597465
Removing intermediate container f0123d597465
 ---> 4b4a3b78efe2
Step 14/19 : FROM scratch
 ---> 
Step 15/19 : COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
 ---> Using cache
 ---> e569bd5eff66
Step 16/19 : COPY --from=builder /etc/passwd /etc/passwd
 ---> Using cache
 ---> cb4ee665795a
Step 17/19 : COPY --from=builder /go/bin/rust-discord-bot /go/bin/rust-discord-bot
 ---> Using cache
 ---> 3732fad79fcd
Step 18/19 : USER rustbot
 ---> Using cache
 ---> b4f655e18993
Step 19/19 : ENTRYPOINT ["/go/bin/rust-discord-bot"]
 ---> Using cache
 ---> 2955612b828a
Successfully built 2955612b828a
Successfully tagged rustbot:latest
```

### Deploying

While you still have not done anything it's time to deploy your application. By running the following command you can deploy your application absolutely.

```
# docker pull rustbot:latest
$ docker run -e "DISCORD_TOKEN=XXXXXXXXXXXXXXXXXXXXXX" rustbot
Bot is now running.  Press CTRL-C to exit.
```

