/*
Handler for DeviantArt
*/
package providers

import (
	"github.com/kuroji-fusky/OmniArt/server/internal/utils"
	"github.com/kuroji-fusky/OmniArt/server/providers/deviantart"
)

const DA_URL = "https://www.deviantart.com"

func DeviantArtSearchQuery(params *deviantart.DASearchQuery) ([]utils.Implementation_WIP, error) {
	validQueryValues := map[deviantart.DAQueryOrder]bool{
		deviantart.QueryOrderWatch:      true,
		deviantart.QueryOrderMostRecent: true,
		deviantart.QueryOrderWeek:       true,
		deviantart.QueryOrderMonth:      true,
		deviantart.QueryOrderYear:       true,
	}

	if queryOrderOK, queryOrderErr := utils.CheckValidMap(validQueryValues, params.Order, ""); !queryOrderOK {
		return nil, queryOrderErr
	}

	return utils.TempData, nil
}
