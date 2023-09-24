package component

import (
	"github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	"github.com/samber/lo"
)

type SelectData struct {
	Options    []string
	HelperText string
	FormName   string
}

func SelectC(dat *SelectData) gomponents.Node {

	return Select(
		Div(
			append([]gomponents.Node{
				Option(
					Selected(),
					gomponents.Text(dat.HelperText),
				),
			}, lo.Map(dat.Options, func(itm string, _ int) gomponents.Node {
				return Option(
					Value(itm),
					gomponents.Text(itm),
				)
			})...)...,
		),
		c.Classes{
			"form-select": true,
		},
		Name(dat.FormName),
	)
}
