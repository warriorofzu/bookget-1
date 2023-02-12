package npmtw

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
	m := regexp.MustCompile(`\?pid=(\d+)`).FindStringSubmatch(taskUrl)
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
	canvases := getImageUrls(bookId, taskUrl)
	if canvases.ImgUrls == nil {
		return
	}
	log.Printf("A total of %d pages.\n", canvases.Size)
	destPath := config.CreateDirectory(taskUrl, bookId)
	util2.CreateShell(destPath, canvases.IiifUrls, nil)
	//用户自定义起始页
	i := util2.LoopIndexStart(canvases.Size)
	for ; i < canvases.Size; i++ {
		uri := canvases.ImgUrls[i] //从0开始
		if uri == "" {
			continue
		}
		ext := util2.FileExt(uri)
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		dest := config.GetDestPath(taskUrl, bookId, sortId+ext)
		curl2.FastGet(uri, dest, nil, true)
	}
}

func getImageUrls(bookId, taskUrl string) (canvases Canvases) {
	var manifest = new(Manifest)
	u := fmt.Sprintf("https://digitalarchive.npm.gov.tw/Painting/setJson?pid=%s&Dept=P", bookId)
	bs, err := curl2.Get(u, nil)
	if err != nil {
		return
	}
	if err = json.Unmarshal(bs, manifest); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	if len(manifest.Sequences) == 0 {
		return
	}

	i := len(manifest.Sequences[0].Canvases)
	canvases.ImgUrls = make([]string, 0, i)
	canvases.IiifUrls = make([]string, 0, i)
	for _, canvase := range manifest.Sequences[0].Canvases {
		for _, image := range canvase.Images {
			u := fmt.Sprintf("%s/info.json", image.Resource.Service.Id)
			canvases.IiifUrls = append(canvases.IiifUrls, u)
			canvases.ImgUrls = append(canvases.ImgUrls, image.Resource.Id)
		}
	}
	canvases.Size = len(canvases.ImgUrls)
	return
}
