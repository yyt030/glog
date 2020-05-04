package main

import (
	"github.com/yyt030/glog"
	"time"
)

func main() {
	glog.SetOpts(
		//glog.LogFile("logs/aaa"),
		glog.AlsoLogToStderr(true),
		glog.Verbosity(1),
	)
	defer glog.Flush()

	go func() {
		glog.Info("this is info message")
	}()

	go func() {
		glog.Warn("this is warning message")
	}()

	go func() {
		glog.Error("this is error message")
	}()
	glog.V(1).Info("this is v1 message")
	glog.V(2).Info("this is v1 message")
	glog.V(3).Info("this is v1 message")
	time.Sleep(time.Second)
}
