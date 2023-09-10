package music

type PlayList struct {
	ID        string
	Title     string
	ImageURLs []string
	Link      string
}

type PlayListReader interface {
	PGetByID(playListID string) (*PlayList, error)
	PSearch(userID string, p *PaginatedRequest) (*OptionallyPaginatedResult[PlayList], error)
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
