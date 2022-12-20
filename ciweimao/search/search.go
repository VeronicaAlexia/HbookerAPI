package search

import (
	"github.com/VeronicaAlexia/HbookerAPI/Template"
	"github.com/VeronicaAlexia/HbookerAPI/pkg/threading"
	"github.com/VeronicaAlexia/HbookerAPI/request"
	"strconv"
)

func GET_SEARCH(KeyWord string, page int) *Template.Search {
	var search Template.Search
	params := map[string]string{"count": "10", "page": strconv.Itoa(page), "category_index": "0", "key": KeyWord}
	request.Post(request.BOOKCITY_GET_FILTER_LIST).Params(params).NewRequests().Unmarshal(&search)
	return &search
}

func AllSearch(KeyWord string) []*Template.Search {
	var searchList []*Template.Search
	thread := threading.InitThreading(20)
	for i := 1; i <= 20; i++ {
		thread.Add()
		go func(i int) {
			defer thread.Done()
			result := GET_SEARCH(KeyWord, i)
			if result != nil {
				searchList = append(searchList, result)
			}
		}(i)
	}
	thread.Wait()
	return searchList
}
