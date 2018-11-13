package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// RustServer represents the information returned from /info/{id}
type RustServer struct {
	Hostname              string `json:"hostname"`
	IP                    string `json:"ip"`
	Port                  string `json:"port"`
	Map                   string `json:"map"`
	OnlineState           string `json:"online_state"`
	Checked               string `json:"checked"`
	PlayersMax            string `json:"players_max"`
	PlayersCurrent        string `json:"players_cur"`
	PlayersAverage        string `json:"players_avg"`
	PlayersMaxMan         string `json:"players_maxman"`
	PlayersMaxForever     string `json:"players_max_forever"`
	PlayersMaxForeverDate string `json:"players_max_forever_date"`
	Bots                  string `json:"bots"`
	Ratings               string `json:"rating"`
	Entities              string `json:"entities"`
	Version               string `json:"version"`
	Seed                  string `json:"seed"`
	Size                  string `json:"size"`
	Uptime                string `json:"uptime"`
	FPS                   string `json:"fps"`
	FPSAverage            string `json:"fps_avg"`
	URL                   string `json:"url"`
	Image                 string `json:"image"`
	OS                    string `json:"os"`
	Memory                string `json:"mem"`
	Country               string `json:"country"`
	CountryFull           string `json:"country_full"`
	ServerMode            string `json:"server_mode"`
	Wipe                  string `json:"wipe_cycle"`
	Queue                 bool
	QueueLine             int
}

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
	dg.AddHandler(lookupLowPop)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Don't finish main goroutine until some sort of system term is received on the sc channel.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Defer to close the Discord session at the end of the main goroutine.
	defer dg.Close()
}

func lookupLowPop(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!lowpop" {

		// this just checks 50 lowpop
		// @TODO make this lookup the ID first and refactor to work
		res, err := http.Get("https://api.rust-servers.info/info/50")
		if err != nil {
			log.Fatal(err)
		}
		// read body and unmarshal it into the RustServer struct
		info, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		res.Body.Close()

		var server RustServer
		json.Unmarshal(info, &server)

		checkQueue(server.PlayersCurrent, server.PlayersMax, &server)

		// Create embedded message with the server information (queue, connection info, etc)
		// based on some conditions, like if a queue exists and if the server is online or not.
		if server.Queue == true && server.OnlineState == "1" {
			line := strconv.Itoa(server.QueueLine)

			embed := &discordgo.MessageEmbed{
				Author:      &discordgo.MessageEmbedAuthor{},
				Color:       0xff0000, // Red
				Description: "Sucks. There's a queue of " + line + " for " + server.Hostname,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Server Information:",
						Value:  "Wipe: " + server.Wipe + "\nMode: " + server.ServerMode + "\nAverage FPS: " + server.FPSAverage + "\nPlayers Online: " + server.PlayersCurrent,
						Inline: true,
					},
					{
						Name:   "Connection String:",
						Value:  "```client.connect " + server.IP + ":" + server.Port + "```",
						Inline: true,
					},
				},
				Image: &discordgo.MessageEmbedImage{
					URL: server.Image,
				},
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
				},
				Timestamp: time.Now().Format(time.RFC3339),
				Title:     server.Hostname,
			}

			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		} else if server.Queue == false && server.OnlineState == "1" {
			embed := &discordgo.MessageEmbed{
				Author:      &discordgo.MessageEmbedAuthor{},
				Color:       0x00ff00, // Green
				Description: "There's NO queue! Leggo zerg! :100:",
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Server Information:",
						Value:  "Wipe: " + server.Wipe + "\nMode: " + server.ServerMode + "\nAverage FPS: " + server.FPSAverage + "\nPlayers Online: " + server.PlayersCurrent,
						Inline: true,
					},
					{
						Name:   "Connection Info:",
						Value:  "```client.connect " + server.IP + ":" + server.Port + "```",
						Inline: true,
					},
				},
				Image: &discordgo.MessageEmbedImage{
					URL: server.Image,
				},
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
				},
				Timestamp: time.Now().Format(time.RFC3339),
				Title:     server.Hostname,
			}

			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		} else if server.OnlineState == "0" {
			embed := &discordgo.MessageEmbed{
				Author:      &discordgo.MessageEmbedAuthor{},
				Color:       0xff0000, // Red
				Description: server.Hostname + " is down. :no_entry_sign: ",
				Image: &discordgo.MessageEmbedImage{
					URL: server.Image,
				},
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
				},
				Timestamp: time.Now().Format(time.RFC3339),
				Title:     server.Hostname,
			}

			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		}
	}
}

func checkQueue(c string, m string, r *RustServer) {
	currentPlayers, err := strconv.Atoi(c)
	if err != nil {
		log.Fatal(err)
	}
	maxPlayers, err := strconv.Atoi(m)
	if err != nil {
		log.Fatal(err)
	}
	diff := (currentPlayers - maxPlayers)
	// if current users minus max users is not a negative number
	// return the difference and set Queue to true.
	if diff > 0 {
		r.Queue = true
		r.QueueLine = diff
	} else {
		r.Queue = false
	}
}
