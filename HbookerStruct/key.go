package HbookerStruct

type Key struct {
	Code string      `json:"code"`
	Tip  interface{} `json:"tip"`
	Data struct {
		Command string `json:"command"`
	} `json:"data"`
}
