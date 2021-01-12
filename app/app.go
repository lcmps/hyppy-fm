package app

import (
	"github.com/lcmps/hippyfm/models"
	"github.com/shkh/lastfm-go/lastfm"
)

// InstanceAPI creates a new authenticated API connection
func InstanceAPI(key, secret string) *lastfm.Api {
	api := lastfm.New(key, secret)
	return api
}

// getLargestCover returns a 300x300 album cover
func getLargestCover(img models.ImgHelper) string {
	s := img[len(img)-1]
	return s.Url
}
