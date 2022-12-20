package request

import (
	"github.com/VeronicaAlexia/HbookerAPI/pkg/config"
	"net/url"
	"strings"
)

func Get(api_url string) *HttpUtils {
	return NewHttpUtils(api_url, "GET")
}

func Post(api_url string) *HttpUtils {
	return NewHttpUtils(api_url, "POST")
}

func NewHttpUtils(api_url, method string) *HttpUtils {
	req := &HttpUtils{method: method, QueryData: &url.Values{}}
	req.url = WEB_SITE + strings.ReplaceAll(api_url, WEB_SITE, "")
	req.Add("login_token", config.AppConfig.LoginToken).
		Add("account", config.AppConfig.Account).
		Add("app_version", config.AppConfig.AppVersion).
		Add("device_token", config.AppConfig.DeviceToken)
	return req
}
