package main

import (
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/baderkha/music-sync/backend/internal/controller/htmx"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	fnMap  = template.FuncMap{
		"add": func(a, b int) string {
			return strconv.FormatInt(int64(a+b), 10)
		},
		"merge_csv": func(slc []string) string {
			return strings.Join(slc, ",")
		},
	}
)

var (
	rscDir    = "./resources/music-sync"
	staticDir = filepath.Join(rscDir, "static")
	tmplatDir = filepath.Join(rscDir, "views", "templates")
)

func main() {
	router.FuncMap = fnMap
	router.Static("/static/", staticDir)
	router.LoadHTMLGlob(filepath.Join(tmplatDir, "*"))
	router.GET("/", htmx.Gin(htmx.HomePage))

	router.POST("/playlists/process", htmx.Gin(htmx.ProcessPlayList))

	router.GET("/playlists", htmx.Gin(htmx.PlayLists))

	router.GET("/tracks", htmx.Gin(htmx.Tracks))

	router.GET("/syncs", htmx.Gin(htmx.Syncs))

	router.Run(":7070")
}
