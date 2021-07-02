package unsplash

type Collection struct {
	CoverPhoto struct {
		Categories  []interface{} `json:"categories"`
		Color       string        `json:"color"`
		CreatedAt   string        `json:"created_at"`
		Description interface{}   `json:"description"`
		Height      int           `json:"height"`
		ID          string        `json:"id"`
		LikedByUser bool          `json:"liked_by_user"`
		Likes       int           `json:"likes"`
		Links       struct {
			Download         string `json:"download"`
			DownloadLocation string `json:"download_location"`
			HTML             string `json:"html"`
			Self             string `json:"self"`
		} `json:"links"`
		UpdatedAt string `json:"updated_at"`
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
			Location     string      `json:"location"`
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
		Width int `json:"width"`
	} `json:"cover_photo"`
	Curated     bool        `json:"curated"`
	Description interface{} `json:"description"`
	Featured    bool        `json:"featured"`
	ID          string         `json:"id"`
	Links       struct {
		HTML    string `json:"html"`
		Photos  string `json:"photos"`
		Related string `json:"related"`
		Self    string `json:"self"`
	} `json:"links"`
	Private     bool   `json:"private"`
	PublishedAt string `json:"published_at"`
	ShareKey    string `json:"share_key"`
	Title       string `json:"title"`
	TotalPhotos int    `json:"total_photos"`
	UpdatedAt   string `json:"updated_at"`
	User        struct {
		Bio            string `json:"bio"`
		FirstName      string `json:"first_name"`
		FollowedByUser bool   `json:"followed_by_user"`
		ID             string `json:"id"`
		LastName       string `json:"last_name"`
		Links          struct {
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
}
