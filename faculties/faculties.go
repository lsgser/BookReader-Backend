package faculties

//Uncomment errors,and strings if you end up using them in the code below
import(
	//"errors"
	//CO "../config"
	//"strings"
)

//Faculty struct
type Faculty struct{
	ID int64 `json:"id,omitempty"`
	School int64 `json:"school,omitempty"`/*University name*/
	Faculty string `json:"faculty,omitempty"`/*string type so this it the name of the faculty i.e Engineering*/
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

//NewFaculty returns a pointer struct of a Faculty type
func NewFaculty() *Faculty{
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
func GetFaculties(s int64) ([]Faculty,error){
	faculties := make([]Faculty,0)
	//CODE has to be generated HERE
	return faculties,nil
}

/*
	GetFaculty accepts the faculty ID as an input and returns that
	faculty data via the faculty struct and it also return an error
	type
*/
func GetFaculty(f int64) (Faculty,error){
	faculty := Faculty{}
	//CODE has to be generated HERE
	return faculty,nil
}