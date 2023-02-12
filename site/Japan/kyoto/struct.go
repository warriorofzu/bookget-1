package kyoto

//ManifestsJson 自动生成 from https://rmda.kulib.kyoto-u.ac.jp/iiif/metadata_manifest/RB00024956/manifest.json
type ManifestsJson struct {
	Label    string `json:"label"`
	Metadata []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"metadata"`
	License          string `json:"license"`
	Id               string `json:"@id"`
	Type             string `json:"@type"`
	ViewingDirection string `json:"viewingDirection"`
	Context          string `json:"@context"`
	Within           string `json:"within"`
	Sequences        []struct {
		Type     string `json:"@type"`
		Canvases []struct {
			Label  string `json:"label"`
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
