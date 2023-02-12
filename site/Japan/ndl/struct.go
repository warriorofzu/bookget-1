package ndl

type Manifest struct {
	Context  string `json:"@context"`
	Type     string `json:"@type"`
	Id       string `json:"@id"`
	Label    string `json:"label"`
	Metadata []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"metadata"`
	ViewingDirection string `json:"viewingDirection"`
	License          string `json:"license"`
	Attribution      string `json:"attribution"`
	Logo             string `json:"logo"`
	SeeAlso          string `json:"seeAlso"`
	Sequences        []struct {
		Type        string `json:"@type"`
		ViewingHint string `json:"viewingHint"`
		Thumbnail   struct {
			Id     string `json:"@id"`
			Type   string `json:"@type"`
			Format string `json:"format"`
		} `json:"thumbnail"`
		Canvases []struct {
			Id        string `json:"@id"`
			Type      string `json:"@type"`
			Label     string `json:"label"`
			Width     int    `json:"width"`
			Height    int    `json:"height"`
			Thumbnail struct {
				Id     string `json:"@id"`
				Type   string `json:"@type"`
				Format string `json:"format"`
			} `json:"thumbnail"`
			Images []struct {
				Type       string `json:"@type"`
				Motivation string `json:"motivation"`
				On         string `json:"on"`
				Resource   struct {
					Id      string `json:"@id"`
					Type    string `json:"@type"`
					Format  string `json:"format"`
					Width   int    `json:"width"`
					Height  int    `json:"height"`
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
