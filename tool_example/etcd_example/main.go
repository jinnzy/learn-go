package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main()  {
	// etcd 连接
	config := clientv3.Config{
		Endpoints:[]string{"192.168.56.122:2399"},
		DialTimeout:10*time.Second,
	}
	cli,err := clientv3.New(config)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect succ")

	defer cli.Close()
	// 存取值
	//kv := clientv3.NewKV(cli)
	ctx,cancel:= context.WithTimeout(context.Background(),5*time.Second)
	// 操作etcd
	_, err = cli.Put(ctx, "/logagent/conf/", "sample_value")
	// 操作完毕，取消etcd
	cancel()
	if err != nil {
		fmt.Println("put err",err)
		return
	}
	//取值，设置超时为1秒
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/conf/")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
