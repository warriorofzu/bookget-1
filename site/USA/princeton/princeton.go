package princeton

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/gohttp"
	util2 "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`catalog/([0-9]+)`).FindStringSubmatch(taskUrl)
	if m != nil {
		bookId = m[1]
		config.CreateDirectory(taskUrl, bookId)
		StartDownload(iTask, taskUrl, bookId)
	}
	return "", err
}

func StartDownload(iTask int, text, bookId string) {
	name := util2.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, text)

	pages := getPages(bookId)
	log.Printf("A total of %d pages.\n", len(pages))

	//用户自定义起始页
	size := len(pages)
	i := util2.LoopIndexStart(size)
	for ; i < size; i++ {
		uri := pages[i] //从0开始
		if uri == "" {
			continue
		}
		ext := util2.FileExt(uri)
		sortId := util2.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		fileName := sortId + ext
		dest := config.GetDestPath(text, bookId, fileName)
		gohttp.FastGet(uri, gohttp.Options{
			Concurrency: config.Conf.Threads,
			DestFile:    dest,
			Overwrite:   false,
			Headers: map[string]interface{}{
				"user-agent": config.UserAgent,
			},
		})
	}

}

func getPages(bookId string) (pages []string) {
	phql := new(Graphql)
	dataRaw := fmt.Sprintf(`{"operationName":"GetResourcesByOrangelightIds","variables":{"ids":["%s"]},"query":"query GetResourcesByOrangelightIds($ids: [String!]!) {\n  resourcesByOrangelightIds(ids: $ids) {\n    id\n    thumbnail {\n      iiifServiceUrl\n      thumbnailUrl\n      __typename\n    }\n    url\n    members {\n      id\n      __typename\n    }\n    ... on ScannedResource {\n      manifestUrl\n      orangelightId\n      __typename\n    }\n    ... on ScannedMap {\n      manifestUrl\n      orangelightId\n      __typename\n    }\n    ... on Coin {\n      manifestUrl\n      orangelightId\n      __typename\n    }\n    __typename\n  }\n}\n"}`,
		bookId)
	header := make(map[string]string)
	header["authority"] = "figgy.princeton.edu"
	bs, err := curl.PostJson("https://figgy.princeton.edu/graphql", []byte(dataRaw), header)
	if err != nil {
		return
	}
	if err = json.Unmarshal(bs, phql); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	var manifestUrl = ""
	for _, v := range phql.Data.ResourcesByOrangelightIds {
		manifestUrl = v.ManifestUrl
	}
	if manifestUrl == "" {
		return
	}

	//查全书分卷URL
	var manifest = new(Manifest)
	body, err := curl.Get(manifestUrl, nil)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, manifest); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	//分卷URL处理
	for _, vol := range manifest.Manifests {
		pages = append(pages, getVolumeByURL(vol.Id)...)
	}
	return
}

func getVolumeByURL(uri string) (images []string) {
	var manifest2 = new(Manifest2)
	body, err := curl.Get(uri, nil)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, manifest2); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	i := len(manifest2.Sequences[0].Canvases)
	images = make([]string, 0, i)
	newWidth := ""
	//此站最大只支持6400
	if config.Conf.FullImageWidth > 6400 {
		newWidth = "full/full/"
	} else if config.Conf.FullImageWidth >= 1000 {
		newWidth = fmt.Sprintf("full/%d,/", config.Conf.FullImageWidth)
	}
	//分卷URL处理
	for _, sequences := range manifest2.Sequences {
		for _, canvase := range sequences.Canvases {
			for _, image := range canvase.Images {
				//JPEG URL
				imgUrl := image.Resource.Id
				if newWidth != "" {
					imgUrl = strings.Replace(image.Resource.Id, "full/1000,/", newWidth, 1)
				}
				images = append(images, imgUrl)
			}
		}
	}
	return
}
