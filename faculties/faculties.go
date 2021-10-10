package faculties

//Uncomment errors,and strings if you end up using them in the code below
import (
	"errors"
	//"log"
	CO "example/BookReader-Backend/config"
	"strings"
)

//Faculty struct
type Faculty struct {
	ID        int64  `json:"id,omitempty"`
	School    int64  `json:"school,omitempty"`  /*University name*/
	Faculty   string `json:"faculty,omitempty"` /*string type so this it the name of the faculty i.e Engineering*/
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Token string `json:"token,omitempty"`
}

//NewFaculty returns a pointer struct of a Faculty type
func NewFaculty() *Faculty {
	return new(Faculty)
}

/*
	Two Functions that will be used in the handler to fetch all the
	faculties of a school/university and also a single faculty of
	a specific school/university
*/

/*
	GetFaculties will accept a school/university ID as an input and return
	all the faculties of that school/universities via the university ID
	it also returns an error type nil if everything went well or a specific
	error if things did not go well
*/
func GetFaculties(s int64) ([]Faculty, error) {
	faculties := make([]Faculty, 0)
	database, err := CO.GetDB()

	if err != nil {
		err = errors.New("DB connection error")
		return faculties, err
	}

	defer database.Close()

	rows, err := database.Query("SELECT * FROM faculties WHERE school_id=?", s)

	if err != nil {
		return faculties, err
	}

	defer rows.Close()

	for rows.Next() {
		faculty := Faculty{}
		rows.Scan(&faculty.ID, &faculty.School, &faculty.Faculty, &faculty.CreatedAt, &faculty.UpdatedAt)
		faculties = append(faculties, faculty)
	}
	//log.Println("Faculties :", faculties)
	return faculties, nil
}

/*
	GetFaculty accepts the faculty ID as an input and returns that
	faculty data via the faculty struct and it also return an error
	type
*/
func GetFaculty(f int64) (Faculty, error) {
	faculty := Faculty{}
	database, err := CO.GetDB()

	if err != nil {
		err := errors.New("DB connection error")
		return faculty, err
	}

	defer database.Close()

	idQuery, err := database.Prepare("SELECT * FROM faculties WHERE id=?")

	if err != nil {
		return faculty, err
	}

	defer idQuery.Close()

	err = idQuery.QueryRow(f).Scan(&faculty.ID, &faculty.School, &faculty.Faculty, &faculty.CreatedAt, &faculty.UpdatedAt)

	if err != nil {
		return faculty, err
	}

	return faculty, nil
}


/*
  SaveFaculty adds a new faculty to
  an institution
*/
func (f *Faculty) SaveFaculty() error{
	db,err := CO.GetDB()

	if err != nil{
		err := errors.New("DB connection error")
		return err
	}

	defer db.Close()

	var school string
	schoolQuery, err := db.Prepare("SELECT school FROM schools WHERE id=?")

	if err != nil {
		return err
	}

	defer schoolQuery.Close()

	err = schoolQuery.QueryRow(f.School).Scan(&school)

	if err != nil{
		err = errors.New("The institution that you've selected does not exist.")
		return err
	}

	stmt,err := db.Prepare("INSERT INTO faculties (school_id,faculty) VALUES(?,?)")

	if err != nil{
		return err
	}
	
	defer stmt.Close()
	/*
		strings.Title makes the first letter of every word a 
		uppercase
	*/
	_,err = stmt.Exec(f.School,strings.Title(strings.ToLower(f.Faculty)))

	if err != nil{
		return err
	}

	return err
}