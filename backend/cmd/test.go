package main

import (
	"os"

	"github.com/baderkha/music-sync/backend/pkg/music/spotify"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	res, err := os.ReadFile("./secrets/tmp_spot_secret")
	if err != nil {
		panic(err)
	}
	api := spotify.NewAPI(string(res))
	spew.Dump(api.SSearch(10, "bulbasor"))
}
