package app

import (
	"fmt"

	"github.com/shkh/lastfm-go/lastfm"
)

// InstanceAPI creates a new authenticated API connection
func InstanceAPI(key, secret string) *lastfm.Api {
	api := lastfm.New(key, secret)
	return api
}

func iterate(key, secret string) {
	ctx := InstanceAPI(key, secret)

	u, _ := ctx.User.GetTopTracks(lastfm.P{"user": "luka1498"})

	for _, ut := range u.Tracks {
		fmt.Println(ut)
	}
}
