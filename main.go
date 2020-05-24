package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/tonyganga/rustbot/commands"
)

func main() {
	// Lookup token from ENV
	Token := os.Getenv("DISCORD_TOKEN")
	if Token == "" {
		fmt.Println("Unable to find token, please make sure DISCORD_TOKEN is set.")
		return
	}

	// Create a new Discord session using Token from DISCORD_TOKEN
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register handlers
	dg.AddHandler(commands.RustHandler)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Don't finish main goroutine until some sort of system term is received on the sc channel.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Defer to close the Discord session at the end of the main goroutine.
	defer dg.Close()
}
