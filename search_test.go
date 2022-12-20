package main

import (
	"fmt"
	"github.com/VeronicaAlexia/HbookerAPI/ciweimao/search"
	"github.com/VeronicaAlexia/HbookerAPI/pkg/config"
	"testing"
)

func TestSearch(t *testing.T) {
	config.AppConfig.AppVersion = "2.9.290"
	config.AppConfig.Account = ""
	config.AppConfig.LoginToken = ""
	config.AppConfig.DeviceToken = "ciweimao_"
	for _, book := range search.AllSearch("刀剑") {
		for index, chapter := range book.Data.BookList {
			fmt.Println("index:", index, "BookName:", chapter.BookName, chapter.BookID)
		}
	}
	fmt.Println("Done: TestSearch")
}
