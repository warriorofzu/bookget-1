package keio

import (
	"bookget/config"
	curl2 "bookget/lib/curl"
	util2 "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`id=([A-Za-z0-9]+)`).FindStringSubmatch(taskUrl)
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
	bookUrls, size := getMultiplebooks(bookId, taskUrl)
	if bookUrls == nil || size == 0 {
		return
	}
	log.Printf("A total of %d books.\n", size)
	imageUrls := make([]string, 0, 100*size)
	iiifUrls := make([]string, 0, 100*size)
	for k, u := range *bookUrls {
		imgs, iiifs := getImageUrls(u)
		fmt.Printf("\rBook %d has a total of %d pages. ", k+1, len(*imgs))
		if imgs != nil {
			imageUrls = append(imageUrls, *imgs...)
		}
		if iiifs != nil {
			iiifUrls = append(iiifUrls, *iiifs...)
		}
	}
	size = len(imageUrls)
	log.Printf("\nA total of %d pages.\n", size)
	destPath := config.CreateDirectory(taskUrl, bookId)
	util2.CreateShell(destPath, iiifUrls, nil)
	//用户自定义起始页
	i := util2.LoopIndexStart(size)
	for ; i < size; i++ {
		uri := imageUrls[i] //从0开始
		if uri == "" {
			continue
		}
		getSingleImage(i, uri, bookId)
	}

}

func getMultiplebooks(bookId string, bookUrl string) (bookUrls *[]string, size int) {
	bs, err := curl2.Get(bookUrl, nil)
	if err != nil {
		return
	}
	text := string(bs)
	//取册数
	matches := regexp.MustCompile(`<p[^>]+data-cid=['|"]([a-zA-Z0-9]+)['|"]`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}
	size = len(matches)
	volumeUrls := make([]string, 0, size)

	for _, v := range matches {
		childId := makeId(v[1], bookId, size)
		fmt.Sprintf("%s\n", childId)
		uri := fmt.Sprintf("https://db2.sido.keio.ac.jp/iiif/manifests/kanseki/%s/%s/manifest.json", bookId, childId)
		volumeUrls = append(volumeUrls, uri)
	}
	return &volumeUrls, size
}

func makeId(childId string, bookId string, iMax int) string {
	childIDfmt := ""
	//i, _ := strconv.Atoi(childId)
	iLen := 3
	if iMax > 999 {
		iLen = 4
	}
	for k := iLen - len(childId); k > 0; k-- {
		childIDfmt += "0"
	}
	childIDfmt += childId
	return bookId + "-" + childIDfmt
}

func getSingleImage(i int, uri, bookId string) {
	ext := util2.FileExt(uri)
	sortId := util2.GenNumberSorted(i + 1)
	log.Printf("Get %s  %s\n", sortId, uri)
	fileName := sortId + ext
	dest := config.GetDestPath(uri, bookId, fileName)
	curl2.FastGet(uri, dest, nil, true)
}

func getImageUrls(bookUrl string) (imgUrls *[]string, iiifUrls *[]string) {
	var manifest = new(Manifest)
	bs, err := curl2.Get(bookUrl, nil)
	if err != nil {
		return
	}

	if bs[0] != 123 {
		for i := 0; i < len(bs); i++ {
			if bs[i] == 123 {
				bs = bs[i:]
				break
			}
		}
	}
	if err = json.Unmarshal(bs, manifest); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	if len(manifest.Sequences) == 0 {
		return
	}
	i := len(manifest.Sequences[0].Canvases)
	imgUri := make([]string, 0, i)
	iiifUri := make([]string, 0, i)
	newWidth := ""
	//>6400使用原图
	if config.Conf.FullImageWidth > 6400 {
		newWidth = "full/full/"
	} else if config.Conf.FullImageWidth >= 1000 {
		newWidth = fmt.Sprintf("full/%d,/", config.Conf.FullImageWidth)
	}
	for _, sequence := range manifest.Sequences {
		for _, canvase := range sequence.Canvases {
			for _, image := range canvase.Images {
				//dezoomify-rs URL
				iiiInfo := fmt.Sprintf("%s/info.json", image.Resource.Service.Id)
				iiifUri = append(iiifUri, iiiInfo)

				//JPEG URL
				imgUrl := image.Resource.Id
				if newWidth != "" {
					imgUrl = strings.Replace(image.Resource.Id, "full/full/", newWidth, 1)
				}
				imgUri = append(imgUri, imgUrl)
			}
		}
	}
	return &imgUri, &iiifUri
}
