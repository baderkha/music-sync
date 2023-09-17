package htmx

import (
	"github.com/baderkha/music-sync/backend/internal/response/view/component"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) (com component.IComponent, err error) {
	return &component.HomePage{
		PageTitle: "🎵 Music Sync",
	}, nil
}
