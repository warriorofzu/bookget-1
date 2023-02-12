package util

import (
	"bookget/config"
	"context"
	"fmt"
	"os"
	"strings"
)

func CreateShell(destPath string, iiifUrls []string, header map[string]string) {
	if string(os.PathSeparator) == "\\" {
		osWin(destPath, iiifUrls, header)
	} else {
		osLinux(destPath, iiifUrls, header)
	}
}

func osWin(destPath string, iiifUrls []string, header map[string]string) {
	if len(iiifUrls) <= 0 {
		return
	}
	text := "@echo on\r\n@echo downloading...\r\nsetlocal enabledelayedexpansion\r\n"
	cookie := ""
	for k, v := range header {
		kk := strings.Replace(k, "-", "", -1)
		text += fmt.Sprintf("set \"%s=%s\"\r\n", kk, v)
		cookie += fmt.Sprintf(" -H \"%s:'%%%s%%'\" ", k, kk)
	}

	for k, v := range iiifUrls {
		sortId := GenNumberSorted(k + 1)
		iifUrl := strings.Replace(v, "%", "%%", -1)
		text += fmt.Sprintf("%s -l --compression 0 %s \"%s\" %s.jpg\r\n", config.Conf.DezoomifyRs, cookie, iifUrl, sortId)
		//cookie, iifUrl, sortId
	}
	text += "\r\n:pause"
	dest := fmt.Sprintf("%s\\dezoomify-rs.urls.bat", destPath)
	FileWrite([]byte(text), dest)
}

func osLinux(destPath string, iiifUrls []string, header map[string]string) {
	if len(iiifUrls) <= 0 {
		return
	}
	//生成dezoomify-rs < urls.txt
	text := "#!/bin/sh\necho downloading...\n"
	cookie := ""
	for k, v := range header {
		kk := strings.Replace(k, "-", "", -1)
		text += fmt.Sprintf("%s=\"%s\"\n", kk, v)
		cookie += fmt.Sprintf(" -H \"%s:'${%s}'\" ", k, kk)
	}
	for k, v := range iiifUrls {
		sortId := GenNumberSorted(k + 1)
		iifUrl := strings.Replace(v, "%", "%%", -1)
		text += fmt.Sprintf("%s -l --compression 0 %s \"%s\" %s.jpg\n", config.Conf.DezoomifyRs, cookie, iifUrl, sortId)
	}

	text += "\n"
	dest := fmt.Sprintf("%s/dezoomify-rs.urls.sh", destPath)
	FileWrite([]byte(text), dest)

	ctx, cancelFunc := context.WithCancel(context.TODO())
	RunCommand(ctx, "chmod +x "+dest)
	cancelFunc()
}
