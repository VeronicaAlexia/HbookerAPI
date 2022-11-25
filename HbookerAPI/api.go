package HbookerAPI

import (
	"github.com/VeronicaAlexia/HbookerAPI/HbookerStruct"
	"github.com/VeronicaAlexia/HbookerAPI/HbookerStruct/bookshelf"
	"github.com/VeronicaAlexia/HbookerAPI/HbookerStruct/division"
	"strconv"
	"time"
)

func GET_DIVISION_LIST_BY_BOOKID(BookId string) *division.VolumeList {
	NewHttpUtils(GET_DIVISION_LIST, "POST").Add("book_id", BookId).NewRequests().Unmarshal(&division.VolumeList{})
	return &division.VolumeList{}
}

func GET_CATALOGUE(DivisionId string) []HbookerStruct.ChapterList {
	NewHttpUtils(GET_CHAPTER_UPDATE, "POST").Add("division_id", DivisionId).NewRequests().Unmarshal(&HbookerStruct.Chapter)
	return HbookerStruct.Chapter.Data.ChapterList
}

func GET_BOOK_SHELF_INDEXES_INFORMATION(shelf_id string) *bookshelf.BookList {
	NewHttpUtils(BOOKSHELF_GET_SHELF_BOOK_LIST, "POST").Add("shelf_id", shelf_id).Add("direction", "prev").
		Add("last_mod_time", "0").NewRequests().Unmarshal(&bookshelf.BookList{})
	return &bookshelf.BookList{}
}

func GET_BOOK_SHELF_INFORMATION() *bookshelf.GetShelfList {
	NewHttpUtils(BOOKSHELF_GET_SHELF_LIST, "POST").NewRequests().Unmarshal(&bookshelf.GetShelfList{})
	return &bookshelf.GetShelfList{}
}
func GET_BOOK_INFORMATION(bid string) *HbookerStruct.Detail {
	NewHttpUtils(BOOK_GET_INFO_BY_ID, "POST").
		Add("book_id", bid).NewRequests().Unmarshal(&HbookerStruct.LoginData{})
	return &HbookerStruct.Detail{}
}

func GET_SEARCH(KeyWord string, page int) *HbookerStruct.Search {
	NewHttpUtils(BOOKCITY_GET_FILTER_LIST, "POST").Add("count", "10").
		Add("page", strconv.Itoa(page)).Add("category_index", "0").Add("key", KeyWord).NewRequests().
		Unmarshal(&HbookerStruct.Search{})
	return &HbookerStruct.Search{}

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
