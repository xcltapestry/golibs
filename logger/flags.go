package logger

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

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

const (
	_default_logfile_maxAge     = 28  //28 days
	_default_logfile_maxSize    = 500 //500 megabytes
	_default_logfile_maxBackups = 3
	_default_logfile_name       = "app.log"
)

var ( //flags
	_logOutput     string
	_logOutputFile string
	_logLevel      string

	_logMaxAge   int
	_logMaxSize  int
	_logCompress bool
)

func parseFlags() {

	flag.StringVar(&_logOutput, "logger.output", "stdout", "仅支持下列选项: file/alsologtostderr/alsologtostdout/stderr/stdout/ ")
	flag.StringVar(&_logOutputFile, "logger.outputfile", "app.log", "输出到指定日志文件")
	flag.StringVar(&_logLevel, "logger.Level", "debug", "日志级别: debug/info/warn/error/fatal")

	flag.IntVar(&_logMaxAge, "logger.maxage", 28, "日志文件保存多长时间")
	flag.IntVar(&_logMaxSize, "logger.maxsize", 500, "单个日志文件大小")
	flag.BoolVar(&_logCompress, "logger.compress", true, "日志文件是否打开压缩")

	// flag.Parse()  //testing时，使用

}

//parseLevelFlag 得到显示日志级别，默认为Debug
func (l *Logx) parseLevelFlag() zapcore.Level {
	var lvl zapcore.Level
	switch strings.ToLower(strings.TrimSpace(_logLevel)) { //_logLevel
	case "debug":
		lvl = zap.DebugLevel
	case "info":
		lvl = zap.InfoLevel
	case "warn":
		lvl = zap.WarnLevel
	case "error":
		lvl = zap.ErrorLevel
	case "fatal":
		lvl = zap.FatalLevel
	default:
		lvl = zap.DebugLevel
	}
	l.lvl = lvl // zapcore.Enabled(level)
	return logx.lvl
}

func (l *Logx) initLogOutFile() *lumberjack.Logger {

	var logFile string

	if _logOutputFile != "" { // _logOutputFile
		logFile = _logOutputFile //_logOutputFile
	} else {
		logDir, _ := os.Getwd()
		logFile = filepath.Join(logDir, _default_logfile_name)
	}

	var maxAge int
	if _logMaxAge > 0 {
		maxAge = _logMaxAge
	} else {
		maxAge = _default_logfile_maxAge
	}

	var maxSize int
	if _logMaxSize > 0 {
		maxSize = _logMaxSize
	} else {
		maxSize = _default_logfile_maxSize
	}

	writeFile := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    maxSize, // megabytes
		MaxBackups: _default_logfile_maxBackups,
		MaxAge:     maxAge, // days
		Compress:   _logCompress,
		LocalTime:  true,
	}
	return writeFile
}
