package shorturl

import (
	"fmt"
	"strconv"
	"crypto/md5"
	"strings"
)

const (
	CHARSET_ALPHANUMERIC = iota
	CHARSET_RANDOM_ALPHANUMERIC
)

func getCharset(t int) string {
	switch t {
	case CHARSET_ALPHANUMERIC:
		return "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case CHARSET_RANDOM_ALPHANUMERIC:
		return "A0a12B3b4CDc56Ede7FGf8Hg9IhJKiLjkMNlmOPnQRopqrSstTuvUVwxWXyYzZ"
	default:
		panic("invalid charset type t:" + strconv.Itoa(t))
	}
}

// 生成6字符短key
// 根据url生成32字符的签名，将其分成4段，每段8位字符
// 循环处理4段8位字符，将每段转换成16进制与0x3FFFFFFF进行逻辑与操作，得到30位的无符号数
// 将30位数分成6段，依次得到6个0-61的数字索引查字符集表获得6位字符串
func generator6(charset ,url, hexMd5 string,  sectionNum int, cb func(url, keyword string) bool) string {
	for i := 0; i < sectionNum; i++ {
		// 将加密的字符按照8位一组变成切片，每次循环是一组， [0:8] [8: 16] 如果生成的keyword没有被使用，就会返回不循环了
		sectionHex := hexMd5[i*8:8+i*8]
		// sectionHex 转换为uint类型
		bits, _ := strconv.ParseUint(sectionHex, 16, 32)
		//把加密字符按照8位一组16进制与0x3FFFFFFF进行位与运算
		bits = bits & 0x3FFFFFFF

		keyword := ""
		for j := 0; j < 6; j++ {
			// 把得到的值与0x3D / 0x0000003D 进行位与运算
			idx := bits & 0x3D
			// 把取得的字符相加
			keyword = keyword + string(charset[idx])
			//每次循环按位右移5位
			bits = bits >> 5
		}

		if cb(url, keyword) {
			return keyword
		}
	}

	return ""
}

// 生成8字符短key
func generator8(charset, url, hexMd5 string,  sectionNum int, cb func(url, keyword string) bool) string {
	for i := 0; i < sectionNum; i++ {
		// [] 生成切片，位置从 i x 8 到 i x 8 + 8
		sectionHex := hexMd5[i*8:i*8+8]
		bits, _ := strconv.ParseUint(sectionHex, 16, 32)

		bits = bits & 0xFFFFFFFF
		keyword := ""
		for j := 0; j < 8; j++ {
			idx := bits & 0x3D
			keyword = keyword + string(charset[idx])
			bits = bits >> 4
		}

		if cb(url, keyword) {
			return keyword
		}
	}

	return ""
}

// 生成6-8字符的短链接，参数t表示字符集类型，回调函数(cb)用于检测短链接是否重复
// 起初生成6位的短链接，当四组6位短链接都重复时，再生成8位的短链接
func Generator(t int, domain string, url string, cb func(url, keyword string) bool) (shorturl string) {
	if domain == "" || url == "" || cb == nil {
		return
	}

	charset := getCharset(t)
	// 对传入网址进行md5加密
	hexMd5 := fmt.Sprintf("%x", md5.Sum([]byte(url)))
	// 获取加密字符长度 除以8 作为循环次数
	sections := len(hexMd5)/8

	keyword := generator6(charset, url, hexMd5, sections, cb)
	if keyword == "" {
		keyword = generator8(charset, url, hexMd5, sections, cb)
		if keyword == "" {
			return ""
		}
	}

	if strings.HasSuffix(domain, "/") {
		shorturl = domain + keyword
	}else {
		shorturl = domain + "/" + keyword
	}
	return
}
