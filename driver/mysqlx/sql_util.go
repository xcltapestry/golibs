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
	_ "github.com/go-sql-driver/mysql"
)

// RunWithTx 在事务中执行一系列操作.
// 可以无需手动管理事务（begin 和 rollback/commit）
// 当 f 未返回错误时, 默认数据库操作成功并且尝试 Commit,
// 有任何错误则将事务 Rollback.
/*
使用例子
err := mysqlx.RunWithTx(ctx, sysdb.SysDB, func(ctx context.Context, tx *sqlx.Tx) error {
	query := "INSERT INTO t1(id,role_code,created_by) VALUES(?,?,?) ON DUPLICATE KEY UPDATE last_updated_by=?"
	for _, v := range ar {
		_, err := tx.ExecContext(ctx, query, v.id, v.rc, ea, ean)
		if err != nil {
			err = errors.Errorf(err, "failed to insert, data:%+v", v)
			return err
		}
	}
	return nil
})
if err != nil {
	err = errors.Errorf(err, "")
	return err
}
*/
func RunWithTx(ctx context.Context, db *sql.DB, f func(context.Context, *sql.Tx) error) error {
	if db == nil {
		return errors.New("[RunWithTx] db is nil")
	}
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("[RunWithTx] db.Begin() exec trans failed err:%s ", err.Error())
	}
	defer tx.Rollback()

	if err := f(ctx, tx); err != nil {
		return fmt.Errorf("[RunWithTx] f(tx) exec trans failed err:%s ", err.Error())
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[RunWithTx] tx.Commit() exec trans failed err:%s ", err.Error())
	}
	return nil
}
