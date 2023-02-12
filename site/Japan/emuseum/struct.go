package emuseum

type Manifest struct {
	Context string `json:"@context"`
	Id      string `json:"@id"`
	Type    string `json:"@type"`
	Label   []struct {
		Value    string `json:"@value"`
		Language string `json:"@language"`
	} `json:"label"`
	Metadata []struct {
		Label []struct {
			Value    string `json:"@value"`
			Language string `json:"@language"`
		} `json:"label"`
		Value []struct {
			Value    string `json:"@value"`
			Language string `json:"@language"`
		} `json:"value"`
	} `json:"metadata"`
	Description []struct {
		Value    string `json:"@value"`
		Language string `json:"@language"`
	} `json:"description"`
	Thumbnail struct {
		Id     string `json:"@id"`
		Type   string `json:"@type"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"thumbnail"`
	ViewingHint      string `json:"viewingHint"`
	ViewingDirection string `json:"viewingDirection"`
	License          string `json:"license"`
	Attribution      []struct {
		Value    string `json:"@value"`
		Language string `json:"@language"`
	} `json:"attribution"`
	Logo []struct {
		Id string `json:"@id"`
	} `json:"logo"`
	Rendering struct {
		Id    string `json:"@id"`
		Label []struct {
			Value    string `json:"@value"`
			Language string `json:"@language"`
		} `json:"label"`
		Format string `json:"format"`
	} `json:"rendering"`
	Within struct {
		Id    string `json:"@id"`
		Type  string `json:"@type"`
		Label []struct {
			Value    string `json:"@value"`
			Language string `json:"@language"`
		} `json:"label"`
	} `json:"within"`
	Sequences []struct {
		Type     string `json:"@type"`
		Canvases []struct {
			Id     string `json:"@id"`
			Type   string `json:"@type"`
			Label  string `json:"label"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
			Images []struct {
				Type       string `json:"@type"`
				Context    string `json:"@context"`
				Motivation string `json:"motivation"`
				Resource   struct {
					Id      string `json:"@id"`
					Type    string `json:"@type"`
					Format  string `json:"format"`
					Service struct {
						Context string        `json:"@context"`
						Id      string        `json:"@id"`
						Profile []interface{} `json:"profile"`
					} `json:"service"`
					Height int `json:"height"`
					Width  int `json:"width"`
				} `json:"resource"`
				On string `json:"on"`
			} `json:"images"`
			Thumbnail struct {
				Id   string `json:"@id"`
				Type string `json:"@type"`
			} `json:"thumbnail"`
		} `json:"canvases"`
	} `json:"sequences"`
}
