package tool

import (
	"io"
	"os"

	"github.com/beego/beego/v2/server/web"
)

var rootPath string

func SetRootPath(path string) {
	rootPath = path
}

func GetRootPath() string {
	if rootPath == "" {
		rootPath = web.AppPath
	}
	return rootPath
}

func ReadFile(path string) ([]byte, error) {
	var ret []byte
	var err error
	file, err := os.Open(path)
	if err == nil {
		content, err := io.ReadAll(file)
		if err == nil {
			ret = content
		}
		defer file.Close()
	}
	return ret, err
}

func ReadFileLast(path string, last int64) ([]byte, error) {
	var ret []byte
	var err error
	if fi, errs := os.Stat(path); errs != nil {
		err = errs
	} else {
		if fi.Size() > int64(last) {
			if file, erro := os.OpenFile(path, os.O_RDONLY, os.ModePerm); erro != nil {
				err = erro
			} else {
				if _, errs := file.Seek(fi.Size()-last, io.SeekStart); errs != nil {
					err = errs
				} else {
					if content, errr := io.ReadAll(file); errr != nil {
						err = errr
					} else {
						ret = content
					}
				}
				defer file.Close()
			}
		} else {
			ret, err = ReadFile(path)
		}
	}
	return ret, err
}

func WriteFile(path string, data string) error {
	return os.WriteFile(path, []byte(data), 0666)
}
