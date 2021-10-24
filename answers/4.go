package answers

import (
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

func NumberFour(c *gin.Context) {
	arrStr := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	arrResult := groupByAnagram(arrStr)

	c.JSON(200, gin.H{"status": 200, "data": arrResult})
}

func groupByAnagram(arrStr []string) [][]string {
	var group [][]string //2D array

	for i := 0; i < len(arrStr); i++ { //words loop
		matched := false //flagging if the word anagram has/hasn't matched already

		if len(group) == 0 {

			group = append(group, []string{arrStr[i]}) //add first group

		} else {

			for j := 0; j < len(group); j++ { //group loop

				for k := 0; k < len(group[j]); k++ { //words in the group loop

					if anagram(arrStr[i], group[j][k]) { //if anagram match

						group[j] = append(group[j], arrStr[i]) //assign word to the group
						matched = true
						break

					}

					if len(group[j])-1 == k { //if it is the last word in the group

						if len(group)-1 == j && !matched { //if it is the last group and the word has not match in any group

							group = append(group, []string{arrStr[i]}) //assign word to new group
							j++                                        //equalize index with new group

						}

					}
				}
			}
		}
	}

	return group
}

func anagram(s1, s2 string) bool {
	return sortAlphabets(s1) == sortAlphabets(s2)
}

func sortAlphabets(str string) string {
	splited := strings.Split(str, "")
	sort.Strings(splited)
	return strings.Join(splited, "")
}
