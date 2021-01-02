package util

import (
	"log"
	"os"
	"path/filepath"
)

var cacheFolder string

func CacheFolder() string {
	if cacheFolder == "" {
		dir, err := os.UserCacheDir()
		if err != nil {
			log.Fatal("Unable to determine user's cache directory. error [", err, "]")
		}
		cacheFolder = filepath.Join(dir, "go-duka")
		if err = os.MkdirAll(cacheFolder, 0755); err != nil {
			log.Fatal("Unable to create go-soft4fx cache directory [", cacheFolder, "] error: [", err, "]")
		}
	}

	return cacheFolder
}
