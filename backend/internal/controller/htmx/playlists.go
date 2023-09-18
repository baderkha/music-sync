package htmx

import (
	"github.com/baderkha/music-sync/backend/internal/response/view/component"
	"github.com/gin-gonic/gin"
)

func PlayLists(c *gin.Context) (com component.IComponent, err error) {
	return component.
		NewTable([]component.PlayListData{
			{
				Title:   "STD_PHONK",
				Service: "Spotify",
				Creator: "Ahmad BAderkhan",
			},
		}).
		WithTableTitle("PlayLists").
		WithActionPostLink("/playlists/process").
		WithBooleanBtnName("Sync").
		WithActionTitle("Sync"), nil
}

func ProcessPlayList(c *gin.Context) (com component.IComponent, err error) {
	return &component.Modal{
		Title:            "Where Do You Want to process?",
		TextBody:         "Choose where you want to process it",
		CloseButtonTitle: "Exit",
	}, nil
}
