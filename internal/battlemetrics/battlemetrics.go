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
	BattleMetricsBase = "https://api.battlemetrics.com"
)

func GetRustServer(id string) RustServer {
	url, err := url.Parse(fmt.Sprintf("%v/servers/%v", BattleMetricsBase, id))
	if err != nil {
		log.Print(err)
	}
	log.Print(url)

	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Print(err)
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

	url, err := url.Parse(fmt.Sprintf("%v/%v", BattleMetricsBase, "servers"))
	if err != nil {
		log.Print(err)
	}

	log.Print(url)

	queryParams := url.Query()
	queryParams.Set("filter[game]", "rust")
	queryParams.Set("filter[countries][0]", "US")
	queryParams.Set("filter[countries][1]", "CA")
	queryParams.Set("page[size]", "25")
	queryParams.Set("filter[search]", q)
	url.RawQuery = queryParams.Encode()

	fmt.Println(url)

	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Print(err)
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
