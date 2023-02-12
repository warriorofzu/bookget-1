package princeton

import (
	"bookget/config"
	"bookget/lib/gohttp"
	"bookget/lib/util"
	"bookget/site/Universal/iiif"
	"log"
	"net/http/cookiejar"
	"net/url"
	"regexp"
)

type ManifestUrls struct {
	Manifests []struct {
		Context string   `json:"@context"`
		Type    string   `json:"@type"`
		Id      string   `json:"@id"`
		Label   []string `json:"label"`
	} `json:"manifests"`
}

func InitDpul(iTask int, taskUrl string) (msg string, err error) {
	name := util.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, taskUrl)
	bookId := getBookId(taskUrl)
	if bookId == "" {
		return "", err
	}
	mfUrls, err := getManifestUrl(taskUrl)
	for i, manifestUrl := range mfUrls {
		tId := util.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", tId, manifestUrl)
		iiif.StartDownload(manifestUrl, bookId)
	}
	return "", err
}

func getBookId(sUrl string) string {
	m := regexp.MustCompile(`catalog/([A-z0-9]+)`).FindStringSubmatch(sUrl)
	if m != nil {
		return m[1]
	}
	return ""
}

func getManifestUrl(sUrl string) (Urls []string, err error) {
	mfUrl := detectManifest(sUrl)
	//cookie 处理
	jar, _ := cookiejar.New(nil)
	resp, err := gohttp.Get(mfUrl, gohttp.Options{
		CookieJar:  jar,
		CookieFile: config.Conf.CookieFile,
		Headers: map[string]interface{}{
			"user-agent": config.UserAgent,
		},
	})
	manifestUrls := ManifestUrls{}
	if err = resp.GetJsonDecodeBody(&manifestUrls); err != nil {
		return
	}
	for _, manifest := range manifestUrls.Manifests {
		Urls = append(Urls, manifest.Id)
	}
	return
}

func detectManifest(sUrl string) string {
	//cookie 处理
	jar, _ := cookiejar.New(nil)
	resp, err := gohttp.Get(sUrl, gohttp.Options{
		CookieJar:  jar,
		CookieFile: config.Conf.CookieFile,
		Headers: map[string]interface{}{
			"user-agent": config.UserAgent,
		},
	})
	bs, err := resp.GetBody()
	if err != nil {
		return ""
	}
	text := string(bs)

	//优先明显是manifest的
	m := regexp.MustCompile(`manifest=([^&]+)&`).FindStringSubmatch(text)
	if m != nil {
		mfUrl, _ := url.QueryUnescape(m[1])
		return mfUrl
	}
	return ""
}
