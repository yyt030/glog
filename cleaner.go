package glog

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	cleanReserve  = time.Hour * 24 * 30 // 30 days
	cleanInterval = time.Hour           // every hour
)

func findAndRemove() {
	fileInfos, err := ioutil.ReadDir(logDir)
	if err != nil {
		Errorf("read dir:%s, failed:%v", logDir, err)
		return
	}

	excludeFiles := make(map[string]struct{}, 0)

	for _, f := range fileInfos {
		if f.IsDir() {
			continue
		}

		if !strings.HasPrefix(f.Name(), fileName) {
			continue
		}

		if f.Mode()&os.ModeSymlink != 0 {
			excludeFiles[f.Name()] = struct{}{}
			readlink, err := os.Readlink(filepath.Join(logDir, f.Name()))
			if err != nil {
				continue
			}
			excludeFiles[readlink] = struct{}{}
		}
	}

	for _, f := range fileInfos {
		// Skip dir
		if f.IsDir() {
			continue
		}

		// Skip not given prefix filename
		if !strings.HasPrefix(f.Name(), fileName) {
			continue
		}

		// Skip symlink file and real file
		if _, ok := excludeFiles[f.Name()]; ok {
			continue
		}

		// Skip not special filename
		sn := strings.Split(f.Name(), ".")
		if len(sn) < 4 || !isNum(sn[len(sn)-1]) {
			continue
		}

		// Drop old files
		if time.Since(f.ModTime()) > cleanReserve {
			if err := os.Remove(filepath.Join(logDir, f.Name())); err != nil {
				Warnf("log cleaner remove:%s/%s, failed:%v", logDir, f.Name(), err)
			} else {
				Infof(">>> drop old file:%s", f.Name())
			}
		}
	}
}

func isNum(s string) bool {
	if _, err := strconv.ParseUint(s, 10, 64); err != nil {
		return false
	}

	return true
}

func RunCleaner() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				Errorf("glog cleaner run failed:%v", err)
			}
		}()
		Infof("logDir:%s, fileName:%s, cleanInterval:%v, cleanReserve:%v", logDir, fileName, cleanInterval, cleanReserve)
		for {
			findAndRemove()
			time.Sleep(cleanInterval)
		}
	}()

}
