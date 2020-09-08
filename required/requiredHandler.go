package required

import(
	"encoding/json"
	"net/http"
	"strconv"
	CO "../config"
	"github.com/julienschmidt/httprouter"
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