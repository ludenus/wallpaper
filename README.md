# Wallpaper
Get images from unsplash.com and set them as wallpaper.
Key differences from the [official unsplash app](https://itunes.apple.com/us/app/unsplash-wallpapers/id1284863847) are:
* This application also works on Linux and Windows
* User can customize image collection and wallpaper switch interval

## How to build
### Prerequisites
Go is installed and the following env vars are set
```
   GOROOT=/usr/local/go # may be different, depends on OS
   GOPATH=$HOME/go
   PATH=$PATH:$GOPATH/bin
```
### Linux
build script asks for sudo password to install dependencies
```
$ ./build_lin.sh
```

### MacOsX
build script asks for sudo password to install dependencies
``` 
$ ./build_mac.sh
```

### Windows
```
build_win.sh
```
## How to run
```
$GOPATH/bin/wallpaper
```
## How it works
### Shortly

* Images are downloaded from [unsplash.com](https://unsplash.com) and saved in local directory.
* Random image from the local directory is set as desktop wallpaper every 5 minutes

### Details

* Folder `$HOME/.wallpaper` is created upon start.
```
├── config.json           # config time intervals, unsplash collection
├── history.json          # image ids  
├── images                # images used as desktop wallpaper
│   ├── VB5tltMV-U8.jpg   # image itself
│   └── VB5tltMV-U8.json  # information about image author and direct image url
├── resources
│   ├── notification.svg  # icon for notification
│   ├── tray_BW.png       # icon for system tray
└── unsplash              # temporary folder for fetching images

```

Default `config.json`
```
{
   "user": "ludenus",
   "collection": "wallpaper",
   "history_limit": 100,
   "switch_wallpaper_interval_seconds": 300,
   "refresh_collection_interval_seconds": 3600
}
```
 
1. Unsplash collection queried via [unsplash API](https://community.unsplash.com/developers). By default [this image collection](https://unsplash.com/collections/1263731/wallpaper) is used. Custom unsplash collection can be set by specifying unsplash user and collection name in `config.json`
1. Image information is saved to json files, images are downloaded and saved as jpg files. Once images are downloaded, application should be able to work without internet connection. 1. `VB5tltMV-U8.jpg` image is embedded into application binary and saved into images folder upon start so that application could work completely offline.
1. In parallel, random image from images folder is set as desktop wallpaper.
1. Wallpaper is switched every 5 minutes by default, interval can be configured in `config.json`
1. Unsplash collection data and images is refreshed every hour, interval can be configured in `config.json`. That means once a new image is added to unsplash collection, it appears in local image cache within an hour.
1. User can trigger setting new random wallpaper using `Random Wallpaper` menu item.
1. User can switch to previous wallpaper by clicking `Previous Wallpaper` menu item. By default only 100 latest wallpapers are saved in history, history items limit can be configured in `config.json`

## Versions
Successfully compiles and works under
* Mac OsX Sierra
* MacOsX High Sierra
* Ubuntu 17.04
* Ubuntu 17.10
* Windows 10

```go version go1.9.1``` 

## P.S.
Code is crappy, there are no tests, this is my first golang experience.
