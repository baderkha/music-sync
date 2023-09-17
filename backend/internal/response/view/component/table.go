package component

import (
	"github.com/tkrajina/go-reflector/reflector"
)

type PlayListData struct {
	Title   string
	Service string
	Creator string
}

type TrackData struct {
	Title    string
	Duration string
	Service  string
	Artists  []string
}

type SyncData struct {
	Label   string
	LastRun string
	From    string
	To      string
	Status  string
}

type Table[T any] struct {
	ActionPostLink     string
	BooleanBtnName     string
	ActionTitle        string
	TableTitle         string
	Columns            []string
	Data               []T
	ActionButtonHidden bool
}

func (t *Table[T]) WithActionButtonHidden(b bool) *Table[T] {
	t.ActionButtonHidden = b
	return t
}

func (t *Table[T]) WithActionTitle(title string) *Table[T] {
	t.ActionTitle = title
	return t
}

func (t *Table[T]) WithBooleanBtnName(name string) *Table[T] {
	t.BooleanBtnName = name
	return t
}

func (t *Table[T]) WithActionPostLink(link string) *Table[T] {
	t.ActionPostLink = link
	return t
}

func (t *Table[T]) WithTableTitle(title string) *Table[T] {
	t.TableTitle = title
	return t
}

func (h *Table[T]) GetTemplate() string {
	return "tables.html"
}

func NewTable[T any](rows []T) *Table[T] {
	var cols []string
	if len(rows) > 0 {
		refl := reflector.New(rows[0])
		fields := refl.FieldsFlattened()
		cols = make([]string, len(fields))
		for i, v := range fields {
			cols[i] = v.Name()
		}
	}
	return &Table[T]{
		Data:    rows,
		Columns: cols,
	}
}
