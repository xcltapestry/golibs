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
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/xcltapestry/golibs/global"
	"github.com/xcltapestry/golibs/security"
)

type Client struct {
	client *http.Client
	cookie []*http.Cookie
}

func NewClient() *Client {
	c := &Client{
		client: http.DefaultClient,
	}
	c.client.Timeout = time.Duration(time.Minute * 1)
	c.client.Transport = &http.Transport{
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	return c
}

func (c *Client) WithClient(cl *http.Client) {
	c.client = cl
}

func (c *Client) GetClient() *http.Client {
	return c.client
}

type HeaderOpion func(*http.Request)

// method 参数: http.MethodGet  http.MethodPost 等， 如为非标准的访问请求会被报400
func (c *Client) Request(ctx context.Context, method, url string, body io.Reader,
	opts ...HeaderOpion) (*http.Response, error) {
	start := time.Now()

	// 检查url的合法性,有通过构建特殊url,攻击本地文件系统的案例，特别是调外部url时
	if safe, err := security.IsURLSafe(url); !safe {
		return nil, fmt.Errorf("[Request] url存在安全风险！ %s: %s err:%v", method, url, err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("[Request] NewRequest %s: %s err:%v", method, url, err)
	}

	// 先执行默认的http头设置
	for k, v := range global.Httpx_DefaultHeader {
		req.Header.Set(k, v)
	}
	// 再执行定制化http头
	if len(opts) > 0 {
		for _, f := range opts {
			if f != nil {
				f(req)
			}
		}
	}

	// 发起请求
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[Request] failed to do %s: %s err:%v ,elapsed:%d",
			method, url, err, time.Since(start))
	}
	if resp == nil {
		return nil, fmt.Errorf("[Request] Do()返回的resp为空 %s: %s err:%v ,elapsed:%d",
			method, url, err, time.Since(start))
	}

	return resp, nil
}

//Response HTTP请求返回的结构体
type Response struct {
	Body          []byte
	StatusCode    int
	RedirectedURL string
	Header        []string
}

//GetByte 相当于http.get()
func (c *Client) GetByte(ctx context.Context, method, url string, body io.Reader,
	opts ...HeaderOpion) (*Response, error) {
	//start := time.Now()
	resp, err := c.Request(ctx, method, url, body, opts...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	var respData []byte
	respData = make([]byte, 0)

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
		respData, err = ioutil.ReadAll(reader)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("[GetByte] gzip ReadAll err:%s", err)
		}
		defer reader.Close()
	case "deflate":
		reader = flate.NewReader(resp.Body)
		respData, err = ioutil.ReadAll(reader)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("[GetByte] deflate ReadAll err:%s", err)
		}
		defer reader.Close()
	case "br":
		respData, err = ioutil.ReadAll(brotli.NewReader(resp.Body))
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("[GetByte] br ReadAll err:%s", err)
		}
	default:
		reader = resp.Body
		respData, err = ioutil.ReadAll(reader)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("[GetByte] ReadAll err:%s", err)
		}
		defer reader.Close()
	}

	res := &Response{}
	res.StatusCode = resp.StatusCode
	res.Body = respData
	res.Header = HeaderToArray(resp.Header)
	res.RedirectedURL = resp.Request.URL.String()
	//logger.Info("[GetByte] 执行完毕! ", method, " : ", url, " elapsed:", time.Since(start))
	return res, nil
}

func HeaderToArray(header http.Header) (res []string) {
	for name, values := range header {
		for _, value := range values {
			res = append(res, fmt.Sprintf("%s: %s", name, value))
		}
	}
	return
}
