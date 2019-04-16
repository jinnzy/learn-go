package main

import (
	"github.com/learn-go/gitlab-webhook/pkg/logging"
	"net/http"
	"github.com/learn-go/gitlab-webhook/webhook"
				)

const (
	path = "/webhooks"
)

func main() {
	logging.Init()
	http.HandleFunc(path,webhook.CreateJob)
	http.ListenAndServe(":9000", nil)
}