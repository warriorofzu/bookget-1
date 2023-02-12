package npmtw

type Canvases struct {
	ImgUrls  []string
	IiifUrls []string
	Size     int
}

type Manifest struct {
	Context  string `json:"@context"`
	Id       string `json:"@id"`
	Type     string `json:"@type"`
	Label    string `json:"label"`
	Logo     string `json:"logo"`
	Metadata []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"metadata"`
	Rendering []struct {
		Id     string `json:"@id"`
		Format string `json:"format"`
		Label  string `json:"label"`
	} `json:"rendering"`
	License          string      `json:"license"`
	ViewingDirection string      `json:"viewingDirection"`
	ViewingHint      interface{} `json:"viewingHint"`
	Sequences        []struct {
		Type     string `json:"@type"`
		Canvases []struct {
			Label  string `json:"label"`
			Id     string `json:"@id"`
			Type   string `json:"@type"`
			Height string `json:"height"`
			Width  string `json:"width"`
			Images []struct {
				Id         string `json:"@id"`
				Type       string `json:"@type"`
				Motivation string `json:"motivation"`
				On         string `json:"on"`
				Resource   struct {
					Description string `json:"description"`
					Id          string `json:"@id"`
					Type        string `json:"@type"`
					Format      string `json:"format"`
					Service     struct {
						Id      string `json:"@id"`
						Context string `json:"@context"`
						Profile string `json:"profile"`
					} `json:"service"`
				} `json:"resource"`
			} `json:"images"`
			OtherContent []struct {
				Id   string `json:"@id"`
				Type string `json:"@type"`
			} `json:"otherContent"`
		} `json:"canvases"`
	} `json:"sequences"`
}
