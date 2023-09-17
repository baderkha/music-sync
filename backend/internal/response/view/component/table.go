package component

import (
	"encoding/json"

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
	LastRun string `header:"Last Run"`
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
		keyMap      map[string]string = make(map[string]string)
	)
	inrec, _ := json.Marshal(rows)
	json.Unmarshal(inrec, &inInterface)

	if len(inInterface) > 0 {
		ref := reflector.New(rows[0])
		fields := ref.FieldsFlattened()
		for _, v := range fields {
			res, err := v.Tag("header")
			if err != nil || res == "" {
				keys = append(keys, v.Name())
				continue
			}
			keys = append(keys, res)
			keyMap[v.Name()] = res

		}

		if len(keyMap) > 0 {
			for _, mp := range inInterface {
				for ogKey, NewKey := range keyMap {
					val := mp[ogKey]
					mp[NewKey] = val
				}
			}
		}

	}
	return &Table{
		Data:    inInterface,
		Columns: keys,
	}
}
