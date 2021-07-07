#!/bin/bash -ex

export GOOS=darwin

echo "refresh go-bindata dependencies..."
go get -u github.com/jteeuwen/go-bindata/...
go build -o go-bindata github.com/jteeuwen/go-bindata/go-bindata

echo "prepare default config ..."
echo "{" > appdir/config_default.json
echo "\"user\":\"ludenus\"," >> appdir/config_default.json
echo "\"collection\":\"wallpaper\"," >> appdir/config_default.json
echo "\"switch_wallpaper_interval_seconds\":300," >> appdir/config_default.json
echo "\"refresh_collection_interval_seconds\":3600," >> appdir/config_default.json
echo "\"history_limit\":100," >> appdir/config_default.json
echo "\"http_timeout_seconds\":30," >> appdir/config_default.json
echo "\"unsplash_api_key\":\"${UNSPLASH_APPLICATION_ID}\"" >> appdir/config_default.json
echo "}" >> appdir/config_default.json

echo "report version..."

echo "{" > appdir/version.json
echo "\"branch\":\"`git rev-parse --abbrev-ref HEAD`\"," >> appdir/version.json
echo "\"commit\":\"`git rev-parse HEAD`\"" >> appdir/version.json
echo "}" >> appdir/version.json

echo "rebuild resources..."
./go-bindata appdir appdir/images appdir/resources

#echo "go get dependencies..."
#sudo go get all

echo "build..."
go build -v github.com/ludenus/wallpaper

echo "install into $GOPATH/bin/"
go install -v github.com/ludenus/wallpaper

echo "should now be able to run $GOPATH/bin/wallpaper"
