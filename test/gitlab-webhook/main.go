package main

import (
	"net/http"
	"github.com/learn-go/test/gitlab-webhook/webhook"
			)

const (
	path = "/webhooks"
)



func main() {

	http.HandleFunc(path,webhook.CreateJob)
	http.ListenAndServe(":9000", nil)
}