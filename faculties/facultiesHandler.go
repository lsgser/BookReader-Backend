package faculties

import(
	//"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	//CO "../config"
)


/*
	Handler that will be responsible for displaying the Facalties list via the route
	that it will be attached to
*/
func ShowFaculties(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	/*
		params will be used to get the school/university ID from the route that 
		the Show faculties function will be attached to

		CHECK line 27 of routes that I commented params.ByName("s") will fetch the value
		that s has in that route.
	*/
}

/*
	Responsible for displaying a single faculty via the route that it will be attached 
	to
*/
func ShowFaculty(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	/*
		params will be used to get the faculty id from the route that 
		the Show a single faculty function will be attached to

		CHECK line 28 of routes that I commented params.ByName("f") will fetch the value
		that f has in that route i.e the f stands for faculty f in this case if a place-holder
		for any faculty id that will provided. 
	*/
}