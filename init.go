package main

import (
	"os/user"
	"strings"
	"github.com/ludenus/wallpaper/util"
)

func user_dir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

func app_dir() string {
	return util.Mkdir(user_dir() + "/.wallpaper")
}

func resource_dir() string {
	return util.Mkdir(app_dir() + "/resources")
}

func images_dir() string {
	return util.Mkdir(app_dir() + "/images")
}

func unsplash_dir() string {
	return util.Mkdir(app_dir() + "/unsplash")
}

func save_resources() {
	app_dir()
	resource_dir()
	images_dir()
	unsplash_dir()

	save_asset_to_app_dir("appdir/config_default.json")

	save_asset_to_app_dir("appdir/version.json")

	save_asset_to_app_dir("appdir/images/VB5tltMV-U8.jpg")
	save_asset_to_app_dir("appdir/images/VB5tltMV-U8.json")
}

func save_asset_to_app_dir(asset string) {
	filename := strings.Replace(asset, "appdir", app_dir(), 1)
	util.Save_bytes_as(asset_bytes(asset), filename)
}

func asset_bytes(name string) []byte {
	bytes, err := Asset(name)
	if err != nil {
		panic(err)
	}
	return bytes
}