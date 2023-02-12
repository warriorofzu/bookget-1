package harvard

import (
	"bookget/lib/util"
	"bookget/site/Universal/iiif"
	"errors"
	"fmt"
	"log"
	"regexp"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := getBookId(taskUrl)
	if bookId == "" {
		return "", errors.New(fmt.Sprintf("Error ID: %s\n", taskUrl))
	}
	return StartDownload(iTask, taskUrl, bookId)
}

func getBookId(sUrl string) string {
	m := regexp.MustCompile(`manifests/view/([A-z0-9-_:]+)`).FindStringSubmatch(sUrl)
	if m != nil {
		return m[1]
	}
	m = regexp.MustCompile(`/manifests/([A-z0-9-_:]+)`).FindStringSubmatch(sUrl)
	if m != nil {
		return m[1]
	}
	return ""
}

func StartDownload(iTask int, text, bookId string) (msg string, err error) {
	name := util.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, text)

	manifestUrl := fmt.Sprintf("https://iiif.lib.harvard.edu/manifests/%s", bookId)
	iiif.StartDownload(manifestUrl, bookId)
	return
}
