package modules

import(
	//CO "../config"
	//errors
)

type Module struct{
	ID int64 `json:"id"`
	School int64 `json:"school"`
	Faculty int64 `json:"faculty"`
	Course int64 `json:"course"`
	Module string `json:"module"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`	
}

func NewModule() *Module{
	return new(Module)
}

/*
	GetModulesBySchool() will accept an int64 input parameter
	this parameter input will represent the school id, with this
	school id we'll return all the modules(i.e a slice of all the modules) that the school/university has
	and we'll also return an error type  
*/
func GetModulesBySchool(school int64) ([]Module,error){
	modules := make([]Module,0)
	/*
		Code
	*/
	return modules,nil
}

/*
	GetModulesByFaculty() will accept an int64 input parameter
	this parameter input will represent the faculty id, with this
	faculty id we'll return all the modules(i.e a slice of all the modules) that the faculty has
	and we'll also return an error type  
*/
func GetModulesByFaculty(faculty int64) ([]Module,error){
	modules := make([]Module,0)
	/*
		Code
	*/
	return modules,nil
}

/*
	GetModulesByFaculty() will accept an int64 input parameter
	this parameter input will represent the course id, with this
	course id we'll return all the modules(i.e a slice of all the modules) that the course has
	and we'll also return an error type

	e.g a course might be maths,maths consists of calculus,algebra,etc as its modules  
*/
func GetModulesByCourse(course int64) ([]Module,error){
	modules := make([]Module,0)
	/*
		Code
	*/
	return modules,nil
}

/*
	Input: module id
	Output: returns the Module data from the database based on the module id
*/
func GetModule(module int64) (Module,error){
	m := Module{}
	/*
		CODE
	*/
	return m,nil
}