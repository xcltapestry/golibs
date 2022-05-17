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
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"sync"
	"time"
)

type MySQLClient struct {
	db                           *sql.DB
	dbName, user, password, host string
	option                       *Option
	closeOnce                    sync.Once
	closeCh                      chan struct{}
}

func NewMySQLClient(dbName, user, password, host string, option *Option) *MySQLClient {
	client := &MySQLClient{}
	client.dbName, client.user, client.password, client.host = dbName, user, password, host
	client.option = option
	client.closeCh = make(chan struct{})
	return client
}

func (self *MySQLClient) Connect() error {
	if strings.TrimSpace(self.dbName) == "" {
		return fmt.Errorf("[Connect] %s dbName is null.", self.dbName)
	}

	if strings.TrimSpace(self.user) == "" {
		return fmt.Errorf("[Connect] %s user is null.", self.user)
	}

	if strings.TrimSpace(self.password) == "" {
		return fmt.Errorf("[Connect] %s password is null.", self.password)
	}

	if strings.TrimSpace(self.host) == "" {
		return fmt.Errorf("[Connect] %s host is null.", self.host)
	}

	var db *sql.DB
	var err error
	db, err = manager.New(self.dbName, self.user, self.password, self.host).Set(
		manager.SetCharset("utf8"),
		manager.SetAllowCleartextPasswords(true),
		manager.SetInterpolateParams(true),
		manager.SetTimeout(1*time.Second),
		manager.SetReadTimeout(1*time.Second),
	).Port(3306).Open(true)
	if err != nil {
		return fmt.Errorf("[Connect] %s Open fail.", err.Error())
	}

	//最大打开连接数(0表示不限制)
	db.SetMaxOpenConns(self.option.maxOpenConns)
	//设置闲置连接数
	db.SetMaxIdleConns(self.option.maxIdleConns)
	//设置连接生命周期（show variables like '%wait_timeout%'; ）
	db.SetConnMaxLifetime(time.Duration(self.option.maxConnLifetimeSec) * time.Second)

	//建立实际连接
	if err = db.Ping(); err != nil {
		return fmt.Errorf("[Connect] %s Ping fail.", err.Error())
	}
	self.db = db
	return nil
}

func (self *MySQLClient) GetInstance() *sql.DB {
	if self.db != nil {
		return self.db
	}
	return nil
}

func (self *MySQLClient) Disconnect() {
	self.closeOnce.Do(func() {
		if self.closeCh != nil {
			close(self.closeCh)
		}
	})

	if self.db != nil {
		self.db.Close()
		self.db = nil
		return
	}
}

//GetOpenConnections 得到当前打开的DB连接数
func (self *MySQLClient) GetOpenConnections() int {
	if self.db == nil {
		return int(0)
	}
	return self.db.Stats().OpenConnections
}
