package app

import (
	"fmt"
	"image"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/deiwin/picasso"
)

//sizes: 3x3 = 900p, 4x4 = 1200p, 5x5 = 1500p, 10x10 = 3000p

// GenerateCollage makes a img with the given proportions and covers
func GenerateCollage(img []image.Image) image.Image {
	layout := picasso.DrawGridLayout(img, 900)
	return layout
}

func handleImage(url string) image.Image {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()

	return img
}

func coverSave(url, u string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()

	n := strings.SplitAfter(url, "300x300/")
	name := u + "_" + n[len(n)-1]

	err = ioutil.WriteFile("img/"+name, data, 0666)
	if err != nil {
		panic(err)
	}
	fmt.Println("Saving file " + name)
}
