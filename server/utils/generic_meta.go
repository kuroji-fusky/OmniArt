package utils

import (
	"google.golang.org/genproto/googleapis/type/date"
)

type GenericArtworkMeta struct {
	Name          string            `json:"art_name"`
	Description   string            `json:"description,omitempty"`
	Provider      string            `json:"provider"`
	URL           string            `json:"url"`
	ImageURL      string            `json:"imageUrl"`
	DateSubmitted date.Date         `json:"dateSubmitted"`
	Artist        GenericArtistMeta `json:"artist"`
}

type GenericArtistMeta struct {
	URL        string    `json:"url"`
	AvatarURL  string    `json:"avatarUrl"`
	DateJoined date.Date `json:"dateJoined"`
	Name       string    `json:"name"`
	Handle     string    `json:"handle,omitempty"`
}
