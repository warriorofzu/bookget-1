package gohttp

import (
	"net/http/cookiejar"
	"time"
)

// Options object
type Options struct {
	Debug       bool
	Concurrency uint //CPU核数
	BaseURI     string
	Timeout     float32
	timeout     time.Duration
	Query       interface{}
	Headers     map[string]interface{}
	Cookies     interface{}
	CookieFile  string
	CookieJar   *cookiejar.Jar
	FormParams  map[string]interface{}
	JSON        interface{}
	XML         interface{}
	Proxy       string
	DestFile    string //保存到本地文件
	Overwrite   bool   //覆蓋文件
}
