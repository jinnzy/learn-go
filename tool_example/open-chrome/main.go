package main

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"io/ioutil"
	"log"
)

func sendkeys() chromedp.Tasks {
	//var buf []byte
	return chromedp.Tasks{
		// url
		chromedp.Navigate("https://trend-mon.foneshare.cn/data/query"),
		chromedp.WaitVisible(`#username`, chromedp.ByID), // ByID 意思是查找对应id标签名为#input1的
		chromedp.WaitVisible(`#password`, chromedp.ByID), // 等待#textarea1 标签渲染成功
		chromedp.WaitVisible(`btn-submit`,chromedp.NodeVisible),
		//chromedp.Sleep(1 * gettime.Second),
		// 输入账号
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
		chromedp.SendKeys(`#endpoint-search`,"vlnx",chromedp.ByID),
		chromedp.SendKeys(`#counter-search`,"cpu.idle",chromedp.ByID),
		//chromedp.Click("order",chromedp.NodeVisible),
		//chromedp.SendKeys("order",kb.ArrowDown,chromedp.BySearch),

		//chromedp.Click("order",chromedp.ByQueryAll),
		chromedp.SendKeys("select[name=order]",kb.ArrowDown), // 选择标签为select 且属性name=order的，执行操作往下选择
		chromedp.SendKeys("select[name=topN]",kb.ArrowUp+kb.ArrowUp+kb.ArrowUp),
		chromedp.Click("#go",chromedp.ByID),
		//chromedp.Click("s2b"),
		// 按照标签截图
		//chromedp.Screenshot(`#row`, &buf,chromedp.ByQuery, chromedp.NodeVisible,chromedp.ByID),
		//chromedp.ActionFunc(func(context.Context, cdp.Executor) error {
		//	return ioutil.WriteFile("test.png", buf, 0644)
		//}),
		//https://medium.com/@ribice/until-recently-i-never-knew-how-simple-it-could-be-to-automate-a-task-in-the-browser-55ff773941a 看下这个里面有下选框
		// 截完整的图
		chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
			_, viewLayout, contentRect, err := page.GetLayoutMetrics().Do(ctxt, h)
			if err != nil {
				return err
			}

			v := page.Viewport{
				X:      contentRect.X,
				Y:      contentRect.Y,
				Width:  viewLayout.ClientWidth, // or contentRect.Width,
				Height: contentRect.Height,
				Scale:  1,
			}
			log.Printf("Capture %#v", v)
			buf, err := page.CaptureScreenshot().WithClip(&v).Do(ctxt, h)
			if err != nil {
				return err
			}
			return ioutil.WriteFile("test.png", buf, 0644)
		}),
	}
}
//func sendkeys2() chromedp.Tasks {
//	return chromedp.Tasks{
//		chromedp.Navigate("https://trend-mon.foneshare.cn/data/query"),
//		chromedp.WaitVisible(`#username`, chromedp.ByID), // ByID 意思是查找对应id标签名为#input1的
//		chromedp.WaitVisible(`#password`, chromedp.ByID), // 等待#textarea1 标签渲染成功
//		chromedp.WaitVisible(`btn-submit`,chromedp.NodeVisible),
//		//chromedp.Sleep(1 * gettime.Second),
//		chromedp.SendKeys(`#username`, "15011424102", chromedp.ByID), // 向指定的html元素内输入内容
//		//chromedp.Value(`#input1`, val1, chromedp.ByID), // 获取#input1的值
//		//chromedp.SetValue(`#input2`, "test3", chromedp.ByID), // 改变对应元素中地value值，例子中会把next input框改为test3 input框
//		chromedp.SendKeys(`#password`, "a15941473777", chromedp.ByID),
//		chromedp.Click("btn-submit",chromedp.NodeVisible),
//	}
//}

func runChrome(ctx context.Context,f func() chromedp.Tasks) {
	var err error

	// create context
	//ctxt, cancel := context.WithCancel(context.Background())
	ctxt, cancel := context.WithCancel(ctx)
	defer cancel()
	//ctxt, _ := context.WithCancel(context.Background())

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	//c, err := chromedp.New(ctxt, chromedp.WithTargets(client.New().WatchPageTargets(ctxt)), chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	err = c.Run(ctxt, f())
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	//err = c.Shutdown(ctxt) // 会关闭chrome进程
	//if err != nil {
	//	log.Fatal(err)
	//}

	// wait for chrome to finish
	//err = c.Wait()
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func main() {
	a := context.Background()
	runChrome(a,sendkeys)


	//var err error
	//// create context
	//ctxt, cancel := context.WithCancel(context.Background())
	//defer cancel()

	//// create chrome instance
	//c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	////c, err := chromedp.New(ctxt, chromedp.WithTargets(client.New().WatchPageTargets(ctxt)), chromedp.WithLog(log.Printf))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// run task list
	//err = c.Run(ctxt, sendkeys())
	//if err != nil {
	//	log.Fatal(err)
	//}

	// shutdown chrome
	//err = c.Shutdown(ctxt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// wait for chrome to finish
	//err = c.Wait()
	//if err != nil {
	//	log.Fatal(err)
	//}
}

