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

func NewRequired() *Required{
	return new(Required)
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