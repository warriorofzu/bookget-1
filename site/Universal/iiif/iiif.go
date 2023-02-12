package iiif

import (
	"bookget/config"
	"bookget/lib/gohttp"
	util2 "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	taskName := util2.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", taskName, taskUrl)

	bookId := getBookId(taskUrl)
	config.CreateDirectory(taskUrl, bookId)

	StartDownload(taskUrl, bookId)
	return "", err
}

func getBookId(taskUrl string) string {
	m := regexp.MustCompile(`/([^/]+)/manifest.json`).FindStringSubmatch(taskUrl)
	if m != nil {
		return m[1]
	}
	return ""
}

func StartDownload(pageUrl, bookId string) {
	canvases := getCanvases(pageUrl)
	if canvases.Size == 0 {
		return
	}
	log.Printf("A total of %d pages.\n", canvases.Size)

	destPath := config.CreateDirectory(pageUrl, bookId)
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
		filename := sortId + ext
		dest := config.GetDestPath(pageUrl, bookId, filename)
		cli := gohttp.NewClient(gohttp.Options{
			DestFile:   dest,
			CookieFile: config.Conf.CookieFile,
			Headers: map[string]interface{}{
				"User-Agent": config.Conf.UserAgent,
			},
		})
		_, err := cli.Get(uri)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCanvases(uri string) (canvases Canvases) {
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
	})
	resp, err := cli.Get(uri)
	if err != nil {
		return
	}
	bs, _ := resp.GetBody()
	var manifest = new(Manifest)
	if err = json.Unmarshal(bs, manifest); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	if len(manifest.Sequences) == 0 {
		return
	}
	newWidth := ""
	//>6400使用原图
	if config.Conf.FullImageWidth > 6400 {
		newWidth = "full/full"
	} else if config.Conf.FullImageWidth >= 1000 {
		newWidth = fmt.Sprintf("full/%d,", config.Conf.FullImageWidth)
	}

	size := len(manifest.Sequences[0].Canvases)
	canvases.ImgUrls = make([]string, 0, size)
	canvases.IiifUrls = make([]string, 0, size)
	for _, canvase := range manifest.Sequences[0].Canvases {
		for _, image := range canvase.Images {
			//iifUrl, _ := url.QueryUnescape(image.Resource.Service.Id)
			//dezoomify-rs URL
			iiiInfo := fmt.Sprintf("%s/info.json", image.Resource.Service.Id)
			canvases.IiifUrls = append(canvases.IiifUrls, iiiInfo)

			//JPEG URL
			imgUrl := fmt.Sprintf("%s/%s/0/default.jpg", image.Resource.Service.Id, newWidth)
			canvases.ImgUrls = append(canvases.ImgUrls, imgUrl)
		}
	}
	canvases.Size = size
	return
}
