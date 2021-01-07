package app

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/deiwin/picasso"
	clg "github.com/ozankasikci/go-image-merge"
)

//sizes: 3x3 = 900p, 4x4 = 1200p, 5x5 = 1500p, 10x10 = 3000p

// GenerateCollage makes a img with the given proportions and covers
func GenerateCollage(img []image.Image) image.Image {
	layout := picasso.DrawGridLayout(img, 900)
	return layout
}

// GenerateCollageGrid a
func GenerateCollageGrid(img []image.Image) {

	grids := []*clg.Grid{}

	for i := 0; i < len(img); i++ {
		data := clg.Grid{
			Image: &img[i],
		}
		fmt.Println(data.Image)
		grids = append(grids, &data)
	}

	rgba, err := clg.New(grids, 3, 3).Merge()
	if err != nil {
		fmt.Println(err.Error())
	}

	file, err := os.Create("testcolage.png")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = png.Encode(file, rgba)
	if err != nil {
		fmt.Println(err.Error())
	}
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
