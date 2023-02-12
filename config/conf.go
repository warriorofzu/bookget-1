package config

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strings"
)

type Input struct {
	DUrl       string //单个输入URL
	UrlsFile   string //输入urls.txt
	CookieFile string //输入cookie.txt
	PageSeq    int    //多图，从第几个开始（只对下载是多张图片的图书馆有效）
	SaveFolder string //下载文件存放目录，默认为当前文件夹下 Downloads 目录下
	//;生成 dezoomify-rs 可用的文件(默认生成文件名 dezoomify-rs.urls.txt）
	// ;0 = 禁用，1=启用 （只对支持的图书馆有效）
	UseNumericFilename int    //下载文件名，是否只使用数字序号？0=否，1=是（目前只对国图生效）
	FullImageWidth     int    //;全高清图下载时，指定宽度像素（16开纸185mm*260mm，像素2185*3071）
	UseCDN             int    //是否使用CDN加速？ 1=是，0=否（目前仅美国国会图书馆 下载的图片类型JP2生效）
	UserAgent          string //自定义UserAgent
	AutoDetect         int    //自动检测下载URL。可选值[0|1|2]，;0=默认;1=通用批量下载（类似IDM、迅雷）;2= IIIF manifest.json 自动检测下载图片
	MergePDFs          int    //;台北故宫博物院 - 善本古籍，是否整册合并一个PDF下载？0=否，1=是。整册合并一个PDF遇到某一册最后一章节【无影像】会导致下载失败。 如：新刊校定集注杜詩 三十六卷 第二十四冊 聞惠子過東溪 无影像
	DezoomifyRs        string //dezoomify-rs.exe 路径
	FileExt            string //指定下载的扩展名
	Threads            uint
	Help               bool
	Version            bool
}

func Init(ctx context.Context) bool {

	//dir, _ := os.Executable()
	dir, _ := os.Getwd()
	//cwd := filepath.Dir(dir)

	flag.StringVar(&Conf.UrlsFile, "i", "", "下载的URLs，指定任意本地文件，例如：urls.txt")
	flag.StringVar(&Conf.SaveFolder, "o", dir, "下载保存到目录")
	flag.IntVar(&Conf.PageSeq, "seq", 0, "图书起始页面数字")
	flag.IntVar(&Conf.FullImageWidth, "w", 7000, "指定图片宽度像素。推荐2400，若>6400为最大图")
	flag.IntVar(&Conf.UseNumericFilename, "fn", 1, "保存文件名规则。可选值[0|1]。0=中文名，1=数字名。仅对 read.nlc.cn 有效。")
	flag.StringVar(&Conf.UserAgent, "ua", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0", "user-agent")
	flag.IntVar(&Conf.MergePDFs, "mp", 0, "合并PDF文件下载，可选值[0|1]。0=否，1=是。仅对 rbk-doc.npm.edu.tw 有效。")
	flag.IntVar(&Conf.UseCDN, "cdn", 0, "使用CDN加速，可选值[0|1]。0=否，1=是。仅对 www.loc.gov 有效。")
	flag.StringVar(&Conf.CookieFile, "c", "", "指定cookie.txt文件路径")
	flag.StringVar(&Conf.FileExt, "ext", "", "指定文件扩展名[.jpg|.tif|.png]等")
	c := uint(runtime.NumCPU() * 2)
	flag.UintVar(&Conf.Threads, "n", c, "最大并发连接数")
	flag.IntVar(&Conf.AutoDetect, "a", 0, "自动检测下载URL。可选值[0|1|2]，;0=默认;\n1=通用批量下载（类似IDM、迅雷）;\n2= IIIF manifest.json 自动检测下载图片")
	flag.BoolVar(&Conf.Help, "h", false, "显示帮助")
	flag.BoolVar(&Conf.Version, "v", false, "显示版本")
	if string(os.PathSeparator) == "\\" {
		flag.StringVar(&Conf.DezoomifyRs, "rs", "dezoomify-rs.exe", "自定义dezoomify-rs路径，例如：C:\\Windows\\dezoomify-rs.exe")
	} else {
		flag.StringVar(&Conf.DezoomifyRs, "rs", "dezoomify-rs", "自定义dezoomify-rs路径，例如：/usr/local/bin/dezoomify-rs")
	}
	flag.Parse()

	k := len(os.Args)
	if k == 2 {
		if os.Args[1] == "-v" || os.Args[1] == "--version" {
			printVersion()
			return false
		}
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			printHelp()
			return false
		}
	}
	v := flag.Arg(0)
	if strings.HasPrefix(v, "http") {
		Conf.DUrl = v
	}
	if Conf.UrlsFile != "" && !strings.Contains(Conf.UrlsFile, string(os.PathSeparator)) {
		Conf.UrlsFile = dir + string(os.PathSeparator) + Conf.UrlsFile
	}
	//fmt.Printf("%+v", Conf)

	//保存目录处理
	_ = os.Mkdir(Conf.SaveFolder, os.ModePerm)
	return true
}

func printHelp() {
	printVersion()
	fmt.Println(`Usage: bookget [OPTION]... [URL]...`)
	flag.PrintDefaults()
	fmt.Println("Email bug reports, questions, discussions to Zhu D.W<zhudwi@foxmail.com>")
	fmt.Println("and/or open issues at https://github.com/deweizhu/bookget/issues")
}

func printVersion() {
	fmt.Printf("bookget v%s\n", version)
}

func CreateDirectory(sUrl, id string) string {
	u, _ := url.Parse(sUrl)
	domain := strings.Replace(u.Host, ":", "", 1)
	sPath := Conf.SaveFolder + string(os.PathSeparator) + domain
	if id != "" {
		sPath += "_" + LetterNumberEscape(id)
	}
	_ = os.Mkdir(sPath, os.ModePerm)
	return sPath
}

func GetDestPath(sUrl, id, filename string) string {
	u, _ := url.Parse(sUrl)
	domain := strings.Replace(u.Host, ":", "", 1)
	sPath := Conf.SaveFolder + string(os.PathSeparator) + domain
	if id != "" {
		sPath += "_" + LetterNumberEscape(id)
	}
	return sPath + string(os.PathSeparator) + filename

}

func LetterNumberEscape(s string) string {
	m := regexp.MustCompile(`([A-Za-z0-9-_]+)`).FindAllString(s, -1)
	if m != nil {
		s = strings.Join(m, "")
	}
	return s
}
