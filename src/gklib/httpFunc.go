package gklib

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetHttpUrl(requestUrl string, data map[string]string) string {
	if _, ok := data["data"]; ok && strings.Contains(requestUrl, "{data") {
		tmpData := strings.Split(data["data"], ",")

		for index, _ := range tmpData {
			requestUrl = strings.Replace(requestUrl, "{data"+strconv.Itoa(index)+"}", tmpData[index], -1)
		}
	}

	if _, ok := data["param"]; ok {
		requestUrl += "?" + data["param"]
	}

	return requestUrl
}

func GetDataByHttpGet(url string, timeoutSec int) (contents string, httpStatusCode string, err error) {
	timeout := time.Duration(time.Duration(timeoutSec) * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36")
	response, err := client.Do(req)

	defer response.Body.Close()

	var httpStatusCodeValue int

	if err == nil && response.StatusCode == 200 {
		contentsBytes, err := ioutil.ReadAll(response.Body)
		if err == nil {
			contents = string(contentsBytes)
			httpStatusCodeValue = response.StatusCode
		}
	} else {
		if err == nil {
			httpStatusCodeValue = response.StatusCode
		} else {
			// timeout or unknown error
			httpStatusCodeValue = 599
		}
	}

	return contents, strconv.Itoa(httpStatusCodeValue), err
}
