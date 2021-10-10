package schools

import(
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	CO "example/BookReader-Backend/config"
	UP "example/BookReader-Backend/uploads"
	A "example/BookReader-Backend/auth"
	//"net/url"
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
	//schoolName,err := url.QueryUnescape(schoolName)
	/*if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}
	*/
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
func AddSchool(w http.ResponseWriter, req *http.Request,_ httprouter.Params){
	CO.AddSafeHeaders(&w)
	school := NewSchool()
	err := req.ParseMultipartForm(3 * UP.GetMB())

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return 
	}

	//Limit upload size to 3MB
	req.Body = http.MaxBytesReader(w,req.Body,3 * UP.GetMB())
	/*
		Check if the admin token is valid
		i.e the admin is logged in
	*/
	if A.CheckAdmin(req.FormValue("token")){
		file,handler,err := req.FormFile("school_icon")

		if err != nil{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"`+err.Error()+`"}`))
			return
		}

		if file == nil{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"Institution icon was not included"}`))
			return
		}else{
			defer file.Close()
		}

		if req.FormValue("school") == ""{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"Institution name not included"}`))
			return
		}
		
		school.School = req.FormValue("school")
		school.SchoolIcon,err = UP.ImageFileUpload(file,handler,"/data/images/institutions/","school")

		if err != nil{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"`+err.Error()+`"}`))
			return
		}

		/*
			Save the data via a controller
		*/
		err = school.SaveSchool()

		if err != nil{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"`+err.Error()+`"}`))
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