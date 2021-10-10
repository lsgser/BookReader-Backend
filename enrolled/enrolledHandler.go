package enrolled

import (
	"encoding/json"
	"net/http"
	"strconv"
	CO "example/BookReader-Backend/config"
	"github.com/julienschmidt/httprouter"
	A "example/BookReader-Backend/auth"
)

func ShowEnrolledByModule(w http.ResponseWriter, req *http.Request, params httprouter.Params){
	CO.AddSafeHeaders(&w)

	module := params.ByName("m")

	if module == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Module was not provided"}`))
		return
	}

	m,err := strconv.ParseInt(module,10,64)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	enrolled, err := GetEnrolledByModule(m)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	err = json.NewEncoder(w).Encode(enrolled)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}
}

func ShowEnrolledByUser(w http.ResponseWriter, req *http.Request, params httprouter.Params){
	CO.AddSafeHeaders(&w)

	user:= params.ByName("u")

	if user == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Student was not provided"}`))
		return
	}

	enrolled, err := GetEnrolledByUser(user)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	err = json.NewEncoder(w).Encode(enrolled)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}
}

func ShowEnrolledUsers(w http.ResponseWriter, req *http.Request, params httprouter.Params){
	CO.AddSafeHeaders(&w)
	module := params.ByName("m")

	if module == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Module was not provided"}`))
		return
	}

	m,err := strconv.ParseInt(module,10,64)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	modules, err := GetEnrolledUsers(m)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	err = json.NewEncoder(w).Encode(modules)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}
}

func ShowEnrolledModules(w http.ResponseWriter, req *http.Request, params httprouter.Params){
	CO.AddSafeHeaders(&w)

	user := params.ByName("u")

	if user == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Student was not provided"}`))
		return
	}

	modules, err := GetEnrolledModules(user)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	err = json.NewEncoder(w).Encode(modules)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}
}

func AddEnrolled(w http.ResponseWriter , req *http.Request , _ httprouter.Params){
	CO.AddSafeHeaders(&w)
	enrol := NewSaveEnrol()
	body := req.Body
	defer body.Close()
	err := json.NewDecoder(body).Decode(enrol)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	if A.CheckAdmin(enrol.Token){
		if enrol.Module != 0 && enrol.User != ""{
			err = enrol.SaveEnrolled()
			if err != nil{
				w.WriteHeader(400)
				w.Write([]byte(`{"status":"`+err.Error()+`"}`))
				return
			}			
		}else{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"Fill in all required fields"}`))
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