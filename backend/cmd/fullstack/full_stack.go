package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

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
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"page_title": "🎵 Music Sync",
		})
	})

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

	router.GET("/playlists", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "tables.html", gin.H{
			"action_post_link": "/playlists/process",
			"boolean_btn_name": "Sync",
			"action_title":     "sync",
			"table_title":      "Playlists",
			"columns":          []string{"title", "service", "creator"},
			"data": []map[string]string{
				{

					"1title":   "R/B",
					"2service": "spotify",
					"3creator": "ahmad",
				},
				{

					"1title":   "Rap",
					"2service": "spotify",
					"3creator": "ahmad",
				},
			},
		})

	})

	router.GET("/tracks", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "tables.html", gin.H{
			"action_post_link": "/tracks/process",
			"boolean_btn_name": "Add",
			"action_title":     "Add",
			"table_title":      "Tracks",
			"columns":          []string{"title", "duration", "service", "artists"},
			"data": []map[string]any{
				{
					"1title":    "R/B",
					"2duration": "1:20",
					"3service":  "spotify",
					"4artists":  "ahmad,james",
				},
				{
					"1title":    "R/B",
					"2duration": "1:20",
					"3service":  "spotify",
					"4artists":  "ahmad,james",
				},
			},
		})
	})

	router.GET("/syncs", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "tables.html", gin.H{
			"table_title":          "Synchronizations ",
			"action_button_hidden": true,
			"columns":              []string{"label", "last_run", "from", "to", "status"},
			"data": []map[string]any{
				{
					"1label":    "my_cool_sync",
					"2last_run": time.Now().Add(-1 * time.Hour * 24).Format(time.DateTime),
					"3from":     "spotify",
					"4to":       "youtube music",
					"5status":   "ok",
				},
			},
		})
	})

	router.Run(":7070")
}
