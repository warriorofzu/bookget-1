package luoyang

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/gohttp"
	util2 "bookget/lib/util"
	"fmt"
	"log"
	"regexp"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`&id=(\d+)`).FindStringSubmatch(taskUrl)
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

	pdfUrls := getMultiplebooks(taskUrl)
	size := len(pdfUrls)
	if pdfUrls == nil || size == 0 {
		return
	}
	log.Printf("A total of %d PDFs.\n", size)
	//用户自定义起始页
	i := util2.LoopIndexStart(size)
	for ; i < size; i++ {
		uri := pdfUrls[i]
		if uri == "" {
			continue
		}
		ext := util2.FileExt(uri)
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		filename := sortId + ext
		dest := config.GetDestPath(taskUrl, bookId, filename)
		gohttp.FastGet(uri, gohttp.Options{
			DestFile:    dest,
			Overwrite:   false,
			Concurrency: config.Conf.Threads,
			Headers: map[string]interface{}{
				"user-agent": config.UserAgent,
			},
		})
	}
}

func getMultiplebooks(bookUrl string) (pdfUrls []string) {
	bs, err := curl.Get(bookUrl, nil)
	if err != nil {
		return
	}
	text := string(bs)
	//取册数
	matches := regexp.MustCompile(`href=["']viewer.php\?pdf=(.+?)\.pdf&`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}
	ids := make([]string, 0, len(matches))
	for _, match := range matches {
		ids = append(ids, match[1])
	}

	hostUrl := util2.GetHostUrl(bookUrl)
	pdfUrls = make([]string, 0, len(ids))
	for _, v := range ids {
		s := fmt.Sprintf("%s%s.pdf", hostUrl, v)
		pdfUrls = append(pdfUrls, s)
	}
	return
}
