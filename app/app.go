package app

import (
	"fmt"

	"github.com/shkh/lastfm-go/lastfm"
)

const (
	key    = "5e6099013422929a716eaba7e3c1ab65"
	secret = "f6ed441cb8b0aee37b73decd65c499d7"
)

//encodeAuth outputs Base64 auth token
func placeHolder() {

	fmt.Println("Holding this in place")
}

func instanceAPI(key, secret string) *lastfm.Api {
	api := lastfm.New(key, secret)
	return api
}

func iterate() {
	ctx := instanceAPI(key, secret)

	u, _ := ctx.User.GetTopTracks(lastfm.P{"user": "luka1498"})

	for _, ut := range u.Tracks {
		fmt.Println(ut)
	}
}
