package HbookerAPI

import (
	"github.com/VeronicaAlexia/HbookerAPI/Template"
	"github.com/VeronicaAlexia/HbookerAPI/request"
	"strconv"
	"time"
)

func GET_BOOK_SHELF_INDEXES_INFORMATION(shelf_id string) *Template.BookList {
	var bookList Template.BookList
	request.Post(request.BOOKSHELF_GET_SHELF_BOOK_LIST).Add("shelf_id", shelf_id).Add("direction", "prev").
		Add("last_mod_time", "0").NewRequests().Unmarshal(&bookList)
	return &bookList
}

func GET_BOOK_SHELF_INFORMATION() *Template.GetShelfList {
	var shelfList Template.GetShelfList
	request.Post(request.BOOKSHELF_GET_SHELF_LIST).NewRequests().Unmarshal(&shelfList)
	return &shelfList
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

func GET_TAG_BOOK(page int) Template.Search {
	var search Template.Search
	params := map[string]string{
		"filter_word":    "",
		"count":          "10",
		"tags":           `[{"filter":"1","tag":""}]`,
		"use_daguan":     "0",
		"page":           strconv.Itoa(page),
		"is_paid":        "",
		"category_index": "0",
		"key":            "",
		"filter_uptime":  "",
		"up_status":      "",
		"order":          "uptime",
	}
	request.Post(request.BOOKCITY_GET_FILTER_LIST).Params(params).NewRequests().Unmarshal(&search)
	return search

}

func GET_LOGIN_TOKEN(account, password string) *Template.Login {
	request.Post(request.MY_SIGN_LOGIN).Add("login_name", account).
		Add("password", password).NewRequests().Unmarshal(&Template.Login{})
	return &Template.Login{}
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

func GET_KET_BY_CHAPTER_ID(chapterId string) *Template.Key {
	request.Post(request.GET_CHAPTER_KEY).Add("chapter_id", chapterId).NewRequests().Unmarshal(&Template.Key{})
	return &Template.Key{}
}

func GET_CHAPTER_CONTENT(chapterId, chapter_key string) *Template.Content {
	request.Post(request.GET_CPT_IFM).Add("chapter_id", chapterId).
		Add("chapter_command", chapter_key).NewRequests().Unmarshal(&Template.Content{})
	return &Template.Content{}
}
