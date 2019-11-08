package gohosts

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Settings struct {
	Hosts  []string `json:"hosts"`
	Output string   `json:"output"`
}

func (sett *Settings) Read() {
	var err error
	fName := "/etc/gohosts.json"
	//read config
	f, err := os.Open(fName)
	if err != nil {
		//create default
		sett.Output = "/tmp/ads_block.list"
		sett.Hosts = []string{"http://winhelp2002.mvps.org/hosts.txt"}
		jsonData, err := json.MarshalIndent(sett, "", "    ")
		if err != nil {
			log.Println(err)
		} else {
			err = ioutil.WriteFile(fName, jsonData, 0755)
			if err != nil {
				log.Println("config : cant write config file", err)
			}
		}

	} else {
		var cfg Settings
		decoder := json.NewDecoder(f)
		err = decoder.Decode(&cfg)
		if err != nil {
		}
		sett.Hosts = cfg.Hosts
		sett.Output = cfg.Output
	}

}
