package htmx

import (
	"net/http"

	"github.com/baderkha/music-sync/backend/internal/response/view/component"
	"github.com/baderkha/music-sync/backend/pkg/http/status"
	"github.com/gin-gonic/gin"
	"github.com/maragudk/gomponents"
)

type ViewHandler func(c *gin.Context) (com component.IComponent, err error)
type GomponentHandler func(c *gin.Context) (node gomponents.Node, err error)

func Gin(v ViewHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := v(ctx)
		if err != nil {
			cast, ok := err.(*status.Error)
			if !ok {
				ctx.String(http.StatusInternalServerError, cast.Error())
				return
			}
			ctx.String(cast.StatusCode, cast.Error())
			return
		}
		ctx.HTML(status.ResolveFromMethod(ctx.Request), res.GetTemplate(), res)
	}
}

func GinC(v GomponentHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := v(ctx)
		if err != nil {
			cast, ok := err.(*status.Error)
			if !ok {
				ctx.String(http.StatusInternalServerError, cast.Error())
				return
			}
			ctx.String(cast.StatusCode, cast.Error())
			return
		}
		res.Render(ctx.Writer)
	}
}
