package main

import (
	"fmt"
	"log"
	"math/rand"
	"os" //"github.com/jteeuwen/go-bindata"
	"runtime"
	"strings"
	"time"

	"github.com/getlantern/systray"
	"github.com/ludenus/wallpaper/config"
	"github.com/ludenus/wallpaper/unsplash"
	"github.com/ludenus/wallpaper/util"
	"github.com/ludenus/wallpaper/version"
	"github.com/reujab/wallpaper"
	"github.com/skratchdot/open-golang/open"
)

var (
	unsplashPhotoUrl  = "https://unsplash.com"
	unsplashAuthorUrl = "https://unsplash.com"

	current_wallpaper *unsplash.Image

	mBack        *systray.MenuItem
	mPhoto       *systray.MenuItem
	mAuthor      *systray.MenuItem
	mRandomImage *systray.MenuItem
	mAbout       *systray.MenuItem
	mQuit        *systray.MenuItem

	conf           *config.Config
	unsplashClient *unsplash.UClient
)

func github_url() string {
	githash := version.Parse(asset_bytes("appdir/version.json")).Commit
	if "unknown" == githash {
		return "https://github.com/ludenus/wallpaper"
	} else {
		return "https://github.com/ludenus/wallpaper/tree/" + githash
	}
}

func version_string() string {
	ver := version.Parse(asset_bytes("appdir/version.json"))
	return fmt.Sprintf("(%s:%s)", ver.Branch, ver.Commit[0:8])
}

func init_menu() {
	mPhoto = systray.AddMenuItem("Photo/Unsplash", "Open photo on Unsplash")
	mAuthor = systray.AddMenuItem("Author/Unsplash", "Author profile on Unsplash")
	mRandomImage = systray.AddMenuItem("Random Wallpaper", "Set random wallpaper")
	mBack = systray.AddMenuItem("Previous Wallpaper", "Previous image")
	mAbout = systray.AddMenuItem("About "+version_string(), "Github Link")
	systray.AddSeparator()
	mQuit = systray.AddMenuItem("Quit", "Quit")
	systray.SetIcon(asset_bytes("appdir/resources/tray_BW.png"))
}

func main() {
	systray.Run(onReady, onExit)
}

// TODO reorganize project into packages properly

func onReady() {

	log.Println("github_url():" + github_url())
	save_resources()
	conf_default := config.Load(app_dir()+"/config_default.json", &config.Config{})
	conf = config.Load(app_dir()+"/config.json", conf_default)
	unsplashClient = unsplash.NewClient(conf.UnsplashApiKey, conf.HttpTimeoutSeconds)
	init_menu()
	switch_wallpaper()
	go refresh_collection()
	loop_menu()
}

func onExit() {
	fmt.Println("on Exit")
}

func loop_menu() {

	wallpaper_tick := time.Tick(time.Duration(conf.SwitchWallpaperIntervalSeconds) * time.Second)
	collection_tick := time.Tick(time.Duration(conf.RefreshCollectionIntervalSeconds) * time.Second)
	for {
		select {
		case <-wallpaper_tick:
			switch_wallpaper()
		case <-collection_tick:
			go refresh_collection()
		case <-mAuthor.ClickedCh:
			open.Run(unsplashAuthorUrl)
		case <-mPhoto.ClickedCh:
			open.Run(unsplashPhotoUrl)
		case <-mRandomImage.ClickedCh:
			switch_wallpaper()
		case <-mBack.ClickedCh:
			set_previous_switch_wallpaper()
		case <-mAbout.ClickedCh:
			open.Run(github_url())
		case <-mQuit.ClickedCh:
			config.Save(conf, app_dir()+"/config.json")
			systray.Quit()
			os.Exit(0)
		}
	}
}

func switch_wallpaper() {
	set_random_switch_wallpaper()
	set_author_info(current_wallpaper)
	add_to_history(current_wallpaper)
}

func refresh_collection() {
	fetch_unsplash_data_from_collection()
	process_fetched_unsplash_data()
}

