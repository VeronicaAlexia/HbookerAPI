package Template

type Recommend struct {
	Code string      `json:"code"`
	Tip  interface{} `json:"tip"`
	Data struct {
		BookList []struct {
			BookID          string `json:"book_id"`
			BookName        string `json:"book_name"`
			Description     string `json:"description"`
			AuthorName      string `json:"author_name"`
			Cover           string `json:"cover"`
			DiscountEndTime string `json:"discount_end_time"`
			UpStatus        string `json:"up_status"`
			TotalWordCount  string `json:"total_word_count"`
			Introduce       string `json:"introduce"`
		} `json:"book_list"`
	} `json:"data"`
}

type Recommend2 struct {
	Code string      `json:"code"`
	Tip  interface{} `json:"tip"`
	Data struct {
		ModuleList []struct {
			ModuleType string `json:"module_type"`
			BossModule struct {
				DesBookList []struct {
					BookID          string `json:"book_id"`
					BookName        string `json:"book_name"`
					CategoryIndex   string `json:"category_index"`
					Description     string `json:"description"`
					AuthorName      string `json:"author_name"`
					Cover           string `json:"cover"`
					DiscountEndTime string `json:"discount_end_time"`
				} `json:"des_book_list"`
			} `json:"boss_module,omitempty"`
		} `json:"module_list"`
	} `json:"data"`
}
