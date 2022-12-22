package book

import (
	"fmt"
	"github.com/VeronicaAlexia/HbookerAPI/pkg/config"
)

func GetContent(chapterId string) string {
	chapterCmd := GET_KET_BY_CHAPTER_ID(chapterId)
	if chapterCmd.Code == "100000" {
		for i := 0; i < 5; i++ {
			content := GET_CHAPTER_CONTENT(chapterId, chapterCmd.Data.Command)
			if content.Code == "100000" {
				return content.Data.ChapterInfo.ChapterTitle + "  " + content.Data.ChapterInfo.Uptime + "\n\n" +
					string(config.Decode(content.Data.ChapterInfo.TxtContent, chapterCmd.Data.Command))
			} else {
				fmt.Println("GetContent Error: ", content.Code, content.Tip)
				fmt.Println("retry: ", i)
			}
		}
	} else {
		fmt.Println("chapterCmd Error: ", chapterCmd.Code, chapterCmd.Tip)
	}
	return ""
}
