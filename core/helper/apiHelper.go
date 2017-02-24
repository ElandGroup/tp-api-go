package helper

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	. "tp-api-go/core"
)

//get url response html
func GetUrlRespHtml(strUrl string, postDict map[string]string) (response string, apiError *APIError) {
	var gCurCookieJar *cookiejar.Jar
	gCurCookieJar, _ = cookiejar.New(nil)
	httpClient := &http.Client{
		Jar: gCurCookieJar,
	}

	var httpReq *http.Request
	if nil == postDict {
		httpReq, _ = http.NewRequest("GET", strUrl, nil)
	} else {
		postValues := url.Values{}
		for postKey, PostValue := range postDict {
			postValues.Set(postKey, PostValue)
		}
		postDataStr := postValues.Encode()
		postDataBytes := []byte(postDataStr)
		postBytesReader := bytes.NewReader(postDataBytes)
		httpReq, _ = http.NewRequest("POST", strUrl, postBytesReader)
		httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	}

	httpResp, err := httpClient.Do(httpReq)

	if err != nil {
		response = ""
		apiError = NewApiError(10001, err.Error())
		return
	}

	defer httpResp.Body.Close()

	body, e := ioutil.ReadAll(httpResp.Body)
	if e != nil {
		response = ""
		apiError = NewApiError(10001, e.Error())
		return
	} else {
		response = string(body)
		apiError = nil
		return
	}
}
