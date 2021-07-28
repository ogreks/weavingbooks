package main

import (
	"books/src/models"
	"fmt"
)

func main() {
	var book = models.Book{
		Names: "test books",
		Author: "test",
		PublishTime: "2021-01-01 10:10:10",
		Price: 7000,
		Desc: "这本书好啊，简直牛逼6666",
	}

	fmt.Println(book)
}
