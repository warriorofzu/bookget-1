package loc

import (
	"bookget/config"
	curl2 "bookget/lib/curl"
	"bookget/lib/gohttp"
	util2 "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http/cookiejar"
	"regexp"
	"strings"
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

func StartDownload(iTask int, taskUrl, bookId string) {
	name := util2.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, taskUrl)

	pages := getPages(bookId)
	size := len(pages)
	log.Printf("A total of %d pages.\n", size)

	//cookie 处理
	jar, _ := cookiejar.New(nil)
	//用户自定义起始页
	i := util2.LoopIndexStart(size)
	for ; i < size; i++ {
		dUrl := pages[i] //从0开始
		if dUrl == "" {
			continue
		}
		ext := util2.FileExt(dUrl)
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, dUrl)
		fileName := sortId + ext
		dest := config.GetDestPath(taskUrl, bookId, fileName)
		gohttp.FastGet(dUrl, gohttp.Options{
			DestFile:    dest,
			Overwrite:   false,
			Concurrency: config.Conf.Threads,
			CookieJar:   jar,
			CookieFile:  config.Conf.CookieFile,
			Headers: map[string]interface{}{
				"user-agent": config.UserAgent,
			},
		})
	}
}

func getPages(bookId string) (pages []string) {
	//读cookie
	header, _ := curl2.GetHeaderFile(config.Conf.CookieFile)
	var manifests = new(ManifestsJson)
	bs, err := curl2.Get(fmt.Sprintf("https://www.loc.gov/item/%s/?fo=json", bookId), header)
	if err != nil {
		return
	}
	if err = json.Unmarshal(bs, manifests); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	//fmt.Println(manifests)
	newWidth := ""
	//限制图片最大宽度
	if config.Conf.FullImageWidth > 6400 {
		newWidth = "full/pct:100/"
	} else if config.Conf.FullImageWidth >= 1000 {
		newWidth = fmt.Sprintf("full/%d,/", config.Conf.FullImageWidth)
	}
	//一本书有N卷
	for _, resource := range manifests.Resources {
		//每卷有P页
		for _, file := range resource.Files {
			//每页有6种下载方式
			imgUrl, ok := getImagePage(file, newWidth)
			if ok {
				pages = append(pages, imgUrl)
			}
		}
	}
	return
}

func getImagePage(fileUrls []ImageFile, newWidth string) (downloadUrl string, ok bool) {
	for _, f := range fileUrls {
		if config.Conf.FileExt == ".jpg" && f.Mimetype == "image/jpeg" {
			if strings.Contains(f.Url, "full/pct:100/") {
				if newWidth != "" && newWidth != "full/pct:100/" {
					downloadUrl = strings.Replace(f.Url, "full/pct:100/", newWidth, 1)
				} else {
					downloadUrl = f.Url
				}
				ok = true
				break
			}
		} else if f.Mimetype != "image/jpeg" {
			if config.Conf.UseCDN == 1 {
				downloadUrl = strings.Replace(f.Url, "https://tile.loc.gov/storage-services/", "http://140.147.239.202/", 1)
			} else {
				downloadUrl = f.Url
			}
			ok = true
			break
		}
	}
	return
}
