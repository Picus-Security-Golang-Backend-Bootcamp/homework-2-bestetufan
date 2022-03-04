package models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Book struct {
	ID         int
	Name       string
	StockCode  string
	ISBN       int
	PageCount  int
	Price      float64
	StockCount int
	IsDeleted  bool
	Author     Author
}

// Constructor
func NewBook(id int, name string, isbn int, pageCount int, price float64,
	stockCount int, author Author) *Book {
	book := &Book{
		ID:         id,
		Name:       name,
		ISBN:       isbn,
		PageCount:  pageCount,
		StockCount: stockCount,
		Price:      price,
		IsDeleted:  false,
		Author:     author,
	}

	// Auto generate stock code in (ID-Name-ISBN) format
	book.StockCode = fmt.Sprintf("%d-%s-%d", book.ID, toAcronym(book.Name), book.ISBN)
	return book
}

// Prints book information
func (b *Book) PrintInfo() {
	fmt.Printf("Book [ID: %d => Name: %s, Author: %s, Pages: %d, Stock Count: %d, ISBN: %d, Stock Code: %s]\n",
		b.ID, b.Name, b.Author.GetFullName(), b.PageCount, b.StockCount, b.ISBN, b.StockCode)
}

// Buys book and decrements the related stock count
func (b *Book) Buy(buyCount int) (bool, error) {
	if buyCount <= 0 {
		return false, errors.New("Invalid buy count")
	}

	if b.StockCount < buyCount {
		return false, errors.New("Insufficient stock amount")
	}

	b.StockCount -= buyCount
	return true, nil
}

// Marks a book deleted
func (b *Book) Delete() (bool, error) {
	if b.IsDeleted {
		return false, errors.New("Book is already deleted")
	}

	b.IsDeleted = true
	return true, nil
}

// Utils
/*
	Converts given string to acronym (Eg. Lord of the Rings: LOTR)
*/
func toAcronym(source string) string {
	reg, _ := regexp.Compile("\\B.|\\P{L}")
	abbr := reg.ReplaceAllString(source, "")
	return strings.ToUpper(abbr)
}
