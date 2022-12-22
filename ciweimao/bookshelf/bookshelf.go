package bookshelf

import (
	"github.com/VeronicaAlexia/HbookerAPI/Template"
	"github.com/VeronicaAlexia/HbookerAPI/request"
)

func GET_BOOK_SHELF_INDEXES_INFORMATION(shelf_id string) *Template.BookList {
	var bookList Template.BookList
	params := map[string]string{"shelf_id": shelf_id, "direction": "prev", "last_mod_time": "0"}
	request.Post(request.BOOKSHELF_GET_SHELF_BOOK_LIST).Params(params).NewRequests().Unmarshal(&bookList)
	return &bookList
}

func GET_BOOK_SHELF_INFORMATION() *Template.ShelfList {
	var shelfList Template.ShelfList
	request.Post(request.BOOKSHELF_GET_SHELF_LIST).NewRequests().Unmarshal(&shelfList)
	return &shelfList
}
