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
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const xForwardFor = "X-Forward-For"

func GetRemoteAddr(r *http.Request) string {
	v := r.Header.Get(xForwardFor)
	if len(v) > 0 {
		return v
	}
	return r.RemoteAddr
}

//GetCookieValue 获取指定名称的cookie
func GetCookieValue(cookies, name string) (ret string) {
	cs := strings.Split(cookies, ";")
	for _, v := range cs {
		tempStr := strings.Split(v, "=")
		if tempStr[0] == name || tempStr[0] == ("["+name) {
			return strings.Split(tempStr[1], "]")[0]
		}
	}
	return
}

// AddHeader 添加一个头
func AddHeader(h *http.Header, name, value string) error {
	if h == nil {
		return errors.New("[AddHeader] header is null. ")
	}
	if h.Get(name) == "" {
		h.Set(name, value)
		return nil
	}
	h.Add(name, value)
	return nil
}

// DelHeader  删除一个header 项
func DelHeader(h *http.Header, name string) error {
	if h == nil {
		return errors.New("[DelHeader] header is null. ")
	}
	if v := h.Get(name); v != "" {
		h.Del(name)
	}
	return nil
}

// CopyHTTPRequest 用于复制新的HTTP Request
//  CopyHTTPRequest(req, ioutil.NopCloser(&bytes.Buffer{}))
func CopyHTTPRequest(r *http.Request, body io.ReadCloser) *http.Request {
	req := new(http.Request)
	*req = *r
	req.URL = &url.URL{}
	*req.URL = *r.URL
	req.Body = body

	req.Header = http.Header{}
	for k, v := range r.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}

	return req
}

//DownloadFile 更多用于下载文件
// func DownloadFile(fileName, method, url string, body io.Reader, opts ...HeaderOpion) error {
// 	start := time.Now()
// 	c := NewClient()
// 	resp, err := c.Request(method, url, body, opts...)
// 	if err != nil {
// 		return fmt.Errorf("[DownloadFile] http.Request error: %s elapsed:%d", err, time.Since(start))
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("[DownloadFile] unexpected statusCode(%d) err:%v elapsed:%d",
// 			resp.StatusCode, err, time.Since(start))
// 	}

// 	var reader io.ReadCloser
// 	switch resp.Header.Get("Content-Encoding") {
// 	case "gzip":
// 		// logger.Debug("[Request] Content-Encoding gzip! url:", url)
// 		reader, _ = gzip.NewReader(resp.Body)
// 	case "deflate":
// 		// logger.Debug("[Request] Content-Encoding deflate! url:", url)
// 		reader = flate.NewReader(resp.Body)
// 	default:
// 		reader = resp.Body
// 	}
// 	defer reader.Close()

// 	//再检查新的路径的安全性
// 	//dir, file := filepath.Split(fileName)
// 	//err = security.CheckFileSecurity(dir, file)
// 	//if err != nil {
// 	//	return fmt.Errorf("[DownloadFile] CheckFileSecurity error: %v elapsed:%d", err, time.Since(start))
// 	//}

// 	// 对于 下载前，文件是否已存在，是覆盖还是其他处理，放在这个函数以外处理
// 	//下载文件
// 	f, err := os.Create(fileName)
// 	if err != nil {
// 		return fmt.Errorf("[DownloadFile] os.Create error: %v elapsed:%d", err, time.Since(start))
// 	}
// 	defer f.Close()

// 	//Write the bytes to the fiel
// 	_, err = io.Copy(f, reader)
// 	if err != nil && err != io.EOF {
// 		return fmt.Errorf("[DownloadFile] file copy error: %v elapsed:%d", err, time.Since(start))
// 	}

// 	return nil
// }
