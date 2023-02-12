package niiac

type Manifest struct {
	Thumbnail struct {
		Id string `json:"@id"`
	} `json:"thumbnail"`
	License string `json:"license"`
	Label   []struct {
		Value    string `json:"@value"`
		Language string `json:"@language"`
	} `json:"label"`
	Attribution string `json:"attribution"`
	Context     string `json:"@context"`
	Within      string `json:"within"`
	Id          string `json:"@id"`
	Type        string `json:"@type"`
	Metadata    []struct {
		Value []struct {
			Value    string `json:"@value"`
			Language string `json:"@language"`
		} `json:"value"`
		Label string `json:"label"`
	} `json:"metadata"`
	Sequences []struct {
		Type     string `json:"@type"`
		Id       string `json:"@id"`
		Canvases []struct {
			Label  string `json:"label"`
			Width  int    `json:"width"`
			Id     string `json:"@id"`
			Type   string `json:"@type"`
			Images []struct {
				Id       string `json:"@id"`
				Resource struct {
					Height  int    `json:"height"`
					Type    string `json:"@type"`
					Format  string `json:"format"`
					Id      string `json:"@id"`
					Service struct {
						Profile string `json:"profile"`
						Context string `json:"@context"`
						Id      string `json:"@id"`
					} `json:"service"`
					Width int `json:"width"`
				} `json:"resource"`
				Motivation string `json:"motivation"`
				On         string `json:"on"`
				Type       string `json:"@type"`
			} `json:"images"`
			Height int `json:"height"`
		} `json:"canvases"`
		Label []struct {
			Language string `json:"@language"`
			Value    string `json:"@value"`
		} `json:"label"`
	} `json:"sequences"`
	ViewingHint      string `json:"viewingHint"`
	ViewingDirection string `json:"viewingDirection"`
	Logo             string `json:"logo"`
	Description      []struct {
		Value    string `json:"@value"`
		Language string `json:"@language"`
	} `json:"description"`
	Related []struct {
		Format string `json:"format"`
		Id     string `json:"@id"`
	} `json:"related"`
}
