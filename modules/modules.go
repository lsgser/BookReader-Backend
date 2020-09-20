package modules

import (
	CO "../config"
	"strings"
	"errors"
)

//Module Struct
type Module struct {
	ID        int64  `json:"id"`
	School    int64  `json:"school"`
	Faculty   int64  `json:"faculty"`
	Course    int64  `json:"course"`
	Module    string `json:"module"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Token 	  string  `json:"token,omitempty"`
}

//NewModule()
func NewModule() *Module {
	return new(Module)
}

/*
	GetModulesBySchool() will accept an int64 input parameter
	this parameter input will represent the school id, with this
	school id we'll return all the modules(i.e a slice of all the modules) that the school/university has
	and we'll also return an error type
*/
func GetModulesBySchool(school int64) ([]Module, error) {
	modules := make([]Module, 0)

	database, err := CO.GetDB()

	if err != nil {
		return modules, err
	}

	rows, err := database.Query("SELECT * FROM modules WHERE school_id =? ", school)

	if err != nil {
		return modules, err
	}

	defer rows.Close()

	for rows.Next() {
		module := Module{}
		rows.Scan(&module.ID, &module.School, &module.Faculty, &module.Course, &module.Module, &module.CreatedAt, &module.UpdatedAt)
		modules = append(modules, module)
	}

	return modules, nil
}

/*
	GetModulesByFaculty() will accept an int64 input parameter
	this parameter input will represent the faculty id, with this
	faculty id we'll return all the modules(i.e a slice of all the modules) that the faculty has
	and we'll also return an error type
*/
func GetModulesByFaculty(faculty int64) ([]Module, error) {
	modules := make([]Module, 0)

	database, err := CO.GetDB()

	if err != nil {
		return modules, err
	}

	rows, err := database.Query("SELECT * FROM modules WHERE faculty_id =?", faculty)

	if err != nil {
		return modules, err
	}
	defer rows.Close()

	for rows.Next() {
		module := Module{}
		rows.Scan(&module.ID, &module.School, &module.Faculty, &module.Course, &module.Module, &module.CreatedAt, &module.UpdatedAt)
		modules = append(modules, module)
	}

	return modules, nil
}

/*
	GetModulesByFaculty() will accept an int64 input parameter
	this parameter input will represent the course id, with this
	course id we'll return all the modules(i.e a slice of all the modules) that the course has
	and we'll also return an error type

	e.g a course might be maths,maths consists of calculus,algebra,etc as its modules
*/
func GetModulesByCourse(course int64) ([]Module, error) {
	modules := make([]Module, 0)

	database, err := CO.GetDB()

	if err != nil {
		return modules, err
	}

	rows, err := database.Query("SELECT * FROM modules WHERE course_id =?", course)

	if err != nil {
		return modules, err
	}
	defer rows.Close()

	for rows.Next() {
		module := Module{}
		rows.Scan(&module.ID, &module.School, &module.Faculty, &module.Course, &module.Module, &module.CreatedAt, &module.UpdatedAt)
		modules = append(modules, module)
	}

	return modules, nil
}

/*
	Input: module id
	Output: returns the Module data from the database based on the module id
*/
func GetModule(m int64) (Module, error) {
	module := Module{}

	database, err := CO.GetDB()

	if err != nil {
		return module, err
	}

	statement, err := database.Prepare("SELECT * FROM modules WHERE id =?")

	if err != nil {
		return module, err
	}
	defer statement.Close()

	err = statement.QueryRow(m).Scan(&module.ID, &module.School, &module.Faculty, &module.Course, &module.Module, &module.CreatedAt, &module.UpdatedAt)

	if err != nil {
		return module, err
	}

	return module, nil
}


/*
	Save a new module to the database
*/
func (m *Module) SaveModule() error{
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return err
	}

	var course string

	courseQuery,err := db.Prepare("SELECT course FROM courses WHERE school_id = ? AND faculty_id = ? AND id = ?")

	if err != nil{
		return err
	}

	defer courseQuery.Close()

	err = courseQuery.QueryRow(m.School,m.Faculty,m.Course).Scan(&course)

	if err != nil{
		err = errors.New("Institution / Faculty / Course does not exist")
		return err
	}

	stmt,err := db.Prepare("INSERT INTO modules (school_id,faculty_id,course_id,module) VALUES(?,?,?,?)")

	if err != nil{
		return err
	}

	_,err = stmt.Exec(m.School,m.Faculty,m.Course,strings.Title(strings.ToLower(m.Module)))

	return err
}