package furrynetwork

type StatsAPIResponse struct {
	Favorites StatFavorites `json:"favorites"`
	Followers StatFollowers `json:"followers"`
	Following StatFollowing `json:"following"`
	Promotes  StatPromotes  `json:"promotes"`
}

type CollatedTotal struct {
	Total int `json:"total"`
}

type (
	StatFavorites struct {
		CollatedTotal
	}

	StatFollowers struct {
		Characters []Characters `json:"characters"`
		CollatedTotal
	}

	StatFollowing struct {
		Characters []Characters `json:"characters"`
		CollatedTotal
	}

	StatPromotes struct {
		CollatedTotal
	}
)

type ProfileStats struct {
	Artwork    int `json:"artwork"`
	Photos     int `json:"photos"`
	Stories    int `json:"stories"`
	Multimedia int `json:"multimedia"`
	Journals   int `json:"journals"`
	Favorites  int `json:"favorites"`
	Followers  int `json:"followers"`
	Following  int `json:"following"`
	Promotes   int `json:"promotes"`
}
