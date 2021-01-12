package app

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"net/http"
	"strings"

	clg "github.com/ozankasikci/go-image-merge"
	"github.com/shkh/lastfm-go/lastfm"
)

//sizes: 3x3 = 900p, 4x4 = 1200p, 5x5 = 1500p, 10x10 = 3000p

// GetAlbumsByPeriod where u- username, p- period : overall | 7day | 1month | 3month | 6month | 12month And l- limit : The number of results to fetch per page. Defaults to 50.
func GetAlbumsByPeriod(api *lastfm.Api, u, p string, l int) ([]image.Image, error) {

	if u == "" {
		return nil, errors.New("User cannot be empty")
	}
	if p == "" {
		return nil, errors.New("Period cannot be empty")
	}
	if l < 1 {
		return nil, errors.New("limit cannot be empty")
	}

	opts := lastfm.P{"user": u, "period": p, "limit": l}

	var covers []image.Image

	res, err := api.User.GetTopAlbums(opts)
	if err != nil {
		return nil, err
	}

	for _, album := range res.Albums {
		img := getLargestCover(album.Images)
		pCover := handleImage(img)
		covers = append(covers, pCover)
	}

	return covers, nil
}

// CreateByteCollage a
func CreateByteCollage(img []image.Image, s int) (*bytes.Buffer, error) {

	if s < 1 {
		return nil, errors.New("Size cannot be empty")
	}

	grids := []*clg.Grid{}

	for i := 0; i < len(img); i++ {
		data := clg.Grid{
			Image: &img[i],
		}
		grids = append(grids, &data)
	}

	rgba, err := clg.New(grids, s, s).Merge()
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	png.Encode(buf, rgba)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func handleImage(url string) image.Image {
	if url == "" {
		i := image.NewRGBA(image.Rect(0, 0, 300, 300))
		c := color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 1,
		}
		i.SetRGBA(300, 300, c)
		return i
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	img, _, err := image.Decode(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	res.Body.Close()

	return img
}

func coverSave(url, u string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	res.Body.Close()

	n := strings.SplitAfter(url, "300x300/")
	name := u + "_" + n[len(n)-1]

	err = ioutil.WriteFile("img/"+name, data, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Saving file " + name)
}
