package courses

import(
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
	CO "../config"
	A "../auth"
	//"log"
)

/*
	Display courses by school/university ID
*/
func ShowCoursesBySchool(w http.ResponseWriter,req *http.Request,params httprouter.Params){
	CO.AddSafeHeaders(&w)

	school := params.ByName("s")

	if school == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"School was not provided"}`))
		return
	}

	s,err := strconv.ParseInt(school,10,64)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	courses,err := GetCoursesBySchool(s)
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(courses)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}
}

/*
	Display courses by faculty ID
*/
func ShowCoursesByFaculty(w http.ResponseWriter,req *http.Request,params httprouter.Params){
	CO.AddSafeHeaders(&w)

	faculty := params.ByName("f")

	if faculty == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Faculty was not provided"}`))
		return
	}

	f,err := strconv.ParseInt(faculty,10,64)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	courses,err := GetCoursesByFaculty(f)
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(courses)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}
}

/*
	Display a single course
*/
func ShowCourse(w http.ResponseWriter,req *http.Request,params httprouter.Params){
	CO.AddSafeHeaders(&w)

	c := params.ByName("c")

	if c == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Course was not provided"}`))
		return
	}

	crs,err := strconv.ParseInt(c,10,64)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	course,err := GetCourse(crs)
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(course)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}	
}

/*
	Adds a new course
*/
func AddCourse(w http.ResponseWriter,req *http.Request,_ httprouter.Params){
	CO.AddSafeHeaders(&w)
	course := NewCourse()
	body := req.Body
	defer body.Close()
	err := json.NewDecoder(body).Decode(course)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	if A.CheckAdmin(course.Token) {
		if course.School != 0 && course.Faculty != 0 && course.Course != ""{
			err = course.SaveCourse()
			if err != nil{
				w.WriteHeader(500)
				w.Write([]byte(`{"status":"`+err.Error()+`"}`))
				return
			}
		}else{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"Fill/Select all the required fields"}`))
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