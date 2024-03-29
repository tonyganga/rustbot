package router

import (
	"log"
	"regexp"
	"time"

	"github.com/Necroforger/dgwidgets"
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

	// show paginated menu when only the BOT_KEYWORD is sent to a channel
	if m.Content == BOT_KEYWORD {
		p := dgwidgets.NewPaginator(s, m.ChannelID)
		p.DeleteMessageWhenDone = true
		p.Widget.Timeout = time.Minute * 2

		p.Add(ReactionMessage(), InfoMessage())
		p.SetPageFooters()

		err := p.Widget.Handle(UPWARD_TREND_EMOJI, func(w *dgwidgets.Widget, r *discordgo.MessageReaction) {
			ids := battlemetrics.GetListOfRustServers()
			p.Add(ids.RankedServerListMessage())
			err := p.Goto(len(p.Pages) - 1)
			if err != nil {
				log.Print(err)
			}
			err = p.Update()
			if err != nil {
				log.Print(err)
			}

		})
		if err != nil {
			log.Print(err)
		}

		err = p.Widget.Handle(ID_EMOJI, func(w *dgwidgets.Widget, r *discordgo.MessageReaction) {
			if msg, err := w.QueryInput("What is the ID of the Rust server you are looking for?", r.UserID, 15*time.Second); err == nil {
				id := msg.Content
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
				p.Add(server.RustServerMessage())
				err = p.Goto(len(p.Pages) - 1)
				if err != nil {
					log.Print(err)
				}
				err = p.Update()
				if err != nil {
					log.Print(err)
				}
			}
		})
		if err != nil {
			log.Print(err)
		}

		err = p.Widget.Handle(SEARCH_EMOJI, func(w *dgwidgets.Widget, r *discordgo.MessageReaction) {
			if msg, err := w.QueryInput("What Rust server are you looking for?", r.UserID, 15*time.Second); err == nil {
				var query string
				content := msg.Content
				for _, v := range content {
					query += string(v)
				}

				ids := battlemetrics.GetListOfRustServers(query)
				p.Add(ids.RankedServerListMessage())
				err := p.Goto(len(p.Pages) - 1)
				if err != nil {
					log.Print(err)
				}
				err = p.Update()
				if err != nil {
					log.Print(err)
				}
			}
		})
		if err != nil {
			log.Print(err)
		}

		err = p.Spawn()
		if err != nil {
			log.Print(err)
		}
		return
	}
}
