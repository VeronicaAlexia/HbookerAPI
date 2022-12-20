package book

import (
	"github.com/VeronicaAlexia/HbookerAPI/Template"
	"github.com/VeronicaAlexia/HbookerAPI/request"
)

func GET_BOOK_INFORMATION(bid string) Template.Detail {
	var book Template.Detail
	request.NewHttpUtils(request.BOOK_GET_INFO_BY_ID, "POST").Add("book_id", bid).NewRequests().Unmarshal(&book)
	return book
}

func GET_DIVISION_LIST_BY_BOOKID(BookId string) Template.VolumeList {
	var divisionList Template.VolumeList
	request.NewHttpUtils(request.GET_DIVISION_LIST_NEW, "POST").Add("book_id", BookId).NewRequests().Unmarshal(&divisionList)
	return divisionList
}

func GET_CATALOGUE(DivisionId string) Template.Chapter {
	var chapterList Template.Chapter
	request.NewHttpUtils(request.GET_CHAPTER_UPDATE, "POST").Add("division_id", DivisionId).NewRequests().Unmarshal(&chapterList)
	return chapterList
}
