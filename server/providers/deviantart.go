/*
Handler for DeviantArt
*/
package providers

import (
	"github.com/kuroji-fusky/OmniArt/server/internal/utils"
	"github.com/kuroji-fusky/OmniArt/server/providers/deviantart"
)

const (
	DA_URL = "https://www.deviantart.com"
)

func DeviantArtQuery(params *deviantart.DASearchQuery) ([]map[string]interface{}, error) {
	validQueryValues := map[string]bool{
		deviantart.QueryOrderWatch:      true,
		deviantart.QueryOrderMostRecent: true,
		deviantart.QueryOrderWeek:       true,
		deviantart.QueryOrderMonth:      true,
		deviantart.QueryOrderYear:       true,
	}

	if queryOrderOK, queryOrderErr := utils.CheckValidMap(validQueryValues, params.Order, ""); !queryOrderOK {
		return nil, queryOrderErr
	}

	tmpData := []map[string]interface{}{{
		"message": "some data here",
	}}

	return tmpData, nil
}
