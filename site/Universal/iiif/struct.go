package iiif

type Canvases struct {
	ImgUrls  []string
	IiifUrls []string
	Size     int
}

//Manifest 自动生成 by view-source:https://iiif.lib.harvard.edu/manifests/drs:53262215
type Manifest struct {
	//Context    string      `json:"@context"`
	//Id         string      `json:"@id"`
	//Type       string      `json:"@type"`
	//Label      string      `json:"label"`
	//License    string      `json:"license"`
	//Logo       string      `json:"logo"`
	Sequences []Sequence `json:"sequences"`
	//Structures []Structure `json:"structures"`
}

type Sequence struct {
	//Id               string    `json:"@id"`
	//Type             string    `json:"@type"`
	Canvases []Canvase `json:"canvases"`
	//ViewingDirection string    `json:"viewingDirection"`
	//ViewingHint      string    `json:"viewingHint"`
}

type Canvase struct {
	//Id        string    `json:"@id"`
	//Type      string    `json:"@type"`
	//Height    int       `json:"height"`
	Images []Image `json:"images"`
	//Label     string    `json:"label"`
	//Thumbnail Thumbnail `json:"thumbnail"`
	//Width     int       `json:"width"`
}

type Thumbnail struct {
	Id   string `json:"@id"`
	Type string `json:"@type"`
}

type Image struct {
	Id   string `json:"@id"`
	Type string `json:"@type"`
	//Motivation string   `json:"motivation"`
	On       string   `json:"on"`
	Resource Resource `json:"resource"`
}

type Resource struct {
	Id      string  `json:"@id"`
	Type    string  `json:"@type"`
	Format  string  `json:"format"`
	Height  int     `json:"height"`
	Service Service `json:"service"`
	Width   int     `json:"width"`
}

type Service struct {
	//Context string `json:"@context"`
	Id string `json:"@id"`
	//Profile string `json:"profile"`
}

type Structure struct {
	Id          string   `json:"@id"`
	Type        string   `json:"@type"`
	Label       string   `json:"label"`
	Ranges      []string `json:"ranges,omitempty"`
	ViewingHint string   `json:"viewingHint,omitempty"`
	Canvases    []string `json:"canvases,omitempty"`
}
