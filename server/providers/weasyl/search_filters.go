package weasyl

type WeasylSearchFilters struct {
	Query       string   `json:"query" query:"q"`
	FindType    string   `json:"find_type" query:"find"`
	Category    int      `json:"category,omitempty" query:"cat"`
	Subcategory int      `json:"subcategory,omitempty" query:"subcat"`
	Within      string   `json:"within" query:"within"`
	Rating      []string `json:"rating" query:"rating"`
}

// Subcategory filters

const (
	SubcategorySketch            = 1010
	SubcategoryTraditional       = 1020
	SubcategoryDigital           = 1030
	SubcategoryAnimation         = 1040
	SubcategoryPhotography       = 1050
	SubcategoryDesignInterface   = 1060
	SubcategoryModelingSculpture = 1070
	SubcategoryCraftsJewelry     = 1075
	SubcategorySewingKnitting    = 1078
	SubcategoryDesktopWallpaper  = 1080
	SubcategoryStory             = 2010 // cat 2000
	SubcategoryPoetryLyrics      = 2020
	SubcategoryScriptScreenplay  = 2030
	SubcategoryOriginalMusic     = 3010 // cat 3000
	SubcategoryCoverVersion      = 3020
	SubcategoryRemixMashup       = 3030
	SubcategorySpeechReading     = 3040
	SubcategoryEmbeddedVideo     = 3500
)

// Category filters

const (
	CategoryVisual     = 1000
	CategoryLiterary   = 2000
	CategoryMultimedia = 3000
)

const (
	RatingGeneral  = "g"
	RatingMature   = "a"
	RatingExplicit = "p"
)
