package config

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/ludenus/wallpaper/util"
)

type Config struct {
	User                      string `json:"user"`
	Collection                string `json:"collection"`
	SwitchWallpaperIntervalSeconds int    `json:"switch_wallpaper_interval_seconds"`
	RefreshCollectionIntervalSeconds int    `json:"refresh_collection_interval_seconds"`
	HistoryLimit int    `json:"history_limit"`
	UnsplashApiKey string    `json:"unsplash_api_key"`
	HttpTimeoutSeconds int `json:"http_timeout_seconds"`
}


func Load(filename string, into *Config) *Config{
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(into)
	cfg := into
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("cannot read config file: " + filename)
	}else{
		err = json.Unmarshal(bytes, cfg)
		if err != nil {
			panic(err)
		}
	}
	log.Println(cfg)
	return cfg
}

func Save(config *Config, filename string) {
	bytes,err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	util.Save_bytes_as(bytes, filename)
}