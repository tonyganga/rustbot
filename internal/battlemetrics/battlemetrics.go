package battlemetrics

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	BattleMetricsURL = "https://api.battlemetrics.com"
	RustFilter       = "?filter[game]=rust"
	CountryFilter    = "&filter[countries][0]=US&filter[countries][1]=CA"
	PageFilter       = "&page[size]=25"
	SearchFilter     = "&filter[search]="
)

func GetRustServer(id string) RustServer {
	urlString := fmt.Sprintf("%s/servers/%s", BattleMetricsURL, id)
	url, err := url.Parse(urlString)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(url)

	client := http.Client{
		Timeout: time.Second * 10,
	}

	req := &http.Request{
		Method: "GET",
		URL:    url,
	}

	res, err := client.Do(req)
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
	search := fmt.Sprintf("%v%v", SearchFilter, q)

	urlString := fmt.Sprintf("%v/servers%v%v%v%v", BattleMetricsURL, RustFilter, CountryFilter, PageFilter, search)
	url, err := url.Parse(urlString)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(url)

	client := http.Client{
		Timeout: time.Second * 10,
	}

	req := &http.Request{
		Method: "GET",
		URL:    url,
	}

	res, err := client.Do(req)
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
