package main

import (
	"crypto/cipher"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var clientHttp *http.Client

// 客户端无论如何都要探测一下服务器那边是否有加密
func testServer(url string, buf []byte) (key string, c cipher.Stream, err error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return
	}
	if useEncrypt != "" {
		useEncrypt = md5str(useEncrypt)
		key, c, err = encryptCipherKey(useEncrypt, buf)
		if err != nil {
			return
		}
		req.Header.Set(encryptFlag, key)
	}
	resp, err := clientHttp.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = NewWebErr("check encryption")
	}
	return
}

// 获取服务器文件大小,用于断点上传功能
func getServerSize(url string) (int64, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set(headerType, janEncoded)
	resp, err := clientHttp.Do(req)
	if err != nil {
		return 0, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return 0, nil // 服务器没有文件
	}
	return parseInt(resp.Header.Get(janbarLength))
}

func clientPost(data, url string, point bool, key string, c cipher.Stream, buf []byte) error {
	var (
		size, cur int64
		path      string
		body      io.ReadSeeker
		err       error
	)

	if point { // 断点上传,获取服务器文件大小
		cur, err = getServerSize(url)
		if err != nil {
			return err
		}
	}

	if len(data) > 1 && data[0] == '@' {
		path = data[1:]
		fr, err := os.Open(path)
		if err != nil {
			return err
		}
		//goland:noinspection GoUnhandledErrorResult
		defer fr.Close()

		fi, err := fr.Stat()
		if err != nil {
			return err
		}
		size, body = fi.Size(), fr
	} else {
		// 不是文件,则上传一段文本内容
		sr := strings.NewReader(data)
		path, size, body = "<string data>", sr.Size(), sr
	}

	if cur > 0 {
		if cur >= size {
			return NewWebErr("file upload is complete")
		}
		// 断点上传时,将文件定位到指定位置
		_, err = body.Seek(cur, io.SeekStart)
		if err != nil {
			return err
		}
	}

	pr := handleWriteReadData(&handleData{
		cur:    cur,
		handle: body.Read,
		cipher: c,
	}, "POST> "+path, size)
	defer pr.Close()

	req, err := http.NewRequest(http.MethodPost, url, pr)
	if err != nil {
		return err
	}

	req.Header.Set(headerType, janEncoded) // 表示使用工具上传
	req.Header.Set(janbarLength, string(strconv.AppendInt(buf[:0], size, 10)))
	if point {
		// 告诉服务器断点续传的上传数据
		req.Header.Set(headPoint, string(strconv.AppendInt(buf[:0], cur, 10)))
	}
	if key != "" {
		// 告诉服务器,加密通信
		req.Header.Set(encryptFlag, key)
	}

	resp, err := clientHttp.Do(req)
	if err != nil {
		return err
	}
	if resp.Body != nil {
		if resp.StatusCode != http.StatusOK {
			_, _ = io.CopyBuffer(os.Stdout, resp.Body, buf)
		} else {
			_, _ = io.CopyBuffer(ioutil.Discard, resp.Body, buf)
		}
		//goland:noinspection GoUnhandledErrorResult
		resp.Body.Close()
	}
	return nil
}

func clientGet(url, output string, point bool, key string, c cipher.Stream, buf []byte) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	if output == "" {
		output = filepath.Base(req.URL.Path)
	}

	fileFlag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	fi, err := os.Stat(output)
	if err == nil {
		if fi.IsDir() {
			return NewWebErr(output + "is dir")
		}
		if point {
			fileFlag = os.O_CREATE | os.O_APPEND
			sSize := string(strconv.AppendInt(buf[:0], fi.Size(), 10))
			// 断点续传,设置规定的header,服务器负责解析并处理
			req.Header.Set("Range", "bytes="+sSize+"-")
			req.Header.Set(janbarLength, sSize) // 告诉服务器,从哪个位置下载
		}
	}
	fw, err := os.OpenFile(output, fileFlag, fileMode)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer fw.Close()

	if key != "" {
		req.Header.Set(encryptFlag, key)
	}

	resp, err := clientHttp.Do(req)
	if err != nil {
		return err
	}
	if resp.Body == nil {
		return NewWebErr("body is null")
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	var size, cur int64
	switch resp.StatusCode {
	case http.StatusOK: // 刚开始下载
		size, err = parseInt(resp.Header.Get(headerLength))
		if err != nil {
			return err
		}
	case http.StatusPartialContent: // 获取断点位置
		cur, _, size, err = scanSize(resp.Header)
		if err != nil {
			return err
		}
	case http.StatusRequestedRangeNotSatisfiable:
		// 已经下载完毕,无需重复下载
		//  相当于实现    /dev/null的操作
		// 就是什么都不去
		size, _ = io.CopyBuffer(io.Discard, resp.Body, buf)
		fmt.Printf("[%d bytes data]\n", size)
		return nil
	default:
		_, _ = io.CopyBuffer(os.Stdout, resp.Body, buf)
		return nil // 打印错误
	}

	pw := handleWriteReadData(&handleData{
		cur:       cur,
		handle:    fw.Write,
		cipher:    c,
		hashAfter: true,
	}, "GET > "+output, size)
	_, err = io.CopyBuffer(pw, resp.Body, buf)
	pw.Close()
	return err
}
