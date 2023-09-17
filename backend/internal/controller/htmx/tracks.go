package htmx

import (
	"github.com/baderkha/music-sync/backend/internal/response/view/component"
	"github.com/gin-gonic/gin"
)

func Tracks(c *gin.Context) (com component.IComponent, err error) {
	return component.
		NewTable([]component.TrackData{
			{
				Title:    "MUSIKA",
				Service:  "Spotify",
				Duration: "1:20",
				Artists:  []string{"Ahmad", "James"},
			},
		}).
		WithTableTitle("Tracks").
		WithActionPostLink("/tracks/process").
		WithBooleanBtnName("Add").
		WithActionTitle("Add"), nil
}
