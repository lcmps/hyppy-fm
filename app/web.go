package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lcmps/hippyfm/models"
	"github.com/shkh/lastfm-go/lastfm"
)

var r *gin.Engine

// Host any
func Host(conn *lastfm.Api, env string) {

	if env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.LoadHTMLGlob("./pages/html/*")

	r.Static("/assets", "./pages/assets/")
	r.Static("/fvc/", "./pages/assets/img/favicon/")

	r.GET("/", home)
	r.GET("/img", serveCollage(conn))

	p := os.Getenv("PORT")
	// p := "8080"
	err := r.Run(":" + p)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func home(ctx *gin.Context) {

	// Call the HTML method of the Context to render a template
	ctx.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Hyppy FM",
		},
	)
}

func serveCollage(conn *lastfm.Api) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var json models.CollageParams
		var query models.UriParams

		if err := ctx.ShouldBindJSON(&json); err == nil {

			img, err := GetAlbumsByPeriod(conn, json.Username, json.Period, json.Limit)
			if err != nil {
				ctx.Data(http.StatusBadRequest, "text/plain", []byte(err.Error()))
				return
			}
			b, err := CreateByteCollage(img, json.Size)
			if err != nil {
				ctx.Data(http.StatusBadRequest, "text/plain", []byte(err.Error()))
				return
			}

			ctx.Data(http.StatusOK, "image/png", b.Bytes())
			return
		} else if err = ctx.ShouldBindQuery(&query); err == nil {

			img, err := GetAlbumsByPeriod(conn, query.Username, query.Period, query.Size*query.Size)
			if err != nil {
				ctx.Data(http.StatusBadRequest, "text/plain", []byte(err.Error()))
				return
			}
			b, err := CreateByteCollage(img, query.Size)
			if err != nil {
				ctx.Data(http.StatusBadRequest, "text/plain", []byte(err.Error()))
				return
			}

			ctx.Data(http.StatusOK, "image/png", b.Bytes())
			return
		}
	}

	return fn
}
