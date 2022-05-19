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
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"github.com/xcltapestry/golibs/logger"
)

type Confd struct {
	viper *viper.Viper
	rmu   sync.RWMutex
}

var Conf *Confd //全局配置信息

func NewConfd() *Confd {
	c := &Confd{}
	// c.appConfig = appConfig
	// c.remoteConfd = NewEtcdConfd(appConfig)
	// c.localConfd = NewConfLocalFile()
	c.viper = viper.New()

	path2, _ := os.Getwd()
	c.viper.AddConfigPath(path2)
	c.viper.AddConfigPath(".")

	c.viper.SetConfigType("yaml")

	return c
}

func (cf *Confd) getConfigType(configType string) string {
	switch configType {
	case "json", "hcl", "prop", "props", "properties", "dotenv", "env", "toml", "yaml", "yml", "ini":
		return configType
	default:
		return "yaml" //_ConfigType
	}
}

func (cf *Confd) LoadConfig(confFile string) error {

	if cf.viper == nil {
		return errors.New("[LoadConfig] viper is null. ")
	}

	path, ext, err := paseConfigFile(confFile)
	if err != nil {
		return fmt.Errorf("[LoadConfig] err:%s confFile:%s", err, confFile)
	}
	ext = strings.ToLower(ext)

	if path != "" {
		cf.viper.AddConfigPath(path)
	}
	path2, _ := os.Getwd()
	cf.viper.AddConfigPath(path2)
	cf.viper.AddConfigPath(".")

	cf.viper.SetConfigName(filepath.Base(confFile))
	cf.viper.SetConfigType(ext) //"yaml")
	// cf.viper.SetConfigType("yaml") //cf.getConfigType(ext))

	//读取配置
	err = cf.viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf(" [LoadConfig] viper 读取本地配置文件失败. err:%s path:%s path2:%s confFile:%s",
			err.Error(), path, path2, confFile)
	}

	logger.Info("[LoadConfig] 已完成本地配置文件读取 confFile:", confFile)

	return nil
}
