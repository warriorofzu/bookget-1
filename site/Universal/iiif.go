package Universal

import (
	"bookget/config"
	"bookget/lib/gohttp"
	"bookget/lib/util"
	"bookget/site/Universal/iiif"
	"errors"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
)

func AutoDetectManifest(iTask int, taskUrl string) (msg string, err error) {
	name := util.GenNumberSorted(iTask)
	log.Printf("Auto Detect %s  %s\n", name, taskUrl)

	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
	})
	resp, err := cli.Get(taskUrl)
	if err != nil {
		return
	}
	bs, _ := resp.GetBody()
	text := string(bs)
	if bs[0] == '{' && strings.Contains(text, "iiif.io") {
		iiif.Init(iTask, taskUrl)
		return
	}
	//href="https://dcollections.lib.keio.ac.jp/ja/kanseki/110x-24-1?manifest=https://dcollections.lib.keio.ac.jp/sites/default/files/iiif/KAN/110X-24-1/manifest.json"
	manifestUrl := getManifestUrl(taskUrl, text)
	if manifestUrl == "" {
		msg = "URL not found: manifest.json"
		err = errors.New(msg)
		return
	}
	iiif.Init(iTask, manifestUrl)
	return
}

func getManifestUrl(pageUrl, text string) string {
	//最后是，相对URI
	u, err := url.Parse(pageUrl)
	if err != nil {
		return ""
	}
	host := fmt.Sprintf("%s://%s/", u.Scheme, u.Host)
	//优先明显是manifest的
	m := regexp.MustCompile(`manifest=(\S+).json["']`).FindStringSubmatch(text)
	if m != nil {
		return padUri(host, m[1]+".json")
	}
	m = regexp.MustCompile(`manifest=(\S+)["']`).FindStringSubmatch(text)
	if m != nil {
		return padUri(host, m[1])
	}
	m = regexp.MustCompile(`href=["'](\S+)/manifest.json["']`).FindStringSubmatch(text)
	if m == nil {
		return ""
	}
	return padUri(host, m[1]+"/manifest.json")
}
func padUri(host, uri string) string {
	//https:// 或 http:// 绝对URL
	if strings.HasPrefix(uri, "https://") || strings.HasPrefix(uri, "http://") {
		return uri
	}
	manifestUri := ""
	if uri[0] == '/' {
		manifestUri = uri[1:]
	} else {
		manifestUri = uri
	}
	return host + manifestUri
}
