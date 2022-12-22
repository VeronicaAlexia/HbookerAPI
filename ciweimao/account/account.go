package account

import (
	"github.com/VeronicaAlexia/HbookerAPI/Template"
	"github.com/VeronicaAlexia/HbookerAPI/request"
	"strconv"
	"time"
)

func GET_LOGIN_TOKEN(account, password string) *Template.Login {
	var login Template.Login
	params := map[string]string{"login_name": account, "password": password}
	request.Post(request.MY_SIGN_LOGIN).Params(params).NewRequests().Unmarshal(&login)
	return &login
}

func GET_AUTO_SIGN(uuid string) {
	params := map[string]string{
		"oauth_type":     "",
		"uuid":           "android" + uuid,
		"oauth_union_id": "",
		"gender":         "1",
		"channel":        "PCdownloadC",
		"oauth_open_id":  "",
	}
	request.Post(request.SIGNUP).NewRequests().Params(params)
}
func GET_USE_GEETEST() *Template.Geetest {
	var geetest Template.Geetest
	request.Post(request.USE_GEETEST).NewRequests().Unmarshal(&geetest)
	return &geetest
}

func GET_GEETEST_REGISTER(UserID string) (string, string) {
	challenge := Template.Challenge
	var params = map[string]string{"user_id": UserID, "t": strconv.FormatInt(time.Now().UnixNano()/1e6, 10)}
	request.Post(request.GEETEST_REGISTER).Params(params).NewRequests().Unmarshal(&challenge)
	return Template.Challenge.Challenge, Template.Challenge.Gt
}
