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

type Tbl struct {
	ActionPostLink     string
	BooleanBtnName     string
	ActionTitle        string
	TableTitle         string
	Columns            []string
	Data               []map[string]any
	ActionButtonHidden bool
}

func (t *Tbl) WithActionButtonHidden(b bool) *Tbl {
	t.ActionButtonHidden = b
	return t
}

func (t *Tbl) WithActionTitle(title string) *Tbl {
	t.ActionTitle = title
	return t
}

func (t *Tbl) WithBooleanBtnName(name string) *Tbl {
	t.BooleanBtnName = name
	return t
}

func (t *Tbl) WithActionPostLink(link string) *Tbl {
	t.ActionPostLink = link
	return t
}

func (t *Tbl) WithTableTitle(title string) *Tbl {
	t.TableTitle = title
	return t
}

func (h *Tbl) GetTemplate() string {
	return "tables.html"
}

func NewTable[T any](rows []T) *Tbl {
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
	return &Tbl{
		Data:    inInterface,
		Columns: keys,
	}
}
