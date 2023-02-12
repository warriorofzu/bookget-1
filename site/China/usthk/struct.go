package usthk

type Canvases struct {
	ImgUrls *[]string
	Size    int
}

type Result struct {
	FileList []string `json:"file_list"`
}
