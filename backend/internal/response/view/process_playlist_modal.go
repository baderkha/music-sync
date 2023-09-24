package view

import (
	"github.com/baderkha/music-sync/backend/internal/response/view/component"
	"github.com/maragudk/gomponents"
)

type ProcessPlayList struct {
	ServiceSelection []string
}

func ProcessPlayListModal(data *ProcessPlayList) gomponents.Node {
	return component.CustomModalC(
		&component.ModalData{
			Title: "Where would you like to Sync your playlist?",
			Body: component.SelectC(&component.SelectData{
				Options:    data.ServiceSelection,
				HelperText: "Select a Service",
				FormName:   "service_select",
			}),
			Footer: component.ModalFooterActionButtonsC(
				&component.ModalActionButtonData{
					ActionButtonTitle: "Process",
					ActionButtonExtraAttributes: []gomponents.Node{
						gomponents.Attr("hx-trigger", "click"),
						gomponents.Attr("hx-post", "/playlists/sync"),
					},
					DismissButtonTitle: "Cancel",
				},
			),
		},
	)
}
