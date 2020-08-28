package books

import(
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	CO "../config"
)

//Handler responsible for displaying a single book based on its ISBN
func ShowBook(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)
	isbn := params.ByName("b")

	book,err := GetBook(isbn)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))		
		return
	}

	err = json.NewEncoder(w).Encode(book)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}
}

//Handler responsible for displaying the books
func ShowBooks(w http.ResponseWriter , req *http.Request , _ httprouter.Params){
	CO.AddSafeHeaders(&w)
	books,err := GetBooks()

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(books)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}
}

//Handler for books searched by title or author or by the ISBN
func ShowBooksByQuery(w http.ResponseWriter,req *http.Request, params httprouter.Params){
	CO.AddSafeHeaders(&w)
	query := params.ByName("q")

	books,err := GetBooksByQuery(query)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status:`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(books)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return	
	}
}