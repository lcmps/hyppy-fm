package app

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"

	"github.com/lcmps/hippyfm/models"
	"github.com/shkh/lastfm-go/lastfm"
)

// InstanceAPI creates a new authenticated API connection
func InstanceAPI(key, secret string) *lastfm.Api {
	api := lastfm.New(key, secret)
	return api
}

// GetAlbumsByPeriod where u- username, p- period : overall | 7day | 1month | 3month | 6month | 12month And l- limit : The number of results to fetch per page. Defaults to 50.
func GetAlbumsByPeriod(api *lastfm.Api, u, p string, l int) {
	opts := lastfm.P{"user": u, "period": p, "limit": l}
	var covers []image.Image

	res, err := api.User.GetTopAlbums(opts)
	if err != nil {
		panic(err)
	}

	for _, album := range res.Albums {
		img := getLargestCover(album.Images)
		pCover := handleImage(img)
		covers = append(covers, pCover)
	}

	collage := GenerateCollage(covers)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, collage, nil)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("testcolage.jpg", buf.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
}

// getLargestCover returns a 300x300 album cover
func getLargestCover(img models.ImgHelper) string {
	s := img[len(img)-1]
	return s.Url
}
