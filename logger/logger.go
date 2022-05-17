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
/*

日志级别:  debug < info < warn < error < fatal < panic

  Debug: 一般程序中的调试信息
  Info: 关键，核心流程的日志
  Warn: 警告信息
  Error: 错误日志
  Fatal: 致命错误，可能会造成程序无法运转
  Panic: 记录日志，然后panic

示例:

 go run main.go  -logger.outputfile="x.log" -logger.output=alsologtostdout

 go run main.go  -logger.output=stdout

 go run main.go  -logger.output=alsologtostdout  -logger.Level=error -logger.outputfile=err.log


调用示例:
	logger.NewLogger(logger.Text) //logger.JSON)
	flag.Parse()
	logger.ParseFlag()

*/

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logx *Logx

func init() {
	logx = &Logx{}
}

type Logx struct {
	flagsOnce      sync.Once    // flags
	formatter      LogFormatter // 日志展示格式
	lvl            zapcore.Level
	zapSugarLogger *zap.SugaredLogger
}

func NewLogger(formatter LogFormatter) {
	logx.flagsOnce.Do(parseFlags)
	logx.setFormatter(formatter)
	// logx.initZap()
}

func NewDefaultLogger() {
	logx.flagsOnce.Do(parseFlags)
	logx.setFormatter(JSON)
	// logx.initZap()
}

func ParseFlag() {
	logx.initZap()
}
