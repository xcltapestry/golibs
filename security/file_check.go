package security

/**
 * Copyright 2022 golibs Author. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @Project golibs
 * @Description
 * @author XiongChuanLiang<br/>(xcl_168@aliyun.com)
 * @license http://www.apache.org/licenses/  Apache v2 License
 * @version 1.0
 */

//import (
//	"fmt"
//	"github.com/xcltapestry/golibs/logger"
//	"os"
//	"path/filepath"
//	"strings"
//)
//
//var fileSeparator = string(filepath.Separator)
//
//var sysDirs = map[string]bool{
//	"/var/":   true,
//	"/data/":  true,
//	"/bin/":   true,
//	"/dev/":   true,
//	"/home/":  true,
//	"/lib/":   true,
//	"/lib64/": true,
//	"/mnt/":   true,
//	"/proc/":  true,
//	"/opt/":   true,
//	"/root/":  true,
//	"/sys/":   true,
//	"/run/":   true,
//	"/sbin/":  true,
//	"/media/": true,
//	"/etc/":   true,
//	"/boot/":  true,
//}
//
//type FileCheck struct {
//	fullName string
//	filePath string
//	fileName string
//	ext      string
//	isDir    bool
//}
//
//func NewFileCheck(fullName string) *FileCheck {
//	if fullName == "" {
//		return &FileCheck{}
//	}
//
//	dir, file := filepath.Split(fullName)
//	ext := filepath.Ext(fullName)
//	var isDir bool
//	isDir = false
//	fileinfo, err := os.Stat(fullName)
//	if err == nil {
//		if fileinfo.IsDir() {
//			isDir = true
//		}
//	}
//	return &FileCheck{
//		fullName: fullName,
//		filePath: dir,
//		fileName: file,
//		ext:      ext,
//		isDir:    isDir,
//	}
//}
//
//func (self *FileCheck) Check() (safe bool) {
//
//	realName, err := filepath.Abs(self.fullName)
//	if err != nil {
//		logger.Error(fmt.Sprintf("[Check] faile to get abs, fullName:%s err:%v",
//			self.fullName, err))
//		return
//	}
//
//	// example : https://xxxx.com?dir=/mydir/../../../../root/.bash_history
//	if !strings.HasSuffix(realName, self.fullName) {
//		logger.Errorf("[Check] security: 文件名包含../ , 全名:%s",
//			self.fullName)
//		return false
//	}
//	var baseDir string
//	if idx := strings.Index(realName[1:], "/"); idx != -1 {
//		if idx = idx + 2; idx < len(realName) {
//			baseDir = realName[:idx]
//		}
//	}
//	return !sysDirs[baseDir]
//}
//
//func (self *FileCheck) Clean(filename string) (cleaned string) {
//	_, cleaned = filepath.Split(filename)
//	// 清理 文件分隔符/ shell特殊字符` $ ? ; | && < >  截断符%00
//	// example: cat /???/passwd
//	cleaned = strings.NewReplacer(
//		fileSeparator, "",
//		"`", "",
//		"$", "",
//		"?", "",
//		";", "",
//		"|", "",
//		"<", "",
//		">", "",
//		"&&", "",
//		"%00", "").Replace(cleaned)
//	cleaned = filepath.Clean(cleaned)
//	if filename != cleaned {
//		logger.Warnf("CleanFilename %s=>%s", filename, cleaned)
//	}
//	return
//}
