package nlc

import (
	"bookget/config"
	curl2 "bookget/lib/curl"
	util2 "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func Init(iTask int, sUrl string) (msg string, err error) {
	dt := new(DownloadTask)
	dt.Domain = util2.GetHostUrl(sUrl)
	dt.Url = sUrl
	dt.Index = iTask
	return Download(dt)
}

func Download(dt *DownloadTask) (msg string, err error) {
	//单册
	if strings.Contains(dt.Url, "OpenObjectBook") {
		bs, err := curl2.Get(dt.Url, nil)
		if err != nil {
			fmt.Printf("Error ID: %s\n", dt.Url)
			return "", err
		}
		body := string(bs)
		dt.BookId = getBookId(body)
		dt.SavePath = config.CreateDirectory(dt.Url, dt.BookId)
		fetchOneBook(dt.Index, dt.Url, dt)
	} else {
		dt.BookId = getBookId(dt.Url)
		dt.SavePath = config.CreateDirectory(dt.Url, dt.BookId)
		fetchBooks(dt)
	}
	return "ok", nil
}

func getBookId(text string) string {
	bookId := ""
	m := regexp.MustCompile(`identifier[\s]+=[\s]+["']([A-Za-z0-9]+)["']`).FindStringSubmatch(text)
	if m != nil {
		bookId = m[1]
	}
	m = regexp.MustCompile(`fid=([A-Za-z0-9]+)`).FindStringSubmatch(text)
	if m != nil {
		bookId = m[1]
	}
	return bookId
}

func fetchBooks(dt *DownloadTask) (msg string, err error) {
	name := util2.GenNumberSorted(dt.Index)
	log.Printf("Get %s  %s\n", name, dt.Url)

	bs, err := curl2.Get(dt.Url, nil)
	if err != nil {
		return
	}
	text := string(bs)
	//取册数
	pdfUrls := regexp.MustCompile(`<a[^>]+class="a1"[^>].+href="/OutOpenBook/([^"]+)"`).FindAllStringSubmatch(text, -1)
	if pdfUrls == nil {
		return
	}

	//用户自定义起始页
	size := len(pdfUrls)
	log.Printf("A total of %d PDFs.\n", size)
	i := util2.LoopIndexStart(size)
	for ; i < size; i++ {
		if pdfUrls[i][1] == "" {
			continue
		}
		pdfUrl := fmt.Sprintf("%sOutOpenBook/%s", dt.Domain, pdfUrls[i][1])
		fetchOneBook(i+1, pdfUrl, dt)
	}
	return "", nil
}

func fetchOneBook(i int, sUrl string, dt *DownloadTask) {
	sortId := util2.GenNumberSorted(i)
	log.Printf("Get %s  %s\n", sortId, sUrl)

	//解析URL
	u, err := url.Parse(sUrl)
	if err != nil {
		log.Printf("URL error: %s \n %s\n", sUrl, err)
		return
	}
	m, _ := url.ParseQuery(u.RawQuery)
	if m["aid"] == nil || m["bid"] == nil {
		log.Printf("URL error: %s \n %s\n", sUrl, err)
		return
	}
	aid := m["aid"][0]
	bid := m["bid"][0]
	//取卷名
	if config.Conf.UseNumericFilename == 0 {
		volumeName := getVolumeTitle(bid, aid, dt.Domain)
		if volumeName != "" {
			sortId += "." + volumeName
		}
	} else {
		sortId = fmt.Sprintf("%s.%s", sortId, bid)
	}
	filename := sortId + ".pdf"
	dest := config.GetDestPath(dt.Url, dt.BookId, filename)
	//文件存在，跳过
	fi, err := os.Stat(dest)
	if err == nil && fi.Size() > 0 {
		return
	}

	tokenKey, timeKey, timeFlag := getToken(sUrl)

	pdfUrl := fmt.Sprintf("%smenhu/OutOpenBook/getReader?aid=%s&bid=%s&kime=%s&fime=%s", dt.Domain, aid, bid, timeKey, timeFlag)

	header := make(map[string]string)
	header["myreader"] = tokenKey
	header["Range"] = "bytes=0-1"
	header["Referer"] = fmt.Sprintf("%sstatic/webpdf/lib/WebPDFJRWorker.js", dt.Domain)

	curl2.FastGet(pdfUrl, dest, header, false)
}

