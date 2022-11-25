package _struct

type Key struct {
	Code string  `json:"code"`
	Data KeyData `json:"data"`
}

type KeyData struct {
	Command string `json:"command"`
}
