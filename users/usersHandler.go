package users

import(
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	CO "../config"
)

//Handler responsible for displaying a single user based on their student id
func ShowUser(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)
	student := params.ByName("s")

	user,err := GetUser(student)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))		
		return
	}

	err = json.NewEncoder(w).Encode(user)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}
}