package main

import (
	"context"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"log"
	"math"
	"sync"
	"time"
)



func sendkeys(endpoint string,counter string,res *[]byte) chromedp.Tasks {

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
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(90).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
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
	}
}
// 执行sendkey函数的，由pool+go 并发调用

type info struct {
	Endpoint string
	Counter string
	Image []byte
	Name string
}

func main() {
	//var err error


	//var item [6]map[string]string
	var item []info
	// cpu初始化
	item =  append(item,info{Endpoint:"vlnx", Counter:"cpu.idle", Name:"CPU TOP 10"})
	// memory
	item = append(item,info{Endpoint:"vlnx",Counter:"mem.memfree", Name:"MEMORY TOP 10"})
	// os ext4
	item = append(item,info{Endpoint:"vlnx",Counter:"df.bytes.free/fstype=ext4,mount=/",Name: "DISK TOP 10（ext4）"})
	// os xfs
	item = append(item,info{Endpoint: "vlnx",Counter: "df.bytes.free/fstype=xfs,mount=/",Name:"CentOS 7 xfs文件系统"})
	// mysql ext4 data
	item = append(item,info{Endpoint: "vlnx022",Counter:"df.bytes.free.percent/fstype=xfs,mount=/opt/mysql",Name:"MySQL数据分区"})
	// mongo xfs data
	item = append(item,info{Endpoint: "vlnx",Counter:"df.bytes.free.percent/fstype=xfs,mount=/opt/mongo", Name:"MongoDB分区"})
	wg := sync.WaitGroup{}
	for n,k := range item{
		wg.Add(1)
		go func(n1 int,k1 info) {
			defer wg.Done()
			ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
			err := chromedp.Run(ctx,sendkeys(k1.Endpoint,k1.Counter,&item[n1].Image))
			if err != nil {
				log.Fatal(err)
			}
			cancel()
		}(n,k)
	}
	wg.Wait()
	NewDocment(item)

	time.Sleep(10 * time.Second)
}

func NewDocment(infos []info)  {
	var err error
	doc := document.New()

	para := doc.AddParagraph()
	run := para.AddRun()
	para.SetStyle("Heading1")
	run.AddText("一、线上linux server OS整体运行状况：")

	run = createParaRunText(doc, "无异常")
	run.Properties().SetBold(true)

	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	run = para.AddRun()
	run.AddText("二、整体系统状态 TOP 10")

	for _,v := range infos {
		run = createParaRunText(doc, v.Name)
		run.Properties().SetBold(true)
		createParaRunDrawing(doc,v.Image)
	}


	err = doc.SaveToFile(time.Now().Format("20060102")+"_"+"Linux os日报.docx")
	if err != nil {
		log.Fatal(err)
	}
}
func createParaRunText(doc *document.Document, s string) document.Run {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(s)
	return run
}
func createParaRunDrawing(doc *document.Document, imgByte []byte)  {
	img, err := common.ImageFromBytes(imgByte)
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}
	imgref, err := doc.AddImage(img)
	if err != nil {
		log.Fatalf("unable to add image to document: %s", err)
	}
	para := doc.AddParagraph()
	inl, err := para.AddRun().AddDrawingInline(imgref)
	if err != nil {
		log.Fatal(err)
	}
	inl.SetSize(7*measurement.Inch, 5*measurement.Inch)
	return
}
