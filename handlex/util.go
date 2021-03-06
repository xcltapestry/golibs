package handlex

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
	"fmt"

	"reflect"
	"runtime"
	"strings"
	"time"
)

// GetHandlerFuncName 得到Handler处理函数名
func GetHandlerFuncName(f interface{}) (handlerFuncName string) {
	handlerFuncName = GetFuncName(f)
	if slash := strings.LastIndex(handlerFuncName, "/"); slash >= 0 {
		handlerFuncName = handlerFuncName[slash+1:]
	}
	return
}

// GetFuncName get function name
// funcPc2, _, _, _ := runtime.Caller(0)
// fmt.Println(" FuncForPC:", runtime.FuncForPC(funcPc2).Name())
func GetFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
