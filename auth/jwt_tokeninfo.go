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

import jwt "github.com/dgrijalva/jwt-go"

type TokenInfo struct {
	Env      string `json:"env"` // 用于区分如产线、测试等不同环境，以防token混用到不同环境。 也可通过不同环境不同key来区分
	Sys      string `json:"sys"` // 用于区分这个Token是用在什么系统上
	Ver      string `json:"ver"` // 即是用哪版的key生成的，用于解决key变更期的检验兼容处理
	UID      int64  `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Reserved string `json:"reserved"` // 预留字段，以便根据实际情况做扩展
}

//CustomClaims 自定义的Name和Email作为有效载荷的一部分
// ps: 就算jwt token验证通过了，也需要再用其它方式检查一遍uid的有效性，
//		以保证如用户id失效时能立即中止访问。
type CustomClaims struct {
	TokenInfo
	jwt.StandardClaims // StandardClaims结构体实现了Claims接口(Valid()函数)
}
