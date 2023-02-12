package niiac

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
	m := regexp.MustCompile(`toyobunko/([^/]+)/([^/]+)`).FindStringSubmatch(taskUrl)
	if m != nil {
		bookId = fmt.Sprintf("%s.%s", m[1], m[2])
		config.CreateDirectory(taskUrl, bookId)
		StartDownload(iTask, taskUrl, bookId)
	}
	return "", err
}

func StartDownload(iTask int, taskUrl, bookId string) {
	name := util2.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, taskUrl)

	imageUrls, iiifUrls := getImageUrls(bookId, taskUrl)
	if imageUrls == nil || iiifUrls == nil {
		return
	}
	size := len(imageUrls)
	log.Printf("A total of %d pages.\n", size)

	destPath := config.CreateDirectory(taskUrl, bookId)
	util2.CreateShell(destPath, iiifUrls, nil)
	//用户自定义起始页
	i := util2.LoopIndexStart(size)
	for ; i < size; i++ {
		uri := imageUrls[i] //从0开始
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
}

func getImageUrls(bookId, bookUrl string) (imgUrls []string, iiifUrls []string) {
	uri := fmt.Sprintf("%s/manifest.json", bookUrl)
	var manifest = new(Manifest)
	bs, err := curl2.Get(uri, nil)
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
	imgUri := make([]string, 0, i)
	iiifUri := make([]string, 0, i)
	newWidth := ""
	//>6400使用原图
	if config.Conf.FullImageWidth > 6400 {
		newWidth = "/full/full/0/default.jpg"
	} else if config.Conf.FullImageWidth >= 1000 {
		newWidth = fmt.Sprintf("/full/%d,/0/default.jpg", config.Conf.FullImageWidth)
	}
	for _, canvase := range manifest.Sequences[0].Canvases {
		for _, image := range canvase.Images {
			//dezoomify-rs URL
			iiiInfo := fmt.Sprintf("%s/info.json", image.Resource.Service.Id)
			iiifUri = append(iiifUri, iiiInfo)

			//JPEG URL
			imgUrl := ""
			if newWidth == "" {
				imgUrl = image.Resource.Id
			} else {
				imgUrl = fmt.Sprintf("%s%s", image.Resource.Service.Id, newWidth)
			}
			imgUri = append(imgUri, imgUrl)
		}
	}
	return imgUri, iiifUri
}
