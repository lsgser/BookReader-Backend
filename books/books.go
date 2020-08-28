package books

import(
	CO "../config"
	"errors"
	"strings"
)

type Book struct{
	ID int64 `json:"-"`
	Title string `json:"title"`
	Author string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN string `json:"isbn"`
	CoverPage string `json:"cover_page"`
	Description string `json:"description,omitempty"`
	Book string `json:"book"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

/*
	NewBook() function returns a Book pointer struct
*/
func NewBook() *Book{
	return new(Book)
}

/*
	GetBook():Gets the book isbn number and returns the book data as
	a struct and also an error type.

	GetBook() accepts a isbn number as an input parameter
*/
func GetBook(isbn string) (Book,error){
	book := Book{}

	isbn = strings.TrimSpace(isbn)

	if isbn == ""{
		err := errors.New("ISBN not provided")
		return book,err
	}

	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return book,err
	}


	stmt,err := db.Prepare("SELECT * FROM books WHERE isbn = ?")

	if err != nil {
		return book,err
	}

	defer stmt.Close()

	err = stmt.QueryRow(isbn).Scan(&book.ID,&book.Title,&book.Author,&book.PublishDate,&book.ISBN,&book.CoverPage,&book.Description,&book.Book,&book.CreatedAt,&book.UpdatedAt)

	if err != nil{
		return book,err
	}

	return book,nil
}


/*
	GetBooks(): Gets all the book from the database and returns
	the book struct slice data and also the error
*/
func GetBooks() ([]Book,error){
	books := make([]Book,0)

	db,err := CO.GetDB()
	
	if err != nil{
		err = errors.New("DB connection error")
		return books,err
	}

	rows,err := db.Query("SELECT * FROM books")

	if err != nil{
		return books,err
	}

	defer rows.Close()

	for rows.Next(){
		book := Book{}
		rows.Scan(&book.ID,&book.Title,&book.Author,&book.PublishDate,&book.ISBN,&book.CoverPage,&book.Description,&book.Book,&book.CreatedAt,&book.UpdatedAt)
		books = append(books,book)
	}

	return books,nil
}

/*
	GetBooksByInfo(): Gets books/book by the title or by Author or by ISBN and returns
	a slice struct Book type and also an error type
	Input:string variable of a query
*/
func GetBooksByQuery(query string) ([]Book,error){
	books := make([]Book,0)

	query = "%"+strings.TrimSpace(query)+"%"

	if query == ""{
		err := errors.New("Empty query")
		return books,err
	}

	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return books,err
	}

	rows,err := db.Query("SELECT * FROM books WHERE title LIKE ? OR author LIKE ? OR isbn LIKE ?",query,query,query)

	if err != nil {
		return books,err
	}

	defer rows.Close()

	for rows.Next(){
		book := Book{}
		rows.Scan(&book.ID,&book.Title,&book.Author,&book.PublishDate,&book.ISBN,&book.CoverPage,&book.Description,&book.Book,&book.CreatedAt,&book.UpdatedAt)
		books = append(books,book)
	}

	return books,nil
}