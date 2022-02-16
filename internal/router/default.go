package router

import "github.com/bwmarrin/discordgo"

// !rustbot
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
				Name:   "Author",
				Value:  "https://github.com/tonyganga/rustbot",
				Inline: false,
			},
		},
	}
}

// !rustbot help
func HelpMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Commands",
		URL:         "https://github.com/tonyganga/rustbot#usage",
		Description: "The following are the available commands for Rustbot",
		Color:       0x93C54B,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "top",
				Value:  "The top command will return the top 25 ranked Rust servers on BattleMetrics.",
				Inline: false,
			},
			{
				Name:   "server",
				Value:  "The server command will lookup the ID you provided and return you detailed information about the Rust server. You can get the ID from the `!rustbot top` command.",
				Inline: false,
			},
			{
				Name:   "search",
				Value:  "The search command will perform a search based off your query. The search is text based, so you can do something like `!rustbot search rustoria` to get a list of all the Rustoria servers, sorted by rank.",
				Inline: false,
			},
			{
				Name:   "commits",
				Value:  "The commits command will link you to the Facepunch official site to view the latest commits for the game.",
				Inline: false,
			},
			{
				Name:   "roadmap",
				Value:  "The roadmap command will link you to the Rust roadmap.",
				Inline: false,
			},
		},
	}
}
