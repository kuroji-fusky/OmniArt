package furaffinity

type FurAffinityUserParams struct {
	User string `json:"user"`
	Tab  string `json:"tab"`
}

const (
	TabHome        = "home"
	TabGallery     = "gallery"
	TabScraps      = "scraps"
	TabFavorites   = "favorites"
	TabJournals    = "journals"
	TabCommissions = "commissions"
)
