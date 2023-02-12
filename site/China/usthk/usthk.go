package usthk

import (
	"bookget/config"
	curl2 "bookget/lib/curl"
	util2 "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`bib/([A-z0-9]+)`).FindStringSubmatch(taskUrl)
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
	canvases := getImageUrls(taskUrl)
	log.Printf("A total of %d pages.\n", canvases.Size)
	if canvases.ImgUrls == nil {
		return
	}
	//用户自定义起始页
	i := util2.LoopIndexStart(canvases.Size)
	for ; i < canvases.Size; i++ {
		uri := (*canvases.ImgUrls)[i] //从0开始
		if uri == "" {
			continue
		}
		ext := util2.FileExt(uri)
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)

		fileName := sortId + ext
		dest := config.GetDestPath(taskUrl, bookId, fileName)
		curl2.FastGet(uri, dest, nil, true)
	}

	return
}

func getImageUrls(taskUrl string) (canvases Canvases) {
	bs, err := curl2.Get(taskUrl, nil)
	if err != nil {
		return
	}
	text := string(bs)

	//view_book('6/o/b1129168/ebook'
	matches := regexp.MustCompile(`view_book\(["'](\S+)["']`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}
	imgUrls := make([]string, 0, 1000)
	for _, m := range matches {
		sPath := m[1]
		uri := fmt.Sprintf("https://lbezone.ust.hk/bookreader/getfilelist.php?path=%s", sPath)
		bs, err = curl2.Get(uri, nil)
		if err != nil {
			return
		}
		result := new(Result)
		if err = json.Unmarshal(bs, result); err != nil {
			log.Printf("json.Unmarshal failed: %s\n", err)
			return
		}
		//imgUrls := make([]string, 0, len(result.FileList))
		for _, v := range result.FileList {
			imgUrl := fmt.Sprintf("https://lbezone.ust.hk/obj/%s/%s", sPath, v)
			imgUrls = append(imgUrls, imgUrl)
		}
	}
	canvases.ImgUrls = &imgUrls
	canvases.Size = len(imgUrls)
	return
}
