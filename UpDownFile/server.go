package main

import (
	"crypto/cipher"
	"crypto/md5"
	"fmt"
	"hash"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

const (
	htmlErrTpl = `<html>
<head><title>info</title></head>
<body>
<div style="text-align: center;">code:%d,msg:%s</div>
</body>
</html>`

	fileMode = fs.FileMode(
		0666,
	)
	headerType   = "Content-Type"
	janEncoded   = "application/janbar" // 使用本工具命令行的头
	headerLength = "Content-Length"
	janbarLength = "Janbar-Length"
	headPoint    = "Point"   // 标识断点上传
	encryptFlag  = "Encrypt" // header秘钥
	limitKeyTime = 10        // 有效秒数,注意客户端和服务器时间不一致问题
)

// 只有一个主路由
func upDownFile(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		//  通过pool 复用缓冲区
		pool = bytePool.Get().(*poolByte)
	)
	defer bytePool.Put(pool)
	// 通过switch 来进行判断
	switch r.Method {
	case http.MethodGet:
		err = handleGetFile(w, r, pool.buf)
	case http.MethodPost:
		err = handlePostFile(w, r, pool.buf)
	case http.MethodHead:
		if enc := r.Header.Get(encryptFlag); useEncrypt == "" {
			if enc != "" { // 服务器不加密,客户端加密,返回授权错误
				err = NewWebErr("check encryption", http.StatusUnauthorized)
			}
		} else {
			_, err = decryptCipherKey(useEncrypt, enc, pool.buf)
		}
	case http.MethodPut:
		err = handlePutFile(w, r, pool.buf)
	default:
		err = NewWebErr(r.Method + " not support")
	}
	if err != nil {
		e, ok := err.(*webErr)
		if !ok {
			e = &webErr{code: http.StatusInternalServerError, msg: err.Error()}
		}
		w.Header().Set(headerType, "text/html;charset=utf-8")
		w.WriteHeader(e.code) // 一定要先设置header,再写code,然后写消息体
		_, _ = fmt.Fprintf(w, htmlErrTpl, e.code, e.msg)
	}
}

// 渲染html模板需要的结构
type lineFileInfo struct {
	Index      int
	Type       string
	Size       string
	Time       string
	Href, Name string
}

func handleGetFile(w http.ResponseWriter, r *http.Request, buf []byte) error {
	path := filepath.Join(basePath, r.URL.Path)
	fi, err := os.Stat(path)
	if err != nil {
		return NewWebErr(path+" not found", http.StatusNotFound)
	}

	if r.Header.Get(headerType) == janEncoded {
		if fi.IsDir() {
			return NewWebErr("unable to get directory size")
		}
		// 获取服务器文件大小,用于断点上传文件,会返回curl断点上传命令
		size := string(strconv.AppendInt(buf[:0], fi.Size(), 10))
		w.Header().Set(janbarLength, size)
		//goland:noinspection HttpUrlsUsage
		_, _ = fmt.Fprintf(w, "curl -C %s -T file http://%s%s\n", size, r.Host, r.URL.Path)
		return nil
	}

	if fi.IsDir() {
		if useEncrypt != "" { // 加密方式不支持浏览目录,懒得写前端代码
			return NewWebErr("encrypt method not support list dir")
		}

		sortNum, _ := strconv.Atoi(r.FormValue("sort"))
		dir, err := sortDir(path, &sortNum) // 根据指定排序得到有序目录内容
		if err != nil {
			return err
		}

		info := make([]lineFileInfo, len(dir))
		for i, v := range dir {
			tmp := lineFileInfo{
				Index: i + 1,
				Size:  convertByte(buf[:0], v.Size()),
				Time:  string(v.ModTime().AppendFormat(buf[:0], "2006-01-02 15:04:05")),
				Name:  v.Name(),
			}

			href := append(buf[:0], url.PathEscape(v.Name())...)
			if v.IsDir() {
				tmp.Type = "D"
				href = append(href, '/')
			} else {
				tmp.Type = "F"
			}
			tmp.Href = string(href)
			info[i] = tmp
		}
		tpl, err := template.New("").Parse(htmlTpl)
		if err != nil {
			return err
		}
		err = tpl.Execute(w, map[string]interface{}{
			"sort": sortNum,
			"info": info,
		})
		if err != nil {
			return err
		}
	} else {
		var c cipher.Stream
		if useEncrypt != "" {
			// 使用加密传输,需要从header中获取秘钥
			c, err = decryptCipherKey(useEncrypt, r.Header.Get(encryptFlag), buf)
			if err != nil {
				return err
			}
		}

		// 尝试获取断点下载的位置,获取不到cur=0
		cur, _ := parseInt(r.Header.Get(janbarLength))
		pw := handleWriteReadData(&handleData{
			cur: cur, cipher: c,
			ResponseWriter: w,
		}, "GET > "+path, fi.Size())
		http.ServeFile(pw, r, path) // 支持断点下载
		pw.Close()
	}
	return nil
}

