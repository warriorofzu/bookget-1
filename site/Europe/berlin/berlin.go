package berlin

import (
	"bookget/config"
	curl2 "bookget/lib/curl"
	util2 "bookget/lib/util"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`PPN=([A-Za-z0-9_-]+)`).FindStringSubmatch(taskUrl)
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

	var canvases Canvases
	for i := 1; i < 10000; i++ {
		sortId := util2.GenNumberSorted(i)
		id := fmt.Sprintf("PHYS_%s", sortId)
		actUrl := singleDziUrl(bookId, id)
		bs, err := curl2.Get(actUrl, nil)
		if err != nil {
			break
		}
		dziUrl := string(bs)
		if dziUrl == "" || !strings.Contains(dziUrl, "/dzi/") {
			break
		}
		fmt.Printf("\rTest page %d ", i)
		canvases.IiifUrls = append(canvases.IiifUrls, dziUrl)
		imgUrl := singleImageUrl(bookId, sortId)
		canvases.ImgUrls = append(canvases.ImgUrls, imgUrl)
	}
	fmt.Println()
	canvases.Size = len(canvases.ImgUrls)

	//用户自定义起始页
	i := util2.LoopIndexStart(canvases.Size)
	log.Printf("A total of %d pages.\n", canvases.Size)
	destPath := config.CreateDirectory(taskUrl, bookId)
	util2.CreateShell(destPath, canvases.IiifUrls, nil)
	for ; i < canvases.Size; i++ {
		uri := canvases.ImgUrls[i] //从0开始
		if uri == "" {
			continue
		}
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)

		fileName := sortId + ".jpg"
		dest := config.GetDestPath(taskUrl, bookId, fileName)
		curl2.FastGet(uri, dest, nil, true)
	}
	return

}

func singleImageUrl(bookId, id string) string {
	uri := fmt.Sprintf("https://content.staatsbibliothek-berlin.de/dms/%s/full/0/0000%s.jpg?original=true",
		bookId, id)
	return uri
}

func singleDziUrl(bookId, id string) string {
	uri := fmt.Sprintf("https://content.staatsbibliothek-berlin.de/?action=metsImage&metsFile=%s&divID=%s&dzi=true", bookId, id)
	//uri := fmt.Sprintf("https://ngcs-core.staatsbibliothek-berlin.de/dzi/%s/%s.dzi", bookId, id)
	return uri
}
