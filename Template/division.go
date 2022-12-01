package Template

type VolumeList struct {
	Code string       `json:"code"`
	Tip  interface{}  `json:"tip"`
	Data DivisionData `json:"data"`
}

type DivisionData struct {
	DivisionList []struct {
		DivisionID    string `json:"division_id"`
		DivisionIndex string `json:"division_index"`
		DivisionName  string `json:"division_name"`
		Description   string `json:"description"`
	} `json:"division_list"`
}
