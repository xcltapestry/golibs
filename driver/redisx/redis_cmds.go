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
)

func (rc *Client) Append(key, value string) (int64, error) {
	return rc.rdb.Append(context.Background(), key, value).Result()
}

// Del -
func (rc *Client) Del(keys ...string) (int64, error) {
	return rc.rdb.Del(context.Background(), keys...).Result()
}

// Close -
func (rc *Client) Close() error {
	return rc.rdb.Close()
}

// Ping -
func (rc *Client) Ping() error {
	_, err := rc.rdb.Ping(context.Background()).Result()
	return err
}

func (rc *Client) Exists(keys ...string) (int64, error) {
	return rc.rdb.Exists(context.Background(), keys...).Result()
}

func (rc *Client) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return rc.rdb.Set(context.Background(), key, value, expiration).Result()
}

func (rc *Client) SetEX(key string, value interface{}, expiration time.Duration) (string, error) {
	return rc.rdb.SetEX(context.Background(), key, value, expiration).Result()
}

func (rc *Client) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return rc.rdb.SetNX(context.Background(), key, value, expiration).Result()
}

func (rc *Client) Get(key string) (string, error) {
	return rc.rdb.Get(context.Background(), key).Result()
}

func (rc *Client) Expire(key string, expiration time.Duration) (bool, error) {
	return rc.rdb.Expire(context.Background(), key, expiration).Result()
}

func (rc *Client) TTL(key string) (time.Duration, error) {
	return rc.rdb.TTL(context.Background(), key).Result()
}

func (rc *Client) Incr(key string) (int64, error) {
	return rc.rdb.Incr(context.Background(), key).Result()
}

func (rc *Client) Decr(key string) (int64, error) {
	return rc.rdb.Decr(context.Background(), key).Result()
}

func (rc *Client) LRange(key string, start, stop int64) ([]string, error) {
	return rc.rdb.LRange(context.Background(), key, start, stop).Result()
}

func (rc *Client) Publish(channel string, message interface{}) (int64, error) {
	return rc.rdb.Publish(context.Background(), channel, message).Result()
}

// reflect.TypeOf(v)
