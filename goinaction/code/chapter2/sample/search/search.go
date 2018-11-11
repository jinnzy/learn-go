package search

import (
	"log"
	"sync"
)

// 注册用于搜索的匹配器的映射
// 声明为Matcher类型的映射（map），这个映射以string类型值作为键，Matcher类型值作为映射后的值

var matchers = make(map[string]Matcher)

// func声明函数，函数名，参数
func Run(searchTerm string)  {
	// 获取需要搜索的数据源列表
	feeds,err := RetrieveFeeds()
	if err != nil{
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接受匹配后的结果
	results := make(chan *Result)

	// 构造一个waitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	// 设置需要等待处理
	// 每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个goroutine来查找结果
	for _,feed := range feeds{
		// 获取一个匹配器用于查找
		matcher, exists := matcher[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		go func(matcher Matcher, feed *Feed) {
			Match(matcher,feed,searchTerm,results)
		}()
	}
}