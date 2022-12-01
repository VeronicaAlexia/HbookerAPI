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
	book_info := book.GET_BOOK_INFORMATION("")
	fmt.Println(book_info.Data.BookInfo.BookName)
	fmt.Println(book_info.Data.BookInfo.AuthorName)
	for index, div := range book.GET_DIVISION_LIST_BY_BOOKID(book_info.Data.BookInfo.BookID).Data.DivisionList {
		for _, chapter := range book.GET_CATALOGUE(div.DivisionID).Data.ChapterList {
			fmt.Println("第", index, "卷", chapter.ChapterTitle)
		}
	}
	fmt.Println("Done: Test_Book")

}
