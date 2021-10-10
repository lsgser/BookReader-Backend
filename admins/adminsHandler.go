package admins

import(
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	CO "example/BookReader-Backend/config"
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
func LoginAdmin(w http.ResponseWriter,req *http.Request,_ httprouter.Params){
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

	token,err := admin.AdminLogin()

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(`{"status":"Success","token":"`+token+`"}`))	
}


/*
	Check if admin is Logged in
*/
func AdminLogged(w http.ResponseWriter , req *http.Request, params httprouter.Params){
	CO.AddSafeHeaders(&w)

	token := params.ByName("t")

	if !AdminIsLoggedIn(token){
		w.WriteHeader(404)
		w.Write([]byte(`{"status":"Not logged in"}`))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(`{"status":"Logged in"}`))
}

/*
	Logout admin
*/
func AdminSignOut(w http.ResponseWriter , req *http.Request, params httprouter.Params){
	CO.AddSafeHeaders(&w)

	token := params.ByName("t")

	err := AdminLogout(token)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(`{"status":"Logged out"}`))
}