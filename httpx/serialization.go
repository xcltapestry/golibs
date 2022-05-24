package httpx

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
	"encoding/json"
	"fmt"
	"github.com/xcltapestry/golibs/utils/jsonx"
	"github.com/xcltapestry/golibs/utils/strutil"
	"io"
	"io/ioutil"
	"net/http"
)

//FromJSON 将body 进行 json解码
func FromJSON(req interface{}, r *http.Request) (body []byte, err error) {
	if r == nil {
		err = fmt.Errorf("[FromJSON] http.Request is null. URL:%s", r.RequestURI)
		return
	}

	if r.Body == nil {
		err = fmt.Errorf("[FromJSON] r.Body is null. URL:%s", r.RequestURI)
		return
	}

	postBody, er := ioutil.ReadAll(r.Body)
	if er != nil {
		err = fmt.Errorf("[FromJSON] ReadAll Failure. URL:%s err:%s", r.RequestURI, er)
		return
	}
	defer r.Body.Close()

	if len(postBody) == 0 {
		err = fmt.Errorf("[FromJSON] Body is null.  URL:%s", r.RequestURI)
		return
	}

	if er = jsonx.Unmarshal(postBody, req); er != nil {
		err = fmt.Errorf("[FromJSON] json.Unmarshal Failure. err:%s, body:%s", er,
			strutil.UnsafeString(postBody))
	}

	body = postBody
	return
}

// example:
//  buf := &bytes.Buffer{}
//  ToJSON(x, buf)
//  buf.Bytes()
// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)
	//处理 & < > = 特殊符号
	// "\\u0026" = &  "\\u003c" = "<" 之类
	// 不然对于C++之类序列化的json处理可能不同
	enc.SetEscapeHTML(false)
	return enc.Encode(i)
}
