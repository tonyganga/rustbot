package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	BOT_KEYWORD = "!rustbot"
)

func rustHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore bot messages
	if m.Author.Bot {
		return
	}

	// !rustbot [command]
	sc := strings.Split(strings.TrimSpace(m.Content), " ")
	if len(sc) < 1 {
		return
	}

	// ignore message that don't start with bot keyword
	if sc[0] != BOT_KEYWORD {
		return
	}

	// send default help message when only keyword is provided
	if m.Content == BOT_KEYWORD {
		s.ChannelMessageSendEmbed(m.ChannelID, infoMessage())
		return
	}

	switch sc[1] {
	case "top":
		{
			ids := getRankedRustServerList()
			s.ChannelMessageSend(m.ChannelID, getListOfRustServerIds(ids))
		}
	case "server":
		// !rustbot server [id]
		if len(sc) >= 2 {
			id := sc[2]
			match, err := regexp.MatchString(`^[0-9]+$`, id)
			if err != nil {
				log.Fatal(err)
			}
			if !match {
				s.ChannelMessageSend(m.ChannelID, "That doesn't look like a valid server ID.")
				return
			}

			server := getRustServer(id)
			s.ChannelMessageSendEmbed(m.ChannelID, server.rustServerMessage())
		}
	case "commits":
		{
			s.ChannelMessageSend(m.ChannelID, "https://rust.facepunch.com/changes/1")
		}
	case "help":
		{
			s.ChannelMessageSend(m.ChannelID, helpMessage())
		}
	}
}

func infoMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "Rustbot",
		Color: 0x93C54B,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "About",
				Value:  "Rustbot is a Discord bot created to help decide what Rust server you might want to play on. To view all available commands type `!rustbot help`.",
				Inline: false,
			},
			{
				Name:   "Author",
				Value:  "https://github.com/tonyganga/rustbot",
				Inline: false,
			},
		},
	}
}

func helpMessage() string {
	var message string
	ticks := "```"
	message =
		`
Rustbot is a Discord bot to help you pick a Rust server to play on.

!rustbot top - The top command will return the top 25 ranked Rust servers on BattleMetrics.
!rustbot server [id] - The server command will lookup the ID you provided and return you detailed information about the Rust server. You can get the ID from the !rustbot top command.
!rustbot commits - The commits command will link you to the Facepunch official site to view the latest commits for the game.
`

	return fmt.Sprintf("%v%v%v", ticks, message, ticks)

}
