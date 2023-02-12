package kotenseki

import (
	"bookget/config"
	curl2 "bookget/lib/curl"
	util2 "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`biblio/([A-Za-z0-9_-]+)`).FindStringSubmatch(taskUrl)
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
	canvases := getCanvases(bookId)
	if canvases.Size == 0 {
		return
	}
	log.Printf("A total of %d pages.\n", canvases.Size)
	destPath := config.CreateDirectory(taskUrl, bookId)
	util2.CreateShell(destPath, canvases.IiifUrls, nil)
	//用户自定义起始页
	i := util2.LoopIndexStart(canvases.Size)
	for ; i < canvases.Size; i++ {
		uri := canvases.ImgUrls[i] //从0开始
		if uri == "" {
			continue
		}
		ext := util2.FileExt(uri)
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		fileName := sortId + ext
		dest := config.GetDestPath(taskUrl, bookId, fileName)
		curl2.FastGet(uri, dest, nil, true)
	}

}

func getManifestUrl(bookId string) (uri string, err error) {
	q := fmt.Sprintf("{\"bool\":{\"must\":[{\"match\":{\"d.bid.keyword\":\"%s\"}}]}}", bookId)
	data := url.Values{}
	data.Set("id", "biblio")
	data.Set("tp", "search")
	data.Set("q", q)
	data.Set("sz", "1")
	s := data.Encode()

	bs, err := curl2.Post("https://kotenseki.nijl.ac.jp/app/ws/search/", []byte(s), nil)
	if err != nil {
		return
	}
	//优先使用IIIFMani
	if strings.Contains(string(bs), "d.IIIFMani") {
		var wssearch2 = new(WsSearch2)
		if err = json.Unmarshal(bs, wssearch2); err == nil {
			//"d.IIIFMani":"https://rmda.kulib.kyoto-u.ac.jp/iiif/metadata_manifest/RB00002861/manifest.json"
			uri = wssearch2.Hits.Hits[0].Source.TIiiflink[0].DIIIFMani
			return
		}
	}
	var wssearch = new(WsSearch)
	if err = json.Unmarshal(bs, wssearch); err == nil {
		//log.Printf("json.Unmarshal failed: %s\n", err)
		if wssearch.Hits.Hits[0].Source.SLicense != nil &&
			wssearch.Hits.Hits[0].Source.SLicense[0] == "not-open" {
			uri = fmt.Sprintf("https://kotenseki.nijl.ac.jp/biblio/%s/privateFlg/manifest", bookId)
		} else {
			uri = fmt.Sprintf("https://kotenseki.nijl.ac.jp/biblio/%s/manifest", bookId)
		}
		return
	}

	return
}

func getCanvases(bookId string) (canvases Canvases) {
	uri, err := getManifestUrl(bookId)
	if err != nil {
		return
	}
	bs, err := curl2.Get(uri, nil)
	if err != nil {
		return
	}
	var manifest = new(Manifest)
	if err = json.Unmarshal(bs, manifest); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	if len(manifest.Sequences) == 0 {
		return
	}
	newWidth := ""
	//>6400使用原图
	if config.Conf.FullImageWidth > 6400 {
		newWidth = "full/full/"
	} else if config.Conf.FullImageWidth >= 1000 {
		newWidth = fmt.Sprintf("full/%d,/", config.Conf.FullImageWidth)
	}

	size := len(manifest.Sequences[0].Canvases)
	canvases.ImgUrls = make([]string, 0, size)
	canvases.IiifUrls = make([]string, 0, size)
	for _, canvase := range manifest.Sequences[0].Canvases {
		for _, image := range canvase.Images {
			//dezoomify-rs URL
			iiiInfo := fmt.Sprintf("%s/info.json", image.Resource.Service.Id)
			canvases.IiifUrls = append(canvases.IiifUrls, iiiInfo)

			//JPEG URL
			imgUrl := image.Resource.Id
			isFind := strings.Contains(imgUrl, "full/full/")
			if isFind && newWidth != "" {
				imgUrl = strings.Replace(image.Resource.Id, "full/full/", newWidth, 1)
			} else if !isFind {
				imgUrl = fmt.Sprintf("%s/%s0/default.jpg", image.Resource.Service.Id, newWidth)
			}

			canvases.ImgUrls = append(canvases.ImgUrls, imgUrl)
		}
	}
	canvases.Size = size
	return
}
