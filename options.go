package glog

import (
	"fmt"
	"path/filepath"
	"time"
)

type Opt func(t *loggingT)

func LogFile(filename string) Opt {
	return func(t *loggingT) {
		fileName = filepath.Base(filename)
		logDir = filepath.Dir(filename)
	}
}

func LogToStderr(flag bool) Opt {
	return func(t *loggingT) {
		t.toStderr = flag
	}
}

func AlsoLogToStderr(flag bool) Opt {
	return func(t *loggingT) {
		t.alsoToStderr = flag
	}
}

func Verbosity(v int) Opt {
	return func(t *loggingT) {
		t.verbosity.Set(fmt.Sprintf("%d", v))
	}
}

func StderrThreshold(v string) Opt {
	return func(t *loggingT) {
		t.stderrThreshold.Set(v)
	}
}

func TraceLocation(v string) Opt {
	return func(t *loggingT) {
		t.traceLocation.Set(v)
	}
}

func FlushInterval(v int) Opt {
	return func(t *loggingT) {
		flushInterval = time.Duration(v) * time.Second
	}
}

func SetOpts(opts ...Opt) {
	for _, opt := range opts {
		opt(&logging)
	}
}
