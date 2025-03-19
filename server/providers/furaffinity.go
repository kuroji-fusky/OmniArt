/*
Handler for Fur Affinity

Domains:
- furaffinity.net
- sfw.furaffinity.net
- d.furaffinity.net
*/
package providers

import (
	"log"

	"google.golang.org/genproto/googleapis/type/date"
)

type FAUserParams struct {
	User string
	Tab  string // "home" | "gallery" | "scraps" | "favorites" | "journals" | "commissions"
}

func FurAffinityUser(params *FAUserParams) (map[string]any, error) {
	validFATabs := map[string]bool{
		"home":        true,
		"gallery":     true,
		"scraps":      true,
		"favorites":   true,
		"journals":    true,
		"commissions": true,
	}

	if !validFATabs[params.Tab] {
		log.Fatalf("Tab %q is not valid", params.Tab)
	}

	// Implementation WIP
	tmpData := map[string]any{}

	return tmpData, nil
}

type FASearchQueryParams struct {
	Page          int    `json:"page,omitempty"`
	OrderBy       string `json:"order-by,omitempty"`        //"popularity" | "date" | "relevancy"
	OrderDirecton string `json:"order-direction,omitempty"` // "asc" | "desc"
	Query         string `json:"q"`

	// Range

	Range     string    `json:"range"` // "1day" | "3days" | "7days" | "30days" | "90days" | "1year" | "3years" | "5years" | "all" | "manual"
	RangeFrom date.Date `json:"range_from,omitempty"`
	RageTo    date.Date `json:"range_to,omitempty"`

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

func FurAffinityQuery(params *FASearchQueryParams) (map[string]any, error) {
	// Implmentation WIP
	// Bools from FASearchQueryParams struct are interpreted as literal `"1"` string from the API

	tmpData := map[string]any{}

	return tmpData, nil
}
