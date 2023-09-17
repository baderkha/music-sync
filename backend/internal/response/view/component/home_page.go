package component

import (
	"os"
	"path/filepath"
)

var _ IComponent = &HomePage{}

var (
	cwd, _    = os.Getwd()
	rscDir    = filepath.Join(cwd, "./resources/music-sync")
	staticDir = filepath.Join(rscDir, "static")
	tmplatDir = filepath.Join(rscDir, "views", "templates")
)

var templateBasePath = tmplatDir

type HomePage struct {
	PageTitle string
}

func (h *HomePage) GetTemplate() string {
	return "index.html"
}
