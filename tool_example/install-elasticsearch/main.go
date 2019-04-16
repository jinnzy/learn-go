package main

import (
	"fmt"
	"github.com/learn-go/tool_example/install-elasticsearch/pkg/logstreamer"
	"log"
	"os"
	"os/exec"
)

//func main()  {
//	var stdoutBuf bytes.Buffer
//
//	cmd := exec.Command("/bin/bash","-c","ls;sleep 2;ls -l")
//
//	cmd.Stdout = &stdoutBuf
//	err := cmd.Start()
//	if err != nil {
//		logging.Error("执行命令错误",zap.Error(err))
//	}
//	go func() {
//		for {
//			if cmd.Stdout != nil {
//				fmt.Println(cmd.Stdout)
//			}
//		}
//	}()
//	_ = cmd.Wait()
//}

func main() {
	// Create a logger (your app probably already has one)
	logger := log.New(os.Stdout, "--> ", log.Ldate|log.Ltime)

	//Setup a streamer that we'll pipe cmd.Stdout to
	logStreamerOut := logstreamer.NewLogstreamer(logger, "stdout", false)
	defer logStreamerOut.Close()
	// Setup a streamer that we'll pipe cmd.Stderr to.
	// We want to record/buffer anything that's written to this (3rd argument true)
	logStreamerErr := logstreamer.NewLogstreamer(logger, "stderr", true)
	defer logStreamerErr.Close()

	// Execute something that succeeds
	cmd := exec.Command("/bin/bash","-c","yum -y install vim;llll;asdasd")

	cmd.Stderr = logStreamerErr
	cmd.Stdout = logStreamerOut

	// Reset any error we recorded
	logStreamerErr.FlushRecord()

	// Execute command
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}