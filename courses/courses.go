package courses

import(
	//"errors"
	CO "../config"
	//"strings"
)

type Course struct{
	ID int64 `json:"id,omitempty"`
	School int64 `json:"school"`
	Faculty int64 `json:"faculty"`
	Course string `json:"course"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"UpdatedAt,omitempty"`
}

/*
	This function returns a pointer Courses struct
*/
func NewCourse() *Course{
	return new(Course)
}

/*
	getCoursesBySchool() will accept a school/university parameter
	and use that to get all the courses of the university then 
	it will return all the courses of that school
*/
func getCoursesBySchool(school int64)([]Course,error){
	courses := make([]Course,0)
	/*
		CODE
	*/
	db,err := CO.GetDB()

	if err != nil{
		return courses,err
	}

	rows,err := db.Query("SELECT * FROM courses WHERE school_id =?",school)

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
	getCoursesByFaculty() will accept a faculty parameter
	and use that to get all the courses of the university then 
	it will return all the courses of that faculty
*/
func getCoursesByFaculty(faculty int64)([]Course,error){
	courses := make([]Course,0)
	
	/*
		CODE
	*/
	db,err := CO.GetDB()

	if err != nil{
		return courses,err
	}

	rows,err := db.Query("SELECT * FROM courses WHERE faculty_id =?",faculty)

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
	getCourse will return take in a course id parameter i.e param called c and return 
	the single course struct based on the course id
*/
func getCourse(c int64) (Course,error){
	course := Course{}
	/*
		CODE
	*/
	db,err := CO.GetDB()

	if err != nil{
		return course,err
	} 

	stmt,err := db.Prepare("SELECT * FROM courses WHERE id=?")

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