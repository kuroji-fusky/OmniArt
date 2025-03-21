/*
Handler for Fur Affinity
*/
package providers

import (
	"github.com/kuroji-fusky/OmniArt/server/internal/utils"
	"github.com/kuroji-fusky/OmniArt/server/providers/furaffinity"
)

// Domains
const (
	FA_URL       = "furaffinity.net"
	FA_SFW_URL   = "sfw.furaffinity.net"
	FA_IMAGE_URL = "d.furaffinity.net"
)

func FurAffinityUser(params furaffinity.FurAffinityUserParams) ([]map[string]interface{}, error) {
	validFATabs := map[string]bool{
		furaffinity.TabHome:        true,
		furaffinity.TabGallery:     true,
		furaffinity.TabScraps:      true,
		furaffinity.TabFavorites:   true,
		furaffinity.TabJournals:    true,
		furaffinity.TabCommissions: true,
	}

	if ok, err := utils.CheckValidMap(validFATabs, params.Tab, ""); !ok {
		return nil, err
	}

	// Implementation WIP
	tmpData := []map[string]interface{}{{
		"message": "some data here",
	}}

	return tmpData, nil
}

func FurAffinityQuery(params furaffinity.FurAffinityQueryParams) ([]map[string]interface{}, error) {
	// Implmentation WIP
	// Bools from FurAffinityQueryParams struct are interpreted as literal `"1"` string from the API

	// "range" param
	validQueryRange := map[string]bool{
		furaffinity.Range1Day:   true,
		furaffinity.Range3Days:  true,
		furaffinity.Range7Days:  true,
		furaffinity.Range30Days: true,
		furaffinity.Range90Days: true,
		furaffinity.Range1Year:  true,
		furaffinity.Range3Years: true,
		furaffinity.Range5Years: true,
		furaffinity.RangeAll:    true,
		furaffinity.RangeManual: true,
	}

	if queryRangeOK, queryRangeErr := utils.CheckValidMap(validQueryRange, params.Range, ""); !queryRangeOK {
		return nil, queryRangeErr
	}

	// "order-by" query
	validQueryOrderBy := map[string]bool{
		furaffinity.OrderByPopularity: true,
		furaffinity.OrderByDate:       true,
		furaffinity.OrderByRelevancy:  true,
	}

	if orderByOK, queryOrderByErr := utils.CheckValidMap(validQueryOrderBy, params.OrderBy, ""); !orderByOK {
		return nil, queryOrderByErr
	}

	// "order-direction" query
	validQueryOrderDirection := map[string]bool{
		furaffinity.OrderDirectionAsc:  true,
		furaffinity.OrderDirectionDesc: true,
	}

	if orderDirectionOK, orderDirectionErr := utils.CheckValidMap(validQueryOrderDirection, params.OrderDirecton, ""); !orderDirectionOK {
		return nil, orderDirectionErr
	}

	validQueryMode := map[string]bool{
		furaffinity.QueryModeAll:      true,
		furaffinity.QueryModeAny:      true,
		furaffinity.QueryModeExtended: true,
	}

	if queryModeOK, queryModeErr := utils.CheckValidMap(validQueryMode, params.QueryMode, ""); !queryModeOK {
		return nil, queryModeErr
	}

	tmpData := []map[string]interface{}{{
		"message": "some data here",
	}}

	return tmpData, nil
}
