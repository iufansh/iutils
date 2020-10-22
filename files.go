package iutils

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 获取文件大小，单位：B
func CalcFileSize(filePath string) (int64, error) {
	fi, err:=os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

// 判断文件或文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, err
}

// 确保文件夹存在，不存在则创建
func EnsurePath(dstPath string) error {
	if flag, _ := PathExists(dstPath); !flag {
		if err := os.MkdirAll(dstPath, 0644); err != nil {
			return err
		}
	}
	return nil
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()
	var dstPath string
	if strings.Contains(dst, "/") {
		dstPath = SubString(dst, 0, strings.LastIndex(dst, "/"))
	} else if strings.Contains(dst, "\\") {
		dstPath = SubString(dst, 0, strings.LastIndex(dst, "\\"))
	}
	if dstPath != "" {
		if flag, _ := PathExists(dstPath); !flag {
			if err := os.MkdirAll(dstPath, 0644); err != nil {
				return 0, err
			}
		}
	}
	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
