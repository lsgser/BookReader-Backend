package schools

import(
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	CO "../config"
)

/*
	Display all schools
*/
func ShowSchools(w http.ResponseWriter , req *http.Request , _ httprouter.Params){
	CO.AddSafeHeaders(&w)
	schools,err := GetSchools()

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))		
		return
	}

	err = json.NewEncoder(w).Encode(schools)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}	
}

/*
	Display a school
*/

func ShowSchool(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)
	schoolName := params.ByName("s")
	
	if schoolName == "" {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Institution was not provided"}`))
		return
	}
	
	school,err := GetSchool(schoolName)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return	
	}

	err = json.NewEncoder(w).Encode(school)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}
}

/*
	AddSchool
*/
