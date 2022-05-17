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
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (l *Logx) initZap() {

	var ws zapcore.WriteSyncer
	var zapCore zapcore.Core

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line", //caller
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 	EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	var encoder zapcore.Encoder
	if logx.formatter == JSON {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	atomLevel := zap.NewAtomicLevelAt(l.parseLevelFlag())

	switch strings.ToLower(strings.TrimSpace(_logOutput)) { //_logOutput
	case "stderr":
		zapCore = zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stderr),
			atomLevel,
		)
	case "stdout":
		zapCore = zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			atomLevel,
		)
	case "file":
		writeFile := l.initLogOutFile()
		zapCore = zapcore.NewCore(
			encoder,
			zapcore.AddSync(writeFile),
			atomLevel,
		)
	case "alsologtostderr":
		writeFile := l.initLogOutFile()
		ws = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stderr),
			zapcore.AddSync(writeFile),
		)
		zapCore = zapcore.NewCore(
			encoder,
			ws,
			atomLevel,
		)
	case "alsologtostdout":
		writeFile := l.initLogOutFile()
		ws = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(writeFile),
		)
		zapCore = zapcore.NewCore(
			encoder,
			ws,
			atomLevel,
		)
	default: //stdout
		zapCore = zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			atomLevel,
		)
	}
	Logger := zap.New(zapCore)
	Logger = Logger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
	l.zapSugarLogger = Logger.Sugar()

}
