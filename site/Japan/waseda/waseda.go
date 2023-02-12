package waseda

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/gohttp"
	util2 "bookget/lib/util"
	"fmt"
	"log"
	"regexp"
	"sort"
)

// 自定义一个类型
type strs []string

func (s strs) Len() int           { return len(s) }
func (s strs) Less(i, j int) bool { return s[i] < s[j] }
func (s strs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`kosho/[A-Za-z0-9_-]+/([A-Za-z0-9_-]+)/`).FindStringSubmatch(taskUrl)
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

	pdfUrls, size := getMultiplebooks(taskUrl)
	if pdfUrls == nil || size == 0 {
		return
	}
	log.Printf("A total of %d pages.\n", size)

	//用户自定义起始页
	i := util2.LoopIndexStart(size)
	for ; i < size; i++ {
		uri := (*pdfUrls)[i] //从0开始
		if uri == "" {
			continue
		}
		ext := util2.FileExt(uri)
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		fileName := sortId + ext
		dest := config.GetDestPath(taskUrl, bookId, fileName)
		gohttp.FastGet(uri, gohttp.Options{
			DestFile:    dest,
			Overwrite:   false,
			Concurrency: config.Conf.Threads,
			Headers: map[string]interface{}{
				"user-agent": config.UserAgent,
			},
		})
	}
}

func getMultiplebooks(bookUrl string) (pdfUrls *[]string, size int) {
	bs, err := curl.Get(bookUrl, nil)
	if err != nil {
		return
	}
	text := string(bs)
	//取册数
	matches := regexp.MustCompile(`href=["'](.+?)\.pdf["']`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}
	ids := make([]string, 0, len(matches))
	for _, match := range matches {
		ids = append(ids, match[1])
	}
	sort.Sort(strs(ids))
	links := make([]string, 0, len(ids))
	for _, v := range ids {
		s := fmt.Sprintf("%s%s.pdf", bookUrl, v)
		links = append(links, s)
	}
	pdfUrls = &links
	size = len(links)
	return
}
