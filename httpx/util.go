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
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const xForwardFor = "X-Forward-For"

func GetRemoteAddr(r *http.Request) string {
	v := r.Header.Get(xForwardFor)
	if len(v) > 0 {
		return v
	}
	return r.RemoteAddr
}

//DownloadFile 更多用于下载文件
func DownloadFile(fileName, method, url string, body io.Reader, opts ...HeaderOpion) error {
	start := time.Now()
	c := NewClient()
	resp, err := c.Request(method, url, body, opts...)
	if err != nil {
		return fmt.Errorf("[DownloadFile] http.Request error: %s elapsed:%d", err, time.Since(start))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[DownloadFile] unexpected statusCode(%d) err:%v elapsed:%d",
			resp.StatusCode, err, time.Since(start))
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		// logger.Debug("[Request] Content-Encoding gzip! url:", url)
		reader, _ = gzip.NewReader(resp.Body)
	case "deflate":
		// logger.Debug("[Request] Content-Encoding deflate! url:", url)
		reader = flate.NewReader(resp.Body)
	default:
		reader = resp.Body
	}
	defer reader.Close()

	//再检查新的路径的安全性
	//dir, file := filepath.Split(fileName)
	//err = security.CheckFileSecurity(dir, file)
	//if err != nil {
	//	return fmt.Errorf("[DownloadFile] CheckFileSecurity error: %v elapsed:%d", err, time.Since(start))
	//}

	// 对于 下载前，文件是否已存在，是覆盖还是其他处理，放在这个函数以外处理
	//下载文件
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("[DownloadFile] os.Create error: %v elapsed:%d", err, time.Since(start))
	}
	defer f.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(f, reader)
	if err != nil && err != io.EOF {
		return fmt.Errorf("[DownloadFile] file copy error: %v elapsed:%d", err, time.Since(start))
	}

	return nil
}
