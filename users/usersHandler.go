package users

import(
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	CO "../config"
	A "../auth"
	//UP "../uploads"
	"github.com/badoux/checkmail"
	"strings"
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

/*
	Add a new user
*/
func AddUser(w http.ResponseWriter , req *http.Request , _ httprouter.Params){
	CO.AddSafeHeaders(&w)
	user := NewUser()
	body := req.Body
	defer body.Close()
	err := json.NewDecoder(body).Decode(user)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	if A.CheckAdmin(user.Token){
		/*
			Check if the email is valid
		*/
		if checkmail.ValidateFormat(user.Email) != nil{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"Invalid email format"}`))
			return	
		}

		user.Name = strings.TrimSpace(strings.Title(strings.ToLower(user.Name)))
		user.Surname = strings.TrimSpace(strings.Title(strings.ToLower(user.Surname)))
		user.Student = strings.TrimSpace(user.Student)
		user.Picture = "/data/users/user_default.jpg"
		user.Password = strings.TrimSpace(user.Password)

		if user.Name != "" && user.Surname != "" && user.Student != "" && user.Password != "" && user.Course != 0 && user.School != 0 && user.Faculty != 0{
			err = user.SaveUser()
			if err != nil{
				w.WriteHeader(500)
				w.Write([]byte(`{"status":"`+err.Error()+`"}`))
				return
			}
		}else{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"Fill in all the required fields"}`))
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

/*
	LoginUser assigns a new login to the 
	user
*/
func LoginUser(w http.ResponseWriter,req *http.Request,_ httprouter.Params){
	CO.AddSafeHeaders(&w)
	body := req.Body
	user := NewLogInUser()
	defer body.Close()
	err := json.NewDecoder(body).Decode(user)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	token,err := user.UserLogin()

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(`{"status":"Success","token":`+token+`}`))	
}

/*
	Check if user is Logged in
*/
func UserLogged(w http.ResponseWriter , req *http.Request, params httprouter.Params){
	CO.AddSafeHeaders(&w)

	token := params.ByName("t")

	if !UserIsLoggedIn(token){
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
func UserSignOut(w http.ResponseWriter , req *http.Request, params httprouter.Params){
	CO.AddSafeHeaders(&w)

	token := params.ByName("t")

	err := UserLogout(token)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(`{"status":"Logged out"}`))
}