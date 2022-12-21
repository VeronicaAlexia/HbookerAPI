package bookshelf

import (
	"github.com/VeronicaAlexia/HbookerAPI/Template"
	"github.com/VeronicaAlexia/HbookerAPI/request"
)

func GET_BOOK_SHELF_INDEXES_INFORMATION(shelf_id string) *Template.BookList {
	var bookList *Template.BookList
	request.Post(request.BOOKSHELF_GET_SHELF_BOOK_LIST).Add("shelf_id", shelf_id).Add("direction", "prev").
		Add("last_mod_time", "0").NewRequests().Unmarshal(bookList)
	return bookList
}

func GET_BOOK_SHELF_INFORMATION() *Template.GetShelfList {
	var shelfList *Template.GetShelfList
	request.Post(request.BOOKSHELF_GET_SHELF_LIST).NewRequests().Unmarshal(shelfList)
	return shelfList
}
