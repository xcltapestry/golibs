package trace

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
	uuid "github.com/nu7hatch/gouuid"
)

// GetTraceID generate a unique ID for the server
func GetTraceID() string {
	id, _ := uuid.NewV4()

	rid := id.String()
	if rid == "" || len(rid) < 8 { // 取毫秒
		return fmt.Sprintf("%d", time.Now().UTC().UnixNano()/1000)
	}
	return rid[0:7] // 取uuid的前8位
}
