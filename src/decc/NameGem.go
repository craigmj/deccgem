package decc

import (	
	"fmt"
	"io"
	"path/filepath"
	"os"
	"regexp"
	"strconv"
	"time"
)

// NameGem returns the name to use for the latest gem build.
// This ensures that we ALWAYS have a different name for each gem.
func NameGem(d string, out io.Writer) (string,error) {
	reg := regexp.MustCompile(`^SACC\.(\d+)\.`)

	spreads, err := os.Open(filepath.Join(d, "spreadsheet"))
	if nil!=err {
		return "", err
	}
	defer spreads.Close()

	files, err := spreads.Readdir(-1)
	if nil!=err {
		return "", err
	}
	latest := ""
	for _, file := range files {
		m := reg.FindStringSubmatch(file.Name())
		if nil==m {
			continue
		}
		if ""==latest {
			latest = m[1]
		} else {
			if m[1]>latest {
				latest = m[1]
			}
		}
	}
	l, err := strconv.Atoi(latest)
	if nil!=err {
		return "", err
	}
	// l -= 141218

	latestTime, err := time.Parse("060102", latest)
	if nil!=err {
		return "", err
	}
	sub := time.Now().Sub(latestTime)
	return fmt.Sprintf("4.%d.%d", l, int(sub.Minutes())), nil
}