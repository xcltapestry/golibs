package filesystem

/*
集成文件处理函数
有借鉴和引用 github.com/88250/gulu
@authr: Xiong Chuan Liang
@mail: xcl_168@aliyun.com
*/

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	//"github.com/xcltapestry/golibs/logger"
)

// GetFileSize get the length in bytes of file of the specified path.
func GetFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if nil != err {
		//logger.Error(err)

		return -1
	}

	return fi.Size()
}

// IsDir 检查路径是否为目录
func IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}

	if nil != err {
		//logger.Warnf("determines whether [%s] is a directory failed: [%v]", path, err)
		return false
	}
	return fio.IsDir()
}

// IsDirExist 检测目录是否存在, true: 为存在  false:为目录不存在
func IsDirExist(dir string) (exist bool) {
	ret, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return
	}

	if ret != nil && ret.IsDir() {
		exist = true
	}
	return
}

// IsExist determines whether the file spcified by the given path is exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsHidden checks whether the file specified by the given path is hidden.
func IsHidden(path string) bool {
	path = filepath.Base(path)
	if 1 > len(path) {
		return false
	}
	return "." == path[:1]
}

// FormatFileSize  字节的单位转换 保留两位小数
//
//	文件大小字节单位换算为 EB TB GB MB KB B
func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2f B", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2f KB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f MB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f GB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f TB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2f EB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

// FileLineCounter Counts line in file
func FileLineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

// ComputeMd5WithSize 计算文件MD5与大小
func ComputeMd5WithSize(filePath string) (string, int64, error) {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", 0, err
	}

	result = hash.Sum(result)
	var str = ""
	for _, b := range result {
		str += strconv.Itoa(int(b))
	}

	stat, err := file.Stat()
	if err != nil {
		return "", 0, err
	}

	return str, stat.Size(), nil
}

// 检测文件所在目录是否存在，如不存在，则创建它
func CheckAndMkFilePath(file string) error {
	fdir, _ := filepath.Split(file)
	if !IsExist(fdir) {
		err := os.MkdirAll(fdir, 0755)
		if err != nil {
			return fmt.Errorf("[CheckAndMkFilePath] 目录创建失败. err:%s file:%s", err, file)
		}
	}
	return nil
}

// CheckAndMkDir 检测目录是否存在，如不存在，则自动创建
func CheckAndMkDir(fdir string) error {
	if !IsExist(fdir) {
		err := os.MkdirAll(fdir, 0755)
		if err != nil {
			return fmt.Errorf("[CheckAndMkDir] 目录创建失败. err:%s dir:%s", err, fdir)
		}
	}
	return nil
}
