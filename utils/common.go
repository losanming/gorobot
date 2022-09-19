package utils

import (
	"encoding/base64"
	"errors"
	"example.com/m/global"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func SendRequest(url string, body io.Reader, addHeaders map[string]string, method string) (resp []byte, err error) {
	// 创建req
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	checkTime := time.Now().Format("2006-01-02 15")
	checkTimeStr := base64.StdEncoding.EncodeToString([]byte(checkTime))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("token", checkTimeStr)
	//设置headers
	if len(addHeaders) > 0 {
		for k, v := range addHeaders {
			request.Header.Add(k, v)
		}
	}
	//发送请求
	c := &http.Client{}
	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = errors.New("http status err")
		fmt.Printf("sendRequest failed, url=%v, response status code=%d", url, response.StatusCode)
		return
	}

	//读取结果
	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return
}

func GetBaseUrl(path string) (url string) {
	url = global.HOSTPORT + path
	return url
}

func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}
