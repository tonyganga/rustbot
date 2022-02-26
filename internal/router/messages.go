package router

import "github.com/bwmarrin/discordgo"

const (
	SEARCH_EMOJI       = "🔎"
	UPWARD_TREND_EMOJI = "📈"
	ID_EMOJI           = "🆔"
)

func InfoMessage() *discordgo.MessageEmbed {
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
				Name:   "Rust Commits",
				Value:  "https://commits.facepunch.com/r/rust_reboot",
				Inline: false,
			},
			{
				Name:   "Rust Roadmap",
				Value:  "https://rust.nolt.io/roadmap",
				Inline: false,
			},
		},
	}
}

func ReactionMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Reactions",
		URL:         "https://github.com/tonyganga/rustbot#usage",
		Description: "The following are the available reactions for Rustbot",
		Color:       0x93C54B,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   UPWARD_TREND_EMOJI,
				Value:  "Reacting with this emoji will return the top 25 ranked Rust servers on BattleMetrics.",
				Inline: false,
			},
			{
				Name:   ID_EMOJI,
				Value:  "Reacting with this emoji will lookup the ID you provided and show you detailed information about the Rust server.",
				Inline: false,
			},
			{
				Name:   SEARCH_EMOJI,
				Value:  "Reacting with this emoji will search for all Rust servers matching your search term.",
				Inline: false,
			},
		},
	}
}
