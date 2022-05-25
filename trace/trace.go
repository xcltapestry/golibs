package trace

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
	"sync"
	"time"
)

/**
* Trace.go
* 用于记录追踪一次完整的请求相关的信息
* 包体不记录，因为包体可能包含敏感信息，要不要打印出来，应取决于相关的研发自己判断,如需要，可在保留trace id的情况下，打印在日志中
*
* 要实现完整的调用链路追踪在实际工程中，其实是比较困难的。
* OpenTracing类项目，会有一个取样率问题。取样太高，trace日志接收方服务会面临很大的压力。
* 且以实际经验看：
*   1. 在层层调用中，如果涉及到不同团队的接口调用，因为不同团队的理解和对这块的态度，很难保证调用链路能一直连续。
*   2. 在实际调用中，可能涉及到MQ,Redis之类，需要相关的协议处理，而在不同开发语言和框架下，要统一也需要做比较多的工作
* 当然，有这个后，收获也会是巨大的
**/

type SQL struct {
	Timestamp   string  `json:"timestamp"`     // 时间，格式：2006-01-02 15:04:05
	Stack       string  `json:"stack"`         // 文件地址和行号
	SQL         string  `json:"sql"`           // SQL 语句
	Rows        int64   `json:"rows_affected"` // 影响行数
	CostSeconds float64 `json:"cost_seconds"`  // 执行时长(单位秒)
}

type Redis struct {
	Timestamp   string  `json:"timestamp"`       // 时间，格式：2006-01-02 15:04:05
	Handle      string  `json:"handle"`          // 操作，SET/GET 等
	Key         string  `json:"key"`             // Key
	Value       string  `json:"value,omitempty"` // Value
	TTL         float64 `json:"ttl,omitempty"`   // 超时时长(单位分)
	CostSeconds float64 `json:"cost_seconds"`    // 执行时间(单位秒)
}

type ThirdPartyRequests struct { //请求其它的http接口
	mu          sync.Mutex
	RequestUrl  string   `json:"request_url"`
	Responses   []string `json:"responses"`    //可能会被重试多次
	Success     bool     `json:"success"`      // 是否成功，true 或 false
	CostSeconds float64  `json:"cost_seconds"` // 执行时长(单位秒)
}

func (a *ThirdPartyRequests) AppendResponse(response string) {
	if response == "" {
		return
	}
	a.mu.Lock()
	a.Responses = append(a.Responses, response)
	a.mu.Unlock()
}

type HttpInfo struct {
	Method     string `json:"method"`
	RequestURL string `json:"request_url"`
	Header     string `json:"header"`
}

type Request struct {
	HttpInfo
}

type Response struct {
	HttpInfo
	HttpCode    int     `json:"http_code"`
	HttpCodeMsg string  `json:"http_code_msg"`
	CostSeconds float64 `json:"cost_seconds"`
}

type Trace struct {
	mu          sync.Mutex
	TraceID     string                `json:"trace_id"`
	Success     bool                  `json:"success"`
	CostSeconds float64               `json:"cost_seconds"`
	SQLs        []*SQL                `json:"sqls"`
	Redis       []*Redis              `json:"redis"`
	Apis        []*ThirdPartyRequests `json:"third_party_requests"`
	Request     *Request              `json:"request"`
	Response    *Response             `json:"response"`
}

func NewTrace(id string) *Trace {
	if id == "" {
		id = fmt.Sprint(time.Now().UnixMilli())
	}
	return &Trace{TraceID: id}
}

func (t *Trace) AppendSQL(sql *SQL) *Trace {
	if sql == nil {
		return t
	}
	t.mu.Lock()
	defer t.mu.Unlock()

	t.SQLs = append(t.SQLs, sql)
	return t
}

func (t *Trace) AppendRedis(redis *Redis) *Trace {
	if redis == nil {
		return t
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	t.Redis = append(t.Redis, redis)
	return t
}

func (t *Trace) AppendAPI(thirdPartyAPI *ThirdPartyRequests) *Trace {
	if thirdPartyAPI == nil {
		return t
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	t.Apis = append(t.Apis, thirdPartyAPI)
	return t
}
