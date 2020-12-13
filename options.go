package glog

import (
	"fmt"
	"path/filepath"
	"time"
)

type Opt func(t *loggingT)

func LogFile(filename string) Opt {
	return func(t *loggingT) {
		logDir = filepath.Dir(filename)
		fileName = filepath.Base(filename)
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

func LogMaxSize(v uint64) Opt {
	return func(t *loggingT) {
		maxSize = v
	}
}

func CleanReserve(v time.Duration) Opt {
	return func(t *loggingT) {
		cleanReserve = v
	}
}

func CleanInterval(v time.Duration) Opt {
	return func(t *loggingT) {
		cleanInterval = v
	}
}

func SetOpts(opts ...Opt) {
	for _, opt := range opts {
		opt(&logging)
	}
}
