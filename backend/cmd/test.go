package main

import (
	"os"

	"github.com/baderkha/music-sync/backend/pkg/music"
	"github.com/baderkha/music-sync/backend/pkg/music/spotify"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	secret, err := os.ReadFile("./secrets/tmp_spot_secret")
	_ = secret
	if err != nil {
		panic(err)
	}
	api := spotify.NewAPI(string(""))
	// cu, err := api.CurrentUser()
	// if err != nil {
	// 	panic(err)
	// }
	// spew.Dump(api.PSearch(cu.ID, &music.PaginatedRequest{
	// 	Offset: 0,
	// 	Limit:  40,
	// }))
	// 	fmt.Println("querying for music")
	// res, err := api.SSearch(10, "enter sandman")
	// if err != nil {
	// 	panic(err)
	// }
	// if len(res) == 0 {
	// 	fmt.Println("no music")
	// 	os.Exit(0)
	// }
	// spew.Dump(api.SGetByID(res[0].ID))

	spew.Dump(api.SGetByPlaylistID("5y0or2fDc6evQ4oETccuPc", &music.PaginatedRequest{
		Offset: 0,
		Limit:  1,
	}))

}
