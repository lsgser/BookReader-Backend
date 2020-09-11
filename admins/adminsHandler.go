package admins

import(
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	CO "../config"
)

/*
	Creates a new admin
*/
func AddAdmin(w http.ResponseWriter,req *http.Request,_ httprouter.Params){
	CO.AddSafeHeaders(&w)
	body := req.Body
	admin := NewAdmin()
	defer body.Close()
	err := json.NewDecoder(body).Decode(admin)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = admin.SaveAdmin()

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(`{"status":"Success"}`))	
}

/*
	Login an admin
*/
