package repository

type SongProvider interface {
	GetSongByID(id string)
}
