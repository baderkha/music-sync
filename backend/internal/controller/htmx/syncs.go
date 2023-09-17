package htmx

import (
	"time"

	"github.com/baderkha/music-sync/backend/internal/response/view/component"
	"github.com/gin-gonic/gin"
)

func Syncs(c *gin.Context) (com component.IComponent, err error) {
	return component.
		NewTable([]component.SyncData{
			{
				Label:   "MY Cool sync",
				LastRun: time.Now().Add(-1 * time.Hour * 24).Format(time.DateTime),
				From:    "Spotify",
				To:      "Youtube Music",
				Status:  "Ok",
			},
		}).
		WithTableTitle("Synchronizations").
		WithActionPostLink("/synchronizations/edit").
		WithActionButtonHidden(true), nil
}
