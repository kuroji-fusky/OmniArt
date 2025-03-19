package furaffinity

import "time"

type FurAffinityQueryParams struct {
	Page          int    `json:"page,omitempty"`
	OrderBy       string `json:"order-by,omitempty"`        // "popularity" | "date" | "relevancy"
	OrderDirecton string `json:"order-direction,omitempty"` // "asc" | "desc"
	Query         string `json:"query" query:"q"`

	// Range

	Range     string    `json:"range" query:"range"` // "1day" | "3days" | "7days" | "30days" | "90days" | "1year" | "3years" | "5years" | "all" | "manual"
	RangeFrom time.Time `json:"rangeFrom,omitempty" query:"range_from"`
	RageTo    time.Time `json:"rangeTo,omitempty" query:"range_to"`

	// Ratings

	RatingGeneral bool `json:"rating-general,omitempty"`
	RatingMature  bool `json:"rating-mature,omitempty"`
	RatingAdult   bool `json:"rating-adult,omitempty"`

	// Submission types

	TypeArt    bool `json:"type-art,omitempty"`
	TypeMusic  bool `json:"type-music,omitempty"`
	TypePhotos bool `json:"type-photos,omitempty"`
	TypeFlash  bool `json:"type-flash,omitempty"`
	TypeStory  bool `json:"type-story,omitempty"`

	// matching keywords from its query param

	QueryMode string `json:"mode"` // "all" | "any" | "extended"
}

const (
	Range1Day   = "1day"
	Range3Days  = "3days"
	Range7Days  = "7days"
	Range30Days = "30days"
	Range90Days = "90days"
	Range1Year  = "1year"
	Range3Years = "3years"
	Range5Years = "5years"
	RangeAll    = "all"
	RangeManual = "manual"
)

const (
	QueryModeAll      = "all"
	QueryModeAny      = "any"
	QueryModeExtended = "extended"
)

const (
	OrderByPopularity = "popularity"
	OrderByDate       = "date"
	OrderByRelevancy  = "relevancy"
)

const (
	OrderDirectionAsc  = "asc"
	OrderDirectionDesc = "desc"
)
