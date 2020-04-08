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
	"getBook",
	"GET",
	"/books/{id}",
	getBook,
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
func getBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	for _, book := range books{
		if book.Id == id{
			if err := json.NewEncoder(w).Encode(book); err != nil{
			log.Print("error getting requested book :: ", err)
			}
		}
	}
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