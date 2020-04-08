package main
import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
const(
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)
type Route struct{
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}
type Routes []Route
var routes = Routes{
	Route{
		"getBook",
		"GET",
		"/books",
		getBooks,
	},
	Route{
		"addBook",
		"POST",
		"/book/add",
		addBook,
	},
	Route{
		"deleteBook",
		"DELETE",
		"/book/delete",
		deleteBook,
	},
}
type Book struct{
	Id string `json:"id"`
	Name string `json:"bookname"`
}
type Books []Book
var books []Book
func init(){
	books = Books{
		Book{Id: "1", Name: "To kill the mocking bird"},
		Book{Id: "2", Name: "English for kids"},
	}
}
func getBooks(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request){
	book := Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil{
		log.Print("error occurred while decoding book, data :: ", err)
		return
	}
	log.Printf("deleting book id :: %s with Name as :: %s ", book.Id, book.Name)
	index := GetIndex(book.Id)
	books = append(books[:index], books[index+1:]...)
	json.NewEncoder(w).Encode(books)
}
func GetIndex(id string) int{
	for i := 0; i < len(books); i++{
		if books[i].Id == id{
			return i
		}
	}
	return -1
}
func addBook(w http.ResponseWriter, r *http.Request){
	book := Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil{
		log.Print("error occurred while decoding book,data :: ", err)
		return
	}
	log.Printf("adding book id :: %s with bookname,as :: %s", book.Id, book.Name)
	books = append(books, Book{Id: book.Id, Name: book.Name})
	json.NewEncoder(w).Encode(books)
}
func AddRoutes(router *mux.Router) *mux.Router{
	for _, route := range routes{
	router.
	Methods(route.Method).
	Path(route.Pattern).
	Name(route.Name).
	Handler(route.HandlerFunc)
	}
	return router
}
func main(){
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil{
	log.Fatal("error starting http server :: ", err)
	return
	}
}