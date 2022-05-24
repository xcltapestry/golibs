package jsonx

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
 ps:
	json的解析在高并发压力下，对资源和速度的影响比较大，
	所以有必要进行替换
*/

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

var JSON = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) ([]byte, error) {
	return JSON.Marshal(&v)
}

func Unmarshal(data []byte, v interface{}) error {
	return JSON.Unmarshal(data, &v)
}

//MarshalJSONWithWriter go默认为转义转=、<、>、&，和其他语言（如C++开发的某支付后端）交互时，会出异常
func MarshalJSONWithWriter(v interface{}, w io.Writer) (err error) {
	enc := JSON.NewEncoder(w)
	//处理 & < > = 特殊符号
	enc.SetEscapeHTML(false)
	err = enc.Encode(v)
	return
}
