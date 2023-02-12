package khirin

type Canvases struct {
	ImgUrls  []string
	IiifUrls []string
	Size     int
}

type Manifest struct {
	Label    string `json:"label"`
	Metadata []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"metadata"`
	Id               string `json:"@id"`
	Type             string `json:"@type"`
	ViewingDirection string `json:"viewingDirection"`
	Context          string `json:"@context"`
	Within           string `json:"within"`
	Sequences        []struct {
		Type     string `json:"@type"`
		Canvases []struct {
			Id     string `json:"@id"`
			Type   string `json:"@type"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
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
		} `json:"canvases"`
	} `json:"sequences"`
}

type Manifest2 struct {
	Context          string `json:"@context"`
	Id               string `json:"@id"`
	Type             string `json:"@type"`
	ViewingDirection string `json:"viewingDirection"`
	ViewingHint      string `json:"viewingHint"`
	Sequences        []struct {
		Id       string `json:"@id"`
		Type     string `json:"@type"`
		Canvases []struct {
			Id     string `json:"@id"`
			Type   string `json:"@type"`
			Label  string `json:"label"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
			Images []struct {
				Id         string `json:"@id"`
				On         string `json:"on"`
				Type       string `json:"@type"`
				Format     string `json:"format"`
				Motivation string `json:"motivation"`
				Resource   struct {
					Id      string `json:"@id"`
					Type    string `json:"@type"`
					On      string `json:"on"`
					Format  string `json:"format"`
					Height  int    `json:"height"`
					Width   int    `json:"width"`
					Service struct {
						Id      string `json:"@id"`
						Profile string `json:"profile"`
						Context string `json:"@context"`
					} `json:"service"`
				} `json:"resource"`
			} `json:"images"`
		} `json:"canvases"`
	} `json:"sequences"`
	Logo     string `json:"logo"`
	Metadata []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"metadata"`
	Label       string `json:"label"`
	Description string `json:"description"`
	License     string `json:"license"`
	Attribution string `json:"attribution"`
}
