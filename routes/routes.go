package routes

import (
	F "../faculties"
	C "../courses"
	M "../modules"
	B "../books"
	U "../users"
	E "../enrolled"
	S "../schools"
	R "../required"
	A "../admins"
	RE "../recommended"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

//NewRouter : new router returns all router
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	//Faculties
	router.GET("/faculties/:s", F.ShowFaculties)
	router.GET("/faculty/:f", F.ShowFaculty)
	router.POST("/new_faculty",F.AddFaculty)

	//Courses
	router.GET("/courses_f/:f",C.ShowCoursesByFaculty)
	router.GET("/courses_s/:s",C.ShowCoursesBySchool)
	router.GET("/course/:c",C.ShowCourse)
	router.POST("/new_course",C.AddCourse)

	//Modules
	router.GET("/modules_f/:f",M.ShowModulesByFaculty)
	router.GET("/modules_s/:s",M.ShowModulesBySchool)
	router.GET("/modules_c/:c",M.ShowModulesByCourse)
	router.GET("/module_n/:c/:m",M.ShowFacultyModuleByName)
	router.GET("/module/:m",M.ShowModule)
	router.POST("/new_module",M.AddModule)

	//Books
	router.GET("/books",B.ShowBooks)
	router.GET("/book_q/:q",B.ShowBooksByQuery)
	router.GET("/book/:b",B.ShowBook)
	router.POST("/new_book",B.AddBook)
	
	//Users
	router.GET("/user/:s",U.ShowUser)
	router.GET("/user_t/:t",U.ShowUserByToken)
	router.GET("/users",U.ShowUsers)
	router.GET("/user_s/:s",U.ShowUsersBySchool)
	router.GET("/user_f/:f",U.ShowUsersByFaculty)
	router.GET("/user_c/:c",U.ShowUsersByCourse)
	router.GET("/users_q/:q",U.ShowUsersByQuery)
	router.POST("/user_login",U.LoginUser)
	router.GET("/user_logged/:t",U.UserLogged)
	router.POST("/new_user",U.AddUser)
	router.DELETE("/user_logout/:t",U.UserSignOut)

	//Enrolled
	router.GET("/enrolled_by_user/:u",E.ShowEnrolledByUser)
	router.GET("/enrolled_by_module/:m",E.ShowEnrolledByModule)
	router.GET("/enrolled_modules/:u",E.ShowEnrolledModules)
	router.GET("/enrolled_users/:m",E.ShowEnrolledUsers)
	router.POST("/new_enrolled",E.AddEnrolled)

	//School
	router.GET("/schools",S.ShowSchools)
	router.GET("/school/:s",S.ShowSchool)
	router.POST("/new_school",S.AddSchool)

	//Required
	router.GET("/required_by_user/:u",R.ShowRequiredByUser)
	router.GET("/required_by_module/:m",R.ShowRequiredByModule)
	router.GET("/required_by_book/:isbn",R.ShowRequiredByBook)
	router.POST("/new_required/",R.AddRequired)

	//Admin
	router.POST("/newadmin",A.AddAdmin)
	router.POST("/admin_login",A.LoginAdmin)
	router.GET("/admin_logged/:t",A.AdminLogged)
	router.DELETE("/admin_logout/:t",A.AdminSignOut)
	
	//Recommended
	router.GET("/recommended_by_book/:isbn",RE.ShowRecommendedByBook)
	router.GET("/recommended_by_module/:m",RE.ShowRecommendedByModule)
	router.GET("/recommended_by_m_and_b/:m/:isbn",RE.ShowRecommendedByModuleAndBook)
	router.GET("/recommended_by_user/:u",RE.ShowRecommendedByUser)
	router.POST("/new_recommend",RE.AddRecommended)

	/*serve institution images*/
	router.ServeFiles("/institution/*filepath",http.Dir("data/images/institutions/"))
	router.ServeFiles("/user_pic/*filepath",http.Dir("data/images/users/"))
	router.ServeFiles("/cover_page/*filepath",http.Dir("data/images/book_covers/"))

	/*PDF book file*/
	router.ServeFiles("/b/*filepath",http.Dir("data/books/"))
	return router
}
