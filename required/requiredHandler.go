package required

import(
	"encoding/json"
	"net/http"
	"strconv"
	CO "example/BookReader-Backend/config"
	"github.com/julienschmidt/httprouter"
	A "example/BookReader-Backend/auth"
)

func ShowRequiredByUser(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)
	
	if params.ByName("u") == "" {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"User was not provided"}`))
		return
	}
	
	required,err := GetRequiredByUser(params.ByName("u"))

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return	
	}

	err = json.NewEncoder(w).Encode(required)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}
}

func ShowRequiredByModule(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)

	if params.ByName("m") == "" {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Module was not provided"}`))
		return
	}

	module,err := strconv.ParseInt(params.ByName("m"),10,64)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	required,err := GetRequiredByModule(module)
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return	
	}

	err = json.NewEncoder(w).Encode(required)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}	
}

func ShowRequiredByBook(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)
	
	if params.ByName("isbn") == "" {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"ISBN was not provided"}`))
		return
	}
	
	required,err := GetRequiredByBook(params.ByName("isbn"))

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return	
	}

	err = json.NewEncoder(w).Encode(required)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}
}

func AddRequired(w http.ResponseWriter , req *http.Request , _ httprouter.Params){
	CO.AddSafeHeaders(&w)
	sRequired := NewSaveRequired()
	body := req.Body
	defer body.Close()
	err := json.NewDecoder(body).Decode(sRequired)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	if A.CheckAdmin(sRequired.Token){
		if sRequired.User == "" || sRequired.Module == 0 || sRequired.ISBN == ""{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"Fill in all the required data."}`))
			return		
		}

		err = sRequired.SaveRequired()

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