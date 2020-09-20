package faculties

import (
	"encoding/json"
	"net/http"
	"strconv"
	CO "../config"
	"github.com/julienschmidt/httprouter"
	A "../auth"
	//"log"
)

/*
	Handler that will be responsible for displaying the Facalties list via the route
	that it will be attached to
*/
func ShowFaculties(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	/*
		params will be used to get the school/university ID from the route that
		the Show faculties function will be attached to

		CHECK line 27 of routes that I commented params.ByName("s") will fetch the value
		that s has in that route.
	*/
	CO.AddSafeHeaders(&w)

	school, err := strconv.Atoi(params.ByName("s"))

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	faculties, err := GetFaculties(int64(school))

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	err = json.NewEncoder(w).Encode(faculties)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}
}

/*
	Responsible for displaying a single faculty via the route that it will be attached
	to
*/
func ShowFaculty(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	/*
		params will be used to get the faculty id from the route that
		the Show a single faculty function will be attached to

		CHECK line 28 of routes that I commented params.ByName("f") will fetch the value
		that f has in that route i.e the f stands for faculty f in this case if a place-holder
		for any faculty id that will provided.
	*/
	CO.AddSafeHeaders(&w)
	school, err := strconv.Atoi(params.ByName("f"))

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	faculties, err := GetFaculty(int64(school))

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	err = json.NewEncoder(w).Encode(faculties)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}
}

func AddFaculty(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	CO.AddSafeHeaders(&w)
	body := req.Body
	faculty := NewFaculty()
	defer body.Close()
	err := json.NewDecoder(body).Decode(faculty)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	if A.CheckAdmin(faculty.Token){
		if faculty.School != 0 && faculty.Faculty != ""{
			err = faculty.SaveFaculty()

			if err != nil{
				w.WriteHeader(500)
				w.Write([]byte(`{"status":"`+err.Error()+`"}`))
				return
			}
		}else{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"All required fields were not provided."}`))
			return
		}
	}else{
		w.WriteHeader(404)
		w.Write([]byte(`{"status":"Invalid login session."}`))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(`{"status":"Success"}`))
}

