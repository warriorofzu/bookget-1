package bluk

import (
	"bookget/config"
	"bookget/lib/curl"
	util2 "bookget/lib/util"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`Viewer.aspx\?ref=([\S]+)`).FindStringSubmatch(taskUrl)
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
	canvases := getImageUrls(taskUrl)
	log.Printf("A total of %d pages.\n", canvases.Size)
	destPath := config.CreateDirectory(taskUrl, bookId)
	util2.CreateShell(destPath, canvases.IiifUrls, nil)
	fmt.Println("Please run the file [dezoomify-rs.urls] to start the download.")
}

func getImageUrls(taskUrl string) (canvases Canvases) {
	bs, err := curl.GetRedirects(taskUrl, nil, 3)
	if err != nil {
		return
	}
	text := string(bs)
	//        <input type="hidden" name="PageList" id="PageList" value="##||or_6814!1_fs001r||or_6814!1_fs001v||or_6814!1_f001r||or_6814!1_f001v||or_6814!1_f002r||or_6814!1_f002v||or_6814!1_f003r||or_6814!1_f003v||or_6814!1_f004r||or_6814!1_f004v||or_6814!1_f005r||or_6814!1_f005v||or_6814!1_f006r||or_6814!1_f006v||or_6814!1_f007r||or_6814!1_f007v||or_6814!1_f008r||or_6814!1_f008v||or_6814!1_f009r||or_6814!1_f009v||or_6814!1_f010r||or_6814!1_f010v||or_6814!1_f011r||or_6814!1_f011v||or_6814!1_f012r||or_6814!1_f012v||or_6814!1_f013r||or_6814!1_f013v||or_6814!1_f014r||or_6814!1_f014v||or_6814!1_f015r||or_6814!1_f015v||or_6814!1_f016r||or_6814!1_f016v||or_6814!1_f017r||or_6814!1_f017v||or_6814!1_f018r||or_6814!1_f018v||or_6814!1_f019r||or_6814!1_f019v||or_6814!1_f020r||or_6814!1_f020v||or_6814!1_f021r||or_6814!1_f021v||or_6814!1_f022r||or_6814!1_f022v||or_6814!1_f023r||or_6814!1_f023v||or_6814!1_f024r||or_6814!1_f024v||or_6814!1_f025r||or_6814!1_f025v||or_6814!1_f026r||or_6814!1_f026v||or_6814!1_f027r||or_6814!1_f027v||or_6814!1_f028r||or_6814!1_f028v||or_6814!1_f029r||or_6814!1_f029v||or_6814!1_f030r||or_6814!1_f030v||or_6814!1_f031r||or_6814!1_f031v||or_6814!1_f032r||or_6814!1_f032v||or_6814!1_f033r||or_6814!1_f033v||or_6814!1_f034r||or_6814!1_f034v||or_6814!1_f035r||or_6814!1_f035v||or_6814!1_f036r||or_6814!1_f036v||or_6814!1_f037r||or_6814!1_f037v||##||or_6814!1_fblefv||or_6814!1_fbrigr||##||or_6814!1_fblefr||or_6814!1_fbrigv||or_6814!1_fbspi" />
	match := regexp.MustCompile(`id="PageList"[\s]+value=["']([\S]+)["']`).FindStringSubmatch(text)
	if match == nil {
		return
	}
	m := strings.Split(match[1], "||")
	size := len(m)
	canvases.ImgUrls = make([]string, 0, size)
	canvases.IiifUrls = make([]string, 0, size)
	for _, v := range m {
		if v == "##" {
			continue
		}
		dziUrl := fmt.Sprintf("http://www.bl.uk/manuscripts/Proxy.ashx?view=%s.xml", v)
		canvases.IiifUrls = append(canvases.IiifUrls, dziUrl)
	}
	canvases.Size = len(canvases.IiifUrls)
	return
}
