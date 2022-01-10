package main

import (
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

type rc4Cipher struct {
	s    [256]uint32
	x, y uint32

	l int

	i, j, i0, j0, tmp uint8
}

func (c *rc4Cipher) reset(key []byte) *rc4Cipher {
	for i := uint32(0); i < 256; i++ {
		c.s[i] = i
	}
	// 初始变量需要做好赋值
	c.i, c.j, c.j0, c.l = 0, 0, 0, len(key)
	for i := 0; i < 256; i++ {
		c.j0 += uint8(c.s[i]) + key[i%c.l]
		c.s[i], c.s[c.j0] = c.s[c.j0], c.s[i]
	}
	c.tmp = uint8(c.s[key[0]])
	return c
}

func (c *rc4Cipher) XORKeyStream(dst, src []byte) {
	c.i0, c.j0 = c.i, c.j
	for k, v := range src {
		c.i0++
		c.x = c.s[c.i0]
		c.j0 += uint8(c.x)
		c.y = c.s[c.j0]
		c.s[c.i0], c.s[c.j0] = c.y, c.x
		dst[k] = v ^ uint8(c.s[uint8(c.x+c.y)]) ^ c.tmp
	}
	c.i, c.j = c.i0, c.j0
}

func md5str(s string) string {
	h := md5.New()
	h.Write([]byte(s)) // 返回16字节的string
	return string(h.Sum(nil))
}

func encryptCipherKey(key string, buf []byte) (string, cipher.Stream, error) {
	// 有点类似与 将一个int64  拆成 多个int8的操作   int64 []int8
	unix := time.Now().Unix()
	buf[60] = byte(unix >> 56)
	buf[61] = byte(unix >> 48)
	buf[62] = byte(unix >> 40)
	buf[63] = byte(unix >> 32)
	buf[64] = byte(unix >> 24)
	buf[65] = byte(unix >> 16)
	buf[66] = byte(unix >> 8)
	buf[67] = byte(unix) // 存入时间戳,控制有效性
	_, err := rand.Read(buf[68:72])
	if err != nil {
		return "", nil, err
	}
	seed := buf[68]
	tmp := append(buf[60:72], key...) // 搞几个随机数混淆
	c := new(rc4Cipher).reset(append([]byte{seed}, key...))
	// 加密和解密需要初始化两个对象
	c.XORKeyStream(buf[1:], tmp)
	buf[0] = seed // 携带1字节随机数
	return base64.RawStdEncoding.EncodeToString(buf[:len(tmp)+1]), c.reset(tmp), nil
}

func decryptCipherKey(key, enc string, buf []byte) (cipher.Stream, error) {
	if enc == "" {
		return nil, NewWebErr("must be encrypted", http.StatusUnauthorized)
	}
	src, err := base64.RawStdEncoding.DecodeString(enc)
	if err != nil {
		return nil, err
	}

	if n := len(src) - 1; n >= len(key)+12 {
		c := new(rc4Cipher).reset(append([]byte{src[0]}, key...))
		c.XORKeyStream(buf, src[1:])
		if string(buf[12:n]) == key {
			t := time.Now().Unix() - (int64(buf[7]) | int64(buf[6])<<8 |
				int64(buf[5])<<16 | int64(buf[4])<<24 | int64(buf[3])<<32 |
				int64(buf[2])<<40 | int64(buf[1])<<48 | int64(buf[0])<<56)
			if t < limitKeyTime && t > -limitKeyTime {
				// 在规定秒时间内的秘钥才有效
				return c.reset(buf[:n]), nil
			}
		}
	}
	return nil, NewWebErr("key decrypt error", http.StatusUnauthorized)
}
