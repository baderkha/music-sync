package htmx

import (
	"github.com/baderkha/music-sync/backend/internal/response/view"
	"github.com/baderkha/music-sync/backend/internal/response/view/component"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/maragudk/gomponents"
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

func ProcessPlayList(c *gin.Context) (node gomponents.Node, err error) {
	spew.Dump("called")
	return view.
		ProcessPlayListModal(&view.ProcessPlayList{
			ServiceSelection: []string{"Spotify", "Youtube Music"},
		}), nil
}
