package weasyl

// Subcategory filters

type (
	WeasylCategory    int
	WeasylSubcategory int
	WeasylRating      string
)

const (
	SubcategorySketch            WeasylSubcategory = 1010
	SubcategoryTraditional       WeasylSubcategory = 1020
	SubcategoryDigital           WeasylSubcategory = 1030
	SubcategoryAnimation         WeasylSubcategory = 1040
	SubcategoryPhotography       WeasylSubcategory = 1050
	SubcategoryDesignInterface   WeasylSubcategory = 1060
	SubcategoryModelingSculpture WeasylSubcategory = 1070
	SubcategoryCraftsJewelry     WeasylSubcategory = 1075
	SubcategorySewingKnitting    WeasylSubcategory = 1078
	SubcategoryDesktopWallpaper  WeasylSubcategory = 1080
	SubcategoryStory             WeasylSubcategory = 2010 // cat 2000
	SubcategoryPoetryLyrics      WeasylSubcategory = 2020
	SubcategoryScriptScreenplay  WeasylSubcategory = 2030
	SubcategoryOriginalMusic     WeasylSubcategory = 3010 // cat 3000
	SubcategoryCoverVersion      WeasylSubcategory = 3020
	SubcategoryRemixMashup       WeasylSubcategory = 3030
	SubcategorySpeechReading     WeasylSubcategory = 3040
	SubcategoryEmbeddedVideo     WeasylSubcategory = 3500
)

// Category filters
const (
	CategoryVisual     WeasylCategory = 1000
	CategoryLiterary   WeasylCategory = 2000
	CategoryMultimedia WeasylCategory = 3000
)

// Ratings
const (
	RatingGeneral  WeasylRating = "g"
	RatingMature   WeasylRating = "a"
	RatingExplicit WeasylRating = "p"
)

type WeasylSearchFilters struct {
	Query       string             `json:"query" query:"q"`
	FindType    string             `json:"find_type" query:"find"`
	Category    *WeasylCategory    `json:"category,omitempty" query:"cat"`
	Subcategory *WeasylSubcategory `json:"subcategory,omitempty" query:"subcat"`
	Within      string             `json:"within" query:"within"`
	Rating      []WeasylRating     `json:"rating" query:"rating"`
}