func handlePostFile(w http.ResponseWriter, r *http.Request, buf []byte) error {
	var (
		path      string
		size, cur int64
		fr        io.ReadCloser
		c         cipher.Stream

		fileFlag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	)

	switch r.Header.Get(headerType) {
	case "application/x-www-form-urlencoded":
		s, err := parseInt(r.Header.Get(headerLength))
		if err != nil {
			return err
		}
		// 普通二进制上传文件,消息体直接是文件内容
		fr, size, path = r.Body, s, filepath.Join(basePath, r.URL.Path)
	case janEncoded:
		s, err := parseInt(r.Header.Get(janbarLength))
		if err != nil {
			return err
		}
		if useEncrypt != "" {
			// 使用加密传输,需要从header中获取秘钥
			c, err = decryptCipherKey(useEncrypt, r.Header.Get(encryptFlag), buf)
			if err != nil {
				return err
			}
		}
		// 判断是断点上传,则cur为断点位置
		cur, err = parseInt(r.Header.Get(headPoint))
		if err == nil {
			fileFlag = os.O_CREATE | os.O_APPEND
		}
		// 本工具命令行上传文件
		fr, size, path = r.Body, s, filepath.Join(basePath, r.URL.Path)
	default:
		rf, rh, err := r.FormFile("file")
		if err != nil {
			return err
		}
		// 使用浏览器上传 或 curl -F "file=@C:\tmp.txt",这两种方式
		fr, size, path = rf, rh.Size, filepath.Join(basePath, r.URL.Path, rh.Filename)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer fr.Close()

	if useEncrypt != "" && c == nil {
		return NewWebErr("encrypt method must use cli")
	}

	fw, err := os.OpenFile(path, fileFlag, fileMode)
	if err != nil {
		return err
	}

	pw := handleWriteReadData(&handleData{
		cur:       cur,
		handle:    fw.Write,
		cipher:    c,
		hashAfter: true,
	}, "POST> "+path, size)
	_, err = io.CopyBuffer(pw, fr, buf)
	_ = fw.Close() // 趁早刷新缓存,因为要计算hash
	pw.Close()
	if err != nil {
		return err
	}
	_, err = w.Write(respOk)
	return err
}

// 只提供curl断点上传处理逻辑
func handlePutFile(w http.ResponseWriter, r *http.Request, buf []byte) error {
	if r.Body == nil {
		return NewWebErr("body is null")
	}
	//goland:noinspection GoUnhandledErrorResult
	defer r.Body.Close()

	if useEncrypt != "" {
		return NewWebErr("encrypt method not support curl")
	}

	var (
		fw        *os.File
		cur, size int64
		path      = filepath.Join(basePath, r.URL.Path)
	)
	fi, err := os.Stat(path)
	if err == nil {
		// 文件存在则校验断点上传的逻辑
		cur, _, size, err = scanSize(r.Header)
		if err == nil {
			fw, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, fileMode)
			//goland:noinspection GoUnhandledErrorResult
			defer fw.Close()
			if err != nil {
				return err
			}

			nSize := fi.Size()
			if nSize == size {
				return NewWebErr("file upload is complete")
			}

			if (cur == 0 && nSize > 0) || cur > nSize {
				//goland:noinspection HttpUrlsUsage
				if nSize == 0 {
					// 返回重建的上传命令
					_, _ = fmt.Fprintf(w, "curl -T file http://%s%s\n", r.Host, r.URL.Path)
				} else {
					// 返回指定位置的上传命令
					_, _ = fmt.Fprintf(w, "curl -C %d -T file http://%s%s\n", nSize, r.Host, r.URL.Path)
				}
				return nil
			}

			if cur > 0 { // 从指定位置继续写文件
				_, err = fw.Seek(cur, io.SeekStart)
				if err != nil {
					return err
				}
			}
		}
	}

	if fw == nil {
		// 没有经过上面断点续传则重新创建文件并上传
		fw, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fileMode)
		//goland:noinspection GoUnhandledErrorResult
		defer fw.Close()
		if err != nil {
			return err
		}
		size, err = parseInt(r.Header.Get(headerLength))
		if err != nil {
			return err
		}
		cur = 0
	}

	pw := handleWriteReadData(&handleData{cur: cur, handle: fw.Write}, "PUT > "+path, size)
	_, err = io.CopyBuffer(pw, r.Body, buf)
	pw.Close() // 关闭相关资源
	if err != nil {
		return err
	}
	_, err = w.Write(respOk)
	return err
}

