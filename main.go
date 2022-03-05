package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/bestetufan/bookstore/helpers"
	"github.com/bestetufan/bookstore/models"
)

// Global book struct slice
var bookStore []*models.Book

func init() {
	// Seeding the random generator
	rand.Seed(time.Now().UnixNano())

	// Test book names & authors (pointer) slice
	var authors = []*models.Author{
		{Name: "Thomas", Surname: "Keneally", Age: 86},
		{Name: "Stephen", Surname: "King", Age: 55},
		{Name: "Ece", Surname: "Temelkuran", Age: 49},
		{Name: "Emrah", Surname: "Serbes", Age: 41},
		{Name: "J.R.R.", Surname: "Tolkien", Age: 72},
	}
	var bookNames = []string{"Schindler's List", "The Shining", "Devir", "Deliduman", "Lord of the Rings"}

	// Populate global book struct slice using local book names & authors
	for i, bookName := range bookNames {
		bookStore = append(bookStore, models.NewBook(
			i+1,                        // ID
			bookName,                   // Name
			rand.Intn(4000)+1000,       // ISBN: 1000-5000 range
			rand.Intn(300)+200,         // PageCount: 200 - 500 range
			20.0+rand.Float64()*(60.0), // Price: 20 - 80 range
			rand.Intn(5),               // StockCount: 0 - 5 range
			*authors[i],                // Author {name, surname, age}
		))
	}
}

func main() {
	args := os.Args
	lowerCaseArgs := helpers.ToLowerSlice(args)

	// Display welcome message in case of no command sent
	if len(lowerCaseArgs) == 1 {
		fmt.Println("Command List")
		fmt.Println("-----------------")
		fmt.Println("Search Operation: \"search {keyword}\" \n",
			"List Operation: \"list\" \n",
			"Buy Operation: \"buy {bookId, count}\" \n",
			"Delete Operation: \"delete {bookId}\"")
		fmt.Println("-----------------")
		return
	}

	// Command logic
	switch lowerCaseArgs[1] {
	case "search":
		if len(lowerCaseArgs) < 3 {
			fmt.Println("Enter a book name to search!")
		} else {
			if booksFound := helpers.SearchStore(bookStore, lowerCaseArgs[2:]); len(booksFound) > 0 {
				helpers.ListBooks(booksFound)
			}
		}
	case "list":
		fmt.Println("List of books:")
		helpers.ListBooks(bookStore)
	case "buy":
		if len(lowerCaseArgs) != 4 {
			fmt.Println("Enter a book id and amount!")
		} else {
			// Convert and check parameters for type int
			bookId, error := strconv.Atoi(lowerCaseArgs[2])
			count, error := strconv.Atoi(lowerCaseArgs[3])

			if error != nil {
				fmt.Println("Parameters must be in correct type!")
				return
			}

			// Perform transaction
			helpers.BuyBookById(bookStore, bookId, count)
		}
	case "delete":
		if len(lowerCaseArgs) != 3 {
			fmt.Println("Enter a book id to delete!")
		} else {
			// Convert and check parameters for type int
			bookId, error := strconv.Atoi(lowerCaseArgs[2])

			if error != nil {
				fmt.Println("Parameters must be in correct type!")
				return
			}

			// Perform transaction
			helpers.DeleteBookById(bookStore, bookId)
		}
	default:
		fmt.Println("Unknown command!")
	}
}
