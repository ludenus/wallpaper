package version

import (
	"encoding/json"
	"log"
)

type Version struct {
	Branch string `json:"branch"`
	Commit string `json:"commit"`
}

func default_version() *Version{
	return &Version{
		Branch: "unknown",
		Commit: "unknown",
	}
}

func Parse(bytes []byte) *Version{

	ver := default_version()
	err := json.Unmarshal(bytes, &ver)
	if err != nil {
		log.Println(err)
	}

	log.Println(ver)
	return ver
}