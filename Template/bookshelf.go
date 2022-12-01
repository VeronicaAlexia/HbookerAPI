package Template

type GetShelfList struct {
	Code string `json:"code"`
	Tip  any    `json:"tip"`
	Data struct {
		ShelfList []struct {
			ShelfID    string `json:"shelf_id"`
			ReaderID   string `json:"reader_id"`
			ShelfName  string `json:"shelf_name"`
			ShelfIndex string `json:"shelf_index"`
			BookLimit  string `json:"book_limit"`
		} `json:"shelf_list"`
	} `json:"data"`
	ScrollChests []interface{} `json:"scroll_chests"`
}

type BookList struct {
	Code string `json:"code"`
	Tip  any    `json:"tip"`
	Data struct {
		BookList []struct {
			IsBuy                     string   `json:"is_buy"`
			BookInfo                  BookInfo `json:"book_info"`
			ModTime                   string   `json:"mod_time"`
			LastReadChapterID         string   `json:"last_read_chapter_id"`
			LastReadChapterUpdateTime string   `json:"last_read_chapter_update_time"`
			IsNotify                  string   `json:"is_notify"`
		} `json:"book_list"`
	} `json:"data"`
	ScrollChests []interface{} `json:"scroll_chests"`
}
