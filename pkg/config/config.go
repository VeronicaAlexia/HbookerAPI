package config

var AppConfig = struct {
	App         bool   `json:"app"`
	Cookie      string `json:"cookie"`
	LoginToken  string `json:"login_token"`
	Account     string `json:"account"`
	AppVersion  string `json:"app_version"`
	DeviceToken string `json:"device_token"`
}{}
