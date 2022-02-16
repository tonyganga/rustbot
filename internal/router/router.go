package router

import (
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tonyganga/rustbot/internal/battlemetrics"
)

const (
	BOT_KEYWORD = "!rustbot"
)

func RustHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore bot messages
	if m.Author.Bot {
		return
	}

	// !rustbot [command] [params]
	sc := strings.Split(strings.TrimSpace(m.Content), " ")

	// ignore message that don't start with bot keyword
	if sc[0] != BOT_KEYWORD {
		return
	}

	// send default help message when only keyword is provided
	if m.Content == BOT_KEYWORD {
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, InfoMessage())
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
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, HelpMessage())
			if err != nil {
				log.Print(err)
			}
		}
	}
}
