package scraper

import (
	"net/url"

	"google.golang.org/genproto/googleapis/type/date"
)

type GenericArtworkMeta struct {
	Name          string
	Description   string `json:"description"`
	Provider      string
	URL           url.URL
	ImageURL      url.URL
	DateSubmitted date.Date
	Artist        GenericArtistMeta
}

type GenericArtistMeta struct {
	URL        url.URL
	AvatarURL  url.URL
	DateJoined date.Date
	Name       string
	Handle     string
}

// Other art metas
type PinterestMeta struct {
	PinID int
}
