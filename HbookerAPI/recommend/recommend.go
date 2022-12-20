package recommend

import (
	"fmt"
	"github.com/VeronicaAlexia/HbookerAPI/Template"
	"github.com/VeronicaAlexia/HbookerAPI/request"
	"strings"
)

type RECOMMEND struct {
	book_list        []string
	recommend_list   [][]string
	book_list_string string
}

func NEW_RECOMMEND() *RECOMMEND {
	var recommend_list [][]string
	var recommend Template.Recommend2
	request.NewHttpUtils(request.BOOKCITY_RECOMMEND_DATA, "POST").
		Add("theme_type", "NORMAL").Add("tab_type", "200").NewRequests().Unmarshal(&recommend)
	if recommend.Code != "100000" {
		fmt.Println(recommend.Tip.(string))
	} else {
		for _, data := range recommend.Data.ModuleList {
			if data.ModuleType == "1" {
				for _, book := range data.BossModule.DesBookList {
					recommend_list = append(recommend_list, []string{book.BookName, book.BookID})
				}
			}
		}
	}
	return &RECOMMEND{recommend_list: recommend_list}

}

func (is *RECOMMEND) InitBookIdList() {
	is.book_list = nil
	for index, value := range is.recommend_list {
		fmt.Println("index:", index, "\t\tbook id:", value[1], "\t\tbook name:", value[0])
		is.book_list = append(is.book_list, value[1])
	}
	is.book_list_string = strings.Join(is.book_list, ",")
}

func (is *RECOMMEND) CHANGE_NEW_RECOMMEND() {
	var recommend Template.Recommend
	request.NewHttpUtils(request.GET_CHANGE_RECOMMEND, "POST").
		Add("book_id", is.book_list_string).Add("from_module_name", "长篇好书").NewRequests().Unmarshal(&recommend)
	is.recommend_list = nil
	if recommend.Code != "100000" {
		fmt.Println(recommend.Tip.(string))
	} else {
		for _, book := range recommend.Data.BookList {
			is.recommend_list = append(is.recommend_list, []string{book.BookName, book.BookID})
		}
	}
}
