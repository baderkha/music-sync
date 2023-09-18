package component

import (
	"github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type ModalData struct {
	Title  string
	Body   gomponents.Node
	Footer gomponents.Node
	IsForm bool
}

type ModalActionButtonData struct {
	ActionButtonTitle            string
	ActionButtonExtraAttributes  []gomponents.Node
	DismissButtonTitle           string
	DismissButtonExtraAttributes []gomponents.Node
}

func ModalFooterActionButtonsC(m *ModalActionButtonData) gomponents.Node {
	return Div(
		Button(
			append(
				m.DismissButtonExtraAttributes,
				Type("button"),
				Class("btn btn-secondary"),
				DataAttr("bs-dismiss", "modal"),
				gomponents.Text(m.DismissButtonTitle),
			)...,
		),
		Button(
			append(
				m.ActionButtonExtraAttributes,
				Class("btn btn-primary"),
				gomponents.Text(m.ActionButtonTitle),
			)...,
		),
		c.Classes{
			"d-flex":                  true,
			"justify-content-between": true,
			"w-100":                   true,
		},
	)
}

func ModalC(d *ModalData) gomponents.Node {
	return Div(
		Div(
			Div(
				FormEl(
					Div(
						H5(
							gomponents.Text(d.Title),
						),
						c.Classes{
							"modal-header": true,
						},
					),
					Div(
						d.Body,
						c.Classes{
							"modal-body": true,
						},
					),
					Div(
						d.Footer,
						c.Classes{
							"modal-footer": true,
						},
					),
					c.Classes{
						"w-100": true,
						"h-100": true,
					},
				),

				c.Classes{
					"modal-content": true,
				},
			),
		),
		c.Classes{
			"modal-dialog":          true,
			"modal-dialog-centered": true,
		},
	)
}
