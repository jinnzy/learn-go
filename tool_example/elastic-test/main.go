package main

import (
	"github.com/olivere/elastic"
	"fmt"
)

func main()  {
	client, err := elastic.NewClient(elastic.SetURL("http://172.23.4.154:30192"))
	if err != nil {
		// Handle error
	}
	fmt.Println(client.Get())
}



