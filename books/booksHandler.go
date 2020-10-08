package books

import(
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	CO "../config"
	A "../auth"
	UP "../uploads"
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

/*
	Adds a new book
*/
func AddBook(w http.ResponseWriter,req *http.Request, _ httprouter.Params){
	CO.AddSafeHeaders(&w)
	book := NewBook()
	err := req.ParseMultipartForm(80 * UP.GetMB())

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
	}

	//Limit upload size to 3MB
	req.Body = http.MaxBytesReader(w,req.Body,80 * UP.GetMB())

	if A.CheckAdmin(req.FormValue("token")){
		if req.FormValue("title") != "" && req.FormValue("author") != "" && req.FormValue("isbn") != "" && req.FormValue("publish_date") != "" {
			bookFile,bookHandler,err := req.FormFile("book")

			if err != nil{
				w.WriteHeader(400)
				w.Write([]byte(`{"status":"`+err.Error()+`"}`))
				return
			}

			if bookFile == nil{
				w.WriteHeader(400)
				w.Write([]byte(`{"status":"Book was not included"}`))
				return
			}else{
				defer bookFile.Close()
			}

			coverPageFile,coverPageHandler,err := req.FormFile("cover_page")

			if err != nil{
				w.WriteHeader(400)
				w.Write([]byte(`{"status":"`+err.Error()+`"}`))
				return
			}

			if coverPageFile == nil{
				w.WriteHeader(400)
				w.Write([]byte(`{"status":"Book cover page was not included"}`))
				return
			}else{
				defer coverPageFile.Close()
			}

			book.Book,err = UP.BookFileUpload(bookFile,bookHandler,"/data/books/","book")
			
			if err != nil{
				w.WriteHeader(400)
				w.Write([]byte(`{"status":"`+err.Error()+`"}`))
				return
			}

			book.CoverPage,err = UP.ImageFileUpload(coverPageFile,coverPageHandler,"/data/images/book_covers/","cover_page")

			if err != nil{
				w.WriteHeader(400)
				w.Write([]byte(`{"status":"`+err.Error()+`"}`))
				return
			}

			book.Author = req.FormValue("author")
			book.Title = req.FormValue("title")
			book.PublishDate = req.FormValue("publish_date")
			book.ISBN = req.FormValue("isbn")
			book.Description = req.FormValue("description")

			err = book.SaveBook()

			if err != nil{
				w.WriteHeader(500)
				w.Write([]byte(`{"status":"`+err.Error()+`"}`))
				return
			}
			}else{
				w.WriteHeader(400)
				w.Write([]byte(`{"status":"Fill in all fields."}`))
				return	
			}		
	}else{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Invalid login session."}`))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(`{"status":"Success"}`))
}