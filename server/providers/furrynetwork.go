package providers

import (
	"net/http"

	"github.com/kuroji-fusky/OmniArt/server/internal/utils"
	"github.com/kuroji-fusky/OmniArt/server/providers/furrynetwork"
)

const (
	BASE_API_URL = "https://furrynetwork.com/api"

	CHARACTER_URL = BASE_API_URL + "/character"
	SEARCH_URL    = BASE_API_URL + "/search"
)

func GetCharacterStats(character string) *http.Response {
	return utils.GenericFetch(http.MethodGet, CHARACTER_URL+character+"/stats", nil)
}

func GetCharacterCollections(character string) *http.Response {
	return utils.GenericFetch(http.MethodGet, CHARACTER_URL+character+"/collections", nil)
}

func GetCharacterProfile(character string) *http.Response {
	return utils.GenericFetch(http.MethodGet, CHARACTER_URL+character+"/profile", nil)
}

func GetCharacterPromotes(character string, params *furrynetwork.FurryNetworkParams) *http.Response {
	return utils.GenericFetch(http.MethodGet, CHARACTER_URL+character+"/promotes", nil)
}

// getting artworks does a search for some reason
// func GetCharacterArtwork(character string, params *furrynetwork.FurryNetworkParams) *http.Response {
// 	paramsDefault := furrynetwork.FurryNetworkParams{
// 		Size:      furrynetwork.QuerySize36,
// 		Time:      "undefined",
// 		From:      0,
// 		Character: character,
// 		Types:     []string{"artwork"},
// 		Sort:      "published",
// 	}
// }

func FurryNetwork() {
}
