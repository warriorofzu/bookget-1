package util

import (
	"bookget/config"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type UriMatch struct {
	Min  string
	Max  string
	IMin int
	IMax int
}

func SubText(text, from, to string) string {
	iPos := strings.Index(text, from)
	if iPos == -1 {
		return ""
	}
	subText := text[iPos:]
	iPos2 := strings.Index(subText, to)
	if iPos2 == -1 {
		return ""
	}
	return subText[:iPos2]
}

func GetUriMatch(uri string) (u UriMatch, ok bool) {
	m := regexp.MustCompile(`\((\d+)-(\d+)\)`).FindStringSubmatch(uri)
	if m == nil {
		return u, false
	}

	u.Min = m[1]
	u.Max = m[2]
	i, _ := strconv.Atoi(u.Min)
	u.IMin = i
	iMax, _ := strconv.Atoi(u.Max)
	u.IMax = iMax

	return u, true
}

func GetHostUrl(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	var hostUrl = fmt.Sprintf("%s://%s/", u.Scheme, u.Host)
	return hostUrl
}

func LoopIndexStart(size int) int {
	//用户自定义起始页
	i := 0

	if config.Conf.PageSeq > i && config.Conf.PageSeq < size {
		i = config.Conf.PageSeq - 1
	}
	return i
}
