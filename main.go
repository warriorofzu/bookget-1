package main

import (
	"bookget/config"
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
)

var Site = SiteMap{}

func main() {
	ctx := context.Background()

	//配置初始化
	if !config.Init(ctx) {
		os.Exit(0)
	}
	//注册站点
	RegisterCommand()

	//单个URL
	if config.Conf.DUrl != "" {
		ExecuteCommand(ctx, 1, config.Conf.DUrl)
		log.Print("Download complete.\n")
		return
	}

	//批量URLs
	if config.Conf.UrlsFile != "" {
		//加载配置文件
		bs, err := ioutil.ReadFile(config.Conf.UrlsFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		mUrls := strings.Split(string(bs), "\n")
		iCount := 0
		for i, sUrl := range mUrls {
			if sUrl == "" {
				continue
			}
			ExecuteCommand(ctx, i+1, sUrl)
			iCount++
		}
		log.Print("Download complete.\n")
		log.Printf("下载完成，共 %d 个任务，请到 %s 目录下查看。\n", iCount, config.Conf.SaveFolder)
		return
	}

	iCount := 0
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter an URL:")
		fmt.Print("-> ")
		sUrl, err := reader.ReadString('\n')
		if err != nil {
			//fmt.Printf("Error: %w \n", err)
			break
		}
		iCount++
		ExecuteCommand(ctx, iCount, sUrl)
	}
	log.Print("Download complete.\n")
	log.Printf("下载完成，共 %d 个任务，请到 %s 目录下查看。\n", iCount, config.Conf.SaveFolder)
}

func ExecuteCommand(ctx context.Context, i int, text string) {
	text = strings.Trim(text, "\r\n")
	u, err := url.Parse(text)
	if err != nil {
		return
	}
	msg, err := Site.Execute(u.Host, i, text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg)
	return
}
