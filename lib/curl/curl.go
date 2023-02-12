package curl

import (
	"bookget/config"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"strings"
)

//var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36"
func Get(uri string, header map[string]string) (body []byte, err error) {
	return request("GET", uri, nil, header)
}

func Post(uri string, data []byte, header map[string]string) (body []byte, err error) {
	if header == nil {
		header = make(map[string]string)
	}
	//application/x-www-form-urlencoded 或 application/json
	header["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	return request("POST", uri, data, header)
}
func PostJson(uri string, data []byte, header map[string]string) (body []byte, err error) {
	if header == nil {
		header = make(map[string]string)
	}
	//application/x-www-form-urlencoded 或 application/json
	header["Content-Type"] = "application/json; charset=UTF-8"
	return request("POST", uri, data, header)
}

func request(method, uri string, data []byte, header map[string]string) (body []byte, err error) {
	// Acquire a request instance
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(uri)
	req.Header.SetMethod(method)
	if header != nil && header["Content-Type"] != "" {
		req.Header.SetContentType(header["Content-Type"]) //application/x-www-form-urlencoded 或 application/json
	}
	req.Header.Set("User-Agent", config.Conf.UserAgent)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	req.SetBody(data)

	// Acquire a response instance
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = fasthttp.Do(req, resp)
	if err != nil {
		log.Printf("Client get failed: %s\n", err)
		return
	}
	//log.Printf("responseBody: %s bytes\n", len(resp.Body()))
	if resp.StatusCode() != fasthttp.StatusOK {
		//log.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
		err = errors.New(fmt.Sprintf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode()))
		return
	}
	body = resp.Body()
	return
}

func GetWithCookie(uri string, header map[string]string) (body []byte, cookies []*http.Cookie, err error) {
	// Acquire a request instance
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(uri)
	req.Header.Set("User-Agent", config.Conf.UserAgent)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	// Acquire a response instance
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = fasthttp.Do(req, resp)
	if err != nil {
		log.Printf("Client get failed: %s\n", err)
		return
	}
	//log.Printf("responseBody: %s bytes\n", len(resp.Body()))
	if resp.StatusCode() != fasthttp.StatusOK {
		//log.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
		err = errors.New(fmt.Sprintf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode()))
		return
	}
	body = resp.Body()
	//返回cookie
	resp.Header.VisitAllCookie(func(key, value []byte) {
		k := string(key)
		v := strings.Split(string(value), ";")
		text := v[0]
		pos := strings.Index(text, "=")
		if pos == -1 {
			return
		}
		cookies = append(cookies, &http.Cookie{Name: k, Value: text[pos+1:]})
	})
	return
}

func PostWithCookie(uri string, data []byte, header map[string]string) (body []byte, cookies *[]http.Cookie, err error) {
	// Acquire a request instance
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(uri)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/x-www-form-urlencoded; charset=UTF-8") //application/x-www-form-urlencoded 或 application/json
	req.Header.Set("User-Agent", config.Conf.UserAgent)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	req.SetBody(data)

	// Acquire a response instance
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = fasthttp.Do(req, resp)
	if err != nil {
		log.Printf("Client get failed: %s\n", err)
		return
	}
	//log.Printf("responseBody: %s\n", resp.Body())
	if resp.StatusCode() != fasthttp.StatusOK {
		log.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
		return
	}
	body = resp.Body()

	//返回cookie
	var rCookies []http.Cookie
	resp.Header.VisitAllCookie(func(key, value []byte) {
		k := string(key)
		v := strings.Split(string(value), ";")
		text := v[0]
		pos := strings.Index(text, "=")
		if pos == -1 {
			return
		}
		rCookies = append(rCookies, http.Cookie{Name: k, Value: text[pos+1:]})
	})
	cookies = &rCookies
	return
}
