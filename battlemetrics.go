package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

const BattleMetricsURL = "https://api.battlemetrics.com"
const RustFilter = "?filter[game]=rust"
const PageFilter = "&page[size]=25"

type ServerList struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			ID         string      `json:"id"`
			Name       string      `json:"name"`
			Address    interface{} `json:"address"`
			IP         string      `json:"ip"`
			Port       int         `json:"port"`
			Players    int         `json:"players"`
			MaxPlayers int         `json:"maxPlayers"`
			Rank       int         `json:"rank"`
			Location   []float64   `json:"location"`
			Status     string      `json:"status"`
			Details    struct {
				Official           bool        `json:"official"`
				RustType           string      `json:"rust_type"`
				Map                string      `json:"map"`
				Environment        string      `json:"environment"`
				RustBuild          string      `json:"rust_build"`
				RustEntCntI        int         `json:"rust_ent_cnt_i"`
				RustFps            int         `json:"rust_fps"`
				RustFpsAvg         float64     `json:"rust_fps_avg"`
				RustGcCl           int         `json:"rust_gc_cl"`
				RustGcMb           int         `json:"rust_gc_mb"`
				RustHash           string      `json:"rust_hash"`
				RustHeaderimage    string      `json:"rust_headerimage"`
				RustMemPv          interface{} `json:"rust_mem_pv"`
				RustMemWs          interface{} `json:"rust_mem_ws"`
				Pve                bool        `json:"pve"`
				RustUptime         int         `json:"rust_uptime"`
				RustURL            string      `json:"rust_url"`
				RustWorldSeed      int         `json:"rust_world_seed"`
				RustWorldSize      int         `json:"rust_world_size"`
				RustDescription    string      `json:"rust_description"`
				RustModded         bool        `json:"rust_modded"`
				RustQueuedPlayers  int         `json:"rust_queued_players"`
				RustLastEntDrop    time.Time   `json:"rust_last_ent_drop"`
				RustLastSeedChange time.Time   `json:"rust_last_seed_change"`
				RustLastWipe       time.Time   `json:"rust_last_wipe"`
				RustLastWipeEnt    int         `json:"rust_last_wipe_ent"`
				ServerSteamID      string      `json:"serverSteamId"`
			} `json:"details"`
			CreatedAt   time.Time `json:"createdAt"`
			UpdatedAt   time.Time `json:"updatedAt"`
			PortQuery   int       `json:"portQuery"`
			Country     string    `json:"country"`
			QueryStatus string    `json:"queryStatus"`
		} `json:"attributes"`
		Relationships struct {
			Game struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
			} `json:"game"`
		} `json:"relationships"`
	} `json:"data"`
	Links struct {
		Next string `json:"next"`
	} `json:"links"`
	Included []interface{} `json:"included"`
}

type RustServer struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			ID         string    `json:"id"`
			Name       string    `json:"name"`
			Address    string    `json:"address"`
			IP         string    `json:"ip"`
			Port       int       `json:"port"`
			Players    int       `json:"players"`
			MaxPlayers int       `json:"maxPlayers"`
			Rank       int       `json:"rank"`
			Location   []float64 `json:"location"`
			Status     string    `json:"status"`
			Details    struct {
				Official          bool        `json:"official"`
				RustType          string      `json:"rust_type"`
				Map               string      `json:"map"`
				Environment       string      `json:"environment"`
				RustBuild         string      `json:"rust_build"`
				RustEntCntI       int         `json:"rust_ent_cnt_i"`
				RustFps           int         `json:"rust_fps"`
				RustFpsAvg        float64     `json:"rust_fps_avg"`
				RustGcCl          int         `json:"rust_gc_cl"`
				RustGcMb          int         `json:"rust_gc_mb"`
				RustHash          string      `json:"rust_hash"`
				RustHeaderimage   string      `json:"rust_headerimage"`
				RustMemPv         interface{} `json:"rust_mem_pv"`
				RustMemWs         interface{} `json:"rust_mem_ws"`
				Pve               bool        `json:"pve"`
				RustUptime        int         `json:"rust_uptime"`
				RustURL           string      `json:"rust_url"`
				RustWorldSeed     int         `json:"rust_world_seed"`
				RustWorldSize     int         `json:"rust_world_size"`
				RustDescription   string      `json:"rust_description"`
				RustModded        bool        `json:"rust_modded"`
				RustQueuedPlayers int         `json:"rust_queued_players"`
				RustLastEntDrop   time.Time   `json:"rust_last_ent_drop"`
				RustLastWipe      time.Time   `json:"rust_last_wipe"`
				RustLastWipeEnt   int         `json:"rust_last_wipe_ent"`
				ServerSteamID     string      `json:"serverSteamId"`
			} `json:"details"`
			CreatedAt   time.Time `json:"createdAt"`
			UpdatedAt   time.Time `json:"updatedAt"`
			PortQuery   int       `json:"portQuery"`
			Country     string    `json:"country"`
			QueryStatus string    `json:"queryStatus"`
		} `json:"attributes"`
		Relationships struct {
			Game struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
			} `json:"game"`
		} `json:"relationships"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}

func (r *RustServer) rustServerMessage() *discordgo.MessageEmbed {
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
				Name:   "Players Online",
				Value:  fmt.Sprintf("%v/%v", r.Data.Attributes.Players, r.Data.Attributes.MaxPlayers),
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
				Value:  fmt.Sprintf("%v:%v", r.Data.Attributes.IP, r.Data.Attributes.Port),
				Inline: false,
			},
		},
	}
}

func getRustServer(id string) RustServer {

	url := fmt.Sprintf("%s/servers/%s", BattleMetricsURL, id)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	info, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var servers RustServer
	err = json.Unmarshal(info, &servers)
	if err != nil {
		log.Fatal(err)
	}
	return servers
}

func getRankedRustServerList() map[string]string {

	results := make(map[string]string)
	url := fmt.Sprintf("%v/servers%v%v", BattleMetricsURL, RustFilter, PageFilter)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	info, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var list ServerList
	err = json.Unmarshal(info, &list)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range list.Data {
		results[v.Attributes.Name] = v.Attributes.ID
	}

	return results
}

func getListOfRustServerIds(ids map[string]string) string {

	var message string
	ticks := "```"

	for k, v := range ids {
		message += fmt.Sprintf("ID: %v Server: %v\n", v, k)
	}

	return fmt.Sprintf("%v\n%v%v", ticks, message, ticks)
}
