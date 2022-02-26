package battlemetrics

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	BattleMetricsURL = "https://api.battlemetrics.com"
	RustFilter       = "?filter[game]=rust"
	CountryFilter    = "&filter[countries][0]=US&filter[countries][1]=CA"
	PageFilter       = "&page[size]=25"
	SearchFilter     = "&filter[search]="
)

func GetRustServer(id string) RustServer {
	url := fmt.Sprintf("%s/servers/%s", BattleMetricsURL, id)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}

	var server RustServer
	err = json.NewDecoder(res.Body).Decode(&server)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()

	return server
}

func GetListOfRustServers(query ...string) RustServers {
	var q string
	for _, v := range query {
		q += string(v)
	}
	search := SearchFilter + q
	url := fmt.Sprintf("%v/servers%v%v%v%v", BattleMetricsURL, RustFilter, CountryFilter, PageFilter, search)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}

	var servers RustServers
	err = json.NewDecoder(res.Body).Decode(&servers)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()

	return servers
}
