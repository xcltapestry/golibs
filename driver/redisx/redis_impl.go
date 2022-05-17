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

import (
	"context"
	"time"

	"github.com/bsm/redislock"
)

type Client struct {
	rdb *redis.Client
}

func NewClient(opt *Options) *Client {
	c := &Client{}
	c.rdb = redis.NewClient(opt.setDriverOptions())
	return c
}

//NewLocker : 从redis连接创建一个分布式锁
func (rc *Client) NewLocker() *Lock {
	lk := &Lock{}
	lk.client = redislock.New(rc.rdb)
	lk.options = nil
	return lk
}

//Lock 分布式锁
type Lock struct {
	client  *redislock.Client
	locker  *redislock.Lock
	options *redislock.Options
}

func (lk *Lock) Obtain(ctx context.Context, key string, ttl time.Duration) error {
	var err error
	lk.locker, err = lk.client.Obtain(context.Background(), key, ttl, lk.options)
	if err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////
//
//
//
//  options 参数初始化
////////////////////////////////////////////
func (lk *Lock) initOptions() {
	if lk.options == nil {
		lk.options = &redislock.Options{}
	}
}

func (lk *Lock) LinearBackoff(backoff time.Duration) {
	lk.initOptions()
	lk.options.RetryStrategy = redislock.LinearBackoff(backoff)
}

func (lk *Lock) LimitRetry(backoff time.Duration, max int) {
	lk.initOptions()
	retryStrategy := redislock.LinearBackoff(backoff)
	lk.options.RetryStrategy = redislock.LimitRetry(retryStrategy, max)
}

func (lk *Lock) ExponentialBackoff(min, max time.Duration) {
	lk.initOptions()
	lk.options.RetryStrategy = redislock.ExponentialBackoff(min, max)
}

////////////////////////////////////////////
//
//
//
//  Lock 相关函数
////////////////////////////////////////////
func (lk *Lock) Key() string {
	return lk.locker.Key()
}

func (lk *Lock) Token() string {
	return lk.locker.Token()
}

func (lk *Lock) Metadata() string {
	return lk.locker.Metadata()
}

func (lk *Lock) TTL(ctx context.Context) (time.Duration, error) {
	return lk.locker.TTL(ctx)
}

func (lk *Lock) Refresh(ctx context.Context, ttl time.Duration) error {
	return lk.locker.Refresh(ctx, ttl, lk.options)
}

func (lk *Lock) Release(ctx context.Context) error {
	return lk.locker.Release(ctx)
}
