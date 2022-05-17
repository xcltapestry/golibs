package logger

/**
 * Copyright 2021  gowk Author. All Rights Reserved.
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
 */

import "testing"

func TestText(t *testing.T) {

	NewLogger(Text)

	Debug("debug msg")
	Info("info msg")
	Warn("warn msg")
	Infow("this is a dummy log1111", "request_id", "I100")
	Infow("this is a dummy log222", "request_id", "I101")
	Errorw("this is a Errorw log", "request_id", "E102")

}

func TestJSON(t *testing.T) {
	NewLogger(JSON)
	Debug("debug msg")
	Info("info msg")
	Warn("warn msg")
	Infow("infow msg", "request_id", "i001")
	Infow("infow msg", "request_id", "i002")
	Errorw("error msg", "request_id", "e003")
	// Fatal("fatal msg")
	// Panic("panic msg")
}

func TestFatal(t *testing.T) {
	NewDefaultLogger()
	Debug("debug msg")
	Info("info msg")
	Warn("warn msg")
	Fatal("fatal msg")
	Infow("infow msg", "request_id", "i001")
	Infow("infow msg", "request_id", "i002")
	Errorw("error msg", "request_id", "e003")
	Panic("panic msg")
}

func TestPanic(t *testing.T) {
	NewDefaultLogger()
	Debug("debug msg")
	Infof("info msg err:%s", "notfound!")
	Warn("warn msg")
	Panic("panic msg") //会退出应用
	Infow("infow msg", "request_id", "i001")
	Infow("infow msg", "request_id", "i002")
	Errorw("error msg", "request_id", "e003")
}
