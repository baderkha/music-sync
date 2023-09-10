package music

type Song struct {
	ID      string
	Title   string
	Authors []string
}

type SongReader interface {
	SGetByID(ID string) (*Song, error)
	SSearch(resMax int, name string) ([]*Song, error)
	SGetByPlaylistID(pID string, p *PaginatedRequest) (*OptionallyPaginatedResult[Song], error)
}

type SongQueryer interface {
	SongReader
}
