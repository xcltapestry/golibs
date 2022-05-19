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
)

/*

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
