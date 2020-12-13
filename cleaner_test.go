package glog

import (
	"strings"
	"testing"
	"time"
)

func TestRunCleaner(t *testing.T) {
	SetOpts(
		LogMaxSize(1024*10),
		//glog.LogFile("C:\\workspaces\\glog\\example\\log/test"),
		LogFile("C:\\workspaces\\glog\\example\\log/abc123ok"),
		FlushInterval(1),
		//glog.LogToStderr(true),
		//glog.AlsoLogToStderr(true),
		CleanInterval(time.Second*5),
		CleanReserve(time.Second*10),
	)
	defer Flush()

	go func() {
		RunCleaner()
	}()

	for i := 0; i < 100; i++ {
		Infof("%d, %s", i, strings.Repeat("this is test,", 3))
		Warnf("%d, %s", i, strings.Repeat("this is test,", 3))
		Errorf("%d, %s", i, strings.Repeat("this is test,", 3))
		//time.Sleep(time.Second)
	}

}
