package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcmps/hippyfm/models"
	"github.com/shkh/lastfm-go/lastfm"
)

var r *gin.Engine

// Host any
func Host(conn *lastfm.Api) {
	r := gin.Default()
	r.LoadHTMLGlob("./html/pages/*")
	r.Static("/css", "./html/css/")
	r.Static("/js", "./html/js/")
	r.Static("/images","./html/images/")

	r.GET("/", home)
	r.GET("/img", serveCollage(conn))

	r.Run()
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
			"title": "Home Page",
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
