package sysutil

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
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

// Signals 用于处理系统消息,以便处理更新或安全中断
func Signals(q chan bool, fsigexit func(), fsigusr1 func()) {
	sigs := make(chan os.Signal)
	defer close(sigs)

EXIT:
	switch runtime.GOOS {
	case "windows":
		for {
			signal.Notify(sigs, syscall.SIGQUIT,
				syscall.SIGTERM,
				syscall.SIGINT)

			switch <-sigs {
			case syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT:
				fsigexit()
				break EXIT
			}
		}
	default:
		//for {
		//	signal.Notify(sigs, syscall.SIGQUIT,
		//		syscall.SIGTERM,
		//		syscall.SIGINT,
		//		syscall.SIGUSR1,
		//		syscall.SIGUSR2)
		//
		//	switch <-sigs {
		//	case syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT:
		//		fsigexit()
		//		break EXIT
		//	case syscall.SIGUSR1:
		//		fsigusr1()
		//	case syscall.SIGUSR2:
		//		fsigexit()
		//		break EXIT
		//	}
		//}

	} //end switch default
	q <- true
}
