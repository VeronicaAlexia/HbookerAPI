package HbookerAPI

import (
	"github.com/VeronicaAlexia/HbookerAPI/HbookerStruct"
	"strconv"
	"time"
)

func GET_DIVISION_LIST_BY_BOOKID(BookId string) HbookerStruct.VolumeList {
	var divisionList HbookerStruct.VolumeList
	NewHttpUtils(GET_DIVISION_LIST, "POST").Add("book_id", BookId).NewRequests().Unmarshal(&divisionList)
	return divisionList
}

func GET_CATALOGUE(DivisionId string) HbookerStruct.Chapter {
	var chapterList HbookerStruct.Chapter
	NewHttpUtils(GET_CHAPTER_UPDATE, "POST").Add("division_id", DivisionId).NewRequests().Unmarshal(&chapterList)
	return chapterList
}

func GET_BOOK_SHELF_INDEXES_INFORMATION(shelf_id string) *HbookerStruct.BookList {
	NewHttpUtils(BOOKSHELF_GET_SHELF_BOOK_LIST, "POST").Add("shelf_id", shelf_id).Add("direction", "prev").
		Add("last_mod_time", "0").NewRequests().Unmarshal(&HbookerStruct.BookList{})
	return &HbookerStruct.BookList{}
}

func GET_BOOK_SHELF_INFORMATION() *HbookerStruct.GetShelfList {
	NewHttpUtils(BOOKSHELF_GET_SHELF_LIST, "POST").NewRequests().Unmarshal(&HbookerStruct.GetShelfList{})
	return &HbookerStruct.GetShelfList{}
}
func GET_BOOK_INFORMATION(bid string) HbookerStruct.Detail {
	var book HbookerStruct.Detail
	NewHttpUtils(BOOK_GET_INFO_BY_ID, "POST").Add("book_id", bid).NewRequests().Unmarshal(&book)
	return book
}

func GET_SEARCH(KeyWord string, page int) HbookerStruct.Search {
	var search HbookerStruct.Search
	params := map[string]string{
		"count":          "10",
		"page":           strconv.Itoa(page),
		"category_index": "0",
		"key":            KeyWord,
	}
	NewHttpUtils(BOOKCITY_GET_FILTER_LIST, "POST").params(params).NewRequests().Unmarshal(&search)
	return search

}

func GET_LOGIN_TOKEN(account, password string) *HbookerStruct.Login {
	// hbooker new version add GEETEST verification, if you enter the wrong information or log in multiple times, GEETEST verification will be triggered.
	// IP address may need to log in again after a few hours to avoid triggering verification, you can try to change the IP to avoid triggering verification.

	NewHttpUtils(MY_SIGN_LOGIN, "POST").Add("login_name", account).
		Add("password", password).NewRequests().Unmarshal(&HbookerStruct.Login{})
	return &HbookerStruct.Login{}
}

func GET_USE_GEETEST() *HbookerStruct.Geetest {
	NewHttpUtils(USE_GEETEST, "POST").NewRequests().Unmarshal(&HbookerStruct.Geetest{})
	return &HbookerStruct.Geetest{}
}

func GET_GEETEST_REGISTER(UserID string) (string, string) {
	NewHttpUtils(GEETEST_REGISTER, "POST").Add("user_id", UserID).
		Add("t", strconv.FormatInt(time.Now().UnixNano()/1e6, 10)).NewRequests().Unmarshal(&HbookerStruct.Challenge)
	return HbookerStruct.Challenge.Challenge, HbookerStruct.Challenge.Gt
}

func GET_KET_BY_CHAPTER_ID(chapterId string) *HbookerStruct.Key {
	NewHttpUtils(GET_CHAPTER_KEY, "POST").Add("chapter_id", chapterId).NewRequests().Unmarshal(&HbookerStruct.Key{})
	return &HbookerStruct.Key{}
}

func GET_CHAPTER_CONTENT(chapterId, chapter_key string) *HbookerStruct.Content {
	NewHttpUtils(GET_CPT_IFM, "POST").Add("chapter_id", chapterId).
		Add("chapter_command", chapter_key).NewRequests().Unmarshal(&HbookerStruct.Content{})
	return &HbookerStruct.Content{}
}
