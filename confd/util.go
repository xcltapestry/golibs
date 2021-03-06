package confd

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
	"github.com/xcltapestry/golibs/utils/strutil"
	"path/filepath"
	"strings"
	"time"
)

func (cfd *Confd) Get(key string) interface{} {
	if cfd.viper == nil {
		return nil
	}
	cfd.rmu.Lock()
	defer cfd.rmu.Unlock()
	return cfd.viper.Get(key)
}

func (cfd *Confd) GetString(key string) string {
	if cfd.viper == nil {
		return ""
	}
	cfd.rmu.Lock()
	defer cfd.rmu.Unlock()
	return cfd.viper.GetString(key)
}

func (cfd *Confd) GetBool(key string) bool {
	cfd.rmu.Lock()
	defer cfd.rmu.Unlock()
	return cfd.viper.GetBool(key)
}

func (cfd *Confd) GetDuration(key string) time.Duration {
	cfd.rmu.Lock()
	defer cfd.rmu.Unlock()
	return cfd.viper.GetDuration(key)
}

func (cfd *Confd) GetInt(key string) int {
	cfd.rmu.Lock()
	defer cfd.rmu.Unlock()
	return cfd.viper.GetInt(key)
}

func (cfd *Confd) AllKeys() []string {
	cfd.rmu.Lock()
	defer cfd.rmu.Unlock()
	return cfd.viper.AllKeys()
}

//paseConfigFile 解析文件，得到path,ext
func paseConfigFile(confFile string) (string, string, error) {
	// if _, err := os.Stat(confFile); os.IsNotExist(err) {
	// 	return "", "", err
	// }
	path, fileName := filepath.Split(confFile)
	ext := strings.ToLower(filepath.Ext(fileName))
	fext := strutil.SubString(ext, 1, len(ext))

	return path, fext, nil
}
