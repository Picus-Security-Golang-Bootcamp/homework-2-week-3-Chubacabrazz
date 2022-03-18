package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Book struct {
	ID     int
	Name   string
	Page   int
	Stock  int
	Price  int
	SCode  int //Book Stock Code
	ISBN   int
	Author string
}

// initializing book datas.
var (
	book1 Book = Book{ID: 22, Name: "Lord Of The Rings: The Two Towers", Page: 450, Stock: 2, Price: 75, SCode: 25, ISBN: 3123213, Author: "J.R.R Tolkien"}
	book2 Book = Book{ID: 23, Name: "Harry Potter", Page: 250, Stock: 4, Price: 50, SCode: 25, ISBN: 3123213, Author: "J.K Rowling"}
	book3 Book = Book{ID: 24, Name: "Dune", Page: 700, Stock: 4, Price: 80, SCode: 25, ISBN: 3123213, Author: "Frank Herbert"}
	book4 Book = Book{ID: 25, Name: "Witcher", Page: 300, Stock: 16, Price: 45, SCode: 25, ISBN: 3123213, Author: "Andrzej Sapkowski"}
	book5 Book = Book{ID: 26, Name: "Lord Of The Rings: The Fellowship Of The Ring", Page: 420, Stock: 5, Price: 100, SCode: 25, ISBN: 3123213, Author: "J.R.R Tolkien"}

	Books = []Book{book1, book2, book3, book4, book5}
)
var (
	idptr   = flag.Int("get", 0, "ID of the Book")
	nameptr = flag.String("search", "", "Name of the Book")
	delete  = flag.Int("delete", 0, "Delete")
	buyptr  = flag.Int("buy", 0, "Buy <ID> <Quantity>")
	listptr = flag.String("list", "", "Lists all the books.")
)

// func List : lists all books in database. -list all
func List() {
	for _, v := range Books {
		fmt.Println(v.Name)
	}
}

// func Search : searchs books by name -search <bookname> (case insensitive)
func Search(nameptr string) {
	succ := 0 //succ for checking if we have the book, if it stays at 0 that means we don't have it.
	for _, v := range Books {
		nameptr = strings.ToLower(nameptr)
		nameptr = strings.Title(nameptr)
		if strings.Contains(v.Name, nameptr) {
			fmt.Printf("We have %s \n", v.Name)
			succ++
		}
	}
	if succ == 0 {
		fmt.Printf("Sorry, there is no such a book named: %s", nameptr)
	}
}

// func Delete : deletes books by ID -delete <ID>
func Delete(delete int) {
	for i, v := range Books {
		if delete == v.ID {
			v.Name = v.Name + " <<This book is deleted.>>"
			Books[i] = v
			fmt.Println("Book deleted")
		}
	}

}

// func Get : Prints books by ID -get <ID>
func Get(idptr int) {
	for _, v := range Books {
		if idptr == v.ID {
			fmt.Printf("Book Name: %s \nPage: %d\nStock: %d\nPrice: %d\nStockCode: %d\nISBN: %d\nAuthor: %s", v.Name, v.Page, v.Stock, v.Price, v.SCode, v.ISBN, v.Author)
		}
	}
}

// func Buy : Buys books by ID -buy <ID> <quantity>
func Buy(buyptr, ints int) {
	for _, v := range Books {
		if buyptr == v.ID {
			if v.Stock > ints {
				v.Stock = v.Stock - ints
				fmt.Printf("Success,thanks for shopping.\nWe have %d more in stock", v.Stock)
			} else {
				fmt.Printf("We don't have that much book in our stock, we have %d.", v.Stock)
				break
			}
		}
	}
}
func main() {

	flag.Parse()
	ints := make([]int, len(flag.Args())) //Converting flag.args string to int for second argument of -buy command(quantity).
	for i, s := range flag.Args() {
		ints[i], _ = strconv.Atoi(s)
	}

	switch {
	case *listptr == "all":
		List()
	case *idptr != 0:
		Get(*idptr)
	case *nameptr != "":
		Search(*nameptr)
	case *buyptr != 0:
		Buy(*buyptr, ints[0])
	case *delete != 0:
		Delete(*delete)
	default:
		fmt.Println("Error!! Enter -list all command for listing, -search <bookname> for searching books.")
	}

}
