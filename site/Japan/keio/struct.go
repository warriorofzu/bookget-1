package keio

type Manifest struct {
	Context  string `json:"@context"`
	Id       string `json:"@id"`
	Type     string `json:"@type"`
	Label    string `json:"label"`
	Metadata []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"metadata"`
	Description      string   `json:"description"`
	Attribution      []string `json:"attribution"`
	License          string   `json:"license"`
	ViewingDirection string   `json:"viewingDirection"`
	Within           string   `json:"within"`
	Sequences        []struct {
		Id       string `json:"@id"`
		Type     string `json:"@type"`
		Canvases []struct {
			Id     string `json:"@id"`
			Type   string `json:"@type"`
			Label  string `json:"label"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
			Images []struct {
				Id         string `json:"@id"`
				Type       string `json:"@type"`
				Motivation string `json:"motivation"`
				On         string `json:"on"`
				Resource   struct {
					Id      string `json:"@id"`
					Type    string `json:"@type"`
					Format  string `json:"format"`
					Service struct {
						Context string `json:"@context"`
						Id      string `json:"@id"`
						Profile string `json:"profile"`
					} `json:"service"`
				} `json:"resource"`
			} `json:"images"`
		} `json:"canvases"`
	} `json:"sequences"`
}
