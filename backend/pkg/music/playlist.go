package music

type PlayList struct {
	ID        string
	Title     string
	ImageURLs []string
	Link      string
}

type PlayListReader interface {
	PGetByID(playListID string) (*PlayList, error)
	PGetByIDs(playListIDs ...string) ([]*PlayList, error)
	PSearch(userID string, p PaginatedRequester) (*OptionallyPaginatedResult[PlayList], error)
}

type PlayListWriter interface {
	PCreate(p *PlayList) error
	PUpdate(p *PlayList) error
	PDeletePlayList(pID string) error
	PAddSongs(pID string, sID ...string) error
	PRemoveSongs(pID string, sID ...string) error
}

type PlayListManager interface {
	PlayListReader
	PlayListWriter
}
