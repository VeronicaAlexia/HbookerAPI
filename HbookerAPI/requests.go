package HbookerAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var HbookerKey = struct {
	LoginToken  string `json:"login_token"`
	Account     string `json:"account"`
	AppVersion  string `json:"app_version"`
	DeviceToken string `json:"device_token"`
}{}

type HttpUtils struct {
	url        string
	method     string
	cookie     []*http.Cookie
	response   *http.Request
	app_type   string
	query_data *url.Values
	content    []byte
}

func MustNewRequest(method, url string, data io.Reader) *http.Request {
	if request, err := http.NewRequest(method, url, data); err != nil {
		fmt.Println("MustNewRequest:", err)
		return nil
	} else {
		return request
	}
}
func (is *HttpUtils) GetEncodeParams() *bytes.Reader {
	return bytes.NewReader([]byte(is.query_data.Encode()))
}
func (is *HttpUtils) GetResultBody() string {
	return string(is.content)
}

func (is *HttpUtils) GetCookie() []*http.Cookie {
	return is.cookie
}

func (is *HttpUtils) GetUrl() string {
	return is.url
}

func (is *HttpUtils) Add(key string, value string) *HttpUtils {
	is.query_data.Add(key, value)
	return is
}
func (is *HttpUtils) params(param map[string]string) *HttpUtils {
	for key, value := range param {
		is.Add(key, value)
	}
	return is
}

func NewHttpUtils(api_url, method string) *HttpUtils {
	req := &HttpUtils{method: method, query_data: &url.Values{}}
	req.url = WEB_SITE + strings.ReplaceAll(api_url, WEB_SITE, "")
	req.Add("login_token", HbookerKey.LoginToken).
		Add("account", HbookerKey.Account).
		Add("app_version", HbookerKey.AppVersion).
		Add("device_token", HbookerKey.DeviceToken)

	return req
}

func (is *HttpUtils) NewRequests() *HttpUtils {
	is.content = nil
	is.response = MustNewRequest(is.method, is.url, is.GetEncodeParams())
	is.response.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	is.response.Header.Set("User-Agent", "Android com.kuangxiangciweimao.novel "+HbookerKey.AppVersion)
	if response, ok := http.DefaultClient.Do(is.response); ok == nil {
		is.cookie = response.Cookies()
		result_body, _ := io.ReadAll(response.Body)
		is.content = Decode(string(result_body), "")
	} else {
		fmt.Println("NewRequests:", ok)
	}
	return is
}

func (is *HttpUtils) Unmarshal(s any) *HttpUtils {
	err := json.Unmarshal(is.content, s)
	if err != nil {
		fmt.Println("Unmarshal:", err)
	}
	return is
}
