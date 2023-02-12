package Universal

import (
	"bookget/config"
	"bookget/lib/file"
	"bookget/lib/gohttp"
	"bookget/lib/util"
	"bookget/lib/zhash"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func StartDownload(iTask int, taskUrl string) (msg string, err error) {
	if taskUrl == "" {
		return
	}
	name := util.GenNumberSorted(iTask)
	log.Printf("download: %s  %s\n", name, taskUrl)
	//通用下载
	downloadUrls, startIndex := getDownloadUrls(taskUrl)

	log.Printf("A total of %d files.\n", len(downloadUrls))
	ext := file.Ext(downloadUrls[0])
	bookId := ""
	if len(downloadUrls) > 1 {
		bookId = strconv.FormatUint(uint64(zhash.CRC32(taskUrl)), 10)
	}
	config.CreateDirectory(taskUrl, bookId)
	for _, dUrl := range downloadUrls {
		sortId := util.GenNumberSorted(startIndex)
		startIndex++
		fileName := ""
		if ext == ".jpg" || ext == ".tif" || ext == ".jp2" || ext == ".png" || ext == ".pdf" {
			fileName = sortId + ext
		} else {
			fileName = file.Name(dUrl)
		}
		log.Printf("Get %s  %s\n", sortId, dUrl)
		dest := config.GetDestPath(dUrl, bookId, fileName)
		cli := gohttp.NewClient()
		_, err = cli.FastGet(dUrl, gohttp.Options{
			DestFile:    dest,
			Concurrency: config.Conf.Threads,
			CookieFile:  config.Conf.CookieFile,
			Headers: map[string]interface{}{
				"User-Agent": config.Conf.UserAgent,
			},
		})
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	return "", err
}

func getDownloadUrls(sUrl string) (downloadUrls []string, startIndex int) {
	matches := regexp.MustCompile(`\((\d+)-(\d+)\)`).FindStringSubmatch(sUrl)
	if matches == nil {
		downloadUrls = append(downloadUrls, sUrl)
		return downloadUrls, startIndex
	}
	i, _ := strconv.ParseInt(matches[1], 10, 64)
	max, _ := strconv.ParseInt(matches[2], 10, 64)
	iMinLen := len(matches[1])
	startIndex = int(i)

	tmpUrl := regexp.MustCompile(`\((\d+)-(\d+)\)`).ReplaceAllString(sUrl, "%s")
	downloadUrls = make([]string, 0, max)
	for ; i <= max; i++ {
		iLen := len(strconv.FormatInt(i, 10))
		if iLen < iMinLen {
			iLen = iMinLen
		}
		sortId := util.GenNumberLimitLen(int(i), iLen)
		dUrl := fmt.Sprintf(tmpUrl, sortId)
		downloadUrls = append(downloadUrls, dUrl)
	}
	return downloadUrls, startIndex
}
