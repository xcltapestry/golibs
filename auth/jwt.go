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

/*

采用哪种方案:
1. 使用RSA证书
 jwt.SigningMethodRS256 :
	//  $ openssl genrsa -out app.rsa 1024
	//  $ openssl rsa -in app.rsa -pubout > app.rsa.pub

2. 增加了 refresh_token 来获取jwttoken. jwttoken的有效期会相对较短（如15分钟）。
	同时可能增加 redis 的知识。  实现起来工作量稍大。

3. 采用 SigningMethodHS256  相对更简便。 但需要补充一个效验去验证是否还有效。
   以防用户已失效，但token仍然访问。

Token存储在哪？
	1. http 头
    2. 放在cookie,主要是因为cookie有个httponly属性

*/

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"os"
)

/*

采用哪种方案:
1. 使用RSA证书
 jwt.SigningMethodRS256 :
	//  $ openssl genrsa -out app.rsa 1024
	//  $ openssl rsa -in app.rsa -pubout > app.rsa.pub

2. 增加了 refresh_token 来获取jwttoken. jwttoken的有效期会相对较短（如15分钟）。
	同时可能增加 redis 的知识。  实现起来工作量稍大。

3. 采用 SigningMethodHS256  相对更简便。 但需要补充一个效验去验证是否还有效。
   以防用户已失效，但token仍然访问。

Token存储在哪？
	1. http 头
    2. 放在cookie,主要是因为cookie有个httponly属性

*/

type JWT struct {
	SigningKey []byte // 签名信息
}

func NewJWT() *JWT {
	j := &JWT{}
	j.SigningKey = j.getDefaultPrivateKey()
	return &JWT{
		[]byte("xcl"),
	}
}

//getDefaultPrivateKey 从默认环境变量中取得默认JWT对应key
func (j *JWT) getDefaultPrivateKey() []byte {
	signingKey := ""
	if signingKey = os.Getenv("JWT-SECRE"); signingKey == "" {
		signingKey = "xcl_168@aliyun.com"
	}
	return []byte(signingKey)
}

//SetPrivateKey 可以直接在外面设置相要的JWT对应key
func (j *JWT) SetPrivateKey(signingKey []byte) {
	j.SigningKey = signingKey
}

//GenerateToken 指定编码的算法为jwt.SigningMethodHS256
func (j *JWT) GenerateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//ParserToken token解码
func (j *JWT) ParserToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.SigningKey, nil
		})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				return nil, errors.New("token不可用")
			case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
				return nil, errors.New("无效的token")
			case ve.Errors&jwt.ValidationErrorExpired != 0:
				return nil, errors.New("token过期")
			default:
				return nil, fmt.Errorf("token不可用. errcode:%d", ve.Errors)
			}
		}
		return nil, errors.WithMessage(err, "发生异常，但不是标准的 jwt.ValidationError 类型异常。")
	}

	// 将token中的claims信息解析出来并断言成用户自定义的有效载荷结构
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token无效")
}
