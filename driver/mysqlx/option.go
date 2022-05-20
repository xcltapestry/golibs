package mysqlx

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
	_ "github.com/go-sql-driver/mysql"
)

type Option struct {
	maxOpenConns, maxIdleConns int
	maxConnLifetimeSec         int
}

func NewOption() *Option {
	return &Option{
		maxConnLifetimeSec: _defaultMaxConnLifetimeSec,
		maxOpenConns:       50,
		maxIdleConns:       5,
	}
}

func (self *Option) MaxConnLifetimeSec(lifetimeSec int) *Option {
	self.maxConnLifetimeSec = lifetimeSec
	return self
}

func (self *Option) MaxOpenConns(conns int) *Option {
	self.maxOpenConns = conns
	return self
}

func (self *Option) MaxIdleConns(conns int) *Option {
	self.maxIdleConns = conns
	return self
}
