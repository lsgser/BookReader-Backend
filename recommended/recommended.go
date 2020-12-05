package recommended;

import(
	CO "../config"
	"errors"
	"database/sql"
	"strings"
	//M "../modules"
	B "../books"
	E "../enrolled"
	"log"
)

type Recommended struct{
	ID int64 `json:"-"`
	Book int64 `json:"book"`
	Module int64 `json:"module"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type RecommendedModuleAndBook struct{
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
	Module string `json:"module"`
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

	defer db.Close()

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

	defer db.Close()

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

	defer db.Close()

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

	defer r.Close()

	err = r.QueryRow(book_id,module).Scan(&recommend_id)

	if err != nil{
		return recommend,err
	}

	recommend.ISBN = isbn
	recommend.Module = module

	return recommend,err
}
/*
	Get all the recommended books via the modules that the user enrolled for
*/
func GetRecommendedByUser(user string) ([]RecommendedModuleAndBook,error){
	log.Println(user)
	recommendedModulesAndBooks := make([]RecommendedModuleAndBook,0)
	recommended := make([]Recommended,0)
	enrollments := make([]E.Enrol,0)

	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return recommendedModulesAndBooks,err
	}

	defer db.Close()

	var user_id int64

	u,err := db.Prepare("SELECT id FROM users WHERE student_nr = ?")

	if err != nil{
		return recommendedModulesAndBooks,err
	}

	defer u.Close()

	err = u.QueryRow(user).Scan(&user_id)
	log.Println(user_id)
	if err != nil{
		return recommendedModulesAndBooks,err
	}

	enrollRows,err := db.Query("SELECT module_id,user_id FROM enrolled WHERE user_id = ?",user_id)

	if err != nil{
		return recommendedModulesAndBooks,err
	}

	defer enrollRows.Close()

	for enrollRows.Next(){
		enr := E.Enrol{}
		enrollRows.Scan(&enr.Module,&enr.User)
		enrollments = append(enrollments,enr)
	}

	/*
		Query the recommended books for a specific module based on 
		what the student enrolled for
	*/
	for _ , enrol := range enrollments{
		rec := Recommended{}
		recommend,err := db.Prepare("SELECT book_id,module_id FROM recommended WHERE module_id = ?")

		if err != nil{
			return recommendedModulesAndBooks,err
		}

		defer recommend.Close()
		
		err = recommend.QueryRow(enrol.Module).Scan(&rec.Book,&rec.Module)
		if err != nil{
			return recommendedModulesAndBooks,err
		}

		recommended = append(recommended,rec)
	}

	for _,book := range recommended{
		moduleAndBook := RecommendedModuleAndBook{}
		bookQuery,err := db.Prepare("SELECT * FROM books WHERE id = ?")

		if err != nil{
			return recommendedModulesAndBooks,err
		}

		defer bookQuery.Close()

		err = bookQuery.QueryRow(book.Book).Scan(&moduleAndBook.ID,&moduleAndBook.Title,&moduleAndBook.Author,&moduleAndBook.PublishDate,&moduleAndBook.ISBN,&moduleAndBook.CoverPage,&moduleAndBook.Description,&moduleAndBook.Book,&moduleAndBook.CreatedAt,&moduleAndBook.UpdatedAt)
		
		if err != nil{
			return recommendedModulesAndBooks,err
		}

		moduleAndBook.CoverPage = B.ImagePath()+strings.Split(moduleAndBook.CoverPage,"/")[4]
		moduleAndBook.Book = B.BookPath()+strings.Split(moduleAndBook.Book,"/")[3]
		/*
			Query a module name that will go hand in hand with the required books
			of the user
		*/
		moduleQuery,err := db.Prepare("SELECT module FROM modules WHERE id=?")
		
		if err != nil{
			return recommendedModulesAndBooks,err
		}

		defer moduleQuery.Close()

		err = moduleQuery.QueryRow(book.Module).Scan(&moduleAndBook.Module)

		recommendedModulesAndBooks = append(recommendedModulesAndBooks,moduleAndBook) 
	}

	return recommendedModulesAndBooks,err
} 

func (r *SaveRecommended) SaveRecommended() error{
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return err
	}

	defer db.Close()

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
		iRecommend,err := db.Prepare("INSERT INTO recommended (book_id,module_id) VALUES(?,?)")

		if err != nil{
			return err
		}

		defer iRecommend.Close()
		
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

	return nil
}