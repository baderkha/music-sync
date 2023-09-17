package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/baderkha/music-sync/backend/internal/controller/htmx"
	"github.com/davecgh/go-spew/spew"
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

	router.POST("/playlists/process", func(ctx *gin.Context) {

		var m map[string]any

		decoder := json.NewDecoder(ctx.Request.Body)

		if err := decoder.Decode(&m); err != nil {
			fmt.Println(err)
			return
		}

		spew.Dump(m)
		ctx.JSON(200, "ok")

	})

	router.GET("/playlists", htmx.Gin(htmx.PlayLists))

	router.GET("/tracks", htmx.Gin(htmx.Tracks))

	router.GET("/syncs", htmx.Gin(htmx.Syncs))

	router.Run(":7070")
}
