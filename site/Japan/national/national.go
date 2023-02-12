package national

import (
	"bookget/config"
	curl2 "bookget/lib/curl"
	util2 "bookget/lib/util"
	"fmt"
	"log"
	"os"
	"regexp"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`BID=([A-Za-z0-9_-]+)`).FindStringSubmatch(taskUrl)
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
	bookIds, size := getMultiplebooks(bookId)
	if bookIds == nil || size == 0 {
		return
	}
	log.Printf("\nA total of %d files.\n", size)
	//用户自定义起始页
	i := util2.LoopIndexStart(size)
	for ; i < size; i++ {
		id := (*bookIds)[i] //从0开始
		if id == "" {
			continue
		}
		ext := ".zip"
		sortId := util2.GenNumberSorted(i + 1)
		extId := "pdf"
		switch config.Conf.FileExt {
		case ".jpg":
			extId = "jpeg"
		case ".jp2":
			extId = "jp2"
		}
		log.Printf("Get %s  %s\n", sortId, extId)
		fileName := sortId + ext
		dest := config.GetDestPath(taskUrl, bookId, fileName)
		//文件存在，跳过
		fi, err := os.Stat(dest)
		if err == nil && fi.Size() > 0 {
			continue
		}
		download(bookId, id, dest, sortId)
	}

}
func getMultiplebooks(bookId string) (bookIDs *[]string, size int) {
	//https://www.digital.archives.go.jp/DAS/meta/listPhoto?LANG=default&BID=F1000000000000095447&ID=&NO=&TYPE=dljpeg&DL_TYPE=jpeg
	downPage := fmt.Sprintf("https://www.digital.archives.go.jp/DAS/meta/listPhoto?LANG=default&BID=%s&ID=&NO=&TYPE=dljpeg&DL_TYPE=jpeg", bookId)
	bs, err := curl2.Get(downPage, nil)
	if err != nil {
		return
	}
	text := string(bs)
	//<input type="checkbox" class="check" name="id_2" posi="2" value="M2016092111023960474"
	//取册数
	matches := regexp.MustCompile(`<input[^>]+posi=["']([0-9]+)["'][^>]+value=["']([A-Za-z0-9]+)["']`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}
	iLen := len(matches)
	ids := make([]string, 0)
	for _, match := range matches {
		//跳过全选复选框
		if iLen > 1 && (match[1] == "0" || match[2] == "") {
			continue
		}
		ids = append(ids, match[2])
	}
	bookIDs = &ids
	size = len(ids)
	return
}

func download(bookId, id string, dest string, sortId string) (size int64, err error) {
	// pdf|jp2|jpeg
	uri := "https://www.digital.archives.go.jp/acv/auto_conversion/download"
	//dataRaw := fmt.Sprintf("DL_TYPE=%s&id_0=%s&page_0=&id_1=%s&page_1=", config.Conf.DownloadImageType, bookId, id)
	extId := "pdf"
	switch config.Conf.FileExt {
	case ".jpg":
		extId = "jpeg"
	case ".jp2":
		extId = "jp2"
	}

	dataRaw := fmt.Sprintf("DL_TYPE=%s&id_1=%s&page_1=", extId, id)
	size, err = curl2.PostDownload(uri, dest, []byte(dataRaw), nil)
	return
}
