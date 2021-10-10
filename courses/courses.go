package courses

import(
	"errors"
	CO "example/BookReader-Backend/config"
	"strings"
	//"log"
)

type Course struct{
	ID int64 `json:"id,omitempty"`
	School int64 `json:"school"`
	Faculty int64 `json:"faculty"`
	Course string `json:"course"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Token string `json:"token,omitempty"`
}

/*
	This function returns a pointer Courses struct
*/
func NewCourse() *Course{
	return new(Course)
}

/*
	GetCoursesBySchool() will accept a school/university parameter
	and use that to get all the courses of the university then 
	it will return all the courses of that school
*/
func GetCoursesBySchool(school int64)([]Course,error){
	courses := make([]Course,0)
	/*
		CODE
	*/
	db,err := CO.GetDB()

	if err != nil{
		return courses,err
	}
	defer db.Close()

	rows,err := db.Query("SELECT id,school_id,faculty_id,course FROM courses WHERE school_id =?",school)

	if err != nil{
		return courses,err
	}

	defer rows.Close()

	for rows.Next(){
		c := Course{}
		rows.Scan(&c.ID,&c.School,&c.Faculty,&c.Course)
		courses = append(courses,c)
	}

	return courses,nil
}

/*
	GetCoursesByFaculty() will accept a faculty parameter
	and use that to get all the courses of the university then 
	it will return all the courses of that faculty
*/
func GetCoursesByFaculty(faculty int64)([]Course,error){
	courses := make([]Course,0)
	
	/*
		CODE
	*/
	db,err := CO.GetDB()

	if err != nil{
		return courses,err
	}
	defer db.Close()

	rows,err := db.Query("SELECT id,school_id,faculty_id,course FROM courses WHERE faculty_id =?",faculty)

	if err != nil{
		return courses,err
	}

	defer rows.Close()

	for rows.Next(){
		c := Course{}
		rows.Scan(&c.ID,&c.School,&c.Faculty,&c.Course)
		courses = append(courses,c)
	}
		
	return courses,nil
}

/*
	GetCourse will return take in a course id parameter i.e param called c and return 
	the single course struct based on the course id
*/
func GetCourse(c int64) (Course,error){
	course := Course{}
	/*
		CODE
	*/
	db,err := CO.GetDB()

	if err != nil{
		return course,err
	}

	defer db.Close()

	stmt,err := db.Prepare("SELECT id,school_id,faculty_id,course FROM courses WHERE id=?")

	if err != nil{
		return course,err
	}

	defer stmt.Close()

	err = stmt.QueryRow(c).Scan(&course.ID,&course.School,&course.Faculty,&course.Course)

	if err != nil{
		return course,err
	}
	return course,nil
}

func (c *Course) SaveCourse() error{
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return err
	}

	defer db.Close()
	
	var faculty string

	facultyQuery,err := db.Prepare("SELECT faculty FROM faculties WHERE id = ? AND school_id = ?")

	if err != nil{
		return err
	}

	defer facultyQuery.Close()

	err = facultyQuery.QueryRow(c.Faculty,c.School).Scan(&faculty)

	if err != nil{
		err = errors.New("Instituion or Faculty does not exist")
		return err
	}
	
	stmt,err := db.Prepare("INSERT INTO courses (school_id,faculty_id,course) VALUES(?,?,?)")

	if err != nil{
		return err
	}

	defer stmt.Close()

	_,err = stmt.Exec(c.School,c.Faculty,strings.Title(strings.ToLower(c.Course)))

	if err != nil{
		return err
	}

	return err
}