package users

import(
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	CO "example/BookReader-Backend/config"
	A "example/BookReader-Backend/auth"
	//UP "../uploads"
	"github.com/badoux/checkmail"
	"strings"
	"strconv"
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

func ShowUserByToken(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)
	token := params.ByName("t")

	if token == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"No token provided"}`))		
		return
	}

	user,err := GetUserByToken(token)

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

func ShowUsers(w http.ResponseWriter , req *http.Request , _ httprouter.Params){
	CO.AddSafeHeaders(&w)
	users,err := GetUsers()

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(users)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}
}

func ShowUsersBySchool(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)
	school := params.ByName("s")

	if school == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"School was not provided"}`))
		return
	}

	s,err := strconv.ParseInt(school,10,64)
	
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	users,err := GetUsersBySchool(s)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(users)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}	
}

func ShowUsersByFaculty(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)
	faculty := params.ByName("f")

	if faculty == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Faculty was not provided"}`))
		return
	}

	f,err := strconv.ParseInt(faculty,10,64)
	
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	users,err := GetUsersByFaculty(f)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(users)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}	
}

func ShowUsersByCourse(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)
	course := params.ByName("c")

	if course == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Course was not provided"}`))
		return
	}

	c,err := strconv.ParseInt(course,10,64)
	
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	users,err := GetUsersByCourse(c)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(users)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}	
}

func ShowUsersByQuery(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)
	query := params.ByName("q")

	if query == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Query was not provided"}`))
		return
	}
	
	users,err := GetUsersByQuery(query)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(users)

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
	w.Write([]byte(`{"status":"Success","token":"`+token+`"}`))	
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