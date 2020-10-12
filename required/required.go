package required

import(
	CO "../config"
	"errors"
)

type Required struct{
	ID int64 `json:"-"`
	Book int64 `json:"book"`
	Module int64 `json:"module"`
	User int64 `json:"-"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type SaveRequired struct{
	ID int64 `json:"-"`
	ISBN string `json:"isbn"`
	Module int64 `json:"module"`
	User string `json:"user"`
	Token string `json:"token,omitempty"`
}

func NewRequired() *Required{
	return new(Required)
}

func NewSaveRequired() *SaveRequired{
	return new(SaveRequired)
}

func GetRequiredByUser(user string) ([]Required,error){
	required := make([]Required,0)
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("Database connection error")
		return required,err
	}

	stmt,err := db.Prepare("SELECT id FROM users WHERE student_nr = ?")

	if err != nil{
		return required,err
	}

	defer stmt.Close()
	var user_id int64

	err = stmt.QueryRow(user).Scan(&user_id)

	if err != nil{
		return required,err
	}

	rows,err := db.Query("SELECT * FROM required WHERE user_id = ?",user_id)

	if err != nil{
		return required,err
	} 

	defer rows.Close()

	for rows.Next(){
		r := Required{}
		rows.Scan(&r.ID,&r.Book,&r.Module,&r.User,&r.CreatedAt,&r.UpdatedAt)
		required = append(required,r)
	}

	return required,nil
}

func GetRequiredByModule(module int64) ([]Required,error){
	required := make([]Required,0)
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("Database connection error")
		return required,err
	}

	rows,err := db.Query("SELECT * FROM required WHERE module_id = ?",module)

	if err != nil{
		return required,err
	} 

	defer rows.Close()

	for rows.Next(){
		r := Required{}
		rows.Scan(&r.ID,&r.Book,&r.Module,&r.User,&r.CreatedAt,&r.UpdatedAt)
		required = append(required,r)
	}

	return required,nil
}

func GetRequiredByBook(isbn string) ([]Required,error){
	required := make([]Required,0)
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("Database connection error")
		return required,err
	}

	stmt,err := db.Prepare("SELECT id FROM books WHERE isbn = ?")

	if err != nil{
		return required,err
	}

	defer stmt.Close()
	var book_id int64

	err = stmt.QueryRow(isbn).Scan(&book_id)

	if err != nil{
		return required,err
	}

	rows,err := db.Query("SELECT * FROM required WHERE book_id = ?",book_id)

	if err != nil{
		return required,err
	} 

	defer rows.Close()

	for rows.Next(){
		r := Required{}
		rows.Scan(&r.ID,&r.Book,&r.Module,&r.User,&r.CreatedAt,&r.UpdatedAt)
		required = append(required,r)
	}

	return required,nil
}

func (r *SaveRequired) SaveRequired() error{
	var (
		book_id int64
		user_id int64
		module_id int64
	)	

	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("Database connection error")
		return err
	}

	stmtBook,err := db.Prepare("SELECT id FROM books WHERE isbn = ?")

	if err != nil{
		return err
	}

	defer stmtBook.Close()

	err = stmtBook.QueryRow(r.ISBN).Scan(&book_id)

	if err != nil{
		return err
	}

	stmtUser,err := db.Prepare("SELECT id FROM users WHERE student_nr = ?")

	if err != nil{
		return err
	}

	defer stmtUser.Close()

	err = stmtUser.QueryRow(r.User).Scan(&user_id)

	if err != nil{
		return err
	}

	stmtModule,err := db.Prepare("SELECT id FROM modules WHERE id = ?")

	if err != nil{
		return err
	}

	defer stmtModule.Close()

	err = stmtModule.QueryRow(r.Module).Scan(&module_id)

	if err != nil{
		return err
	}

	stmtRequired,err := db.Prepare("INSERT INTO required (book_id,module_id,user_id) VALUES(?,?,?)")

	if err != nil{
		return err
	}

	_,err = stmtRequired.Exec(book_id,r.Module,user_id)

	if err != nil{
		return err
	}

	return nil
}