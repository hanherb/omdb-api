package controllerRest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UnitTest(c *gin.Context) {
	searchWord := "Lion King"
	pagination := "1"

	fmt.Println("Test 1 : ")
	fmt.Println(RunSearch(c, &searchWord, &pagination))

	searchWord = "Lion King"
	pagination = "100"

	fmt.Println("Test 2 : ")
	fmt.Println(RunSearch(c, &searchWord, &pagination))

	searchWord = "LionKing"
	pagination = "1"

	fmt.Println("Test 3 : ")
	fmt.Println(RunSearch(c, &searchWord, &pagination))

	searchWord = "liOn kiNG"
	pagination = "1"

	fmt.Println("Test 4 : ")
	fmt.Println(RunSearch(c, &searchWord, &pagination))

	c.JSON(200, gin.H{"status": 200, "message": "Unit test completed"})
}
