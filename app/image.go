package app

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"time"

	clg "github.com/ozankasikci/go-image-merge"
	"github.com/shkh/lastfm-go/lastfm"
)

//sizes: 3x3 = 900p, 4x4 = 1200p, 5x5 = 1500p, 10x10 = 3000p

// GetAlbumsByPeriod where u- username, p- period : overall | 7day | 1month | 3month | 6month | 12month And l- limit : The number of results to fetch per page. Defaults to 50.
func GetAlbumsByPeriod(api *lastfm.Api, u, p string, l int) ([]image.Image, error) {

	ex := existCollage(filenameGenerator(u, p, int(math.Sqrt(float64(l)))))
	if ex {
		return nil, nil
	}

	if u == "" {
		return nil, errors.New("user empty or invalid")
	}
	if p == "" {
		return nil, errors.New("period empty or invalid")
	}
	if l < 1 {
		return nil, errors.New("size empty or invalid")
	}

	opts := lastfm.P{"user": u, "period": p, "limit": l}

	var covers []image.Image

	res, err := api.User.GetTopAlbums(opts)
	if err != nil {
		return nil, err
	} else if res.Albums == nil {
		return nil, errors.New("User " + u + " has no top albums")
	}

	for _, album := range res.Albums {
		img := getLargestImage(album.Images)
		pCover := handleImage(img)
		covers = append(covers, pCover)
	}

	return covers, nil
}

// CreateByteCollage generates a album collage with the sourced images
func CreateByteCollage(img []image.Image, s int) (*bytes.Buffer, error) {

	if s < 1 {
		return nil, errors.New("size cannot be empty")
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

	enc := &png.Encoder{
		CompressionLevel: png.BestSpeed,
	}
	enc.Encode(buf, rgba)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func saveCollage(clg []byte, u, p string, s int) (fileName string) {

	name := filenameGenerator(u, p, s)

	err := ioutil.WriteFile("pages/assets/img/clg/"+name, clg, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Saving file " + name)
	return name
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

func filenameGenerator(u, p string, s int) string {
	t := time.Now()
	today := t.Format("2006-01-02-T15")
	name := fmt.Sprintf(`%s_%s_%s_%dx%d.png`, u, today, p, s, s)
	return name
}

func existCollage(fileName string) bool {
	_, err := os.Open("pages/assets/img/clg/" + fileName)
	if err != nil {
		return false
	}
	return true
}
