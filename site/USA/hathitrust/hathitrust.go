package hathitrust

import (
	"bookget/config"
	curl2 "bookget/lib/curl"
	util2 "bookget/lib/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`id=([^&]+)`).FindStringSubmatch(taskUrl)
	if m != nil {
		bookId = m[1]
		config.CreateDirectory(taskUrl, bookId)
		StartDownload(iTask, taskUrl, bookId)
	}
	return "", err
}

func StartDownload(num int, uri, bookId string) {
	name := util2.GenNumberSorted(num)
	log.Printf("Get %s  %s\n", name, uri)

	bs, err := curl2.Get(uri, nil)
	if err != nil {
		return
	}
	text := string(bs)
	//取页数
	// <input id="range-seq" class="navigator-range" type="range" min="1" max="1036" value="2" aria-label="Progress" dir="rtl" />
	matches := regexp.MustCompile(`<input(?:[^>]+)id="range-seq"(?:[^>]+)max="([0-9]+)"`).FindStringSubmatch(text)
	if matches == nil {
		return
	}
	size := 0
	if matches[1] != "" {
		size, _ = strconv.Atoi(matches[1])
	}
	log.Printf("A total of %d pages.\n", size)
	ext := ".jpeg"
	//用户自定义起始页，特殊站点seq=1是第一页
	i := util2.LoopIndexStart(size) + 1
	for ; i <= size; i++ {
		for true {
			sortId := util2.GenNumberSorted(i)
			imgurl := fmt.Sprintf("https://babel.hathitrust.org/cgi/imgsrv/image?id=%s&attachment=1&size=full&format=image/jpeg&seq=%d", bookId, i)
			log.Printf("Get %s  %s\n", sortId, imgurl)

			fileName := sortId + ext
			dest := config.GetDestPath(uri, bookId, fileName)

			header := make(map[string]string)
			_, err = curl2.FastGet(imgurl, dest, header, true)
			if err != nil {
				fmt.Println(err)
				//log.Println("images (1 file per page, watermarked,  max. 20 MB / 1 min), image quality:Full")
				for t := 60; t > 0; t-- {
					seconds := strconv.Itoa(t)
					if t < 10 {
						seconds = fmt.Sprintf("0%d", t)
					}
					fmt.Printf("\rServer: maximum download limit exceeded...please wait.... [00:%s of appr. Max 1 min]", seconds)
					time.Sleep(time.Second)
				}
				fmt.Println()
				continue
			}
			break
		}
	}

}
