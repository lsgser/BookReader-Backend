package modules

import(
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
	CO "../config"
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