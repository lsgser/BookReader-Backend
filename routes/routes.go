package routes

import (
	F "../faculties"
	C "../courses"
	M "../modules"
	B "../books"
	"github.com/julienschmidt/httprouter"
)

//NewRouter : new router returns all router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	//router.GET("/",Index)
	/*
		router.GET("/users",U.ShowUsers)
		router.POST("/adduser",U.AddUser)
		router.GET("/user/:u",U.ShowUser)
		router.PUT("/edituser",U.UpdateUser)
		router.DELETE("/removeuser/:u",U.RemoveUser)
	*/
	//Profiles
	/*
		router.GET("/profiles/",P.ShowProfiles)
		router.GET("/profile/:u",P.ShowProfile)
		router.POST("/addprofile",P.AddProfile)
	*/

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
	return router
}
