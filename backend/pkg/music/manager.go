package music

import (
	"errors"

	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/errgroup"
)

const (
	MaxLargeIteration = 1000
	PaginationLimit   = 50
)

var (
	InfiniteLoopError = errors.New("Infinite loop occured during music Sync or iteration exceeded max allowed")
)

func New(p StreamingServiceProvider) *StreamingServiceManager {
	return &StreamingServiceManager{
		StreamingServiceProvider: p,
	}
}

// StreamingServiceManager : abstraction above the regular provider that allows you to
type StreamingServiceManager struct {
	StreamingServiceProvider
	ConcurrentLimit int
}

func (a *StreamingServiceManager) FetchPlayListsWithSongs(SourcePlayListIDs ...string) ([]*PlayListWithMusic, error) {
	var (
		plistMusic []*PlayListWithMusic = make([]*PlayListWithMusic, len(SourcePlayListIDs))
		wg         errgroup.Group
	)
	// grab playlists
	plists, err := a.PGetByIDs(SourcePlayListIDs...)
	if err != nil {
		return nil, err
	}

	for i, v := range plists {
		wg.Go(func(index int, v *PlayList) func() error {
			return func() error {
				res, err := a.resolveMusicForPlayList(v)
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

func (a *StreamingServiceManager) CreatePlayListWithMusic(p ...*PlayListWithMusic) error {

	return nil
}

// AtomicBatchCreateEmptyPlayLists : create empty playlists but they must be an all or nothing
func (a *StreamingServiceManager) AtomicBatchCreateEmptyPlayLists(p ...*PlayList) error {
	var wg errgroup.Group
	wg.SetLimit(a.ConcurrentLimit)
	for _, v := range p {
		wg.Go(func(plist *PlayList) func() error {
			return func() error {
				err := a.PCreate(plist)
				if err != nil {
					return &PlayListCreateFailed{
						Item:  plist,
						OgErr: err,
					}
				}
				return nil
			}
		}(v))
	}
	return wg.Wait()
}

func (a *StreamingServiceManager) BatchCreateEmptyPlayLists(p ...*PlayList) error {
	var (
		wg       errgroup.Group
		finalErr error
		errList  []error = make([]error, len(p))
	)
	wg.SetLimit(a.ConcurrentLimit)
	for index, v := range p {

		wg.Go(func(idx int, plist *PlayList) func() error {
			return func() error {
				err := a.PCreate(plist)
				if err != nil {
					errList[idx] = &PlayListCreateFailed{
						Item:  plist,
						OgErr: err,
					}
				}
				return nil
			}
		}(index, v))
	}
	wg.Wait()

	finalErr = multierror.Append(finalErr, errList...)
	return finalErr
}

func (a *StreamingServiceManager) resolveMusicForPlayList(p *PlayList) ([]*Song, error) {
	// this will be running on a machine that will be pay as you go
	//so i don't want to cause an infinite loop and risk taxing myself with a large bill :)
	var tracks []*Song
	var offsetOrToken string
Free:
	for i := 0; i < MaxLargeIteration; i++ {
		if i+1 >= MaxLargeIteration {
			return nil, errors.New("Infinite loop occured")
		}
		res, err := a.
			SGetByPlaylistID(
				p.ID,
				a.
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
