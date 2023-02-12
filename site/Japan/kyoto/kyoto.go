package kyoto

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/gohttp"
	util2 "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`item/([A-Za-z0-9]+)`).FindStringSubmatch(taskUrl)
	if m != nil {
		bookId = m[1]
		config.CreateDirectory(taskUrl, bookId)
		StartDownload(iTask, taskUrl, bookId)
	}
	return "", err
}

func StartDownload(iTask int, text, bookId string) {
	name := util2.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, text)

	pages, iiifInfo := getPages(bookId)
	log.Printf("A total of %d pages.\n", len(pages))

	destPath := config.CreateDirectory(text, bookId)
	util2.CreateShell(destPath, iiifInfo, nil)
	//用户自定义起始页
	size := len(pages)
	i := util2.LoopIndexStart(size)
	for ; i < size; i++ {
		uri := pages[i] //从0开始
		if uri == "" {
			continue
		}
		ext := util2.FileExt(uri)
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		fileName := sortId + ext
		dest := config.GetDestPath(text, bookId, fileName)
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

func getPages(bookId string) (pages []string, iiifInfo []string) {
	var manifests = new(ManifestsJson)
	bs, err := curl.Get(fmt.Sprintf("https://rmda.kulib.kyoto-u.ac.jp/iiif/metadata_manifest/%s/manifest.json", bookId), nil)
	if err != nil {
		return
	}
	if err = json.Unmarshal(bs, manifests); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	if len(manifests.Sequences) == 0 {
		return
	}

	i := len(manifests.Sequences[0].Canvases)
	pages = make([]string, 0, i)
	newWidth := config.Conf.FullImageWidth
	//此站最大只支持3000
	if newWidth < 1000 || newWidth > 3000 {
		newWidth = 3000
	}
	for _, sequence := range manifests.Sequences {
		for _, canvase := range sequence.Canvases {
			for _, image := range canvase.Images {
				//dezoomify-rs URL
				iiifUrl := fmt.Sprintf("%s/info.json", image.Resource.Service.Id)
				iiifInfo = append(iiifInfo, iiifUrl)
				//JPEG URL
				imgurl := fmt.Sprintf("%s/full/%d,/0/default.jpg", image.Resource.Service.Id, newWidth)
				pages = append(pages, imgurl)
			}
		}
	}
	return
}
