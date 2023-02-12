package main

import (
	"bookget/config"
	"bookget/site/China/cuhk"
	"bookget/site/China/gzlib"
	"bookget/site/China/idp"
	"bookget/site/China/luoyang"
	"bookget/site/China/nlc"
	"bookget/site/China/npmtw"
	"bookget/site/China/rbkdocnpmtw"
	"bookget/site/China/szlib"
	"bookget/site/China/twnlc"
	"bookget/site/China/usthk"
	"bookget/site/China/wzlib"
	"bookget/site/Europe/bavaria"
	"bookget/site/Europe/berlin"
	"bookget/site/Europe/bluk"
	"bookget/site/Europe/oxacuk"
	"bookget/site/Japan/emuseum"
	"bookget/site/Japan/kanjikyoto"
	"bookget/site/Japan/keio"
	"bookget/site/Japan/khirin"
	"bookget/site/Japan/kotenseki"
	"bookget/site/Japan/kyoto"
	"bookget/site/Japan/national"
	"bookget/site/Japan/ndl"
	"bookget/site/Japan/niiac"
	"bookget/site/Japan/utokyo"
	"bookget/site/Japan/waseda"
	"bookget/site/Japan/yonezawa"
	"bookget/site/USA/familysearch"
	"bookget/site/USA/harvard"
	"bookget/site/USA/hathitrust"
	"bookget/site/USA/loc"
	"bookget/site/USA/princeton"
	"bookget/site/USA/stanford"
	"bookget/site/Universal"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Init func(iTask int, uri string) (string, error)
type SiteMap map[string]Init

func (b SiteMap) RegisterCommand(command string, f Init) error {
	if _, exists := b[command]; exists {
		return errors.New("command already exists")
	}
	b[command] = f
	return nil
}

func (b SiteMap) Execute(command string, iTask int, uri string) (string, error) {
	if config.Conf.AutoDetect > 0 {
		if config.Conf.AutoDetect == 1 {
			return Universal.StartDownload(iTask, uri)
		}
		if config.Conf.AutoDetect == 2 {
			return Universal.AutoDetectManifest(iTask, uri)
		}
	}
	if com, exists := b[command]; exists {
		return com(iTask, uri)
	}
	urlType := contentType(uri)
	if urlType == "json" {
		return Universal.AutoDetectManifest(iTask, uri)
	} else if urlType != "html" {
		return Universal.StartDownload(iTask, uri)
	}
	msg := fmt.Sprintf("Unsupported URL: %s", uri)
	return "", errors.New(msg)
}

func RegisterCommand() (err error) {
	//001.中国国家图书馆
	err = Site.RegisterCommand("read.nlc.cn", nlc.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//001.中国国家图书馆
	err = Site.RegisterCommand("mylib.nlc.cn", nlc.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//002.哈佛大学图书馆
	err = Site.RegisterCommand("iiif.lib.harvard.edu", harvard.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//003.中国台北图书馆
	err = Site.RegisterCommand("rbook.ncl.edu.tw", twnlc.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//004.hathitrust 数字图书馆
	err = Site.RegisterCommand("babel.hathitrust.org", hathitrust.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//005.普林斯顿大学图书馆
	err = Site.RegisterCommand("catalog.princeton.edu", princeton.Init)
	err = Site.RegisterCommand("dpul.princeton.edu", princeton.InitDpul)
	if err != nil {
		fmt.Println(err)
		return
	}
	//006.京都大学图书馆
	err = Site.RegisterCommand("rmda.kulib.kyoto-u.ac.jp", kyoto.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//007.美国国会图书馆
	err = Site.RegisterCommand("www.loc.gov", loc.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//008.日本国立国会图书馆
	err = Site.RegisterCommand("dl.ndl.go.jp", ndl.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//009.日本E国宝eMuseum
	err = Site.RegisterCommand("emuseum.nich.go.jp", emuseum.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//010.日本宫内厅书陵部（汉籍集览）
	err = Site.RegisterCommand("db2.sido.keio.ac.jp", keio.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//011.东京大学东洋文化研究所（汉籍善本资料库）
	err = Site.RegisterCommand("shanben.ioc.u-tokyo.ac.jp", utokyo.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//012.香港中文大学图书馆
	err = Site.RegisterCommand("repository.lib.cuhk.edu.hk", cuhk.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//013.牛津大学博德利图书馆
	err = Site.RegisterCommand("digital.bodleian.ox.ac.uk", oxacuk.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//014.日本国立公文书馆（内阁文库）
	err = Site.RegisterCommand("www.digital.archives.go.jp", national.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//015.日本东洋文库
	err = Site.RegisterCommand("dsr.nii.ac.jp", niiac.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//016.日本早稻田大学图书馆
	err = Site.RegisterCommand("archive.wul.waseda.ac.jp", waseda.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	////017.韩国国家图书馆（已删除）
	//err = Site.RegisterCommand("lod.nl.go.kr", nlgokr.Init)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//018.新日本古典籍综合数据库
	err = Site.RegisterCommand("kotenseki.nijl.ac.jp", kotenseki.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//019.京都大学人文科学研究所 东方学数字图书博物馆
	err = Site.RegisterCommand("kanji.zinbun.kyoto-u.ac.jp", kanjikyoto.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//020.德国柏林国立图书馆
	err = Site.RegisterCommand("digital.staatsbibliothek-berlin.de", berlin.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//021.英国图书馆文本手稿
	err = Site.RegisterCommand("www.bl.uk", bluk.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//022.香港科技大学图书馆
	err = Site.RegisterCommand("lbezone.ust.hk", usthk.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//023.台北故宫博物院 - 古籍善本
	err = Site.RegisterCommand("rbk-doc.npm.edu.tw", rbkdocnpmtw.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//024.台北故宫博物院 - 典藏资料
	err = Site.RegisterCommand("digitalarchive.npm.gov.tw", npmtw.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//025.日本国立历史民俗博物馆
	err = Site.RegisterCommand("khirin-a.rekihaku.ac.jp", khirin.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//026.日本市立米泽图书馆
	err = Site.RegisterCommand("www.library.yonezawa.yamagata.jp", yonezawa.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//027.日本庆应义塾大学图书馆
	err = Site.RegisterCommand("dcollections.lib.keio.ac.jp", Universal.AutoDetectManifest)
	if err != nil {
		fmt.Println(err)
		return
	}
	//028.日本关西大学图书馆
	err = Site.RegisterCommand("www.iiif.ku-orcas.kansai-u.ac.jp", Universal.AutoDetectManifest)
	if err != nil {
		fmt.Println(err)
		return
	}
	//029.洛阳市图书馆
	err = Site.RegisterCommand("221.13.137.120:8090", luoyang.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//030.温州市图书馆
	err = Site.RegisterCommand("oyjy.wzlib.cn", wzlib.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//031.巴伐利亞州立圖書館東亞數字資源庫
	err = Site.RegisterCommand("ostasien.digitale-sammlungen.de", bavaria.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//032.小利兰·斯坦福大学图书馆
	err = Site.RegisterCommand("searchworks.stanford.edu", stanford.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//033.深圳市图书馆-古籍
	err = Site.RegisterCommand("yun.szlib.org.cn", szlib.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//034.美国犹他州家谱
	err = Site.RegisterCommand("www.familysearch.org", familysearch.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//035.广州大典
	err = Site.RegisterCommand("gzdd.gzlib.gov.cn", gzlib.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	//036.國際敦煌項目
	err = Site.RegisterCommand("idp.nlc.cn", idp.Init)
	err = Site.RegisterCommand("idp.bl.uk", idp.Init)
	err = Site.RegisterCommand("idp.orientalstudies.ru", idp.Init)
	err = Site.RegisterCommand("idp.afc.ryukoku.ac.jp", idp.Init)
	err = Site.RegisterCommand("idp.bbaw.de", idp.Init)
	err = Site.RegisterCommand("idp.bnf.fr", idp.Init)
	err = Site.RegisterCommand("idp.korea.ac.kr", idp.Init)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func contentType(sUrl string) string {
	if strings.Contains(sUrl, ".json") {
		return "json"
	}
	m := regexp.MustCompile(`\((\d+)-(\d+)\)`).FindStringSubmatch(sUrl)
	if m != nil {
		return "octet-stream"
	}

	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}
	client := http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}
	req, _ := http.NewRequest("GET", sUrl, nil)
	req.Header.Set("User-Agent", config.Conf.UserAgent)
	req.Header.Set("Range", "bytes=0-0")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	ret := ""
	//application/ld+json
	bodyType := resp.Header.Get("content-type")
	m = strings.Split(bodyType, ";")
	switch m[0] {
	case "application/ld+json":
		ret = "json"
		break
	case "application/json":
		ret = "json"
		break
	case "text/html":
		ret = "html"
		break
	}
	return ret
}
