package auth

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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJWTToken(t *testing.T) {

	info := TokenInfo{}
	info.Env = "dev"
	info.Name = "name"
	info.Ver = "v0.1"
	info.UID = 12567

	auth := NewAuthService()
	tkstr, err := auth.GenerateToken(info)
	if err != nil {
		t.Log(" err:", err)
		return
	}
	t.Log(" tkstr:", tkstr)

	ret, err := auth.VerifyToken(tkstr)
	if err != nil {
		t.Log(" err:", err)
		return
	}
	t.Log(ret)

	t.Log(" dev:", ret.Env)
	t.Log(" Name:", ret.Name)
	t.Log(" Ver:", ret.Ver)
	t.Logf(" ID: %d", ret.UID)

	assert.Equal(t, info.Env, ret.Env, "they should be equal")

}
