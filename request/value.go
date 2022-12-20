package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (is *HttpUtils) Unmarshal(s any) *HttpUtils {
	err := json.Unmarshal(is.content, s)
	if err != nil {
		fmt.Println("Unmarshal:", err)
	}
	return is
}
func (is *HttpUtils) WriteJson() *HttpUtils {
	info, err := json.MarshalIndent(string(is.content), "", "  ")
	if err != nil {
		fmt.Println("WriteJson:", err)
	} else {
		os.WriteFile("result.json", info, 0666)
	}
	return is
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
func (is *HttpUtils) Params(param map[string]string) *HttpUtils {
	for key, value := range param {
		is.Add(key, value)
	}
	return is
}
