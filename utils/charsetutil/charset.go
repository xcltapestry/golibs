package charsetutil

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
)

func GbkToUtf8(s []byte) ([]byte, error) {
	//第二个参数为“transform.Transformer”接口，simplifiedchinese.GBK.NewDecoder()包含了该接口
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//DeterminEncoding 会根据 HTML 页面中的 meta 元信息猜测网页编码
func DeterminEncoding(r io.Reader) (encoding.Encoding, string, error) {
	//这里的r读取完得保证resp.Body还可读
	body, err := bufio.NewReader(r).Peek(1024)
	if err != nil && err != io.EOF { //ErrBufferFull
		return nil, "", fmt.Errorf("[DeterminEncoding] err:%v", err)
	}
	// DetermineEncoding会截取1024个字符进行编码格式的推断
	// encoding, name, certain := charset.DetermineEncoding(gbk, "text/html")
	e, name, _ := charset.DetermineEncoding(body, "")
	return e, name, nil
}
