package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//GenNumberSorted 生成文件名排序
func GenNumberSorted(i int) string {
	text := strconv.Itoa(i)
	s := ""
	for k := 4 - len(text); k > 0; k-- {
		s += "0"
	}
	return fmt.Sprintf("%s%d", s, i)
}

func GenNumberLimitLen(i int, iLen int) string {
	text := strconv.Itoa(i)
	s := ""
	for k := iLen - len(text); k > 0; k-- {
		s += "0"
	}
	return fmt.Sprintf("%s%d", s, i)
}

func LetterNumberEscape(s string) string {
	m := regexp.MustCompile(`([A-Za-z0-9-_]+)`).FindAllString(s, -1)
	if m != nil {
		s = strings.Join(m, "")
	}
	return s
}
