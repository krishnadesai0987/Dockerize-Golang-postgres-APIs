// GO packages
package main

// Import required libraries

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

const (
	host        = "10.105.168.35"
	port        = 5432
	DB_USER     = "krishna"
	DB_PASSWORD = "db@123"
	DB_NAME     = "library"
)

// DB set up
func init() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=disable", host, port, DB_USER, DB_PASSWORD, DB_NAME)
	db, err = sql.Open("postgres", dbinfo)

	checkerr(err)

}

type Book struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Book `json:"data"`
	Message string `json:"message"`
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkerr(err error) {
	if err != nil {
		panic(err)
	} else {

		fmt.Println("We are connected to Postgress database!!")
	}
}

//Get All Books
// response and request handlers
func Getbooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "we are connected to a browser\n")
	//fetch all movies from movies table
	rows, err := db.Query("SELECT * FROM books")

	//check errors
	checkerr(err)
	printMessage("fetching books ............")

	// prepare response
	var books []Book
	for rows.Next() {
		//book := Book{}
		var id int
		var title string
		var author string
		var description string

		err := rows.Scan(&id, &title, &author, &description)
		checkerr(err)

		books = append(books, Book{Id: id, Title: title, Author: author, Description: description})

	}

	var response = JsonResponse{Type: "success", Data: books}

	json.NewEncoder(w).Encode(response)
}

//Get a book
func Getbook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "we are connected to a browser\n")
	params := mux.Vars(r)
	var book_id = params["id"]

	printMessage("Getting book details from DB")
	row, err := db.Query("select * from books where id=$1", book_id)

	checkerr(err)
	defer row.Close()
	var book []Book
	if row.Next() {

		var id int
		var title string
		var author string
		var description string

		err := row.Scan(&id, &title, &author, &description)
		checkerr(err)

		book = append(book, Book{Id: id, Title: title, Author: author, Description: description})
	}
	var response = JsonResponse{Type: "success", Data: book}

	json.NewEncoder(w).Encode(response)

}

// Create a new book
func Createbook(w http.ResponseWriter, r *http.Request) { //done

	id := r.FormValue("id")
	title := r.FormValue("title")
	author := r.FormValue("author")
	description := r.FormValue("description")

	var response = JsonResponse{}

	if id == "" || title == "" {
		response = JsonResponse{Type: "error", Message: "You are missing bookID or bookName parameter."}
	} else {
		printMessage("Inserting book into DB")
		fmt.Println("Inserting new book details with book id"+id, "and title "+title)
		var lastInsertID int
		err := db.QueryRow("INSERT INTO books(id, title,author,description) VALUES($1,$2,$3,$4) returning id;", id, title, author, description).Scan(&lastInsertID)
		// check errors
		checkerr(err)

		response = JsonResponse{Type: "success", Message: "The book details has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Update a book
func Updatebook(w http.ResponseWriter, r *http.Request) { //done

	id := r.FormValue("id")
	title := r.FormValue("title")
	author := r.FormValue("author")
	description := r.FormValue("description")
	var response = JsonResponse{}

	if id == "" || title == "" {
		response = JsonResponse{Type: "error", Message: "You are missing bookID or bookName parameter."}
	} else {

		// create the update sql query
		//sqlStatement := "UPDATE books SET id=$1, title=$2, author=$3 description=$4 WHERE userid=$1"

		printMessage("Updating the data.....")
		_, err := db.Exec("UPDATE books SET id=$1, title=$2, author=$3, description=$4 WHERE id=$1;", id, title, author, description)
		// check errors
		checkerr(err)

		response = JsonResponse{Type: "success", Message: "The book details has been inserted successfully!"}

	}

	json.NewEncoder(w).Encode(response)

}

// Delete a book
func Deletebook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID := params["id"]

	var response = JsonResponse{}

	if bookID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing bookID parameter."}
	} else {

		printMessage("Deleting book from Database.....")

		_, err := db.Exec("DELETE FROM books where id=$1", bookID)

		// check errors
		checkerr(err)

		response = JsonResponse{Type: "success", Message: "The book has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)

}

//Delete All Books
func DeleteAllBooks(w http.ResponseWriter, r *http.Request) {

	printMessage("Deleting details of the all books.....")

	_, err = db.Exec("DELETE FROM books")
	checkerr(err)

	printMessage("All the books details have been deleted successfully..!")

	var response = JsonResponse{Type: "success", Message: "All books have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)

}

func main() {

	//Init router
	router := mux.NewRouter()

	//Endpoints
	router.HandleFunc("/api/books", Getbooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", Getbook).Methods("GET")
	router.HandleFunc("/api/books", Createbook).Methods("POST")
	router.HandleFunc("/api/books/{id}", Updatebook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", Deletebook).Methods("DELETE")
	router.HandleFunc("/api/books", DeleteAllBooks).Methods("DELETE")
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
