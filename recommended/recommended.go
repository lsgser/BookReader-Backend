package recommended;

import(
	CO "../config"
	"errors"
	"database/sql"
)

type Recommended struct{
	ID int64 `json:"-"`
	Book int64 `json:"book"`
	Module int64 `json:"module"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type SaveRecommended struct{
	ID int64 `json:"-"`
	ISBN string `json:"book"`
	Module int64 `json:"module"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Token string `json:"token,omitempty"`
}

type RecommendedISBN struct{
	ID int64 `json:"-"`
	ISBN string `json:"book"`
	Module int64 `json:"module"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func NewRecommended() *Recommended{
	return new(Recommended)
}

func NewSaveRecommended() *SaveRecommended{
	return new(SaveRecommended)
} 

func NewRecommendedISBN() *RecommendedISBN{
	return new(RecommendedISBN)
}

func GetRecommendedByBook(isbn string) ([]RecommendedISBN,error){
	recommended := make([]RecommendedISBN,0)
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return recommended,err
	}

	var book_id int64

	book, err := db.Prepare("SELECT id FROM books WHERE isbn=?")

	if err != nil {
		return recommended, err
	}

	defer book.Close()

	err = book.QueryRow(isbn).Scan(&book_id)

	if err != nil{
		return recommended,err
	}

	rows,err := db.Query("SELECT module_id FROM recommended WHERE book_id = ?",book_id)

	if err != nil{
		return recommended,err
	}

	defer rows.Close()

	for rows.Next(){
		recommend := RecommendedISBN{}
		rows.Scan(&recommend.Module)
		recommend.ISBN = isbn
		recommended = append(recommended,recommend)
	}

	return recommended,err
}

func GetRecommendedByModule(module int64) (RecommendedISBN,error){
	recommend := RecommendedISBN{}

	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return recommend,err
	}

	var (
		book_id int64
	)

	m, err := db.Prepare("SELECT id FROM modules WHERE id=?")

	if err != nil {
		return recommend, err
	}

	defer m.Close()

	err = m.QueryRow(module).Scan(&recommend.Module)

	if err != nil{
		return recommend,err
	}

	b,err := db.Prepare("SELECT book_id FROM recommended WHERE module_id = ?")

	if err != nil{
		return recommend,err
	}

	defer b.Close()

	err = b.QueryRow(module).Scan(&book_id)

	book, err := db.Prepare("SELECT isbn FROM books WHERE id=?")

	if err != nil {
		return recommend, err
	}

	defer book.Close()

	err = book.QueryRow(book_id).Scan(&recommend.ISBN)
	
	if err != nil{
		return recommend,err
	}

	return recommend,err
}

func GetRecommendedByModuleAndBook(isbn string,module int64) (RecommendedISBN,error){
	recommend := RecommendedISBN{}
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return recommend,err
	}

	var (
		book_id int64
		recommend_id int64
	)

	b,err := db.Prepare("SELECT id FROM books WHERE isbn = ?")

	if err != nil{
		return recommend,err
	}

	defer b.Close()
	err = b.QueryRow(isbn).Scan(&book_id)

	if err != nil{
		return recommend,err
	}

	r, err := db.Prepare("SELECT id FROM recommended WHERE book_id = ? AND module_id = ?")

	if err != nil{
		return recommend,err
	}

	err = r.QueryRow(book_id,module).Scan(&recommend_id)

	if err != nil{
		return recommend,err
	}

	recommend.ISBN = isbn
	recommend.Module = module

	return recommend,err
}

func (r *SaveRecommended) SaveRecommended() error{
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return err
	}
	var (
		book_id int64
		recommend_id int64
		module_id int64
	)

	b,err := db.Prepare("SELECT id FROM books WHERE isbn = ?")

	if err != nil{
		return err
	}

	defer b.Close()
	err = b.QueryRow(r.ISBN).Scan(&book_id)

	if err != nil{
		return err
	}

	/*Check if the module exists*/
	m,err := db.Prepare("SELECT id FROM modules WHERE id = ?")

	if err != nil{
		return err
	}

	defer m.Close()
	err = m.QueryRow(r.Module).Scan(&module_id)

	if err != nil{
		return err
	}

	recommend,err := db.Prepare("SELECT id FROM recommended WHERE book_id = ? AND module_id = ?")
	
	if err != nil{
		return err
	}

	defer recommend.Close()

	err = recommend.QueryRow(book_id,r.Module).Scan(&recommend_id)

	if err == sql.ErrNoRows {
		iRecommend,err := db.Prepare("INSERT (book_id,module_id) INTO recommended VALUES(?,?)")

		if err != nil{
			return err
		}

		_,err = iRecommend.Exec(book_id,r.Module)

		if err != nil{
			return err
		}
	}else if err != nil{
		return err
	}else{
		err = errors.New("Recommendation for this module already exists.")
		return err
	}

	return err
}