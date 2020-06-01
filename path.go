package utils

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// PathItem struct
type PathItem struct {
	name     string
	fullpath string
	modtime  time.Time
}

// PathItemList slice
type PathItemList []*PathItem

// for PathItemList sort
func (list PathItemList) Len() int {
	return len(list)
}
func (list PathItemList) Less(i, j int) bool {
	if list[i].modtime.UnixNano() < list[j].modtime.UnixNano() {
		return true
	}
	return false
}
func (list PathItemList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// PathExists func
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CheckAndCreatePath func
func CheckAndCreatePath(path string) error {
	if rc, _ := PathExists(path); rc {
		return nil
	}
	err := os.MkdirAll(path, os.ModePerm)
	return err
}

// GetFiles func
func GetFiles(path string, incDir bool) (list PathItemList) {
	filepath.Walk(path, func(p string, f os.FileInfo, err error) error {
		if f == nil {
			return nil
		}
		//fmt.Println(p)
		if incDir == false && f.IsDir() {
			return nil
		}
		item := &PathItem{}
		item.fullpath = p
		item.name = f.Name()
		item.modtime = f.ModTime()
		list = append(list, item)
		return nil
	})
	return list
}

// GetFilesOrderByTime func
func GetFilesOrderByTime(path string, incDir bool, r bool) (list PathItemList) {
	list = GetFiles(path, incDir)
	sort.Sort(list)
	if r {
		reverse(list)
	}
	return list
}

func reverse(s PathItemList) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// GetCurrentDirectory func
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
