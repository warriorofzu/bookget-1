package princeton

//Graphql 查manifestUrl
type Graphql struct {
	Data struct {
		ResourcesByOrangelightIds []struct {
			Id        string      `json:"id"`
			Thumbnail interface{} `json:"thumbnail"`
			Url       string      `json:"url"`
			Members   []struct {
				Id       string `json:"id"`
				Typename string `json:"__typename"`
			} `json:"members"`
			Typename      string `json:"__typename"`
			ManifestUrl   string `json:"manifestUrl"`
			OrangelightId string `json:"orangelightId"`
		} `json:"resourcesByOrangelightIds"`
	} `json:"data"`
}

type Manifest struct {
	Context     string   `json:"@context"`
	Type        string   `json:"@type"`
	Id          string   `json:"@id"`
	Label       []string `json:"label"`
	Description []string `json:"description"`
	ViewingHint string   `json:"viewingHint"`
	Metadata    []struct {
		Label string        `json:"label"`
		Value []interface{} `json:"value"`
	} `json:"metadata"`
	Manifests []struct {
		Context   string   `json:"@context"`
		Type      string   `json:"@type"`
		Id        string   `json:"@id"`
		Label     []string `json:"label"`
		Thumbnail struct {
			Id      string `json:"@id"`
			Service struct {
				Context string `json:"@context"`
				Id      string `json:"@id"`
				Profile string `json:"profile"`
			} `json:"service"`
		} `json:"thumbnail"`
	} `json:"manifests"`
	SeeAlso []struct {
		Id     string `json:"@id"`
		Format string `json:"format"`
	} `json:"seeAlso"`
	License   string `json:"license"`
	Thumbnail struct {
		Id      string `json:"@id"`
		Service struct {
			Context string `json:"@context"`
			Id      string `json:"@id"`
			Profile string `json:"profile"`
		} `json:"service"`
	} `json:"thumbnail"`
	Rendering struct {
		Id     string `json:"@id"`
		Format string `json:"format"`
	} `json:"rendering"`
}

//Manifest 查info.json
type Manifest2 struct {
	Context          string   `json:"@context"`
	Type             string   `json:"@type"`
	Id               string   `json:"@id"`
	Label            []string `json:"label"`
	ViewingHint      string   `json:"viewingHint"`
	ViewingDirection string   `json:"viewingDirection"`
	Metadata         []struct {
		Label string   `json:"label"`
		Value []string `json:"value"`
	} `json:"metadata"`
	Sequences []struct {
		Type      string `json:"@type"`
		Id        string `json:"@id"`
		Rendering []struct {
			Id     string `json:"@id"`
			Label  string `json:"label"`
			Format string `json:"format"`
		} `json:"rendering"`
		Canvases []struct {
			Type      string `json:"@type"`
			Id        string `json:"@id"`
			Label     string `json:"label"`
			Rendering []struct {
				Id     string `json:"@id"`
				Label  string `json:"label"`
				Format string `json:"format"`
			} `json:"rendering"`
			Width  int `json:"width"`
			Height int `json:"height"`
			Images []struct {
				Type       string `json:"@type"`
				Motivation string `json:"motivation"`
				Resource   struct {
					Type    string `json:"@type"`
					Id      string `json:"@id"`
					Height  int    `json:"height"`
					Width   int    `json:"width"`
					Format  string `json:"format"`
					Service struct {
						Context string `json:"@context"`
						Id      string `json:"@id"`
						Profile string `json:"profile"`
					} `json:"service"`
				} `json:"resource"`
				Id string `json:"@id"`
				On string `json:"on"`
			} `json:"images"`
		} `json:"canvases"`
		ViewingHint string `json:"viewingHint"`
	} `json:"sequences"`
	Structures []struct {
		Type        string   `json:"@type"`
		Id          string   `json:"@id"`
		Label       string   `json:"label"`
		ViewingHint string   `json:"viewingHint,omitempty"`
		Ranges      []string `json:"ranges"`
		Canvases    []string `json:"canvases"`
	} `json:"structures"`
	SeeAlso struct {
		Id     string `json:"@id"`
		Format string `json:"format"`
	} `json:"seeAlso"`
	License   string `json:"license"`
	Thumbnail struct {
		Id      string `json:"@id"`
		Service struct {
			Context string `json:"@context"`
			Id      string `json:"@id"`
			Profile string `json:"profile"`
		} `json:"service"`
	} `json:"thumbnail"`
	Rendering struct {
		Id     string `json:"@id"`
		Format string `json:"format"`
	} `json:"rendering"`
	Logo string `json:"logo"`
}
