package enrolled

import (
	"encoding/json"
	"net/http"
	"strconv"
	CO "../config"
	"github.com/julienschmidt/httprouter"
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