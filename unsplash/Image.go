package unsplash


type Image struct {

	Bytes  []byte `json:"-"`
	Json   []byte `json:"-"`

	// generated via http://json2struct.mervine.net/
	Categories             []interface{} `json:"categories"`
	Color                  string        `json:"color"`
	CreatedAt              string        `json:"created_at"`
	CurrentUserCollections []interface{} `json:"current_user_collections"`
	Description            interface{}   `json:"description"`
	Downloads              int           `json:"downloads"`
	Exif                   struct {
		Aperture     string `json:"aperture"`
		ExposureTime string `json:"exposure_time"`
		FocalLength  string `json:"focal_length"`
		Iso          int    `json:"iso"`
		Make         string `json:"make"`
		Model        string `json:"model"`
	} `json:"exif"`
	Height      int    `json:"height"`
	ID          string `json:"id"`
	LikedByUser bool   `json:"liked_by_user"`
	Likes       int    `json:"likes"`
	Links       struct {
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
		HTML             string `json:"html"`
		Self             string `json:"self"`
	} `json:"links"`
	Location struct {
		City     string `json:"city"`
		Country  string `json:"country"`
		Name     string `json:"name"`
		Position struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"position"`
		Title string `json:"title"`
	} `json:"location"`
	Slug      interface{} `json:"slug"`
	UpdatedAt string      `json:"updated_at"`
	Urls      struct {
		Full    string `json:"full"`
		Raw     string `json:"raw"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	User struct {
		Bio       string `json:"bio"`
		FirstName string `json:"first_name"`
		ID        string `json:"id"`
		LastName  string `json:"last_name"`
		Links     struct {
			Followers string `json:"followers"`
			Following string `json:"following"`
			HTML      string `json:"html"`
			Likes     string `json:"likes"`
			Photos    string `json:"photos"`
			Portfolio string `json:"portfolio"`
			Self      string `json:"self"`
		} `json:"links"`
		Location     interface{} `json:"location"`
		Name         string      `json:"name"`
		PortfolioURL interface{} `json:"portfolio_url"`
		ProfileImage struct {
			Large  string `json:"large"`
			Medium string `json:"medium"`
			Small  string `json:"small"`
		} `json:"profile_image"`
		TotalCollections int         `json:"total_collections"`
		TotalLikes       int         `json:"total_likes"`
		TotalPhotos      int         `json:"total_photos"`
		TwitterUsername  interface{} `json:"twitter_username"`
		UpdatedAt        string      `json:"updated_at"`
		Username         string      `json:"username"`
	} `json:"user"`
	Views int `json:"views"`
	Width int `json:"width"`
}
