package main

import (
	"fmt"
	"github.com/VeronicaAlexia/HbookerAPI/HbookerAPI/book"
	"github.com/VeronicaAlexia/HbookerAPI/config"
	"testing"
)

func Test_Book(t *testing.T) {
	config.AppConfig.AppVersion = "2.9.290"
	config.AppConfig.Account = ""
	config.AppConfig.LoginToken = ""
	config.AppConfig.DeviceToken = "ciweimao_"
	book_info := book.GET_BOOK_INFORMATION("100280239")
	fmt.Println(book_info.Data.BookInfo.BookName)
	fmt.Println(book_info.Data.BookInfo.AuthorName)
	for _, div := range book.GET_DIVISION_LIST_BY_BOOKID(book_info.Data.BookInfo.BookID).Data.ChapterList {
		for _, chapter := range div.ChapterList {
			fmt.Println("第", div.DivisionIndex, "卷:", div.DivisionName, chapter.ChapterTitle)
		}
	}
	fmt.Println("Done: Test_Book")

}
