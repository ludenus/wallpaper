echo "refresh go-bindata dependencies..."
go get -u github.com/jteeuwen/go-bindata/...
go build -o go-bindata.exe github.com/jteeuwen/go-bindata/go-bindata


echo "prepare config ..."
echo "{" > appdir/config.json
echo "\"user\":\"ludenus\"," >> appdir/config.json
echo "\"collection\":\"wallpaper\"," >> appdir/config.json
echo "\"switch_wallpaper_interval_seconds\":300," >> appdir/config.json
echo "\"refresh_collection_interval_seconds\":3600," >> appdir/config.json
echo "\"history_limit\":100," >> appdir/config.json
echo "\"unsplash_api_key\":\"${UNSPLASH_API_KEY}\"" >> appdir/config.json
echo "}" >> appdir/config.json

echo "report version..."

echo "{" > appdir/version.json
echo "\"branch\":\"`git rev-parse --abbrev-ref HEAD`\"," >> appdir/version.json
echo "\"commit\":\"`git rev-parse HEAD`\"" >> appdir/version.json
echo "}" >> appdir/version.json


echo "rebuild resources..."
./go-bindata.exe appdir/**

echo "go get dependencies..."
go get all

echo "build..."
go build -v github.com/ludenus/wallpaper

echo "install..."
go install github.com/ludenus/wallpaper

echo "should now be able to run wallpaper"
