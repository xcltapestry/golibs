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
	"time"
)

type Options struct {
	Network      string
	Addr         string
	Username     string
	Password     string
	DB           int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	MinIdleConns int
	MaxConnAge   time.Duration
	PoolTimeout  time.Duration
	IdleTimeout  time.Duration
}

func (copt *Options) setDriverOptions() *redis.Options {
	opt := &redis.Options{}
	opt.Network = copt.Network
	opt.Addr = copt.Addr
	opt.Username = copt.Username
	opt.Password = copt.Password
	opt.DB = copt.DB
	opt.DialTimeout = copt.DialTimeout
	opt.ReadTimeout = copt.ReadTimeout
	opt.WriteTimeout = copt.WriteTimeout
	opt.PoolSize = copt.PoolSize
	opt.MinIdleConns = copt.MinIdleConns
	opt.MaxConnAge = copt.MaxConnAge
	opt.PoolTimeout = copt.PoolTimeout
	opt.IdleTimeout = copt.IdleTimeout

	return opt
}
