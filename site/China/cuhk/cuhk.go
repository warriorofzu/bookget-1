package cuhk

import (
	"bookget/config"
	curl2 "bookget/lib/curl"
	util2 "bookget/lib/util"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := getBookId(taskUrl)
	if bookId == "" {
		return "", errors.New("not found BookId")
	}
	StartDownload(iTask, taskUrl, bookId)
	return "", nil
}

func getBookId(bookUrl string) (bookId string) {
	m := regexp.MustCompile(`item/cuhk-([A-Za-z0-9]+)`).FindStringSubmatch(bookUrl)
	if m != nil {
		bookId = m[1]
	}
	return
}

func StartDownload(iTask int, taskUrl, bookId string) {
	name := util2.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, taskUrl)
	bookUrls := getMultiplebooks(taskUrl)
	if bookUrls == nil || len(bookUrls) == 0 {
		log.Printf("Not found URLs.\n")
		return
	}
	size := len(bookUrls)
	log.Printf("A total of %d volumes \n", size)
	for i := 0; i < size; i++ {
		uri := bookUrls[i]
		log.Printf("Test volume %d ... \n", i+1)
		id := fmt.Sprintf("%s.%s", bookId, util2.GenNumberSorted(i+1))
		config.CreateDirectory(uri, bookId)
		do(id, uri)
	}
}

func do(bookId, bookUrl string) {
	iiifUri, manifest, useToken, cookies := getImageUrls(bookUrl)
	if manifest == nil {
		log.Printf("Not found URL.")
		return
	}
	//用户自定义起始页
	size := len(manifest.Pages)
	i := util2.LoopIndexStart(size)
	log.Printf("A total of %d pages.\n", size)
	//访问带cookie
	sCookie := curl2.HttpCookie2String(cookies)
	for ; i < size; i++ {
		page := manifest.Pages[i]
		imgUrl := fmt.Sprintf("%s/%s/full/full/0/default.jpg", iiifUri, page.Identifier)
		ext := util2.FileExt(imgUrl)
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, imgUrl)

		filename := sortId + ext
		dest := config.GetDestPath(bookUrl, bookId, filename)

		header := make(map[string]string, 2)
		header["Cookie"] = sCookie
		header["Referer"] = bookUrl
		if useToken {
			header["X-ISLANDORA-TOKEN"] = page.Token
		}
		curl2.FastGet(imgUrl, dest, header, true)
	}
}

func getMultiplebooks(bookUrl string) (uri []string) {
	bs, err := curl2.Get(bookUrl, nil)
	if err != nil {
		return
	}
	text := string(bs)
	subText := util2.SubText(text, "id=\"block-islandora-compound-object-compound-navigation-select-list\"", "id=\"book-viewer\">")
	matches := regexp.MustCompile(`value=['"]([A-z\d:_-]+)['"]`).FindAllStringSubmatch(subText, -1)
	if matches == nil {
		return
	}
	for _, m := range matches {
		//value='ignore'
		if m[1] == "ignore" {
			continue
		}
		id := strings.Replace(m[1], ":", "-", 1)
		uri = append(uri, fmt.Sprintf("https://repository.lib.cuhk.edu.hk/sc/item/%s#page/1/mode/2up", id))
	}
	return
}

func getImageUrls(bookUrl string) (iiifUri string, manifest *iiifManifest, useToken bool, cookies []*http.Cookie) {
	bs, c, err := curl2.GetWithCookie(bookUrl, nil)
	if err != nil {
		return
	}
	text := string(bs)
	iiifUri = getIiifUri(&text)
	useToken = getTokenHeader(&text)
	bsPage := getBody(&text)
	manifest = new(iiifManifest)
	if err = json.Unmarshal(bsPage, manifest); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	return iiifUri, manifest, useToken, c
}

// header.Set("X-ISLANDORA-TOKEN")
func getTokenHeader(text *string) bool {
	useToken := false
	matchToken := regexp.MustCompile(`"tokenHeader":([a-zA-z]+)`).FindStringSubmatch(*text)
	if matchToken != nil {
		//请求头信息是否包含token
		if strings.ToLower(matchToken[1]) == "true" {
			useToken = true
		}
	}
	return useToken
}

func getIiifUri(text *string) string {
	// 还有一种方式，不用cookie可以下载JP2
	//https://repository.lib.cuhk.edu.hk/en/islandora/object/cuhk%3A412226/datastream/JP2/view/
	match := regexp.MustCompile(`"iiifUri":"([^"]+)"`).FindStringSubmatch(*text)
	iiifUri := ""
	if match != nil {
		iiifUri = "https://repository.lib.cuhk.edu.hk" + strings.ReplaceAll(match[1], "\\/", "/")
	}
	return iiifUri
}

func getBody(text *string) (bs []byte) {
	matches := regexp.MustCompile(`"pages":([^]]+)]`).FindStringSubmatch(*text)
	if matches == nil {
		return
	}
	sText := []byte("{\"pages\":" + matches[1] + "]}")
	return sText
}
