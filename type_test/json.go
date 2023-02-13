package types

import (
	"encoding/json"
	"flag"
	"os"
)

type configs struct {
	Port     int    `json:"port"`
	Url      string `json:"url"`
	Database database
}

type database struct {
	Name       string `json:"name"`
	Collection string `json:"collection"`
}

func loadConfigs() (*configs, error) {
	c := new(configs)
	var configFlag = flag.String("config", "./config.json", "configuration json file path")
	if file, err := os.Open(*configFlag); err != nil {
		return &configs{}, err
	} else {
		defer file.Close()
		if err := json.NewDecoder(file).Decode(c); err != nil {
			return &configs{}, err
		} else {
			// c.sanitize()
			return c, nil
		}
	}
}
