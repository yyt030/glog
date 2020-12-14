package glog

import (
	"strings"
	"testing"
	"time"
)

func TestRunCleaner(t *testing.T) {
	SetOpts(
		LogMaxSizeMB(10),
		//glog.LogFile("C:\\workspaces\\glog\\example\\log/test"),
		LogFile("C:\\workspaces\\glog\\example\\log\\abc123ok"),
		FlushIntervalSecond(1),
		//glog.LogToStderr(true),
		//glog.AlsoLogToStderr(true),
		CleanIntervalSecond(5),
		CleanReserveDay(1),
	)
	defer Flush()

	RunCleaner()

	for i := 0; i < 100; i++ {
		Infof("%d, %s", i, strings.Repeat("this is test,", 3))
		Warnf("%d, %s", i, strings.Repeat("this is test,", 3))
		Errorf("%d, %s", i, strings.Repeat("this is test,", 3))
		//time.Sleep(time.Second)
	}

	time.Sleep(time.Second * 1)
}
