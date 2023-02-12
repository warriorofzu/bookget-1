package idp

import (
	"bookget/config"
	"bookget/lib/gohttp"
	util2 "bookget/lib/util"
	"fmt"
	"log"
	"net/http/cookiejar"
	"net/url"
	"regexp"
)

func Init(iTask int, sUrl string) (msg string, err error) {
	dt := new(DownloadTask)
	dt.CookieJar, _ = cookiejar.New(nil)
	dt.ParsedUrl, _ = url.Parse(sUrl)
	dt.Url = sUrl
	dt.Index = iTask
	dt.BookId = getBookId(sUrl)
	return StartDownload(dt)
}

func getBookId(sUrl string) string {
	bookId := ""
	m := regexp.MustCompile(`uid=([A-Za-z0-9]+)`).FindStringSubmatch(sUrl)
	if m != nil {
		bookId = m[1]
	}
	return bookId
}

func StartDownload(dt *DownloadTask) (msg string, err error) {
	canvases := getCanvases(dt.Url, dt)
	if canvases.Size == 0 {
		return
	}
	log.Printf("A total of %d pages.\n", canvases.Size)

	config.CreateDirectory(dt.Url, dt.BookId)
	//用户自定义起始页
	i := util2.LoopIndexStart(canvases.Size)
	ext := ".jpg"
	for ; i < canvases.Size; i++ {
		dUrl := canvases.ImgUrls[i] //从0开始
		if dUrl == "" {
			continue
		}
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, dUrl)
		filename := sortId + ext
		dest := config.GetDestPath(dt.Url, dt.BookId, filename)
		//_, err = wget.FastGetv2(dUrl, dest, dt.CookieJar, config.Conf.CookieFile, true)
		cli := gohttp.NewClient(gohttp.Options{
			DestFile:   dest,
			CookieJar:  dt.CookieJar,
			CookieFile: config.Conf.CookieFile,
			Headers: map[string]interface{}{
				"User-Agent": config.Conf.UserAgent,
			},
		})
		_, err = cli.Get(dUrl)
		if err != nil {
			fmt.Println(err)
		}
	}
	return "", nil
}

func getCanvases(sUrl string, dt *DownloadTask) (canvases Canvases) {
	cli := gohttp.NewClient(gohttp.Options{
		Timeout:    0,
		CookieFile: config.Conf.CookieFile,
		CookieJar:  dt.CookieJar,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
	})
	resp, err := cli.Get(sUrl)
	if err != nil {
		log.Fatalln(err)
	}
	bs, _ := resp.GetBody()
	//imageUrls[0] = "/image_IDP.a4d?type=loadRotatedMainImage;recnum=31305;rotate=0;imageType=_M";
	//imageRecnum[0] = "31305";
	m := regexp.MustCompile(`imageRecnum\[\d+\][ \S]?=[ \S]?"(\d+)";`).FindAllSubmatch(bs, -1)
	if m == nil {
		return
	}
	canvases.ImgUrls = make([]string, 0, len(m))
	for _, v := range m {
		id := string(v[1])
		imgUrl := fmt.Sprintf("%s://%s/image_IDP.a4d?type=loadRotatedMainImage;recnum=%s;rotate=0;imageType=_L",
			dt.ParsedUrl.Scheme, dt.ParsedUrl.Host, id)
		canvases.ImgUrls = append(canvases.ImgUrls, imgUrl)
	}
	canvases.Size = len(canvases.ImgUrls)
	return
}
