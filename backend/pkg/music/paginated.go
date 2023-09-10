package music

type PagintatedResult[T any] struct {
	Data         []*T
	TotalRecords int
	CurrentPage  int
	NextPage     int
	LastPage     int
}

type PaginatedRequest struct {
	Page int
	Size int
}
