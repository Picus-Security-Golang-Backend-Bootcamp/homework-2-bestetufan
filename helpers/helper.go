package helpers

import (
	"fmt"
	"strings"

	"github.com/bestetufan/bookstore/interfaces"
	"github.com/bestetufan/bookstore/models"
)

/*
   Converts all items of a string slice to lowercase equivalent.
*/
func ToLowerSlice(slc []string) []string {
	loweredSlc := []string{}
	for i := 0; i < len(slc); i++ {
		loweredSlc = append(loweredSlc, strings.ToLower(slc[i]))
	}
	return loweredSlc
}

/*
   Searches and returns all the books satisfying given search query.
*/
func SearchStore(bookStore []*models.Book, args []string) []*models.Book {
	query := strings.Join(args, " ")
	booksFound := []*models.Book{}

	for _, book := range bookStore {
		name, sku, author :=
			strings.ToLower(book.Name),
			strings.ToLower(book.StockCode),
			strings.ToLower(book.Author.Name)

		if strings.Contains(name, query) ||
			strings.Contains(sku, query) ||
			strings.Contains(author, query) {
			booksFound = append(booksFound, book)
		}
	}
	return booksFound
}

/*
   Lists all active books (isDeleted: false) of given store slice.
*/
func ListBooks(bookStore []*models.Book) {
	for _, book := range bookStore {
		if !book.IsDeleted {
			book.PrintInfo()
		}
	}
}

/*
   Searches for the book related to given bookId. Returns nil if book
   is not found. Else, returns the reference of the book.
*/
func SearchBookById(bookStore []*models.Book, bookId int) *models.Book {
	for _, book := range bookStore {
		if book.ID == bookId {
			return book
		}
	}
	return nil
}

/*
   Searches for the book related to given bookId. If the book exists
   performs a buy transaction and decrement the related stock count.
*/
func BuyBookById(bookStore []*models.Book, bookId int, count int) {
	book := SearchBookById(bookStore, bookId)
	if book == nil {
		fmt.Println("Book could not be found by entered id!")
	} else {

		fmt.Print("Book information before buy operation: ")
		book.PrintInfo()

		result, error := book.Buy(count)

		if result {
			fmt.Print("Operation successful, current book information: ")
			book.PrintInfo()
		} else {
			fmt.Println(error)
		}
	}
}

/*
   Marks the books related to given bookId parameters as deleted.
*/
func DeleteBookById(bookStore []*models.Book, bookId int) {
	book := SearchBookById(bookStore, bookId)
	if book == nil {
		fmt.Println("Book could not be found by entered id!")
	} else {
		result, error := delete(book)

		if result {
			fmt.Println("Operation successful, current book list: ")
			ListBooks(bookStore)
		} else {
			fmt.Println(error)
		}
	}
}

// Helper for delete through interface method
func delete(deleteable interfaces.Deleteable) (bool, error) {
	return deleteable.Delete()
}
