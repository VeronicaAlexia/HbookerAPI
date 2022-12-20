package search

import (
	"github.com/VeronicaAlexia/HbookerAPI/Template"
	"github.com/VeronicaAlexia/HbookerAPI/request"
	"strconv"
)

func GET_SEARCH(KeyWord string, page int) Template.Search {
	var search Template.Search
	params := map[string]string{"count": "10", "page": strconv.Itoa(page), "category_index": "0", "key": KeyWord}
	request.NewHttpUtils(request.BOOKCITY_GET_FILTER_LIST, "POST").Params(params).NewRequests().Unmarshal(&search)
	return search

}
