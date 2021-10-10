package modules

import(
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
	CO "example/BookReader-Backend/config"
	A "example/BookReader-Backend/auth"
)

func ShowModulesBySchool(w http.ResponseWriter,req *http.Request,params httprouter.Params){
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

	modules,err := GetModulesBySchool(s)
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(modules)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}
}

func ShowModulesByFaculty(w http.ResponseWriter,req *http.Request,params httprouter.Params){
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

	modules,err := GetModulesByFaculty(f)
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(modules)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}
}

func ShowModulesByCourse(w http.ResponseWriter,req *http.Request,params httprouter.Params){
	CO.AddSafeHeaders(&w)

	course := params.ByName("c")

	if course == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Course was not provided"}`))
		return
	}

	c,err := strconv.ParseInt(course,10,64)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	modules,err := GetModulesByCourse(c)
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(modules)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}
}

func ShowFacultyModuleByName(w http.ResponseWriter,req *http.Request,params httprouter.Params){
	CO.AddSafeHeaders(&w)

	course := params.ByName("c")
	moduleName := params.ByName("m")

	if course == "" || moduleName == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Course was not provided"}`))
		return
	}

	c,err := strconv.ParseInt(course,10,64)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	modules,err := GetCourseModuleByName(c,moduleName)
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(modules)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}	
}

func ShowModule(w http.ResponseWriter,req *http.Request,params httprouter.Params){
	CO.AddSafeHeaders(&w)

	module := params.ByName("m")

	if module == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Module was not provided"}`))
		return
	}

	m,err := strconv.ParseInt(module,10,64)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	mdl,err := GetModule(m)
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(mdl)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}
}

/*
	Adds a new module to the system
*/
func AddModule(w http.ResponseWriter,req *http.Request,_ httprouter.Params){
	CO.AddSafeHeaders(&w)
	body := req.Body
	module:= NewModule()
	defer body.Close()

	err := json.NewDecoder(body).Decode(module)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	if A.CheckAdmin(module.Token){
		if module.School != 0 && module.Faculty != 0 && module.Course != 0 && module.Module != ""{
			err = module.SaveModule()
			if err != nil{
				w.WriteHeader(500)
				w.Write([]byte(`{"status":"`+err.Error()+`"}`))
				return
			}

		}else{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"Fill/Select all required fields"}`))
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