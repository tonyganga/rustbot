package battlemetrics

import "time"

type RustServerMetadata struct {
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
			RustFps            int         `json:"rust_fps"`
			RustBuild          string      `json:"rust_build"`
			RustQueuedPlayers  int         `json:"rust_queued_players"`
			RustGcMb           int         `json:"rust_gc_mb"`
			RustFpsAvg         float64     `json:"rust_fps_avg"`
			Official           bool        `json:"official"`
			RustURL            string      `json:"rust_url"`
			RustWorldSeed      int         `json:"rust_world_seed"`
			RustLastEntDrop    time.Time   `json:"rust_last_ent_drop"`
			RustGcCl           int         `json:"rust_gc_cl"`
			Map                string      `json:"map"`
			RustHeaderimage    string      `json:"rust_headerimage"`
			RustMemPv          interface{} `json:"rust_mem_pv"`
			RustLastWipe       time.Time   `json:"rust_last_wipe"`
			RustLastWipeEnt    int         `json:"rust_last_wipe_ent"`
			RustEntCntI        int         `json:"rust_ent_cnt_i"`
			RustBorn           time.Time   `json:"rust_born"`
			RustMemWs          interface{} `json:"rust_mem_ws"`
			RustDescription    string      `json:"rust_description"`
			ServerSteamID      string      `json:"serverSteamId"`
			RustUptime         int         `json:"rust_uptime"`
			RustType           string      `json:"rust_type"`
			Environment        string      `json:"environment"`
			Pve                bool        `json:"pve"`
			RustWorldSize      int         `json:"rust_world_size"`
			RustHash           string      `json:"rust_hash"`
			RustModded         bool        `json:"rust_modded"`
			RustLastSeedChange time.Time   `json:"rust_last_seed_change"`
		} `json:"details"`
		Private   bool      `json:"private"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		PortQuery int       `json:"portQuery"`
		Country   string    `json:"country"`
	} `json:"attributes"`
	Relationships struct {
		Game struct {
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"game"`
	} `json:"relationships"`
}

type RustServer struct {
	Data RustServerMetadata
}

type RustServers struct {
	Data []RustServerMetadata
}
