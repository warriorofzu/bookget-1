package nlc

import (
	"sync"
)

type DownloadTask struct {
	Index    int
	Url      string
	Domain   string
	SavePath string
	BookId   string
}

//Catalog 目录
type Catalog struct {
	Success bool      `json:"success"`
	Msg     string    `json:"msg"`
	Obj     []Chapter `json:"obj"`
}

//Chapter 章节
type Chapter struct {
	ChapterName1 string `json:"chapter_name1"` //章节1名称
	ChapterName2 string `json:"chapter_name2"` //章节2名称
	ChapterNum1  string `json:"chapter_num1"`  //章节1数字（卷数1）
	ChapterNum2  string `json:"chapter_num2"`  //章节2数字（卷数2）
}

//最近一次的章节
var LastChapter SafeLastChapter

type SafeLastChapter struct {
	ChapterName1 string `json:"chapter_name1"` //章节1名称
	ChapterName2 string `json:"chapter_name2"` //章节2名称
	ChapterNum1  string `json:"chapter_num1"`  //章节1数字（卷数1）
	ChapterNum2  string `json:"chapter_num2"`  //章节2数字（卷数2）
	Mux          sync.Mutex
}
