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
	"github.com/pkg/errors"
	"time"
)

type AuthService struct {
	j           *JWT
	tokenExpiry time.Duration // Token有效期
	issuer      string
}

func NewAuthService() *AuthService {
	a := &AuthService{}
	a.j = NewJWT()
	a.issuer = "xcl"
	a.tokenExpiry = time.Duration(24 * 7 * time.Hour)
	return a
}

//GenerateToken 生成token
func (a *AuthService) GenerateToken(info TokenInfo) (string, error) {
	j := NewJWT()

	// 构造用户claims信息(负荷)
	claims := CustomClaims{TokenInfo: info}
	claims.ExpiresAt = time.Now().Add(a.tokenExpiry).Unix()
	claims.Issuer = a.issuer // 签名颁发者

	// 根据claims生成token对象
	token, err := j.GenerateToken(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

//VerifyToken 效验有效性
func (a *AuthService) VerifyToken(token string) (*CustomClaims, error) {
	j := NewJWT()

	claims, err := j.ParserToken(token)
	if err != nil {
		return nil, err
	}
	if claims == nil {
		return nil, errors.New("claims is null")
	}
	return claims, nil
}

//GetTokenMetadata 解析得到tokeninfo 数据
func (a *AuthService) GetTokenMetadata(data *CustomClaims) (*TokenInfo, error) {

	if data == nil {
		return nil, errors.New(" data is null")
	}

	tkInfo := &TokenInfo{}
	tkInfo.Env = data.Env
	tkInfo.Name = data.Name
	tkInfo.Ver = data.Ver
	tkInfo.UID = data.UID
	tkInfo.Email = data.Email
	tkInfo.Reserved = data.Reserved
	tkInfo.Sys = data.Sys

	return tkInfo, nil
}
