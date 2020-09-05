package enrolled

import(
	CO "../config"
	M "../modules"
	U "../users"
	//"strings"
	"errors"
)

type Enrol struct{
	ID int64 `json:"-"`
	Module int64 `json:"module"`
	User int64 `json:"user"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func NewEnrol() *Enrol{
	return new(Enrol)
}

func GetEnrolledByModule(module int64) ([]Enrol,error){
	enrolled := make([]Enrol,0)
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return enrolled,err
	}

	rows,err := db.Query("SELECT * FROM enrolled WHERE module_id = ?",module)

	if err != nil{
		return enrolled,err
	}

	defer rows.Close()

	for rows.Next(){
		enrol := Enrol{}
		rows.Scan(&enrol.ID,&enrol.Module,&enrol.User,&enrol.CreatedAt,&enrol.UpdatedAt)
		enrolled = append(enrolled,enrol)
	}

	return enrolled,nil
}

func GetEnrolledByUser(user string) ([]Enrol,error){
	enrolled := make([]Enrol,0)
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return enrolled,err
	}
	
	stmt,err := db.Prepare("SELECT id FROM users WHERE student_nr=?")

	if err != nil{
		return enrolled,err
	}

	defer stmt.Close()
	var user_id int64

	err = stmt.QueryRow(user).Scan(&user_id)
	if err != nil{
		return enrolled,err
	}

	rows,err := db.Query("SELECT * FROM enrolled WHERE user_id = ?",user_id)

	if err != nil{
		return enrolled,err
	}

	defer rows.Close()

	for rows.Next(){
		enrol := Enrol{}
		rows.Scan(&enrol.ID,&enrol.Module,&enrol.User,&enrol.CreatedAt,&enrol.UpdatedAt)
		enrolled = append(enrolled,enrol)
	}

	return enrolled,nil
}

/*
	This function will get the collection
	of all modules and return the module data
	through the module struct
*/
func GetEnrolledModules(user string) ([]M.Module,error){
	modules := make([]M.Module,0)

	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return modules,err			
	}
	stmt,err := db.Prepare("SELECT id FROM users WHERE student_nr=?")

	if err != nil{
		return modules,err
	}

	defer stmt.Close()
	var user_id int64
	err = stmt.QueryRow(user).Scan(&user_id)
	
	if err != nil{
		return modules,err
	}

	rows,err := db.Query("SELECT * FROM enrolled WHERE user_id = ?",user_id)

	if err != nil{
		return modules,err
	}

	defer rows.Close()

	for rows.Next(){
		enrol := Enrol{}
		module := M.Module{}
		rows.Scan(&enrol.ID,&enrol.Module,&enrol.User,&enrol.CreatedAt,&enrol.UpdatedAt)
		stmt,err := db.Prepare("SELECT * FROM modules WHERE id = ?")

		if err != nil{
			return modules,err
		}

		defer stmt.Close()

		err = stmt.QueryRow(enrol.Module).Scan(&module.ID,&module.School,&module.Faculty,&module.Course,&module.Module,&module.CreatedAt,&module.UpdatedAt)

		if err != nil{
			return modules,err
		}

		modules = append(modules,module)
	}

	return modules,nil
}

/*
	Returns a list of students that are enrolled for
	a specific module via the module id that is
	provided
*/
func GetEnrolledUsers(module int64) ([]U.User,error){
	users := make([]U.User,0)

	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return users,err
	}

	rows,err := db.Query("SELECT * FROM enrolled WHERE module_id = ?",module)

	if err != nil{
		return users,err
	}

	defer rows.Close()

	for rows.Next(){
		enrol := Enrol{}
		user := U.User{}
		rows.Scan(&enrol.ID,&enrol.Module,&enrol.User,&enrol.CreatedAt,&enrol.UpdatedAt)
		stmt,err := db.Prepare("SELECT * FROM users WHERE id = ?")

		if err != nil{
			return users,err
		}

		defer stmt.Close()

		err = stmt.QueryRow(enrol.User).Scan(&user.ID,&user.School,&user.Faculty,&user.Course,&user.Student,&user.Name,&user.Surname,&user.Email,&user.Picture,&user.Password,&user.CreatedAt,&user.UpdatedAt)

		if err != nil{
			return users,err
		}

		users = append(users,user)
	}

	return users,nil	
}