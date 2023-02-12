package utokyo

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
	m := regexp.MustCompile(`nu=([A-Za-z0-9]+)`).FindStringSubmatch(taskUrl)
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

	bookUrls := getMultiplebooks(bookId, taskUrl)
	if bookUrls == nil || len(bookUrls) == 0 {
		return
	}
	//用户自定义起始页
	size := len(bookUrls)
	i := util2.LoopIndexStart(size)
	for ; i < size; i++ {
		uri := bookUrls[i] //从0开始
		if uri == "" {
			continue
		}
		//ext := util.FileExt(uri)
		fName := util2.FileName(uri)
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		fileName := sortId + fName
		dest := config.GetDestPath(taskUrl, bookId, fileName)
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

func getMultiplebooks(bookId string, taskUrl string) (bookUrls []string) {
	bs, err := curl.Get(taskUrl, nil)
	if err != nil {
		return
	}
	text := string(bs)
	//取册数
	matches := regexp.MustCompile(`<a href="pdf/([^"]+)"`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}
	bookUrls = make([]string, 0, len(matches))
	for _, v := range matches {
		uri := fmt.Sprintf("http://shanben.ioc.u-tokyo.ac.jp/pdf/%s", v[1])
		bookUrls = append(bookUrls, uri)
	}
	return
}
