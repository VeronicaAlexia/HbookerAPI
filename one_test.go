package main

import (
	"fmt"
	"github.com/VeronicaAlexia/HbookerAPI/HbookerAPI"
	"testing"
)

func Test_Book(t *testing.T) {
	HbookerAPI.HbookerKey.AppVersion = "2.9.290"
	HbookerAPI.HbookerKey.Account = ""
	HbookerAPI.HbookerKey.LoginToken = ""
	HbookerAPI.HbookerKey.DeviceToken = "ciweimao_"

	for index, div := range HbookerAPI.GET_DIVISION_LIST_BY_BOOKID("100337744").Data.DivisionList {
		for _, chapter := range HbookerAPI.GET_CATALOGUE(div.DivisionID).Data.ChapterList {
			fmt.Println("第", index, "卷", chapter.ChapterTitle)
		}
	}

}