type handleData struct {
	http.ResponseWriter

	cur       int64
	rate      chan int64
	sumHex    chan []byte
	pool      *poolByte
	hash      hash.Hash
	hashAfter bool // true表示加解密后数据计算hash
	cipher    cipher.Stream
	handle    func([]byte) (int, error)
}

func handleWriteReadData(p *handleData, prefix string, size int64) *handleData {
	if p.ResponseWriter != nil {
		// 这个是http服务的写入操作
		p.handle = p.ResponseWriter.Write
	}

	p.hash = md5.New()
	p.rate = make(chan int64)
	p.sumHex = make(chan []byte)
	p.pool = bytePool.Get().(*poolByte)
	go func() {
		pCur := "\r" + prefix + " %3d%%"
		for cur := range p.rate {
			fmt.Printf(pCur, cur*100/size)
		}
		fmt.Printf(pCur+" %02x\n", 100, <-p.sumHex)
		p.sumHex <- nil // 打印完成才能退出
	}()
	return p
}

func (p *handleData) add(n int) {
	p.cur += int64(n)
	select {
	case p.rate <- p.cur:
	default:
	}
}

func (p *handleData) grow(n int) []byte {
	if n > len(p.pool.buf) {
		p.pool.buf = make([]byte, n)
	}
	return p.pool.buf[:n] // 获取足够缓存
}

func (p *handleData) Write(b []byte) (n int, err error) {
	if p.cipher != nil {
		tmp := p.grow(len(b))
		p.cipher.XORKeyStream(tmp, b)
		if n, err = p.handle(tmp); n > 0 {
			if p.hashAfter {
				// 使用解密后数据计算hash
				p.hash.Write(tmp[:n])
			} else {
				// 使用加密前数据计算hash
				p.hash.Write(b[:n])
			}
		}
	} else if n, err = p.handle(b); n > 0 {
		p.hash.Write(b[:n])
	}
	p.add(n)
	return
}

func (p *handleData) Read(b []byte) (n int, err error) {
	if p.cipher != nil {
		tmp := p.grow(len(b))
		if n, err = p.handle(tmp); n > 0 {
			p.hash.Write(tmp[:n]) // 使用加密前数据计算hash
			p.cipher.XORKeyStream(b[:n], tmp[:n])
		}
	} else if n, err = p.handle(b); n > 0 {
		p.hash.Write(b[:n])
	}
	p.add(n)
	return
}

func (p *handleData) Close() {
	bytePool.Put(p.pool)
	close(p.rate)
	p.sumHex <- p.hash.Sum(nil)
	<-p.sumHex // 发送hash结果,确保打印结束
}
func convertByte(buf []byte, b int64) string {
	tmp, unit := float64(b), "B"
	for i := 1; i < len(unitByte); i++ {
		if tmp < unitByte[i].byte {
			tmp /= unitByte[i-1].byte
			unit = unitByte[i].unit
			break
		}
	}
	return string(strconv.AppendFloat(buf, tmp, 'f', 2, 64)) + unit
}
