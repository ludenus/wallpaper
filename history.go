package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/ludenus/wallpaper/util"
)


func load_history(filename string) []string {
	ids := make([]string, 0)

	text, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("cannot read history file")
	}else{
		err = json.Unmarshal(text, &ids)
		if err != nil {
			panic(err)
		}
	}
	return ids
}

func save_history(ids []string, filename string) {
	bytes,err := json.Marshal(ids)
	if err != nil {
		panic(err)
	}
	util.Save_bytes_as(bytes, filename)
}
