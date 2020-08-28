package routes

import (
	F "../faculties"
	C "../courses"
	M "../modules"
	B "../books"
	U "../users"
	"github.com/julienschmidt/httprouter"
)

//NewRouter : new router returns all router
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	//Faculties
	router.GET("/faculties/:s", F.ShowFaculties)
	router.GET("/faculty/:f", F.ShowFaculty)

	//Courses
	router.GET("/courses/:f",C.ShowCoursesByFaculty)
	router.GET("/courses/:s",C.ShowCoursesBySchool)
	router.GET("/course/:c",C.ShowCourse)

	//Modules
	router.GET("/modules_f/:f",M.ShowModulesByFaculty)
	router.GET("/modules_s/:s",M.ShowModulesBySchool)
	router.GET("/modules_c/:c",M.ShowModulesByCourse)
	router.GET("/module/:m",M.ShowModule)

	//Books
	router.GET("/books",B.ShowBooks)
	router.GET("/book_q/:q",B.ShowBooksByQuery)
	router.GET("/book/:b",B.ShowBook)

	//Users
	router.GET("/user/:s",U.ShowUser)
	return router
}
