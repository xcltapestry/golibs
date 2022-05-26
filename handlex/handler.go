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
	"context"
	"github.com/xcltapestry/golibs/trace"
	"net/http"
)

/**
ps :
  1. 建议不要在context放太多key,维护很麻烦容易漏,且业务变化很快，用一个结构体收拢就好
  2. context下的key ，各种包用多了，是有可能重复的，所以建议改成 struct{} 方式

*/

//HandlerInfoContextKey 业务上的相关信息
type HandlerInfoContextKey struct{}

//TraceInfoContextKey 追踪的key
type TraceInfoContextKey struct{}

//TraceIDContextKey 追踪id
type TraceIDContextKey struct{}

// TraceIDFromContext 从ctx中得到消息的TraceID
func TraceIDFromContext(ctx context.Context) string {
	id, ok := ctx.Value(TraceIDContextKey{}).(string)
	if ok {
		return id
	}
	return ""
}

//WithTraceID 写入TraceID
func WithTraceID(ctx context.Context, key TraceInfoContextKey, val string) context.Context {
	return context.WithValue(ctx, TraceIDContextKey{}, val)
}

//WithTraceInfo  记录Trace信息
func WithTraceInfo(ctx context.Context, key TraceInfoContextKey, val *trace.Trace) context.Context {
	return context.WithValue(ctx, key, val)
}

//TraceInfoFromContext 从上下文得到Trace信息
func TraceInfoFromContext(r *http.Request) *trace.Trace {
	if t, ok := r.Context().Value(TraceInfoContextKey{}).(*trace.Trace); ok {
		return t
	}
	return nil
}

type HandlerInfo struct {
	Data interface{} // 依业务情况去自定义
}

//WithHandlerInfo 写入业务信息
func WithHandlerInfo(ctx context.Context, key HandlerInfoContextKey, val *HandlerInfo) context.Context {
	return context.WithValue(ctx, key, val)
}

//HandlerInfoFormContext 获得业务信息
func HandlerInfoFormContext(r *http.Request) *HandlerInfo {
	if h, ok := r.Context().Value(HandlerInfoContextKey{}).(*HandlerInfo); ok {
		return h
	}
	return nil
}
