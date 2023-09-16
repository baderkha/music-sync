package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

var (
	rscDir    = "./resources/music-sync"
	staticDir = filepath.Join(rscDir, "static")
	tmplatDir = filepath.Join(rscDir, "views", "templates")
)

func main() {

	router.Static("/static/", staticDir)
	router.LoadHTMLGlob(filepath.Join(tmplatDir, "*"))
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"page_title": "ahmad",
		})
	})
	router.GET("/playlists", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "playlists.html", gin.H{
			"playlists": []map[string]string{
				{
					"title": "R/B",
				},
				{
					"title": "Jazz",
				},
				{
					"title": "Phonk",
				},
			},
		})
	})
	router.Run(":7070")
}
