package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)
const TIME_LAYOUT = "2006-01-02 15:04:05"

func main() {
	//GetAllFile("./config-passwd")
	//walkFunc := func(path string, info os.FileInfo, err error) error{
	//	fmt.Println(path)
	//	fmt.Println(info.Name())
	//	return nil
	//}
	//filepath.Walk("./config-passwd", walkFunc)
	//ret,err := listExpireFilesNew("./config-passwd", 30)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(ret)
	//is.New()
	//a := []int{1,2,3,4,5}
	//fmt.Println(a[3/2])
	//fmt.Println(3/2)
	//fmt.Println(1/2)
	//fmt.Println(1 << 29)
	//fmt.Println( 1 ^ 1)
	//fmt.Println(uint64(14695981039346656037) ^ uint64(116))
	//Sum64("testKey")
	//const (
	//	prime64 = 1099511628211
	//)
	//var (
	//	a uint64 = 14695981039346656081
	//)
	//
	//a = a * prime64
	//b := uint64(16158402040730025883278554301091)
	//fmt.Println(14695981039346656081 * prime64)
    //fmt.Println(a)
	x := 19
	fmt.Println(x & 1)
	fmt.Println(x%2 == (x&(2-1)))
}
const (
	offset64 = 14695981039346656037
)
//func Sum64(key string) uint64 {
//	var hash uint64 = offset64
//	for i := 0; i < len(key); i++ {
//		fmt.Println("uint64 key[i]:", uint64(key[i]))
//		hash ^= uint64(key[i])
//		fmt.Println("^:", hash)
//		hash = hash * prime64
//		fmt.Println("*:", hash)
//	}
//	return hash
//}
func parseWithLocation(name string, timeStr string) (time.Time, error) {
	locationName := name
	if l, err := time.LoadLocation(locationName); err != nil {
		println(err.Error())
		return time.Time{}, err
	} else {
		lt, _ := time.ParseInLocation(TIME_LAYOUT, timeStr, l)
		fmt.Println(locationName, lt)
		return lt, nil
	}
}

func GetAllFile(pathname string) (files []string, err error) {
	rd, err := ioutil.ReadDir(pathname)
	// 操作系统路径分隔符
	//pathSep := string(os.PathSeparator)

	for _, fi := range rd {
		if fi.IsDir() {

			 fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())

		GetAllFile(pathname + fi.Name() + "\\")

		} else {
			fmt.Println(pathname+fi.Name())

		files = append(files,pathname)
		 }

		 }

	 return

}
func listExpireFilesNew(dir string, expireSeconds int64) (ret []os.FileInfo, err error) {
	nowUnix := time.Now().Unix()

	walkFunc := func(path string, info os.FileInfo, err error) error{
		if nowUnix-info.ModTime().Unix() > expireSeconds {
			ret = append(ret, info)
		}
		return nil
	}

	err = filepath.Walk(dir, walkFunc)
	return
}
