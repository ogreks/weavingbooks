package models

import (
	"strconv"
)

type Book struct {
	Names       string
	Desc        string
	Author      string
	PublishTime string
	Price       int64
}

func (books *Book) SetNames(names string) {
	books.Names = names
}

func (books Book) String() string {
	var desc = books.Desc
	if len(desc) > 12 {
		desc = desc[:15] + "..."
	}
	return "\t---------- "+ books.Names +" -----------\t\n" +
		"\t|    desc:" + desc + "\t\t|\n" +
		"\t|    author:" + books.Author + "\t\t|\n" +
		"\t|    times:" + books.PublishTime + "\t|\n" +
		"\t|    price:" + strconv.FormatInt(books.Price/100, 10) + "\t\t\t|\n" +
		"\t---------------------------------\n"
}
