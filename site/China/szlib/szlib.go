package szlib

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/gohttp"
	util2 "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`book_id=([A-z\d]+)`).FindStringSubmatch(taskUrl)
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

	rstVolumes := getMultiplebooks(bookId)
	if rstVolumes == nil {
		return
	}
	log.Printf("A total of %d volumes.\n", len(rstVolumes.Volumes))
	for j, item := range rstVolumes.Volumes {
		fmt.Printf("\r Test volume %d ... ", j+1)
		canvases := getCanvases(bookId, item)
		if canvases.ImgUrls == nil || canvases.Size == 0 {
			return
		}
		fmt.Println()
		log.Printf("A total of %d pages.\n", canvases.Size)

		//用户自定义起始页
		//i := util.LoopIndexStart(canvases.Size)
		for i := 0; i < canvases.Size; i++ {
			imgUrl := canvases.ImgUrls[i]
			if imgUrl == "" {
				continue
			}
			ext := util2.FileExt(imgUrl)
			sortId := util2.GenNumberSorted(i + 1)
			log.Printf("Get %s  %s\n", sortId, imgUrl)
			fileName := fmt.Sprintf("vol%d_%s%s", j+1, sortId, ext)
			dest := config.GetDestPath(taskUrl, bookId, fileName)
			gohttp.FastGet(imgUrl, gohttp.Options{
				DestFile:    dest,
				Overwrite:   false,
				Concurrency: config.Conf.Threads,
				Headers: map[string]interface{}{
					"user-agent": config.UserAgent,
				},
			})
		}
	}
}

func getMultiplebooks(bookId string) (rstVolumes *ResultVolumes) {
	uri := fmt.Sprintf("https://yun.szlib.org.cn/stgj2021/book_view/%s", bookId)
	bs, err := curl.Get(uri, nil)
	if err != nil {
		return
	}
	rstVolumes = new(ResultVolumes)
	if err = json.Unmarshal(bs, rstVolumes); err != nil {
		return
	}
	return rstVolumes
}

func getCanvases(bookId string, item Directory) (canvases Canvases) {
	p1 := getOnePage(bookId, item.Volume, item.Children[0].Page)
	pos := strings.LastIndex(p1, "/")
	urlPre := p1[:pos]
	ext := util2.FileExt(p1)
	for _, child := range item.Children {
		imgUrl := fmt.Sprintf("%s/%s%s", urlPre, child.Page, ext)
		canvases.ImgUrls = append(canvases.ImgUrls, imgUrl)
	}
	canvases.Size = len(canvases.ImgUrls)
	return
}

func getOnePage(bookId string, volumeId string, page string) (imgUrl string) {
	uri := fmt.Sprintf("https://yun.szlib.org.cn/stgj2021/book_page/%s/%s/%s", bookId, volumeId, page)
	bs, err := curl.Get(uri, nil)
	if err != nil {
		return
	}
	rstPage := new(ResultPage)
	if err = json.Unmarshal(bs, rstPage); err != nil {
		return
	}
	imgUrl = rstPage.BookImageUrl + rstPage.PicInfo.Path
	return imgUrl
}
