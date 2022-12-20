package tag

import (
	"github.com/VeronicaAlexia/HbookerAPI/Template"
	"github.com/VeronicaAlexia/HbookerAPI/request"
	"strconv"
)

func GET_TAG_BOOK(page int) Template.Search {
	var search Template.Search
	params := map[string]string{
		"filter_word":    "",
		"count":          "10",
		"tags":           `[{"filter":"1","tag":""}]`,
		"use_daguan":     "0",
		"page":           strconv.Itoa(page),
		"is_paid":        "",
		"category_index": "0",
		"key":            "",
		"filter_uptime":  "",
		"up_status":      "",
		"order":          "uptime",
	}
	request.Post(request.BOOKCITY_GET_FILTER_LIST).Params(params).NewRequests().Unmarshal(&search)
	return search

}
