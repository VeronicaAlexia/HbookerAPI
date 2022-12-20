package account

import (
	"github.com/VeronicaAlexia/HbookerAPI/Template"
	"github.com/VeronicaAlexia/HbookerAPI/request"
	"strconv"
	"time"
)

func GET_LOGIN_TOKEN(account, password string) *Template.Login {
	request.Post(request.MY_SIGN_LOGIN).Add("login_name", account).
		Add("password", password).NewRequests().Unmarshal(&Template.Login{})
	return &Template.Login{}
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
	request.Post(request.USE_GEETEST).NewRequests().Unmarshal(&Template.Geetest{})
	return &Template.Geetest{}
}

func GET_GEETEST_REGISTER(UserID string) (string, string) {
	request.Post(request.GEETEST_REGISTER).Add("user_id", UserID).
		Add("t", strconv.FormatInt(time.Now().UnixNano()/1e6, 10)).NewRequests().Unmarshal(&Template.Challenge)
	return Template.Challenge.Challenge, Template.Challenge.Gt
}
