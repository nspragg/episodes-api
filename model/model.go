package model

/* Model definition */
type Root struct {
	Version  string    `json:"version"`
	Schema   string    `json:"schema"`
	Episodes []Episode `json:"episodes"`
}

type MasterBrand struct {
	Id     string `json:"id"`
	Titles struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"titles"`
	Attribution string `json:"attribution"`
	IdentId     string `json:"ident_id"`
}

type Episode struct {
	Id                string `json:"id"`
	Type              string `json:"type"`
	LexicalSortLetter string `json:"lexical_sort_letter"`
	Title             string `json:"title"`
	Subtitle          string `json:"subtitle"`
	Synopses          struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"synopses"`
	Images struct {
		Standard string `json:"standard"`
		Type     string `json:"type"`
	} `json:"images"`
	TleoId       string        `json:"tleo_id"`
	TleoType     string        `json:"tleo_type"`
	ParentId     string        `json:"parent_id"`
	Guidance     bool          `json:"guidance,omitempty"`
	ReleaseDate  string        `json:"release_date"`
	Masterbrand  MasterBrand   `json:"master_brand,omitempty"`
	Status       string        `json:"status"`
	Versions     []Version     `json:"versions"`
	RelatedLinks []interface{} `json:"related_links"`
}

type Version struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	Kind         string `json:"kind"`
	Availability struct {
		Remaining struct {
			Text string `json:"text"`
		} `json:"remaining"`
		End   string `json:"end"`
		Start string `json:"start"`
	} `json:"availability"`
	Hd       bool `json:"hd"`
	Duration struct {
		Value string `json:"value"`
		Text  string `json:"text"`
	} `json:"duration"`
	FirstBroadcast string `json:"first_broadcast"`
	Guidance       struct {
		Id   string `json:"id"`
		Text struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
			Large  string `json:"large"`
		} `json:"text"`
	} `json:"guidance"`
}
