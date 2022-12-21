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

func GET_KET_BY_CHAPTER_ID(chapterId string) *Template.Key {
	get_chapter_cmd := &Template.Key{}
	request.Post(request.GET_CHAPTER_KEY).Add("chapter_id", chapterId).NewRequests().Unmarshal(get_chapter_cmd)
	return get_chapter_cmd
}

func GET_CHAPTER_CONTENT(chapterId, chapter_key string) *Template.Content {
	content := &Template.Content{}
	params := map[string]string{"chapter_id": chapterId, "chapter_key": chapter_key}
	request.Post(request.GET_CPT_IFM).Params(params).NewRequests().Unmarshal(content)
	return content
}

func GET_CATALOGUE_OLD(DivisionId string) Template.Chapter {
	var chapterList Template.Chapter
	request.Post(request.GET_CHAPTER_UPDATE).Add("division_id", DivisionId).NewRequests().Unmarshal(&chapterList)
	return chapterList
}

func GET_DIVISION_LIST_BY_BOOKID_OLD(BookId string) Template.VolumeList {
	var divisionList Template.VolumeList
	request.Post(request.GET_DIVISION_LIST_NEW).Add("book_id", BookId).NewRequests().Unmarshal(&divisionList)
	return divisionList
}
