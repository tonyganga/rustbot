package battlemetrics

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// !rustbot server [id]
func (r *RustServer) RustServerMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       r.Data.Attributes.Name,
		Description: r.Data.Attributes.Details.RustDescription,
		URL:         r.Data.Attributes.Details.RustURL,
		Color:       0x93C54B,
		Image: &discordgo.MessageEmbedImage{
			URL: r.Data.Attributes.Details.RustHeaderimage,
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Server Rank",
				Value:  fmt.Sprintf("%v", r.Data.Attributes.Rank),
				Inline: false,
			},
			{
				Name:   "Last Wipe",
				Value:  fmt.Sprintf("%v", r.Data.Attributes.Details.RustLastWipe.Format("2006-01-02 15:04:05")),
				Inline: false,
			},
			{
				Name:   "Players Online/Queue",
				Value:  fmt.Sprintf("%v/%v (%v)", r.Data.Attributes.Players, r.Data.Attributes.MaxPlayers, r.Data.Attributes.Details.RustQueuedPlayers),
				Inline: false,
			},
			{
				Name:   "Average FPS",
				Value:  fmt.Sprintf("%v", r.Data.Attributes.Details.RustFpsAvg),
				Inline: false,
			},
			{
				Name:   "Map Size",
				Value:  fmt.Sprintf("%v", r.Data.Attributes.Details.RustWorldSize),
				Inline: false,
			},
			{
				Name:   "Connection Information",
				Value:  fmt.Sprintf("client.connect %v:%v", r.Data.Attributes.IP, r.Data.Attributes.Port),
				Inline: false,
			},
		},
	}
}
