package ctxutil

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
	"errors"
	"fmt"
	"time"
)

/**

example:
	import "github.com/google/uuid"

	ctx := context.Background()
	key, val := "traceId", uuid.New().String()

	subCtx := AddCtxValue[string, string](ctx, key, val)
	traceId, _ := ReadCtxValue[string, string](subCtx, key)
	fmt.Println(" traceId = ", traceId)

*/
func AddCtxValue[T any, V any](ctx context.Context, key T, val V) context.Context {
	return context.WithValue(ctx, key, val)
}

func ReadCtxValue[T any, V any](ctx context.Context, key T) (V, error) {
	if ret, ok := ctx.Value(key).(V); ok {
		return ret, nil
	} else {
		return ret, errors.New(fmt.Sprint(" 在context中没有找到此Key: ", key))
	}
}

//ShrinkDeadline 得到剩余TTL
func ShrinkDeadline(ctx context.Context, timeout time.Duration) time.Time {
	var timeoutTime = time.Now().Add(timeout)
	if ctx == nil {
		return timeoutTime
	}
	if deadline, ok := ctx.Deadline(); ok && timeoutTime.After(deadline) {
		return deadline
	}
	return timeoutTime
}

// 赋值 context 对象
// 通过r.context 获取上下文，会在当前 goroutine 结束时调用 release
// 重新生成新的对象，脱离当前生命周期
func ContextDup(ctx context.Context, ctxParams map[int]string) context.Context {
	if ctx == nil {
		return nil
	}

	ctxDup := context.Background()
	for k, v := range ctxParams {
		ctxDup = AddCtxValue(ctxDup, k, v)
	}
	return ctxDup
}
