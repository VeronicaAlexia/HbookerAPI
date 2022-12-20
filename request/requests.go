package request

import (
	"fmt"
	"github.com/VeronicaAlexia/HbookerAPI/config"
	"io"
	"net/http"
	"net/url"
)

type HttpUtils struct {
	url        string
	method     string
	cookie     []*http.Cookie
	response   *http.Request
	app_type   string
	query_data *url.Values
	content    []byte
}

func (is *HttpUtils) NewRequests() *HttpUtils {
	var err error
	is.content = nil
	if is.method == "GET" {
		is.response, err = http.NewRequest(is.method, is.url+"?"+is.query_data.Encode(), nil)
	} else {
		is.response, err = http.NewRequest(is.method, is.url, is.GetEncodeParams())
	}
	if err != nil {
		fmt.Println("NewRequests:", err)
		return nil
	}
	is.response.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	is.response.Header.Set("User-Agent", "Android com.kuangxiangciweimao.novel "+config.AppConfig.AppVersion)
	if response, ok := http.DefaultClient.Do(is.response); ok == nil {
		is.cookie = response.Cookies()
		result_body, _ := io.ReadAll(response.Body)
		is.content = config.Decode(string(result_body), "")
	} else {
		fmt.Println("NewRequests:", ok)
	}
	return is
}
