package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// LogLevel  0-no log, 1-error, 2-info, 3-debug
var LogLevel = 3

// LogVerbose  console print
var LogVerbose = 1

// LogFileMaxSize define
var LogFileMaxSize int64 = 10000000

// LogFileLimits define
var LogFileLimits = 10

// Tomlog struct
type Tomlog struct {
	logpath    string
	f          *os.File
	w          *bufio.Writer
	totalBytes int64
}

var tomlog Tomlog
var loglocker sync.Mutex

func init() {
	openLogFile()

	//log.SetPrefix("【TOM】")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(tomlog)
}

// Open log file
func openLogFile() {
	loglocker.Lock()
	defer loglocker.Unlock()

	tomlog.totalBytes = 0
	tomlog.w = nil
	if tomlog.f != nil {
		tomlog.f.Close()
		tomlog.f = nil
	}
	tomlog.logpath = GetCurrentDirectory() + "/log"
	CleanLogPath(tomlog.logpath)

	t := time.Now()
	tstr := fmt.Sprintf("/%04d%02d%02d_%02d%02d%02d.log", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	tomlog.logpath += tstr
	//log.Println(ipmlog)

	tomlog.w = nil
	tomlog.f = nil
}

func checkCreateLogFile() {
	if tomlog.f != nil {
		return
	}
	path := GetCurrentDirectory() + "/log"
	CheckAndCreatePath(path)

	f, err := os.Create(tomlog.logpath)
	if err != nil {
		fmt.Printf("open log file: %s failed.\n", tomlog.logpath)
		return
	}
	//syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	tomlog.w = bufio.NewWriter(f)
	tomlog.f = f
}

// Write func
func (ipm Tomlog) Write(p []byte) (n int, err error) {
	if LogVerbose > 0 {
		fmt.Print(string(p))
	}

	loglocker.Lock()
	checkCreateLogFile()
	if tomlog.w != nil {
		n, err = tomlog.w.Write(p)
		tomlog.w.Flush()
		tomlog.totalBytes += int64(n)
	} else {
		n = len(p)
		err = nil
	}
	loglocker.Unlock()

	if tomlog.totalBytes > LogFileMaxSize {
		openLogFile()
	}

	return n, err
}

// CleanLogPath func
func CleanLogPath(logpath string) {
	list := GetFilesOrderByTime(logpath, false, true)
	for i, item := range list {
		if i >= LogFileLimits-1 {
			os.Remove(item.fullpath)
		}
	}
}
