package familysearch

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/gohttp"
	util2 "bookget/lib/util"
	"fmt"
	"log"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//家谱图像 https://www.familysearch.org/records/images/
func ImagesDownload(t *Downloader) (msg string, err error) {

	name := util2.GenNumberSorted(t.Index)
	log.Printf("Get %s  %s\n", name, t.Url)

	canvases, err := getCanvases(t.BookId, config.Conf.CookieFile)
	if err != nil {
		return "", err
	}
	//用户自定义起始页
	log.Printf("A total of %d Pages.\n", canvases.Size)

	//cookie 处理
	jar, _ := cookiejar.New(nil)
	header, _ := curl.GetHeaderFile(config.Conf.CookieFile)
	util2.CreateShell(t.SavePath, canvases.IiifUrls, header)
	//用户自定义起始页
	i := util2.LoopIndexStart(canvases.Size)
	for ; i < canvases.Size; i++ {
		dUrl := canvases.ImageUrls[i] //从0开始
		if dUrl == "" {
			continue
		}
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, dUrl)
		fileName := sortId + ".jpg"
		dest := config.GetDestPath(t.Url, t.BookId, fileName)
		for {
			_, err = gohttp.FastGet(dUrl, gohttp.Options{
				DestFile:    dest,
				Overwrite:   false,
				Concurrency: config.Conf.Threads,
				CookieJar:   jar,
				CookieFile:  config.Conf.CookieFile,
				Headers: map[string]interface{}{
					"user-agent": config.UserAgent,
				},
			})
			if err != nil {
				fmt.Println(err)
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

	return "", nil
}

func getCanvases(bookId, cookieFile string) (Canvases, error) {
	canvases := Canvases{}
	apiUrl := fmt.Sprintf("https://www.familysearch.org/records/images/api/imageDetails/groups/%s?properties&changeLog&coverageIndex=null", bookId)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: cookieFile,
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
		},
	})
	resp, err := cli.Get(apiUrl)
	if err != nil {
		return canvases, err
	}
	imageGroups := ImageGroups{}
	if err = resp.GetJsonDecodeBody(&imageGroups); err != nil {
		return canvases, err
	}
	canvases.IiifUrls = make([]string, 0, imageGroups.VolumeSet.ChildCount)
	canvases.ImageUrls = make([]string, 0, imageGroups.VolumeSet.ChildCount)
	for _, group := range imageGroups.Groups {
		for _, v := range group.ImageUrls {
			dzUrl := fmt.Sprintf("%s/image.xml", v)
			u, _ := url.Parse(v)
			i := strings.LastIndex(u.Path, "/")
			id := u.Path[i+1:]
			imgUrl := fmt.Sprintf("%s://%s/service/records/storage/dascloud/das/v2/%s/dist.jpg?proxy=true", u.Scheme, u.Host, id)
			canvases.IiifUrls = append(canvases.IiifUrls, dzUrl)
			canvases.ImageUrls = append(canvases.ImageUrls, imgUrl)
		}
	}
	canvases.Size = len(canvases.IiifUrls)
	return canvases, nil
}
