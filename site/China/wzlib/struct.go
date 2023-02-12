package wzlib

type Result []item

type item struct {
	Items []struct {
		Id          string `json:"_id"`
		DcPublisher string `json:"dc_publisher"`
		DcTitle     string `json:"dc_title"`
		WzlPdfUrl   string `json:"wzl_pdf_url"`
	} `json:"items"`
	Title string `json:"title"`
}

type PdfUrls []PdfUrl
type PdfUrl struct {
	Url  string
	Name string
}

type ResultPdf struct {
	Data struct {
		Id         string `json:"_id"`
		DcTitle    string `json:"dc_title"`
		ModelId    string `json:"model_id"`
		RelateName string `json:"relate_name"`
		WzlPdfUrl  string `json:"wzl_pdf_url"`
	} `json:"Data"`

	RelateList []interface{} `json:"RelateList"`
}
