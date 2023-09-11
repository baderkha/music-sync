package service

import (
	"errors"

	"github.com/baderkha/music-sync/backend/pkg/music"
	"golang.org/x/sync/errgroup"
)

const (
	LargeIteration  = 1000000
	PaginationLimit = 50
)

var (
	InfiniteLoopError = errors.New("Infinite loop occured during music Sync")
)

type streamingProvider = music.StreamingServiceProvider

type SourceToTargetProivder struct {
	Source music.StreamingServiceProvider
	Target music.StreamingServiceProvider
}

type SynchronizationReq struct {
	SourceToTargetProivder
	SourcePlayListIDs []string
}

func fetchUserInfo(s SourceToTargetProivder) (sourceUsr *music.User, target *music.User, err error) {
	var wg errgroup.Group
	wg.Go(func() error {
		usr, err := s.Source.CurrentUser()
		sourceUsr = usr
		return err
	})
	wg.Go(func() error {
		usr, err := s.Target.CurrentUser()
		target = usr
		return err
	})

	err = wg.Wait()
	if err != nil {
		return nil, nil, err
	}
	return sourceUsr, target, nil
}

type PlayListWithMusic struct {
	music.PlayList
	Songs []*music.Song
}

func fetchPlayListsWithSongs(r *SynchronizationReq) ([]*PlayListWithMusic, error) {
	var (
		plistMusic []*PlayListWithMusic = make([]*PlayListWithMusic, len(r.SourcePlayListIDs))
		wg         errgroup.Group
	)
	// grab playlists
	plists, err := r.Source.PGetByIDs(r.SourcePlayListIDs...)
	if err != nil {
		return nil, err
	}

	for i, v := range plists {
		wg.Go(func(index int, v *music.PlayList) func() error {
			return func() error {
				res, err := resolveMusicForPlayList(v, r.Source)
				if err != nil {
					return err
				}

				plistMusic[i] = &PlayListWithMusic{
					PlayList: *v,
					Songs:    res,
				}
				return nil
			}
		}(i, v))
	}
	err = wg.Wait()
	if err != nil {
		return nil, err
	}
	return plistMusic, nil
}

func resolveMusicForPlayList(p *music.PlayList, source music.StreamingServiceProvider) ([]*music.Song, error) {
	// this will be running on a machine that will be pay as you go
	//so i don't want to cause an infinite loop and risk taxing myself with a large bill :)
	var tracks []*music.Song
	var offsetOrToken string
Free:
	for i := 0; i < LargeIteration; i++ {
		if i+1 >= LargeIteration {
			return nil, errors.New("Infinite loop occured")
		}
		res, err := source.
			SGetByPlaylistID(
				p.ID,
				source.
					NewPaginator().
					SetLimit(PaginationLimit).
					SetOffsetOrToken(offsetOrToken),
			)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, res.Data...)
		if !res.IsPaginated || len(res.Data) == 0 {
			break Free
		}
		offsetOrToken = res.NextPage
		break Free
	}

	return tracks, nil
}

func createAllPlayLists()

func SynchronizePlaylist(r *SynchronizationReq) error {

	// grab users
	sourceUser, targetUser, err := fetchUserInfo(r.SourceToTargetProivder)
	if err != nil {
		return err
	}
	_ = sourceUser
	_ = targetUser

	playLists, err := fetchPlayListsWithSongs(r)
	if err != nil {
		return err
	}

}
