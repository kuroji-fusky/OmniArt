/*
Handler for Fur Affinity
*/
package providers

import (
	"github.com/kuroji-fusky/OmniArt/server/utils"
	"google.golang.org/genproto/googleapis/type/date"
)

// Domains

const (
	FA_URL       = "furaffinity.net"
	FA_SFW_URL   = "sfw.furaffinity.net"
	FA_IMAGE_URL = "d.furaffinity.net"
)

type FurAffinityUserParams struct {
	User string
	Tab  string // "home" | "gallery" | "scraps" | "favorites" | "journals" | "commissions"
}

func FurAffinityUser(params FurAffinityUserParams) (map[string]any, error) {
	validFATabs := map[string]bool{
		"home":        true,
		"gallery":     true,
		"scraps":      true,
		"favorites":   true,
		"journals":    true,
		"commissions": true,
	}

	if ok, err := utils.CheckValidStringMap(validFATabs, params.Tab, ""); !ok {
		return nil, err
	}

	// Implementation WIP
	tmpData := map[string]any{}

	return tmpData, nil
}

type FurAffinityQueryParams struct {
	Page          int    `json:"page,omitempty"`
	OrderBy       string `json:"order-by,omitempty"`        // "popularity" | "date" | "relevancy"
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

func FurAffinityQuery(params *FurAffinityQueryParams) (map[string]any, error) {
	// Implmentation WIP
	// Bools from FurAffinityQueryParams struct are interpreted as literal `"1"` string from the API

	// "range" param
	validQueryRange := map[string]bool{
		"1day":   true,
		"3days":  true,
		"7days":  true,
		"30days": true,
		"90days": true,
		"1year":  true,
		"3years": true,
		"5years": true,
		"all":    true,
		"manual": true,
	}

	if queryRangeOK, queryRangeErr := utils.CheckValidStringMap(validQueryRange, params.Range, ""); !queryRangeOK {
		return nil, queryRangeErr
	}

	// "order-by" query
	validQueryOrderBy := map[string]bool{
		"popularity": true,
		"date":       true,
		"relevancy":  true,
	}

	if orderByOK, queryOrderByErr := utils.CheckValidStringMap(validQueryOrderBy, params.OrderBy, ""); !orderByOK {
		return nil, queryOrderByErr
	}

	// "order-direction" query
	validQueryOrderDirection := map[string]bool{
		"asc":  true,
		"desc": true,
	}

	if orderDirectionOK, orderDirectionErr := utils.CheckValidStringMap(validQueryOrderDirection, params.OrderDirecton, ""); !orderDirectionOK {
		return nil, orderDirectionErr
	}

	tmpData := map[string]any{}

	return tmpData, nil
}
