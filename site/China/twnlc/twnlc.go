package twnlc

import (
	"bookget/config"
	curl2 "bookget/lib/curl"
	util2 "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`item=([A-Za-z0-9]+)`).FindStringSubmatch(taskUrl)
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

	bs, cookies, err := curl2.GetWithCookie(taskUrl, nil)
	if err != nil {
		return
	}
	text := string(bs)
	canvases := getImageUrls(text)
	log.Printf("A total of %d pages.\n", canvases.Size)
	if canvases.Size == 0 {
		return
	}
	requestVerificationToken := getRequestToken(text)

	//循环下载图片
	ext := ".jpg"
	//用户自定义起始页
	i := util2.LoopIndexStart(canvases.Size)
	for ; i < canvases.Size; i++ {
		u := canvases.ImgUrls[i]
		if u == "" {
			continue
		}
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, u)
		fileName := sortId + ext
		dest := config.GetDestPath(taskUrl, bookId, fileName)
		//文件存在，跳过
		fi, err := os.Stat(dest)
		if err == nil && fi.Size() > 0 {
			continue
		}
		token := getToken(requestVerificationToken, cookies)
		imgurl := fmt.Sprintf("https://rbook.ncl.edu.tw%s&token=%s", u, token)
		curl2.FastGet(imgurl, dest, nil, false)
	}
	return
}

func getImageUrls(text string) (canvases Canvases) {
	//取页数
	matches := regexp.MustCompile(`name="ImageCheck" value="([^>]+)"`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}
	canvases.ImgUrls = make([]string, 0, len(matches))
	for _, v := range matches {
		href := strings.Replace(v[1], "&amp;", "&", -1)
		canvases.ImgUrls = append(canvases.ImgUrls, href)
	}
	canvases.Size = len(canvases.ImgUrls)
	return
}

func getRequestToken(text string) string {
	//取请求token
	// <input name="__RequestVerificationToken" type="hidden" value="ayk-lqrk1RrbJb1xB6FM2-cALjxxYUHAapQoPBSLuVQFSmJQQ-DQSAhzcE7IciaEw3GZBs_irf71OGFXZxUctQeJaSBfU2V1TvI5vijRjMA1" />
	matchesToken := regexp.MustCompile(`name="__RequestVerificationToken(?:.+)value="(\S+)"`).FindStringSubmatch(text)
	if matchesToken == nil {
		return ""
	}
	//reqToken
	return matchesToken[1]
}

func getToken(requestVerificationToken string, cookies []*http.Cookie) string {
	uri := "https://rbook.ncl.edu.tw/NCLSearch/Watermark/getToken"
	data := "__RequestVerificationToken=" + requestVerificationToken

	header := make(map[string]string)
	header["Cookie"] = curl2.HttpCookie2String(cookies)
	bs, err := curl2.Post(uri, []byte(data), header)
	if err != nil {
		return ""
	}
	resToken := new(ResponseToken)
	if err = json.Unmarshal(bs, resToken); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
	}
	return resToken.Token
}
