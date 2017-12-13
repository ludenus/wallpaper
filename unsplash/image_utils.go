package unsplash

import (
	"log"
	"encoding/json"
	"sort"
	"github.com/ludenus/wallpaper/util"
	"fmt"
)

func Persist_images(images []Image, dir string, client *UClient) {
	for _, image := range images {
		Persist_image(&image, dir, client)
	}
}

func Persist_image(image *Image, dir string, client *UClient) (error) {

	err:= Persist_image_jpg(image,dir,client)
	if err != nil {
		log.Println(fmt.Sprintf("Persist_image_jpg failed %s", err))
		return err
	}

	err = Persist_image_json(image, dir)
	if err != nil {
		log.Println(fmt.Sprintf("Persist_image_json failed %s", err))
		return err
	}
	return nil
}

func Persist_image_jpg(image *Image, dir string, client *UClient) (error) {
	jpg_file := dir + "/" + image.ID + ".jpg"
	if util.File_exists(jpg_file) {
		//log.Println("jpeg exists: " + jpg_filename)
	} else {
		if nil == image.Bytes {
			// according to https://medium.com/unsplash/unsplash-api-guidelines-triggering-a-download-c39b24e99e02
			bytes, err := client.DownloadImage(image)
			if err != nil {
				log.Println(fmt.Sprintf("DownloadImage failed %s", err))
				return err
			}
			image.Bytes = bytes
		}
		util.Save_bytes_as(image.Bytes, jpg_file)
	}
	return nil
}

func Persist_image_json(image *Image, dir string) (error) {
	json_file := dir + "/" + image.ID + ".json"
	err := Save_image_json_as(image, json_file)
	if err != nil {
		log.Println(fmt.Sprintf("Save_image_json_as failed %s", err))
		return err
	}
	return nil
}


func Save_image_json_as(image *Image, json_filename string) (error) {
	if nil == image.Json {
		bytes, err := json.Marshal(image)
		if err != nil {
			log.Println(fmt.Sprintf("json.Marshal failed %s", err))
			return err
		}
		image.Json = bytes
	}
	util.Save_bytes_as(image.Json, json_filename)
	return nil
}

func Parse_images_from_dir(dir string) []Image {
	ls := util.List_dir(dir + "/*.json")
	images := make([]Image, 0)
	for _, f := range ls {
		json := util.Read_bytes_from(f)
		parsed, err := ParseImages(json)
		if err != nil {
			log.Println(fmt.Sprintf("ParseImages failed %s", err))
			log.Println("failed to parse json:" + string(json))
		} else {
			images = append(images, parsed...)
		}
	}

	sort.Sort(sort.Reverse(byIDs(images)))
	for i := 0; i < len(images)-1; i++ {
		// duplicates are neighbours when sorted by id
		if images[i].ID == images[i+1].ID {
			// remove duplicate
			images = append(images[:i], images[i+1:]...)
		}
	}

	return images
}


func ParseCollections(response []byte) ([]Collection, error) {

	collections := make([]Collection, 0)

	err := json.Unmarshal(response, &collections)
	if err != nil {
		log.Println(fmt.Sprintf("json.Unmarshal failed %s", err))
		return nil, err
	}

	return collections, nil
}

func ParseImage(response []byte) *Image {
	var image Image
	err := json.Unmarshal(response, &image)
	if err != nil {
		log.Println(fmt.Sprintf("json.Unmarshal failed %s", err))
		panic(err)
	}

	return &image
}

func ParseImages(response []byte) ([]Image, error) {

	images := make([]Image, 0)

	err := json.Unmarshal(response, &images)
	if err != nil {
		log.Println(fmt.Sprintf("json.Unmarshal failed %s", err))
		return nil, err
	}

	return images, nil
}

// Returns a new slice containing all items in the slice that satisfy the predicate f.
func FindAll(images []Image, f func(Image) bool) []Image {
	filtered := make([]Image, 0)
	for _, it := range images {
		if f(it) {
			filtered = append(filtered, it)
		}
	}
	return filtered
}

type byDownloads []Image

func (a byDownloads) Len() int      { return len(a) }
func (a byDownloads) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byDownloads) Less(i, j int) bool {
	if a[i].Downloads < a[j].Downloads {
		return true
	}
	if a[i].Downloads > a[j].Downloads {
		return false
	}
	return a[i].Downloads < a[j].Downloads
}

type byLikes []Image

func (a byLikes) Len() int      { return len(a) }
func (a byLikes) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byLikes) Less(i, j int) bool {
	if a[i].Likes < a[j].Likes {
		return true
	}
	if a[i].Likes > a[j].Likes {
		return false
	}
	return a[i].Likes < a[j].Likes
}

type byIDs []Image

func (a byIDs) Len() int      { return len(a) }
func (a byIDs) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byIDs) Less(i, j int) bool {
	if a[i].ID < a[j].ID {
		return true
	}
	if a[i].ID > a[j].ID {
		return false
	}
	return a[i].ID < a[j].ID
}

//sort.Sort(byDownloads(images))
