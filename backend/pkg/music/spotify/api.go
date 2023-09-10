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
	_ music.PlayListManager         = &API{}
	_ music.SongManager             = &API{}
	_ music.UserInformationProvider = &API{}
	t                               = fmt.Sprintf
)

func TrimQ(s string) string {
	return strings.Trim(s, `"`)
}

func NewAPI(token string) *API {
	token = strings.Trim(token, "\n")
	spew.Dump(token)
	return &API{
		client: resty.New().SetHeader("Authorization", "Bearer "+token).SetBaseURL(BASEURL),
	}
}

type URIBuilder struct {
	ResourceType string
	ID           string
}

func (u *URIBuilder) WithResourceType(r string) *URIBuilder {
	u.ResourceType = r
	return u
}

func (u *URIBuilder) AsTrack() *URIBuilder {
	return u.WithResourceType("track")
}

func (u *URIBuilder) WithID(ID string) *URIBuilder {
	u.ID = ID
	return u
}

func (u *URIBuilder) ToString() string {
	return fmt.Sprintf("spotify:%s:%s", u.ResourceType, u.ID)
}

func NewURI() *URIBuilder {
	return new(URIBuilder)
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

func (s *API) CurrentUser() (*music.User, error) {
	res, err := H(s.client.R().Get(t("me")))
	if err != nil {
		return nil, err
	}
	container, err := gabs.ParseJSON(res.Body())
	if err != nil {
		return nil, err
	}
	return &music.User{
		ID:        strings.ReplaceAll(container.Path("id").String(), `"`, ""),
		Name:      strings.ReplaceAll(container.Path("display_name").String(), `"`, ""),
		Followers: int(container.Path("followers.total").Data().(float64)),
	}, nil
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

func (s *API) PSearch(userID string, p music.PaginatedRequester) (*music.OptionallyPaginatedResult[music.PlayList], error) {
	// Implement logic to search for playlists by user ID on Spotify
	// Use s.client to make HTTP GET request
	res, err := H(s.client.R().Get(t("users/%s/playlists?limit=%d&offset=%s", userID, p.GetLimit(), p.GetOffsetOrToken())))
	if err != nil {
		return nil, err
	}

	container, err := gabs.ParseJSON(res.Body())
	if err != nil {
		return nil, err
	}

	playlists := make([]*music.PlayList, 0, p.GetLimit())
	items, _ := container.Path("items").Children()

	for _, item := range items {
		playlist := music.PlayList{
			ID:    TrimQ(item.Path("id").Data().(string)),
			Title: TrimQ(item.Path("name").Data().(string)),
			Link:  item.Path("href").String(),
		}

		playlists = append(playlists, &playlist)
	}

	return &music.OptionallyPaginatedResult[music.PlayList]{
		Data:         playlists,
		IsPaginated:  true,
		TotalRecords: int(container.Path("total").Data().(float64)),
	}, nil
}

func (s *API) SGetByPlaylistID(pID string, p music.PaginatedRequester) (*music.OptionallyPaginatedResult[music.Song], error) {
	res, err := H(s.client.R().Get(t("playlists/%s/tracks?limit=%d&offset=%s", pID, p.GetLimit(), p.GetOffsetOrToken())))
	if err != nil {
		return nil, err
	}

	// "href": "https://api.spotify.com/v1/playlists/5y0or2fDc6evQ4oETccuPc/tracks",
	fmt.Println(string(res.Body()))
	container, err := gabs.ParseJSON(res.Body())
	_ = container
	if err != nil {
		return nil, err
	}
	items, _ := container.Path("items").Children()
	songs := make([]*music.Song, len(items))
	for i, v := range items {
		artists, err := v.Path("track.artists").Children()
		if err != nil {
			return nil, err
		}

		arts := make([]string, len(artists))

		for i, child := range artists {
			arts[i] = child.Path("name").String()
		}
		songs[i] = &music.Song{
			ID:      v.Path("track.id").String(),
			Title:   v.Path("track.name").String(),
			Authors: arts,
		}
	}
	return &music.OptionallyPaginatedResult[music.Song]{
		Data:         songs,
		TotalRecords: int(container.Path("total").Data().(float64)),
		IsPaginated:  true,
	}, nil
}

func (s *API) PGetByID(playListID string) (*music.PlayList, error) {

	res, err := H(s.client.R().Get(t("playLists/%s", playListID)))
	if err != nil {
		return nil, err
	}

	container, err := gabs.ParseJSON(res.Body())
	if err != nil {
		return nil, err
	}
	return &music.PlayList{
		ID:    container.Path("id").String(),
		Title: container.Path("name").String(),
	}, nil
}

func (s *API) PCreate(p *music.PlayList) error {
	// Implement logic to create a new playlist on Spotify
	// Use s.client to make s calls
	return nil
}

func (s *API) PUpdate(p *music.PlayList) error {
	// Implement logic to update an existing playlist on Spotify
	// Use s.client to make s calls
	return nil
}

func (s *API) PDeletePlayList(pID string) error {
	// Implement logic to delete a playlist on Spotify by its ID
	// Use s.client to make s calls
	return nil
}

func (s *API) PAddSongs(pID string, sID ...string) error {
	return nil
}
func (s *API) PRemoveSongs(pID string, sID ...string) error {
	return nil
}
