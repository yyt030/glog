# glog - the extended version for [glog](https://github.com/golang/glog)

### Overview

This library support for deleting old logs and go mod !

### Install
```go
go get  github.com/yyt030/glog
```

### Usage
Here is a example usage that will do check and clean the log files that creation time older than variables.

```go
package main

import (
	"flag"
	"strings"
	"time"

	"github.com/yyt030/glog"
)

var (
	logSize       = flag.Uint64("s", 1024*10, "log file size")
	logName       = flag.String("f", "abc", "log file name")
	cleanInterval = flag.Uint64("ci", 5, "second of clean interval")
	cleanReserve  = flag.Int64("cr", -1, "second of clean reserve")
)

func main() {
	flag.Parse()
	glog.SetOpts(
		glog.LogMaxSizeMB(*logSize),
		//glog.LogFile("C:\\workspaces\\glog\\example\\log/test"),
		glog.LogFile(*logName),
		//glog.FlushIntervalSecond(1),
		//glog.LogToStderr(true),
		//glog.AlsoLogToStderr(true),
		glog.CleanIntervalSecond(*cleanInterval),
		glog.CleanReserveDay(*cleanReserve),
	)
	defer glog.Flush()

	glog.Infof("arg:%d, flag:%d", flag.NArg(), flag.NFlag())

	glog.RunCleaner()

	for i := 0; i < 100; i++ {
		glog.Infof("%d, %s", i, strings.Repeat("this is test,", 9))
		//time.Sleep(time.Second)
	}

	time.Sleep(time.Second * 30)
}
```

