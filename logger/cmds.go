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

func Debug(args ...interface{}) {
	logx.zapSugarLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logx.zapSugarLogger.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	logx.zapSugarLogger.Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	logx.zapSugarLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logx.zapSugarLogger.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	logx.zapSugarLogger.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	logx.zapSugarLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logx.zapSugarLogger.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	logx.zapSugarLogger.Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	logx.zapSugarLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logx.zapSugarLogger.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	logx.zapSugarLogger.Errorw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	logx.zapSugarLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logx.zapSugarLogger.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	logx.zapSugarLogger.Fatalw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	logx.zapSugarLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logx.zapSugarLogger.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	logx.zapSugarLogger.Panicw(msg, keysAndValues...)
}

func Flush() {

	if logx.zapSugarLogger != nil {
		logx.zapSugarLogger.Sync()
	}
}
