package security

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
	"fmt"
	"net/url"
)

// IsURLSafe
// 判断请求的URL是否安全, 如果安全, 返回 true, nil
// 如果不安全, 返回false, error.Error() 为相应的描述
func IsURLSafe(rawURL string) (safe bool, err error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false, fmt.Errorf("secutils: url not safe, invalid url:%s", rawURL)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false, fmt.Errorf("secutils: url not safe, invalid schema, url:%s", rawURL)
	}
	return true, nil
}
