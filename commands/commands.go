package commands

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tonyganga/rustbot/battlemetrics"
)

const (
	BOT_KEYWORD = "!rustbot"
)

func RustHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
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
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, infoMessage())
		if err != nil {
			log.Print(err)
		}
		return
	}

	switch sc[1] {
	case "top":
		{
			ids := battlemetrics.GetListOfRustServers("")
			_, err := s.ChannelMessageSend(m.ChannelID, battlemetrics.GetRankedListOfRustServers(ids))
			if err != nil {
				log.Print(err)
			}
		}
	case "server":
		if len(sc) == 2 {
			_, err := s.ChannelMessageSend(m.ChannelID, "No options passed to !rustbot server.")
			if err != nil {
				log.Print(err)
			}
			return
		}
		// !rustbot server [id]
		if len(sc) > 2 {
			id := sc[2]
			match, err := regexp.MatchString(`^[0-9]+$`, id)
			if err != nil {
				log.Print(err)
			}
			if !match {
				_, err := s.ChannelMessageSend(m.ChannelID, "That doesn't look like a valid server ID.")
				if err != nil {
					log.Print(err)
				}
				return
			}

			server := battlemetrics.GetRustServer(id)
			_, err = s.ChannelMessageSendEmbed(m.ChannelID, server.RustServerMessage())
			if err != nil {
				log.Print(err)
			}
		}
	case "commits":
		{
			_, err := s.ChannelMessageSend(m.ChannelID, "https://commits.facepunch.com/r/rust_reboot")
			if err != nil {
				log.Print(err)
			}
		}
	case "roadmap":
		{
			_, err := s.ChannelMessageSend(m.ChannelID, "https://rust.nolt.io/roadmap")
			if err != nil {
				log.Print(err)
			}
		}
	case "search":
		if len(sc) == 2 {
			_, err := s.ChannelMessageSend(m.ChannelID, "No options passed to !rustbot search.")
			if err != nil {
				log.Print(err)
			}
			return
		}
		// !rustbot search [query]
		if len(sc) > 2 {
			var query string
			for _, v := range sc[2:] {
				query += v + "+"
			}
			ids := battlemetrics.GetListOfRustServers(query)
			_, err := s.ChannelMessageSend(m.ChannelID, battlemetrics.GetRankedListOfRustServers(ids))
			if err != nil {
				log.Print(err)
			}
		}
	case "help":
		{
			_, err := s.ChannelMessageSend(m.ChannelID, helpMessage())
			if err != nil {
				log.Print(err)
			}
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
!rustbot search [query] - The search command will perform a search based off your query. It will return you a list of servers matching the search sorted by rank.
!rustbot commits - The commits command will link you to the Facepunch official site to view the latest commits for the game.
!rustbot roadmap - The roadmap command will link you to the Rust roadmap.
`

	return fmt.Sprintf("%v%v%v", ticks, message, ticks)

}
