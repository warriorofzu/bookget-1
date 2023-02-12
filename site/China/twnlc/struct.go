package twnlc

type ResponseToken struct {
	Token string `json:"token"`
}

type Canvases struct {
	ImgUrls []string
	Size    int
}
