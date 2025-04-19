package furaffinity

type FATab string

const (
	TabHome        FATab = "home"
	TabGallery     FATab = "gallery"
	TabScraps      FATab = "scraps"
	TabFavorites   FATab = "favorites"
	TabJournals    FATab = "journals"
	TabCommissions FATab = "commissions"
)

type FurAffinityUserParams struct {
	User string `json:"user"`
	Tab  FATab  `json:"tab"`
}
