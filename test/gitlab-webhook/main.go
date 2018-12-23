package main

import (
		"net/http"
		"fmt"
	"io/ioutil"
	"bytes"
)

const (
	path = "/webhooks"
)

func main() {
	//hook, _ := gitlab.New(gitlab.Options.Secret("123"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

		result, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("123")
			fmt.Println(bytes.NewBuffer(result).String())
			fmt.Println("123")
		}

		//switch payload.(type) {
		//
		//case gitlab.ReleasePayload:
		//	release := payload.(gitlab.ReleasePayload)
		//	// Do whatever you want from here...
		//	fmt.Printf("%+v", release)
		//
		//case gitlab.PullRequestPayload:
		//	pullRequest := payload.(github.PullRequestPayload)
		//	// Do whatever you want from here...
		//	fmt.Printf("%+v", pullRequest)
		//}
	})
	http.ListenAndServe(":9000", nil)
}