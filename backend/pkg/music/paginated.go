package music

import "strconv"

// OptionallyPaginatedResult : resources that can be paginated are listed here
// Make sure to access the isPaginated Element
type OptionallyPaginatedResult[T any] struct {
	Data         []*T
	TotalRecords int
	IsPaginated  bool
	NextPage     string
}

type PaginatedRequester interface {
	GetLimit() int
	GetOffsetOrToken() string
}

type PaginatedLimitOffsetReq struct {
	Offset int
	Limit  int
}

func (p *PaginatedLimitOffsetReq) GetOffsetOrToken() string {
	return strconv.Itoa(p.Offset)
}

func (p *PaginatedLimitOffsetReq) GetLimit() int {
	return p.Limit
}

type PaginatedTokenReq struct {
	Token string
	Limit int
}

func (p *PaginatedTokenReq) GetLimit() int {
	return p.Limit
}

func (p *PaginatedTokenReq) GetOffsetOrToken() string {
	return p.Token
}
