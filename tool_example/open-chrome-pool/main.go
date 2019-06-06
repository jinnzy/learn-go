package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"log"
	"sync"
	"time"
)

func sendkeys(endpoint string,counter string) chromedp.Tasks {
	//var buf []byte
	return chromedp.Tasks{
		// url
		chromedp.Navigate("https://trend-mon.foneshare.cn/data/query"),
		chromedp.WaitVisible(`#username`, chromedp.ByID), // ByID 意思是查找对应id标签名为#input1的
		chromedp.WaitVisible(`#password`, chromedp.ByID), // 等待#textarea1 标签渲染成功
		chromedp.WaitVisible(`btn-submit`,chromedp.NodeVisible),
		//chromedp.Sleep(1 * gettime.Second),
		// 输入falcon账号
		chromedp.SendKeys(`#username`, "15011424102", chromedp.ByID), // 向指定的html元素内输入内容
		//chromedp.Value(`#input1`, val1, chromedp.ByID), // 获取#input1的值
		//chromedp.SetValue(`#input2`, "test3", chromedp.ByID), // 改变对应元素中地value值，例子中会把next input框改为test3 input框
		// 输入密码
		chromedp.SendKeys(`#password`, "a15941473777", chromedp.ByID),
		// 点击登录
		chromedp.Click("btn-submit",chromedp.NodeVisible),
		//chromedp.Submit("btn-submit",chromedp.NodeVisible),
		// 等待
		chromedp.WaitVisible(`#endpoint-search`,chromedp.ByID),
		chromedp.WaitVisible(`#counter-search`,chromedp.ByID),
		chromedp.SendKeys(`#endpoint-search`,endpoint,chromedp.ByID),
		chromedp.SendKeys(`#counter-search`,counter,chromedp.ByID),
		//chromedp.Click("order",chromedp.NodeVisible),
		//chromedp.SendKeys("order",kb.ArrowDown,chromedp.BySearch),

		//chromedp.Click("order",chromedp.ByQueryAll),
		chromedp.SendKeys("select[name=order]",kb.ArrowDown), // 选择标签为select 且属性name=order的，执行操作往下选择
		chromedp.SendKeys("select[name=topN]",kb.ArrowUp+kb.ArrowUp+kb.ArrowUp), // 使用的默认搜错 chromedp.querysearch
		chromedp.Click("#go",chromedp.ByID),
		chromedp.WaitVisible(".table",chromedp.ByQuery),
		//chromedp.Click("s2b"),
		// 按照标签截图
		//chromedp.Screenshot(`#row`, &buf,chromedp.ByQuery, chromedp.NodeVisible,chromedp.ByID),
		//chromedp.ActionFunc(func(context.Context, cdp.Executor) error {
		//	return ioutil.WriteFile("test.png", buf, 0644)
		//}),
		//https://medium.com/@ribice/until-recently-i-never-knew-how-simple-it-could-be-to-automate-a-task-in-the-browser-55ff773941a 看下这个里面有下选框
		// 截完整的图
		//chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
		//	_, viewLayout, contentRect, err := page.GetLayoutMetrics().Do(ctxt, h)
		//	if err != nil {
		//		return err
		//	}
		//
		//	v := page.Viewport{
		//		X:      contentRect.X,
		//		Y:      contentRect.Y,
		//		Width:  viewLayout.ClientWidth, // or contentRect.Width,
		//		Height: contentRect.Height,
		//		Scale:  1,
		//	}
		//	log.Printf("Capture %#v", v)
		//	buf, err := page.CaptureScreenshot().WithClip(&v).Do(ctxt, h)
		//	if err != nil {
		//		return err
		//	}
		//	return ioutil.WriteFile("test.png", buf, 0644)
		//}),
		chromedp.Sleep(60000 * time.Second),
	}
}
// 执行sendkey函数的，由pool+go 并发调用
func taskSendkey(ctxt context.Context,wg *sync.WaitGroup, pool *chromedp.Pool,item map[string]string) {
	defer wg.Done()
	c,err := pool.Allocate(ctxt)
	if err != nil {
		//log.Printf("url (%d) `%s` error: %v", id, urlstr, err)
		log.Printf("url (error: %v", err)
		return
	}
	defer c.Release()

	err = c.Run(ctxt,sendkeys(item["endpoint"],item["counter"]))
	if err != nil {
		//log.Printf("url (%d) `%s` error: %v", id, urlstr, err)
		log.Printf(" error: %v", err)
		return
	}
}

func main() {
	var err error
	ctxt,cancel := context.WithCancel(context.Background())
	defer cancel()

	pool,err := chromedp.NewPool()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	var item [6]map[string]string
	// cpu初始化
	item[0] = map[string]string{"endpoint": "vlnx","counter": "cpu.idle"}
	// memory
	item[1] = map[string]string{"endpoint": "vlnx","counter": "mem.memfree"}
	// os ext4
	//item[2] = map[string]string{"endpoint": "vlnx","counter": "df.bytes.free.percent/fstype=ext4,mount=/"}
	item[2] = map[string]string{"endpoint": "vlnx","counter": "df.bytes.free/fstype=ext4,mount=/"}
	// os xfs
	//item[3] = map[string]string{"endpoint": "vlnx","counter": "df.bytes.free.percent/fstype=xfs,mount=/"}
	item[3] = map[string]string{"endpoint": "vlnx","counter": "df.bytes.free/fstype=xfs,mount=/"}
	// mysql ext4 data
	item[4] = map[string]string{"endpoint": "vlnx022","counter": "df.bytes.free.percent/fstype=ext4,mount=/opt/mysql/data"}
	// mongo xfs data
	item[5] = map[string]string{"endpoint": "vlnx","counter": "df.bytes.free.percent/fstype=xfs,mount=/opt/mongo"}

	for _,k := range item{

		wg.Add(1)
		go taskSendkey(ctxt,&wg, pool,k)
	}
	wg.Wait()
	time.Sleep(10 * time.Second)
}