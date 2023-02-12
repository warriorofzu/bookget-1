package rbkdocnpmtw

import (
	"bookget/config"
	curl2 "bookget/lib/curl"
	util2 "bookget/lib/util"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bs, err := curl2.Get(taskUrl, nil)
	if err != nil {
		return "", err
	}
	text := string(bs)
	bookId := getBookId(text)
	if bookId == "" {
		return
	}
	config.CreateDirectory(taskUrl, bookId)
	StartDownload(iTask, taskUrl, bookId, text)
	return "", err
}

func getBookId(text string) string {
	bookId := ""
	//<tr><th>統一編號</th><td>平圖018255-018299</td></tr>
	//<tr><th>題名</th><td><span class=red>增補六臣注文選</span> 存四十五卷</td></tr>
	m := regexp.MustCompile(`<tr><th>統一編號</th><td>(.+?)</td></tr>`).FindStringSubmatch(text)
	if m == nil {
		m = regexp.MustCompile(`<tr><th>題名</th><td>(.+?)</td></tr>`).FindStringSubmatch(text)
		if m == nil {
			return ""
		}
	}
	re := regexp.MustCompile(`<([^>])*>`)
	bookId = re.ReplaceAllString(m[1], "")
	//bookId = strings.ReplaceAll(strings.ReplaceAll(s, "  ", ""), " ", ".")
	bookId = regexp.MustCompile(`[<>\\/?:|\s$]*`).ReplaceAllString(bookId, "")
	return bookId
}

func StartDownload(iTask int, taskUrl, bookId, text string) {
	name := util2.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, taskUrl)
	links := getVolumeUrls(text)
	if links == nil {
		return
	}
	size := len(links)
	var canvases Canvases
	canvases.ImgUrls = make([]string, 0, size)

	//用户自定义起始页
	k := util2.LoopIndexStart(size)
	for j, link := range links {
		if j <= k {
			continue
		}
		if j > 0 {
			fmt.Printf("\r Test volume %d ... ", j)
		}
		imgUrls := getImageUrls(link)
		canvases.ImgUrls = append(canvases.ImgUrls, imgUrls...)
	}
	fmt.Println()
	canvases.Size = len(canvases.ImgUrls)
	log.Printf("A total of %d files.\n", canvases.Size)
	if canvases.ImgUrls == nil {
		return
	}
	ext := ".pdf"
	for i := 0; i < canvases.Size; i++ {
		uri := canvases.ImgUrls[i] //从0开始
		if uri == "" {
			continue
		}
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)

		filename := sortId + ext
		dest := config.GetDestPath(taskUrl, bookId, filename)
		curl2.FastGet(uri, dest, nil, true)
	}
	return
}

func getVolumeUrls(text string) (links []string) {
	start := strings.Index(text, "id=tree_title")
	if start == -1 {
		start = strings.Index(text, "id=\"tree_title\"")
		if start == -1 {
			return
		}
	}
	subText1 := text[start:]
	end := strings.Index(subText1, "id=\"footer-wrapper\"")
	if end == -1 {
		return
	}
	subText := subText1[:end]
	//'/npmtpc/npmtpall?ID=2581&SECU=1113437189&PAGE=rbmap/2ND_rbmap&ACTION=
	matches := regexp.MustCompile(`href=['"](\S+)["']`).FindAllStringSubmatch(subText, -1)
	if matches == nil {
		return
	}
	for _, match := range matches {
		v := match[1]
		if v[0] != '/' {
			v = "/" + v
		}
		link := fmt.Sprintf("https://rbk-doc.npm.edu.tw%s", v)
		links = append(links, link)
	}
	return
}

func getImageUrls(uri string) (imgUrls []string) {
	bs, err := curl2.Get(uri, nil)
	if err != nil {
		return
	}

	text := string(bs)
	start := strings.Index(text, "id=\"rbmeta\"")
	if start == -1 {
		start = strings.Index(text, "id=\"rbmap_rbmeta\"")
		if start == -1 {
			return
		}
	}
	subText1 := text[start:]
	end := strings.Index(subText1, "id=tree_title")
	if end == -1 {
		end = strings.Index(subText1, "id=\"tree_title\"")
		if end == -1 {
			return
		}
	}
	subText := subText1[:end]
	matches := regexp.MustCompile(`href=['"]?(\S+)["']?`).FindAllStringSubmatch(subText, -1)
	if matches == nil {
		return
	}
	id, sec, _ := getIdSecu(text)

	//整册合并为一个PDF？
	if config.Conf.MergePDFs == 1 {
		max := len(matches)
		first := matches[0][1]
		last := matches[max-1][1]
		newId := fmt.Sprintf("%s%s", last[:4], first[4:])
		link := fmt.Sprintf("https://rbk-doc.npm.edu.tw/npmtpc/npmtpall?ID=%s&SECU=%s&ACTION=UI,%s", id, sec, newId)
		imgUrls = append(imgUrls, link)
		return
	}
	// 按台北故宫官方提供的URL下载若干个PDF
	for _, match := range matches {
		link := fmt.Sprintf("https://rbk-doc.npm.edu.tw/npmtpc/npmtpall?ID=%s&SECU=%s&ACTION=UI,%s", id, sec, match[1])
		imgUrls = append(imgUrls, link)
	}
	return
}

func getIdSecu(text string) (id string, sec string, tphc string) {
	//<input type=hidden name=ID value=3408><input type=hidden name=SECU value=536673644>
	//<input type=hidden name=TPHC value=1 size=30>
	matches := regexp.MustCompile(`<input\s+type=hidden\s+name=(ID|SECU|TPHC)\s+value=(\d+)`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}

	for _, v := range matches {
		if v[1] == "ID" {
			id = v[2]
		} else if v[1] == "SECU" {
			sec = v[2]
		} else if v[1] == "TPHC" {
			tphc = v[2]
		}
	}
	return
}

func getLoginForm() (id string, sec string, tphc string) {
	rand.Seed(time.Now().Unix())
	uri := fmt.Sprintf("https://rbk-doc.npm.edu.tw/npmtpc/npmtpall?@@%d", rand.Int())
	bs, err := curl2.GetRedirects(uri, nil, 3)
	if err != nil {
		return
	}
	return getIdSecu(string(bs))
}

func userPass(name, pwd string) string {
	/*
		sys/00/userid: aaa
		sys/00/passwd: bbb
		_BTN_登入^^^SI: 登入
		local/title:
		local/problemDes:
		local/name:
		local/email:
		local/company:
		local/phone:
		ID: 2626
		SECU: 1678849616
		TPHC: 11
	*/
	id, sec, tphc := getLoginForm()
	data := url.Values{}
	data.Set("sys/00/userid", name)
	data.Set("sys/00/passwd", pwd)
	data.Set("ID", id)
	data.Set("SECU", sec)
	data.Set("TPHC", tphc)
	data.Set("_BTN_登入^^^SI", "登入")
	data.Set("local/title", "")
	data.Set("local/problemDes", "")
	data.Set("local/name", "")
	data.Set("local/email", "")
	data.Set("local/company", "")
	data.Set("local/phone", "")
	return data.Encode()
}
