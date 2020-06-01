package utils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// ReadLineCallback func
type ReadLineCallback func(lineNum int64, line string)

// TextFileReadLine func
func TextFileReadLine(fn string, cb ReadLineCallback) (totalLines int64, err error) {
	if cb == nil {
		return 0, errors.New("callback ptr is nil")
	}

	f, err := os.Open("legacy.txt")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for totalLines = 0; ; totalLines++ {
		line, err := br.ReadBytes('\n')
		if err != nil && len(line) == 0 {
			break
		}
		strLine := string(line)
		strLine = strings.ReplaceAll(strLine, "\r", "")
		strLine = strings.ReplaceAll(strLine, "\n", "")
		if len(strLine) == 0 {
			continue
		}
		cb(totalLines, strLine)
	}
	return totalLines, nil
}
