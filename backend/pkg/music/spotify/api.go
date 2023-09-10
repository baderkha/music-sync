package spotify

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/baderkha/music-sync/backend/pkg/music"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
)

const (
	BASEURL  = "https://api.spotify.com/v1/"
	Provider = "SPOTIFY"
)

var (
	_ music.PlayListQueryer = &API{}
	_ music.SongQueryer     = &API{}
	t                       = fmt.Sprintf
)

func NewAPI(token string) *API {
	token = strings.Trim(token, "\n")
	spew.Dump(token)
	return &API{
		client: resty.New().SetHeader("Authorization", "Bearer "+token).SetBaseURL(BASEURL),
	}
}

type API struct {
	client *resty.Client
}

func H(res *resty.Response, err error) (*resty.Response, error) {
	if err != nil {
		return nil, &music.UnexpectedFatalError{
			OgError:  err,
			Provider: Provider,
		}
	}

	switch res.StatusCode() {
	case http.StatusNotFound:
		return nil, &music.ResourceNotFoundError{
			ResourceType: res.Request.URL,
			ID:           res.Request.URL,
			Provider:     Provider,
		}
	default:
		if res.StatusCode() >= 200 && res.StatusCode() <= 300 {
			return res, nil
		}
		return nil, &music.UnexpectedFatalError{
			OgError:  errors.New(string(res.Body())),
			Provider: Provider,
		}
	}
}

func (s *API) SGetByID(ID string) (*music.Song, error) {
	res, err := H(s.client.R().Get(t("tracks/%s", ID)))
	if err != nil {
		return nil, err
	}
	container, err := gabs.ParseJSON(res.Body())
	if err != nil {
		return nil, err
	}
	artists, err := container.Path("artists").Children()
	if err != nil {
		return nil, err
	}

	arts := make([]string, len(artists))

	for i, child := range artists {
		arts[i] = child.Path("name").String()
	}

	return &music.Song{
		ID:      container.Path("id").String(),
		Title:   container.Path("name").String(),
		Authors: arts,
	}, nil

}

func (s *API) SSearch(resMax int, name string) ([]*music.Song, error) {
	// Implement logic to search for songs by name on Spotify
	// Use s.client to make HTTP GET request
	res, err := H(s.client.R().SetQueryParam("q", name).SetQueryParam("type", "track").SetQueryParam("limit", t("%d", resMax)).Get("search"))
	if err != nil {
		return nil, err
	}

	container, err := gabs.ParseJSON(res.Body())
	if err != nil {
		return nil, err
	}

	songs := []*music.Song{}
	items, _ := container.Path("tracks.items").Children()

	for _, item := range items {
		artists, err := item.Path("artists").Children()
		if err != nil {
			return nil, err
		}
		arts := make([]string, len(artists))

		for i, child := range artists {
			arts[i] = child.Path("name").Data().(string)
		}

		song := &music.Song{
			ID:      item.Path("id").Data().(string),
			Title:   item.Path("name").Data().(string),
			Authors: arts,
		}

		songs = append(songs, song)
	}

	return songs, nil
}

func (s *API) PSearch(userID string, p *music.PaginatedRequest) (*music.PagintatedResult[music.PlayList], error) {
	// Implement logic to search for playlists by user ID on Spotify
	// Use s.client to make HTTP GET request
	res, err := H(s.client.R().Get(t("users/%s/playlists", userID)))
	if err != nil {
		return nil, err
	}

	container, err := gabs.ParseJSON(res.Body())
	if err != nil {
		return nil, err
	}

	playlists := make([]*music.PlayList, 0, p.Size)
	items, _ := container.Path("items").Children()

	for _, item := range items {
		playlist := music.PlayList{
			ID:    item.Path("id").Data().(string),
			Title: item.Path("name").Data().(string),
		}

		playlists = append(playlists, &playlist)
	}

	return &music.PagintatedResult[music.PlayList]{Data: playlists}, nil
}

func (s *API) PGetByID(playListID string) (*music.PlayList, error) {
	// Implement logic to fetch a playlist by its ID from Spotify
	// Use s.client to make s calls
	// Parse the response into a PlayList struct
	return nil, nil
}

func (s *API) PCreatePlayList(p *music.PlayList) error {
	// Implement logic to create a new playlist on Spotify
	// Use s.client to make s calls
	return nil
}

func (s *API) PUpdatePlayList(p *music.PlayList) error {
	// Implement logic to update an existing playlist on Spotify
	// Use s.client to make s calls
	return nil
}

func (s *API) PDeletePlayList(pID string) error {
	// Implement logic to delete a playlist on Spotify by its ID
	// Use s.client to make s calls
	return nil
}
