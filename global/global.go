package global

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

////////////////////////////
//  httpx
////////////////////////////
const (
	_UA_macos = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36"
)

//DefaultHeader 默认必设的http头，可在应用初始化时配置
var Httpx.DefaultHeader = map[string]string{
	"User-Agent": _UA_macos,
}

////////////////////////////
//  errorx
////////////////////////////
//GeneralErrorMap  全局变量，在应用服务启动时，初始化将相关的错误信息(code与message)赋值进来。
// 一个业务系统建议只有一个统一的全局error map.以防系统复杂后，不能有效的归纳整理错误信息
var errorx.GeneralErrorMap map[int]string

