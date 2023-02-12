package cuhk

type iiifManifest struct {
	Pages []Page `json:"pages"`
}

type Page struct {
	Pid        string `json:"pid"`
	Page       string `json:"page"`
	Label      string `json:"label"`
	Width      string `json:"width"`
	Height     string `json:"height"`
	Dsid       string `json:"dsid"`
	Token      string `json:"token"`
	Identifier string `json:"identifier"`
}
