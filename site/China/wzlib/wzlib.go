package wzlib

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/gohttp"
	util2 "bookget/lib/util"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`\?id=([A-z\d]+)`).FindStringSubmatch(taskUrl)
	if m != nil {
		bookId = m[1]
		config.CreateDirectory(taskUrl, bookId)
		StartDownload(iTask, taskUrl, bookId)
	}
	return "", err
}

func StartDownload(iTask int, taskUrl, bookId string) {
	name := util2.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, taskUrl)

	pdfUrls := PdfUrls{}
	pdfUrl, err := getBook(bookId)
	if err == nil {
		pdfUrls = append(pdfUrls, pdfUrl)
	} else {
		pdfUrls = getMultiplebooks(bookId)
	}
	size := len(pdfUrls)
	if size == 0 {
		return
	}
	log.Printf("A total of %d PDFs.\n", size)
	//用户自定义起始页
	i := util2.LoopIndexStart(size)
	ext := ".pdf"
	for ; i < size; i++ {
		pdfUrl := pdfUrls[i]
		if pdfUrl.Url == "" {
			continue
		}
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, pdfUrl.Url)
		fileName := pdfUrl.Name + ext
		dest := config.GetDestPath(taskUrl, bookId, fileName)
		gohttp.FastGet(pdfUrl.Url, gohttp.Options{
			DestFile:    dest,
			Overwrite:   false,
			Concurrency: config.Conf.Threads,
			Headers: map[string]interface{}{
				"user-agent": config.UserAgent,
			},
		})
	}
}

func getBook(bookId string) (pdf PdfUrl, err error) {
	uri := fmt.Sprintf("https://oyjy.wzlib.cn/api/search/v1/resource/%s", bookId)
	bs, err := curl.Get(uri, nil)
	if err != nil {
		return
	}
	var result ResultPdf
	if err = json.Unmarshal(bs, &result); err != nil {
		return
	}
	if result.Data.WzlPdfUrl == "" {
		return pdf, errors.New("Not found pdfUrl.")
	}
	m := regexp.MustCompile(`file=(\S+)`).FindStringSubmatch(result.Data.WzlPdfUrl)
	if m == nil {
		return pdf, errors.New("Not found pdfUrl.")
	}
	pdf.Url = fmt.Sprintf("https://db.wzlib.cn%s", m[1])
	pdf.Name = result.Data.DcTitle
	return pdf, nil
}

func getMultiplebooks(bookId string) (pdfUrls PdfUrls) {
	relatedUri := fmt.Sprintf("https://oyjy.wzlib.cn/api/search/v1/resource_related/%s", bookId)
	bs, err := curl.Get(relatedUri, nil)
	if err != nil {
		return
	}
	var result Result
	if err = json.Unmarshal(bs, &result); err != nil {
		return
	}
	pdfUrls = make([]PdfUrl, 0, len(result[0].Items))
	for _, v := range result[0].Items {
		if v.WzlPdfUrl == "" {
			continue
		}
		var pdfUrl PdfUrl
		pdfUrl.Url = fmt.Sprintf("https://db.wzlib.cn%s", v.WzlPdfUrl)
		pdfUrl.Name = v.DcTitle
		pdfUrls = append(pdfUrls, pdfUrl)
	}

	return
}
