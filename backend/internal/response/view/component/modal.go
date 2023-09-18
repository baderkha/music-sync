package component

type Modal struct {
	Title            string
	TextBody         string
	CloseButtonTitle string
}

func (m *Modal) GetTemplate() string {
	return "modal.html"
}
