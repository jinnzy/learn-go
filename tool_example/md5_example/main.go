package main

import (
	"fmt"
	"github.com/learn-go/tool_example/md5_example/shorturl"
)

func main()  {
	url := "https://oss.foneshare.cn/prometheus/graph?g0.range_input=1h&g0.expr=up%7Bjob!~%22fs-sql-server%7Cfs-ats-exporter%7C.*service-endpoints%22%7D%20%3D%3D%200&g0.tab=1"
	cb := func(url, keyword string) bool {
		// todo 查db或缓存判断keyword是否重复
		return true
	}

	domain := "http://shorturl.cn"
	surl := shorturl.Generator(shorturl.CHARSET_ALPHANUMERIC, domain, url, cb)
	fmt.Print(surl)

}
