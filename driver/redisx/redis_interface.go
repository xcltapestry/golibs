package redisx

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
备注:
	redis连接库之前使用redigo，这次选择了go-redis
	分布式锁推荐 redislock
	通过自定义Interface限制所能调用的redis方法，以防误用破坏性比较大的命令
*/

import (
	"context"
	"time"
)

type Redis interface {
	Close() error

	Ping() (int64, error)
	Exists(keys ...string) (int64, error)
	Append(key, value string) (int64, error)
	Set(key string, value interface{}, expiration time.Duration) (string, error)
	SetEX(key string, value interface{}, expiration time.Duration) (string, error)
	SetNX(key string, value interface{}, expiration time.Duration) (bool, error)

	Get(key string) (string, error)
	Del(key string) (int64, error)
	TTL(key string) (time.Duration, error)
	Expire(key string, expiration time.Duration) (bool, error)
	Incr(key string) (int64, error)
	Decr(key string) (int64, error)

	LRange(key string, start, stop int64) ([]string, error)
	Publish(channel string, message interface{}) (int64, error)

	NewLocker() *RedisLock //分布式锁
}

//RedisLock 分布式锁
type RedisLock interface {
	LinearBackoff(backoff time.Duration)
	LimitRetry(backoff time.Duration, max int)
	ExponentialBackoff(min, max time.Duration)

	Key() string
	Token() string
	Metadata() string
	TTL(ctx context.Context) (time.Duration, error)
	Refresh(ctx context.Context, ttl time.Duration)
	Release(ctx context.Context) error
}
