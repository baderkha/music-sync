package music

type PlayList struct {
	ID               string
	Title            string
	Creator          string
	ImageURL         string
	OptionaImageMeta map[string]interface{}
	Link             string
	Songs            []*Song
}

type PlayListReader interface {
	PGetByID(playListID string) (*PlayList, error)
	PSearch(userID string, p *PaginatedRequest) (*PagintatedResult[PlayList], error)
}

type PlayListManager interface {
	PCreatePlayList(p *PlayList) error
	PUpdatePlayList(p *PlayList) error
	PDeletePlayList(pID string) error
}

type PlayListQueryer interface {
	PlayListReader
	PlayListManager
}
