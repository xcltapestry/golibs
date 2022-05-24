package errorx

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
  备注:
    通常项目可能是由多个端(android,iOS,web,mwb等)，及多种开发语言下的微服务构成。
	对于错误或业务基础信息等，在早期就必须定义一个统一的数据来源。
	否则在后期代价会很大！
*/

import (
	"github.com/pkg/errors"
	"github.com/xcltapestry/golibs/global"
)

//////////////////////////////////////////////
//
// 在应用服务启动时，需先初始化GeneralErrorMap。
// 在错误处理时，如需附加更多信息，可以用  errors.WithMessage
// 错误类型判断可以用 error.Is
//
//const (
//	SYS_ServerError   = 10101 // Internal Server Error
//	SYS_CallHTTPError = 10102 // 调用第三方HTTP接口失败
//)
//
//  errorx.GeneralErrorMap =
//  map[int]string{
//	SYS_ServerError:   "Internal Server Error",
//	SYS_CallHTTPError: "Too Many Requests",
// }
//////////////////////////////////////////////

type GeneralError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *GeneralError) Error() string {
	return e.Message
}

func (self *GeneralError) GetMessage(code int) string {
	errMessage, found := global.Errorx_GeneralErrorMap[code]
	if found == false {
		return "Unknown error code"
	}
	return errMessage
}

func NewGeneralError(code int) error {
	gerror := new(GeneralError)
	gerror.Code = code
	gerror.Message = gerror.GetMessage(code)
	return errors.Wrap(gerror, "")
}

// 使用:
// var myError = new(GeneralError)
//  if errors.As(err, &myError) { myError.Code, myError.Message }

//////////////////////////////////////////////
//
// 自动化方式，建议错误码相关的放在应用本地，
// 不放在这个公共库，公共库不要绑定具体业务信息。
// 权责，边界要清晰
//
// 自动生成前需安装stringer
// go get golang.org/x/tools/cmd/stringer
//////////////////////////////////////////////
//type ErrCode int64
//
//func (self *ErrCode) String() string {
//	code := int64(*self)
//	return strconv.FormatInt(code, 10)
//}
//
//type CustomError struct {
//	Code    ErrCode `json:"code"`
//	Message string  `json:"message"`
//}
//
//func (self CustomError) Error() string {
//	return self.Code.String()
//}

//func NewCustomError(code ErrCode) error {
//	return errors.Wrap(&CustomError{
//		Code:    code,
//		Message: code.String(),
//	}, "")
//}

// 自动化生成例子
////go:generate stringer -type ErrCode -linecomment
//const (
//	SYS_ServerError   ErrCode = 10101 // Internal Server Error
//	SYS_CallHTTPError ErrCode = 10102 // 调用第三方HTTP接口失败
//)
