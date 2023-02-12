package curl

import (
	"bookget/config"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type clientDoer interface {
	Do(req *fasthttp.Request, resp *fasthttp.Response) error
}

func GetRedirects(uri string, header map[string]string, maxRedirectsCount int) (body []byte, err error) {
	// Acquire a request instance
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	//req.SetRequestURI(uri)
	req.Header.Set("User-Agent", config.Conf.UserAgent)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	// Acquire a response instance
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	c := fasthttp.Client{}
	statusCode, _, err := doRequestFollowRedirects(req, resp, uri, maxRedirectsCount, &c)
	if err != nil {
		log.Printf("Client get failed: %s\n", err)
		return
	}
	//log.Printf("responseBody: %s bytes\n", len(resp.Body()))
	if statusCode != fasthttp.StatusOK {
		//log.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
		err = errors.New(fmt.Sprintf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode()))
		return
	}
	body = resp.Body()
	return
}

func doRequestFollowRedirects(req *fasthttp.Request, resp *fasthttp.Response, url string, maxRedirectsCount int, c clientDoer) (statusCode int, body []byte, err error) {
	redirectsCount := 0
	var rCookies []http.Cookie
	for {
		req.SetRequestURI(url)
		for _, v := range rCookies {
			req.Header.SetCookie(v.Name, v.Value)
		}
		if err = c.Do(req, resp); err != nil {
			break
		}
		statusCode = resp.Header.StatusCode()
		if !fasthttp.StatusCodeIsRedirect(statusCode) {
			break
		}
		//返回cookie
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

		redirectsCount++
		if redirectsCount > maxRedirectsCount {
			err = fasthttp.ErrTooManyRedirects
			break
		}
		location := resp.Header.Peek(fasthttp.HeaderLocation)
		if len(location) == 0 {
			err = fasthttp.ErrMissingLocation
			break
		}
		url = getRedirectURL(url, location)
	}
	return statusCode, body, err
}

func getRedirectURL(baseURL string, location []byte) string {
	u := fasthttp.AcquireURI()
	u.Update(baseURL)
	u.UpdateBytes(location)
	redirectURL := regexp.MustCompile(`(:80|:443)/`).ReplaceAllString(u.String(), "/")
	fasthttp.ReleaseURI(u)
	return redirectURL
}
