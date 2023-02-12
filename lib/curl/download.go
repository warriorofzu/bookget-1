package curl

import (
	"bookget/config"
	"bookget/lib/util"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

//FastGet 下载器
func FastGet(uri, dest string, header map[string]string, ignore bool) (size int64, err error) {
	if ignore {
		//文件存在，跳过
		fi, err := os.Stat(dest)
		if err == nil && fi.Size() > 0 {
			return 0, nil
		}
	}
	return execDownload(uri, dest, "GET", nil, header)
}

func PostDownload(uri, dest string, data []byte, header map[string]string) (size int64, err error) {
	return execDownload(uri, dest, "POST", data, header)
}

func execDownload(uri, dest, method string, data []byte, header map[string]string) (size int64, err error) {
	var destTemp = fmt.Sprintf("%s.downloading", dest)
	file, err := os.Create(destTemp)
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
		if err == nil {
			os.Rename(destTemp, dest)
		}
	}()

	dl := &Download{}
	dl.startedAt = time.Now()
	dl.ctx = context.Background()
	client := &http.Client{}
	dl.Client = client

	var req = new(http.Request)
	if method == "POST" {
		body := bytes.NewReader(data)
		req, err = http.NewRequest(method, uri, body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = http.NewRequest(method, uri, nil)
	}
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", config.Conf.UserAgent)
	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	//处理返回结果
	rsp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func(d *Download) {
		_ = rsp.Body.Close()
		dl.StopProgress = true
		fmt.Fprintf(os.Stdout, "\r100%%[================================================>]  %s/%s  %s/s    in %s", util.ByteUnitString(int64(d.Size())),
			util.ByteUnitString(int64(d.TotalSize())), util.ByteUnitString(int64(d.AvgSpeed())), d.TotalCost())
		fmt.Println()
		log.Printf("Save as  %s  (%s)\n", dest, util.ByteUnitString(size))
	}(dl)

	dl.totalSize = uint64(rsp.ContentLength)
	// Allocate the file completely so that we can write concurrently
	file.Truncate(rsp.ContentLength)

	go dlProgressBar(dl)

	size, err = io.Copy(file, io.TeeReader(rsp.Body, dl))
	return size, err
}

func dlProgressBar(d *Download) {
	// Set default interval.
	if d.Interval == 0 {
		d.Interval = uint64(400 / runtime.NumCPU())
	}
	sleepd := time.Duration(d.Interval) * time.Millisecond
	for {
		if d.StopProgress {
			break
		}
		// Context check.
		select {
		case <-d.ctx.Done():
			return
		default:
		}

		// Run progress func.
		if d.TotalSize() <= 0 {
			return
		}
		pd := d.Size() * 100 / d.TotalSize()
		if pd == 100 {
			return
		}
		speed := "="
		max := int(pd)
		for k := 0; k < max; k += 2 {
			speed += "="
		}
		speed += ">"
		after := 50 - len(speed)
		for k := 0; k < after; k++ {
			speed += " "
		}
		fmt.Fprintf(os.Stdout, "\r%d%%[%s]  %s/%s  %s/s    in %s", pd, speed, util.ByteUnitString(int64(d.Size())),
			util.ByteUnitString(int64(d.TotalSize())), util.ByteUnitString(int64(d.AvgSpeed())), d.TotalCost())

		// Update last size
		atomic.StoreUint64(&d.lastSize, atomic.LoadUint64(&d.size))
		// Interval.
		time.Sleep(sleepd)
	}
}
