package answers

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func NumberThree(c *gin.Context) {
	str := "Today is a happy day (just kidding) hahah"

	firstString := findFirstStringInBracket(str)

	firstStringV2 := findFirstStringInBracketV2(str)

	c.JSON(200, gin.H{"status": 200, "before_fix": firstString, "after_fix": firstStringV2})
}

//function from the question
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1]) //this is gonna crash the code if string contains "()"
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
}

//answer function (refactored and cleaner code, fixed missing last character on string and prevent crashes)
func findFirstStringInBracketV2(str string) string {
	if len(str) <= 0 {
		return ""
	}

	regex := regexp.MustCompile(`\((.*?)\)`)
	if !regex.MatchString(str) {
		return ""
	}
	strResult := regex.FindStringSubmatch(str)

	return strResult[1]
}
