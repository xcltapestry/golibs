package strutil

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
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

// 部份函数来自 https://github.com/duke-git/lancet 它的函数很赞
// 可直接引用 lancet

var isIntStrRegexMatcher *regexp.Regexp = regexp.MustCompile(`^[\+-]?\d+$`)

func IsIntStr(str string) bool {
	return isIntStrRegexMatcher.MatchString(str)
}

var isAlphaRegexMatcher *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z]+$`)

func IsAlpha(str string) bool {
	return isAlphaRegexMatcher.MatchString(str)
}

func IsAllUpper(str string) bool {
	for _, r := range str {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return str != ""
}

func IsAllLower(str string) bool {
	for _, r := range str {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return str != ""
}

var isEmailRegexMatcher *regexp.Regexp = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)

func IsEmail(email string) bool {
	return isEmailRegexMatcher.MatchString(email)
}

var isUrlRegexMatcher *regexp.Regexp = regexp.MustCompile(`^((ftp|http|https?):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(([a-zA-Z0-9]+([-\.][a-zA-Z0-9]+)*)|((www\.)?))?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?))(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`)

func IsUrl(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}

	return isUrlRegexMatcher.MatchString(str)
}

func IsEmptyString(str string) bool {
	return len(str) == 0
}

var containChineseRegexMatcher *regexp.Regexp = regexp.MustCompile("[\u4e00-\u9fa5]")

func ContainChinese(s string) bool {
	return containChineseRegexMatcher.MatchString(s)
}

var isCreditCardRegexMatcher *regexp.Regexp = regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11}|6[27][0-9]{14})$`)

func IsCreditCard(creditCart string) bool {
	return isCreditCardRegexMatcher.MatchString(creditCart)
}

var isBase64RegexMatcher *regexp.Regexp = regexp.MustCompile(`^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`)

func IsBase64(base64 string) bool {
	return isBase64RegexMatcher.MatchString(base64)
}

var isChineseIdNumRegexMatcher *regexp.Regexp = regexp.MustCompile(`^[1-9]\d{5}(18|19|20|21|22)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)

func IsChineseIdNum(id string) bool {
	return isChineseIdNumRegexMatcher.MatchString(id)
}

var isChineseMobileRegexMatcher *regexp.Regexp = regexp.MustCompile("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$")

func IsChineseMobile(mobileNum string) bool {
	return isChineseMobileRegexMatcher.MatchString(mobileNum)
}

var isChinesePhoneRegexMatcher *regexp.Regexp = regexp.MustCompile(`\d{3}-\d{8}|\d{4}-\d{7}`)

func IsChinesePhone(phone string) bool {
	return isChinesePhoneRegexMatcher.MatchString(phone)
}
