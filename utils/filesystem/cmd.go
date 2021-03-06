package filesystem

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
	"bytes"
	"github.com/xcltapestry/golibs/utils/sysutil"
	"os/exec"
)

//CheckCmdExists 检测命令是否存在
func CheckCmdExists(cmd string) bool {
	_, err := exec.LookPath(cmd) //ls
	if err != nil {
		return false //fmt.Printf("didn't find 'ls' executable\n")
	}
	return true //	fmt.Printf("'ls' executable is in '%s'\n", path)
}

func ExecCommand(command string) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errout bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	if sysutil.IsWindows() {
		cmd = exec.Command("cmd")
	}
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err = cmd.Run()

	if err != nil {
		stderr = string(errout.Bytes())
	}
	stdout = string(out.Bytes())

	return
}
