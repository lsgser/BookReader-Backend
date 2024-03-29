package users

import(
	CO "example/BookReader-Backend/config"
	"strings"
	"errors"
	A "example/BookReader-Backend/auth"
	P "example/BookReader-Backend/processing"
)

type User struct{
	ID int64 `json:"-"`
	School int64 `json:"school,omitempty"`
	Faculty int64 `json:"faculty,omitempty"`
	Course int64 `json:"course,omitempty"`
	Student string `json:"student,omitempty"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Picture string `json:"picture"`
	Password string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Token string `json:"token,omitempty"`
}

var (
	imagePath = "/user_pic/"
)
type LogInUser struct{
	UserText string `json:"user"`
	Password string `json:"password"`
}

func NewUser() *User{
	return new(User)
}

func NewLogInUser() *LogInUser{
	return new(LogInUser)
}

func GetUser(student string) (User,error){
	user := User{}

	student = strings.TrimSpace(student)

	if student == "" {
		err := errors.New("Student number not provided")
		return user,err
	} 

	db,err := CO.GetDB()

	if err != nil {
		err = errors.New("DB connection error")
		return user,err
	}

	defer db.Close()

	stmt,err := db.Prepare("SELECT school_id,faculty_id,course_id,student_nr,name,surname,email,picture FROM users WHERE student_nr = ?")

	if err != nil{
		return user,err
	}

	defer stmt.Close()

	err = stmt.QueryRow(student).Scan(&user.School,&user.Faculty,&user.Course,&user.Student,&user.Name,&user.Surname,&user.Email,&user.Picture)

	if err != nil{
		return user,err
	}

	user.Picture = imagePath+strings.Split(user.Picture,"/")[3]

	return user,nil
}

func GetUserByToken(token string) (User,error){
	user := User{}
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return user,err
	}

	defer db.Close()

	var(
		user_id int64
	)

	hashedToken := CO.HashData(token)

	stmtID,err := db.Prepare("SELECT user_id FROM login_tokens WHERE token = ?")

	if err != nil{
		return user,err
	}

	defer stmtID.Close()

	err = stmtID.QueryRow(hashedToken).Scan(&user_id)

	if err != nil{
		return user,err
	}

	stmt,err := db.Prepare("SELECT student_nr,name,surname,email,picture FROM users WHERE id = ?")

	if err != nil{
		return user,err
	}

	defer stmt.Close()

	err = stmt.QueryRow(user_id).Scan(&user.Student,&user.Name,&user.Surname,&user.Email,&user.Picture)

	if err != nil{
		return user,err
	}

	user.Picture = imagePath+strings.Split(user.Picture,"/")[3]

	return user,nil	
}

/************************************************************************************
					These functions will be for strictly the admin user

					Students won't have access to this data
*************************************************************************************/
/*
	Retrieves all users in the database
*/
func GetUsers() ([]User,error){
	users := make([]User,0)
	/*
		CODE will be added here
	*/
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return users,err
	}

	defer db.Close()

	rows,err := db.Query("SELECT school_id,faculty_id,course_id,student_nr,name,surname,email,picture FROM users")

	if err != nil{
		return users,err
	}

	defer rows.Close()

	for rows.Next(){
		user := User{}
		rows.Scan(&user.School,&user.Faculty,&user.Course,&user.Student,&user.Name,&user.Surname,&user.Email,&user.Picture)
		user.Picture = imagePath+strings.Split(user.Picture,"/")[3]
		users = append(users,user)
	}

	return users,nil
}	

/*
	Retrieves all users of a specific institution
*/
func GetUsersBySchool(school int64) ([]User,error){
	users := make([]User,0)
	/*
		CODE will be added here
	*/
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return users,err
	}

	defer db.Close()

	rows,err := db.Query("SELECT school_id,faculty_id,course_id,student_nr,name,surname,email,picture FROM users where school_id = ?",school)

	if err != nil{
		return users,err
	}

	defer rows.Close()

	for rows.Next(){
		user := User{}
		rows.Scan(&user.School,&user.Faculty,&user.Course,&user.Student,&user.Name,&user.Surname,&user.Email,&user.Picture)
		user.Picture = imagePath+strings.Split(user.Picture,"/")[3]
		users = append(users,user)
	}

	return users,nil
}

/*
	Retrieves all users by a specific faculty of a school
*/
func GetUsersByFaculty(faculty int64) ([]User,error){
	users := make([]User,0)
	/*
		CODE will be added here
	*/
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return users,err
	}

	defer db.Close()

	rows,err := db.Query("SELECT school_id,faculty_id,course_id,student_nr,name,surname,email,picture FROM users WHERE faculty_id=?",faculty)

	if err != nil{
		return users,err
	}

	defer rows.Close()

	for rows.Next(){
		user := User{}
		rows.Scan(&user.School,&user.Faculty,&user.Course,&user.Student,&user.Name,&user.Surname,&user.Email,&user.Picture)
		user.Picture = imagePath+strings.Split(user.Picture,"/")[3]
		users = append(users,user)
	}

	return users,nil
}

/*
	Retrieves all users by a specific course
*/
func GetUsersByCourse(course int64) ([]User,error){
	users := make([]User,0)
	/*
		CODE will be added here
	*/
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return users,err
	}

	defer db.Close()

	rows,err := db.Query("SELECT school_id,faculty_id,course_id,student_nr,name,surname,email,picture FROM users WHERE course_id = ?",course)

	if err != nil{
		return users,err
	}

	defer rows.Close()

	for rows.Next(){
		user := User{}
		rows.Scan(&user.School,&user.Faculty,&user.Course,&user.Student,&user.Name,&user.Surname,&user.Email,&user.Picture)
		user.Picture = imagePath+strings.Split(user.Picture,"/")[3]
		users = append(users,user)
	}
	
	return users,nil
}

/*
	Retrieves users by name,surname,student number or email
*/
func GetUsersByQuery(query string) ([]User,error){
	users := make([]User,0)
	/*
		CODE will be added here
	*/
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return users,err
	}

	defer db.Close()

	query = "%"+query+"%"

	rows,err := db.Query("SELECT school_id,faculty_id,course_id,student_nr,name,surname,email,picture FROM users WHERE student_nr LIKE ? OR name LIKE ? OR surname LIKE ? OR email LIKE ?",query,query,query,query)

	if err != nil{
		return users,err
	}

	defer rows.Close()

	for rows.Next(){
		user := User{}
		rows.Scan(&user.School,&user.Faculty,&user.Course,&user.Student,&user.Name,&user.Surname,&user.Email,&user.Picture)
		user.Picture = imagePath+strings.Split(user.Picture,"/")[3]
		users = append(users,user)
	}

	return users,nil
}

/*
	Saves a new user to the database
*/
func (u *User) SaveUser() error{
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return err
	}

	defer db.Close()

	ToText := "Student "+u.Student+" :("+CO.MakeTimeStamp()+")\n"
	ToText = ToText+"Password : "+u.Password	
	err = P.WriteToTextFile(ToText,"./text/passwords.txt")
	if err != nil{
		return err 
	}

	hashedPass,err := CO.HashPassword(u.Password)

	if err != nil {
		return err
	}

	stmt,err := db.Prepare("INSERT INTO users (school_id,faculty_id,course_id,student_nr,name,surname,email,picture,password) VALUES (?,?,?,?,?,?,?,?,?)")

	if err != nil{
		return err
	}

	defer stmt.Close()
	
	_,err = stmt.Exec(u.School,u.Faculty,u.Course,u.Student,u.Name,u.Surname,u.Email,u.Picture,hashedPass)

	return err
}

/*
	Uses AuthAdmin to login 
	an administrator and then
	allocates a token along with the 
	admins email
*/
func (u *LogInUser) UserLogin() (string,error){
	token,err := A.AuthUser(u.UserText,u.Password)

	if err != nil {
		return "",err 
	}
	
	return token,err
}
/*
	Check if the user is logged in
*/
func UserIsLoggedIn(token string) bool{
	isLogged := A.CheckUser(token)

	return isLogged
}

/*
	Logs out the user
*/
func UserLogout(token string) error{
	err := A.DeleteUserAuth(token)

	return err
}