func getVolumeTitle(id, indexName, domain string) string {
	arr := strings.Split(id, ".")
	dataRaw := fmt.Sprintf("id=%s&indexName=data_%s", arr[0], indexName)
	//{"success":true,"msg":"","obj":[{"chapter_name2":"經濟彙編戎政典","chapter_name1":"經濟彙編戎政典","chapter_num1":"第一百四十五卷","chapter_num2":"第一百四十六卷"}]}
	sUrl := fmt.Sprintf("%sallSearch/formatCatalog", domain)
	var cata = new(Catalog)
	bs, err := curl2.Post(sUrl, []byte(dataRaw), nil)
	if err != nil {
		return ""
	}
	if err = json.Unmarshal(bs, cata); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
	}
	text := ""
	if cata.Obj == nil || len(cata.Obj) == 0 {
		cnNumber := strings.Replace(LastChapter.ChapterNum2, "第", "", 1)
		cnNumber = strings.Replace(cnNumber, "卷", "", 1)
		numInt := util2.ChineseToNumber(cnNumber)
		numInt++
		var chapterNum2 = util2.NumberToChinese(int64(numInt))
		chapterNum2 = "第" + chapterNum2 + "卷"
		text += LastChapter.ChapterName2 + "." + chapterNum2

		LastChapter.Mux.Lock()
		LastChapter.ChapterNum2 = chapterNum2
		LastChapter.Mux.Unlock()
		return text
	}
	for _, v := range cata.Obj {
		if v.ChapterName1 == "" && v.ChapterNum1 == "" {
			continue
		}
		//使用卷数字，当卷名为空时
		if v.ChapterNum1 != "" && v.ChapterName1 == "" {
			text += v.ChapterNum1 + "-" + v.ChapterNum2
		} else if v.ChapterName1 == v.ChapterName2 { //使用卷名
			text += v.ChapterName1 + "." + v.ChapterNum1 + "-" + v.ChapterNum2
		} else {
			text += v.ChapterName1 + "." + v.ChapterNum1 + "-" + v.ChapterName2 + "." + v.ChapterNum2
		}
		LastChapter.Mux.Lock()
		LastChapter.ChapterName1 = v.ChapterName1
		LastChapter.ChapterName2 = v.ChapterName2
		LastChapter.ChapterNum1 = v.ChapterNum1
		LastChapter.ChapterNum2 = v.ChapterNum2
		LastChapter.Mux.Unlock()
	}
	return text
}

func getToken(uri string) (tokenKey, timeKey, timeFlag string) {
	body, err := curl2.Get(uri, nil)
	if err != nil {
		log.Printf("Server unavailable: %s", err.Error())
		return
	}
	//<iframe id="myframe" name="myframe" src="" width="100%" height="100%" scrolling="no" frameborder="0" tokenKey="4ADAD4B379874C10864990817734A2BA" timeKey="1648363906519" timeFlag="1648363906519" sflag=""></iframe>
	params := regexp.MustCompile(`(tokenKey|timeKey|timeFlag)="([a-zA-Z0-9]+)"`).FindAllStringSubmatch(string(body), -1)
	//tokenKey := ""
	//timeKey := ""
	//timeFlag := ""
	for _, v := range params {
		if v[1] == "tokenKey" {
			tokenKey = v[2]
		} else if v[1] == "timeKey" {
			timeKey = v[2]
		} else if v[1] == "timeFlag" {
			timeFlag = v[2]
		}
	}
	return
}
