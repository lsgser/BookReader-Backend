package users

import(
	CO "../config"
	"strings"
	"errors"
	A "../auth"
)

type User struct{
	ID int64 `json:"-"`
	School int64 `json:"school"`
	Faculty int64 `json:"faculty"`
	Course int64 `json:"course"`
	Student string `json:"student"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Picture string `json:"picture"`
	Password string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Token string `json:"token,omitempty"`
}

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

	stmt,err := db.Prepare("SELECT * FROM users WHERE student_nr = ?")

	if err != nil{
		return user,err
	}

	defer stmt.Close()

	err = stmt.QueryRow(student).Scan(&user.ID,&user.School,&user.Faculty,&user.Course,&user.Student,&user.Name,&user.Surname,&user.Email,&user.Picture,&user.Password,&user.CreatedAt,&user.UpdatedAt)

	if err != nil{
		return user,err
	}

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
	return users,nil
}

/*
	Retrieves users by name,surname,or student number
*/
func GetUsersByQuery(query string) ([]User,error){
	users := make([]User,0)
	/*
		CODE will be added here
	*/
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

	hashedPass,err := CO.HashPassword(u.Password)

	if err != nil {
		return err
	}

	stmt,err := db.Prepare("INSERT INTO users (school_id,faculty_id,course_id,student_nr,name,surname,email,picture,password) VALUES (?,?,?,?,?,?,?,?,?)")

	if err != nil{
		return err
	}

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