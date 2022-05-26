package sysutil

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
	"log"
	"runtime/debug"
	"strings"
)

//CatchPanic 专用于各种panic下的recover
func CatchPanic(msgTitle ...string) {

	var title string
	title = "panic"
	if len(msgTitle) > 0 {
		for _, v := range msgTitle {
			title += "-" + v
		}
	}

	// 例外处理
	if r := recover(); r != nil {
		stack := strings.Join(strings.Split(string(debug.Stack()), "\n")[2:], "\n")
		log.Printf("[%s] recover: %+v stack: %s", title, r, stack)
	}
}
