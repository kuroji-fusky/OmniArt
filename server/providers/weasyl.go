/*
Handler for Weasyl
*/
package providers

import (
	"github.com/kuroji-fusky/OmniArt/server/internal/utils"
	"github.com/kuroji-fusky/OmniArt/server/providers/weasyl"
)

func WeasylSearchQuery(query weasyl.WeasylSearchFilters) ([]utils.Implementation_WIP, error) {
	validSubcategories := map[int]bool{
		weasyl.SubcategorySketch:            true,
		weasyl.SubcategoryTraditional:       true,
		weasyl.SubcategoryDigital:           true,
		weasyl.SubcategoryAnimation:         true,
		weasyl.SubcategoryPhotography:       true,
		weasyl.SubcategoryDesignInterface:   true,
		weasyl.SubcategoryModelingSculpture: true,
		weasyl.SubcategoryCraftsJewelry:     true,
		weasyl.SubcategorySewingKnitting:    true,
		weasyl.SubcategoryDesktopWallpaper:  true,
		weasyl.SubcategoryStory:             true,
		weasyl.SubcategoryPoetryLyrics:      true,
		weasyl.SubcategoryScriptScreenplay:  true,
		weasyl.SubcategoryOriginalMusic:     true,
		weasyl.SubcategoryCoverVersion:      true,
		weasyl.SubcategoryRemixMashup:       true,
		weasyl.SubcategorySpeechReading:     true,
		weasyl.SubcategoryEmbeddedVideo:     true,
	}

	if subcatOK, subcatErr := utils.CheckValidMap(validSubcategories, query.Subcategory, ""); !subcatOK {
		return nil, subcatErr
	}

	// Category filters
	validCategories := map[int]bool{
		weasyl.CategoryVisual:     true,
		weasyl.CategoryLiterary:   true,
		weasyl.CategoryMultimedia: true,
	}

	if catOK, catErr := utils.CheckValidMap(validCategories, query.Category, ""); !catOK {
		return nil, catErr
	}

	// Ratings
	validRatings := map[string]bool{
		weasyl.RatingGeneral:  true,
		weasyl.RatingMature:   true,
		weasyl.RatingExplicit: true,
	}

	ratingErrorMsg := "Rating must be one of the following: g, a, p, but got %s"
	for _, rating := range query.Rating {
		if ratingOK, ratingErr := utils.CheckValidMap(validRatings, rating, ratingErrorMsg); !ratingOK {
			return nil, ratingErr
		}
	}

	return utils.TempData, nil
}
