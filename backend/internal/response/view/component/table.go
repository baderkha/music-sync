package component

import "encoding/json"

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

type Table struct {
	ActionPostLink     string
	BooleanBtnName     string
	ActionTitle        string
	TableTitle         string
	Columns            []string
	Data               []map[string]any
	ActionButtonHidden bool
}

func (t *Table) WithActionButtonHidden(b bool) *Table {
	t.ActionButtonHidden = b
	return t
}

func (t *Table) WithActionTitle(title string) *Table {
	t.ActionTitle = title
	return t
}

func (t *Table) WithBooleanBtnName(name string) *Table {
	t.BooleanBtnName = name
	return t
}

func (t *Table) WithActionPostLink(link string) *Table {
	t.ActionPostLink = link
	return t
}

func (t *Table) WithTableTitle(title string) *Table {
	t.TableTitle = title
	return t
}

func (h *Table) GetTemplate() string {
	return "tables.html"
}

func NewTable[T any](rows []T) *Table {
	var (
		inInterface []map[string]interface{}
		keys        []string
	)
	inrec, _ := json.Marshal(rows)
	json.Unmarshal(inrec, &inInterface)

	if len(inInterface) > 0 {
		for k := range inInterface[0] {
			keys = append(keys, k)
		}
	}
	return &Table{
		Data:    inInterface,
		Columns: keys,
	}
}
