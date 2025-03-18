package scraper

import (
	"net/url"

	"google.golang.org/genproto/googleapis/type/date"
)

type GenericArtworkMeta struct {
	Name          string            `json:"art_name"`
	Description   string            `json:"description,omitempty"`
	Provider      string            `json:"provider"`
	URL           url.URL           `json:"url"`
	ImageURL      url.URL           `json:"imageUrl"`
	DateSubmitted date.Date         `json:"dateSubmitted"`
	Artist        GenericArtistMeta `json:"artist"`
}

type GenericArtistMeta struct {
	URL        url.URL   `json:"url"`
	AvatarURL  url.URL   `json:"avatarUrl"`
	DateJoined date.Date `json:"dateJoined"`
	Name       string    `json:"name"`
	Handle     string    `json:"handle,omitempty"`
}

// Other art metas
type PinterestMeta struct {
	PinID int `json:"pinterestId,omitempty"`
}
