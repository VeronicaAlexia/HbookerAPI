package book

import (
	"github.com/VeronicaAlexia/HbookerAPI/Template"
	"github.com/VeronicaAlexia/HbookerAPI/request"
)

func GET_BOOK_INFORMATION(bid string) Template.Detail {
	var book Template.Detail
	request.Post(request.BOOK_GET_INFO_BY_ID).Add("book_id", bid).NewRequests().Unmarshal(&book)
	return book
}

func GET_DIVISION_LIST_BY_BOOKID(BookId string) Template.NewVolumeList {
	var divisionList Template.NewVolumeList
	request.Post(request.GET_DIVISION_LIST_NEW).Add("book_id", BookId).NewRequests().Unmarshal(&divisionList).WriteJson()
	return divisionList
}

func GET_CATALOGUE(DivisionId string) Template.Chapter {
	var chapterList Template.Chapter
	request.Post(request.GET_CHAPTER_UPDATE).Add("division_id", DivisionId).NewRequests().Unmarshal(&chapterList)
	return chapterList
}
