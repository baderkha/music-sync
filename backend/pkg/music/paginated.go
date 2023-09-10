package music

type OptionallyPaginatedResult[T any] struct {
	Data         []*T
	TotalRecords int
	IsPaginated  bool
}

type PaginatedRequest struct {
	Offset int
	Limit  int
}
