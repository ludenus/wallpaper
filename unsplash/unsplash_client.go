package unsplash

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"errors"
	"log"
	"encoding/json"
	"time"
)

type UClient struct {
	unsaplshApiKey string
	timeoutSeconds int
}

func NewClient(apiKey string, timeoutSeconds int) *UClient {
	return &UClient{
		unsaplshApiKey: apiKey,
		timeoutSeconds: timeoutSeconds,
	}
}

func (uClient *UClient) get(url string, adjust func(*http.Request) *http.Request) ([]byte, error) {
	// setup httpClient
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr}
	httpClient.Timeout = time.Duration(uClient.timeoutSeconds) * time.Second

	// create request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// adjust request
	request = adjust(request)

	// send request
	response, err := httpClient.Do(request)
	if err != nil {
		log.Println(fmt.Sprintf("failed to send request: %s", err))
		return nil, err
	}

	// read response
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		if err.Error()=="net/http: request canceled (Client.Timeout exceeded while reading body)"{
			return nil,err
		} else{
			panic(err)
		}
	}

	log.Println(fmt.Sprintf("[ %d ] %s", response.StatusCode, url))
	if 200 != response.StatusCode {
		log.Println(string(contents))
		err = errors.New(fmt.Sprintf("received status code [ %d ] for url %s", response.StatusCode, url))
		return nil, err
	}

	return contents, nil
}

func (uClient *UClient) Get(url string) ([]byte, error) {

	var adjust = func(request *http.Request) *http.Request {
		request.Header.Set("Accept-Version", "v1")
		request.Header.Set("Authorization", "uClient-ID "+uClient.unsaplshApiKey)
		return request
	}
	contents, nil := uClient.get(url, adjust)
	return contents, nil
}

func (uClient *UClient) DownloadImage(image *Image) ([]byte, error) {
	// according to https://medium.com/unsplash/unsplash-api-guidelines-triggering-a-download-c39b24e99e02?vero_id=882806
	bytes, err := uClient.Get(image.Links.DownloadLocation)

	if err != nil {
		log.Println(fmt.Sprintf("DownloadLocation failed: %s", err))
		return nil, err
	}

	url := &Url{}
	err = json.Unmarshal(bytes, url)
	if err != nil {
		log.Println(fmt.Sprintf("json.Unmarshal failed: %s", err))
		return nil, err
	}

	//bytes, err := Get(image.Urls.Raw)
	bytes, err = uClient.Get(url.URL)
	if err != nil {
		log.Println(fmt.Sprintf("Get raw failed: %s", err))
		return nil, err
	}
	return bytes, nil
}

func (uClient *UClient) GetUserCollections(username string) ([]byte, error) {
	bytes, err := uClient.Get("https://api.unsplash.com/users/" + username + "/collections")
	if err != nil {
		log.Println(fmt.Sprintf("GetUserCollections failed: %s", err))
		return nil, err
	}
	return bytes, nil
}

func (uClient *UClient) GetPhotosByUserCollection(username string, collection_name string) ([]Image, error) {
	collection, err := uClient.GetUserCollectionByName(username, collection_name)
	if err != nil {
		log.Println(fmt.Sprintf("GetUserCollectionByName failed: %s", err))
		return nil, err
	}
	photos, err := uClient.ParsePhotosFromCollection(collection)
	if err != nil {
		log.Println(fmt.Sprintf("ParsePhotosFromCollection failed %s", err))
		return nil, err
	}
	return photos, nil
}

func (uClient *UClient) GetUserCollectionByName(username string, collection_name string) (*Collection, error) {
	bytes, err := uClient.GetUserCollections(username)
	if err != nil {
		log.Println(fmt.Sprintf("GetUserCollections failed %s", err))
		return nil, err
	}
	collections, err := ParseCollections(bytes)
	if err != nil {
		log.Println(fmt.Sprintf("ParseCollections failed %s", err))
		return nil, err
	}
	var result *Collection
	for _, coll := range collections {
		if coll.Title == collection_name {
			result = &coll
			break
		}
	}
	return result, nil
}

func (uClient *UClient) ParsePhotosFromCollection(collection *Collection) ([]Image, error) {
	images := make([]Image, 0)
	per_page := 30
	number_of_requests := 1 + collection.TotalPhotos/per_page
	var resErr error = nil
	for page := 1; page <= number_of_requests; page++ {
		url := fmt.Sprintf("https://api.unsplash.com/collections/%d/photos?per_page=%d&page=%d", collection.ID, per_page, page)
		bytes, err := uClient.Get(url)
		if err != nil {
			log.Println(fmt.Sprintf("uClient.Get failed %s", err))
			resErr = err
			break
		}
		images_from_page, err := ParseImages(bytes)
		if err != nil {
			log.Println(fmt.Sprintf("ParseImages failed %s", err))
			resErr = err
			break
		}
		images = append(images, images_from_page...)
	}

	return images, resErr
}
