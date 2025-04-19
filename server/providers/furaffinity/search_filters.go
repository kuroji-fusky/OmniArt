package furaffinity

import "time"

type (
	FARange          string
	FAQueryMode      string
	FAOrderBy        string
	FAOrderDirection string
)

const (
	Range1Day   FARange = "1day"
	Range3Days  FARange = "3days"
	Range7Days  FARange = "7days"
	Range30Days FARange = "30days"
	Range90Days FARange = "90days"
	Range1Year  FARange = "1year"
	Range3Years FARange = "3years"
	Range5Years FARange = "5years"
	RangeAll    FARange = "all"
	RangeManual FARange = "manual"
)

const (
	QueryModeAll      FAQueryMode = "all"
	QueryModeAny      FAQueryMode = "any"
	QueryModeExtended FAQueryMode = "extended"
)

const (
	OrderByPopularity FAOrderBy = "popularity"
	OrderByDate       FAOrderBy = "date"
	OrderByRelevancy  FAOrderBy = "relevancy"
)

const (
	OrderDirectionAsc  FAOrderDirection = "asc"
	OrderDirectionDesc FAOrderDirection = "desc"
)

type FurAffinityQueryParams struct {
	Page          *int             `json:"page,omitempty"`
	OrderBy       FAOrderBy        `json:"order-by,omitempty"`
	OrderDirecton FAOrderDirection `json:"order-direction,omitempty"`
	Query         string           `json:"query" query:"q"`

	// Range

	Range     FARange   `json:"range" query:"range"`
	RangeFrom time.Time `json:"rangeFrom,omitempty" query:"range_from"`
	RageTo    time.Time `json:"rangeTo,omitempty" query:"range_to"`

	// Ratings

	RatingGeneral *bool `json:"rating-general,omitempty"`
	RatingMature  *bool `json:"rating-mature,omitempty"`
	RatingAdult   *bool `json:"rating-adult,omitempty"`

	// Submission types

	TypeArt    *bool `json:"type-art,omitempty"`
	TypeMusic  *bool `json:"type-music,omitempty"`
	TypePhotos *bool `json:"type-photos,omitempty"`
	TypeFlash  *bool `json:"type-flash,omitempty"`
	TypeStory  *bool `json:"type-story,omitempty"`

	// matching keywords from its query param

	QueryMode FAQueryMode `json:"mode"` // "all" | "any" | "extended"
}
