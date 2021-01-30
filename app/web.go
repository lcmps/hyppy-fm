package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcmps/hippyfm/models"
	"github.com/shkh/lastfm-go/lastfm"
)

var r *gin.Engine

// Host runs the webapplication and related endpoints
func Host(conn *lastfm.Api, p string) {
	r := gin.Default()
	r.LoadHTMLGlob("./pages/html/*")

	r.Static("/assets", "./pages/assets/")
	r.Static("/fvc/", "./pages/assets/img/favicon/")
	r.Static("/hy", "./pages/assets/img/clg/")

	// Pages
	r.GET("/", home)
	r.GET("/dash", dashboards)

	// APIs
	r.GET("/usr", usrInfoEndpoint(conn))
	r.POST("/img", serveImageClg(conn))
	r.POST("/clg", serveLink(conn))

	err := r.Run(":" + p)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func home(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Hyppy FM",
		},
	)
}

func dashboards(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"dashboard.html",
		gin.H{
			"title": "Hyppy FM | Dashboard",
		},
	)
}

func usrInfoEndpoint(conn *lastfm.Api) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var uri models.UsrURI

		if err := ctx.ShouldBindQuery(&uri); err == nil {
			info, err := GetUserInfo(conn, uri.Username)
			if err != nil {
				errResp := models.InternalError{
					Reason: err.Error(),
				}
				ctx.JSON(http.StatusBadRequest, errResp)
				return
			}
			ctx.JSON(http.StatusOK, info)
			return
		}
	}
	return fn
}

func serveLink(conn *lastfm.Api) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var usrPayload models.CollageParams

		if err := ctx.ShouldBindJSON(&usrPayload); err == nil {

			img, err := GetAlbumsByPeriod(conn, usrPayload.Username, usrPayload.Period, usrPayload.Size*usrPayload.Size)
			if err != nil {
				errResp := models.InternalError{
					Reason: err.Error(),
				}

				ctx.JSON(http.StatusBadRequest, errResp)
				return
			}
			if img == nil {
				resp := models.ResponseClg{
					Path: "/hy/" + filenameGenerator(usrPayload.Username, usrPayload.Period, usrPayload.Size),
				}
				ctx.JSON(http.StatusOK, resp)
				return
			}

			b, err := CreateByteCollage(img, usrPayload.Size)
			if err != nil {
				errResp := models.InternalError{
					Reason: err.Error(),
				}

				ctx.JSON(http.StatusBadRequest, errResp)
				return
			}
			fileName := saveCollage(b.Bytes(), usrPayload.Username, usrPayload.Period, usrPayload.Size)
			resp := models.ResponseClg{
				Path: "/hy/" + fileName,
			}

			ctx.JSON(http.StatusOK, resp)
			return
		}
	}
	return fn
}

func serveImageClg(conn *lastfm.Api) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var usrPayload models.CollageParams
		var query models.URIParams

		if err := ctx.ShouldBindJSON(&usrPayload); err == nil {

			img, err := GetAlbumsByPeriod(conn, usrPayload.Username, usrPayload.Period, usrPayload.Size*usrPayload.Size)
			if err != nil {
				errResp := models.InternalError{
					Reason: err.Error(),
				}

				ctx.JSON(http.StatusBadRequest, errResp)
				return
			}
			b, err := CreateByteCollage(img, usrPayload.Size)
			if err != nil {
				errResp := models.InternalError{
					Reason: err.Error(),
				}

				ctx.JSON(http.StatusBadRequest, errResp)
				return
			}

			ctx.Data(http.StatusOK, "image/png", b.Bytes())
			return
		} else if err = ctx.ShouldBindQuery(&query); err == nil {

			img, err := GetAlbumsByPeriod(conn, query.Username, query.Period, query.Size*query.Size)
			if err != nil {
				errResp := models.InternalError{
					Reason: err.Error(),
				}

				ctx.JSON(http.StatusBadRequest, errResp)
				return
			}
			b, err := CreateByteCollage(img, query.Size)
			if err != nil {
				errResp := models.InternalError{
					Reason: err.Error(),
				}

				ctx.JSON(http.StatusBadRequest, errResp)
				return
			}

			ctx.Data(http.StatusOK, "image/png", b.Bytes())
			return
		}
	}

	return fn
}