func add_to_history(image *unsplash.Image) {
	ids := load_history(app_dir() + "/history.json")
	ids = append(ids, image.ID)
	// limit number of history records
	if len(ids) > conf.HistoryLimit {
		ids = ids[len(ids)-conf.HistoryLimit:]
	}
	save_history(ids, app_dir()+"/history.json")
}

func set_previous_switch_wallpaper() {
	ids := load_history(app_dir() + "/history.json")

	if len(ids) < 2 {
		log.Println("not enough items in history to set previous wallpaper")
		return
	} else {
		// remove last item from history
		ids = ids[:len(ids)-1]
	}

	previous := ids[len(ids)-1]
	set_wallpaper_by_id(previous)
	set_author_info(current_wallpaper)

	// limit number of history records
	if len(ids) > conf.HistoryLimit {
		ids = ids[len(ids)-conf.HistoryLimit:]
	}

	// owerwrite history
	save_history(ids, app_dir()+"/history.json")
}

func fetch_unsplash_data_from_collection() {
	images, err := unsplashClient.GetPhotosByUserCollection(conf.User, conf.Collection)
	if err != nil {
		log.Println(fmt.Sprintf("GetPhotosByUserCollection failed %s", err))
		return
	}
	unsplash.Persist_images(images, images_dir(), unsplashClient)
}

func process_fetched_unsplash_data() {
	images := unsplash.Parse_images_from_dir(unsplash_dir())
	unsplash.Persist_images(images, images_dir(), unsplashClient)
}

func set_random_switch_wallpaper() *unsplash.Image {
	image := random_image_local()
	return set_wallpaper_by_id(image.ID)
}

func set_wallpaper_by_id(id string) *unsplash.Image {
	json := util.Read_bytes_from(images_dir() + "/" + id + ".json")
	image := unsplash.ParseImage(json)
	set_wallpaper(images_dir() + "/" + id + ".jpg")
	current_wallpaper = image
	return image
}

func random_image_local() *unsplash.Image {

	images := util.List_dir(images_dir() + "/*.json")

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(images))

	json := util.Read_bytes_from(images[r])
	image := unsplash.ParseImage(json)
	jpg_filename := strings.Replace(images[r], "json", "jpg", 1)
	image.Bytes = util.Read_bytes_from(jpg_filename)

	return image
}

func set_wallpaper(filename string) {

	switch os := runtime.GOOS; os {
	case "darwin":
		util.Exec_command("osascript", []string{"-e", "tell application \"System Events\" to tell every desktop to set picture to \"" + filename + "\""})

	case "linux":
		util.Exec_command("/usr/bin/gsettings", []string{"set", "org.gnome.desktop.background", "picture-uri", "file://" + filename})
		util.Exec_command("/usr/bin/gsettings", []string{"set", "org.gnome.desktop.background", "picture-options", "stretched"})
		//gsettings set org.gnome.desktop.background picture-options 'stretched'
		//gsettings set org.gnome.desktop.background picture-options 'none'
		//gsettings set org.gnome.desktop.background picture-options 'wallpaper'
		//gsettings set org.gnome.desktop.background picture-options 'centered'
		//gsettings set org.gnome.desktop.background picture-options 'scaled'
		//gsettings set org.gnome.desktop.background picture-options 'zoom'
		//gsettings set org.gnome.desktop.background picture-options 'spanned'
	case "windows":
		// https://github.com/reujab/wallpaper/
		// wallpaper.SetFromURL("https://i.imgur.com/pIwrYeM.jpg")
		wallpaper.SetFromFile(filename)
	default:
		panic("do not know how to set wallpaper for os: " + os)
	}

}

func set_author_info(image *unsplash.Image) {
	mPhoto.SetTitle("Photo / Unsplash")
	mAuthor.SetTitle(image.User.Name + " / Unsplash")
	// according to https://medium.com/unsplash/unsplash-api-guidelines-attribution-4d433941d777
	unsplashPhotoUrl = image.Links.HTML + "?utm_source=desktop_slideshow&utm_medium=referral"
	unsplashAuthorUrl = image.User.Links.HTML + "?utm_source=desktop_slideshow&utm_medium=referral"
}
