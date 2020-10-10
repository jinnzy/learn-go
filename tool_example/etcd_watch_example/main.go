package main


import (
"go.etcd.io/etcd/clientv3"
"time"
"fmt"
"context"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.56.122:2399"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect failed ,err ", err)
		return
	}

	defer client.Close()

	background := context.Background()
	client.Put(background, "/logagent/conf/", "123456")
	if err != nil {
		fmt.Println("err :", err)
		return
	}

	fmt.Println("connec success !!")
	for {
		watch := client.Watch(context.Background(), "/logagent/conf/")
		for wresp := range watch {
			for _, v := range wresp.Events {
				fmt.Println(v)
				fmt.Printf("%s %q : %q \n", v.Type, v.Kv.Key, v.Kv.Value)
			}
		}
	}

}
