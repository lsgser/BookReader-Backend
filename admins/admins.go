package admins

import(
	CO "example/BookReader-Backend/config"
	A "example/BookReader-Backend/auth"
	"errors"
	"strings"
	"github.com/badoux/checkmail"
	"database/sql"
)

type Admin struct{
	ID int64 `json:"-"`
	Email string `json:"email"`
	Password string `json:"password,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewAdmin() *Admin{
	return new(Admin)
}

/*
	Uses AuthAdmin to login 
	an administrator and then
	allocates a token along with the 
	admins email
*/
func (a *Admin) AdminLogin() (string,error){
	if strings.TrimSpace(a.Email) == "" || strings.TrimSpace(a.Password) == ""{
		err := errors.New("Fill in all fields")
		return "",err
	}

	token,err := A.AuthAdmin(a.Email,a.Password)

	if err != nil {
		return "",err 
	}

	return token,err
}

/*
	Create a new admin
*/
func (a *Admin) SaveAdmin() (err error){
	admin := Admin{}
	
	if strings.TrimSpace(a.Email) == "" || strings.TrimSpace(a.Password) == ""{
		err = errors.New("Fill in all fields")
		return err
	}
	
	if checkmail.ValidateFormat(a.Email) != nil{
		err = errors.New("Invalid email format")
		return err
	}

	/*
		Check if the email address already exists
	*/
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return err
	}
	
	defer db.Close()
	emailStmt,err := db.Prepare("SELECT * FROM admins WHERE email=?")

	if err != nil{
		return err
	}
	defer emailStmt.Close()

	err = emailStmt.QueryRow(a.Email).Scan(&admin.ID,&admin.Email,&admin.Password,&admin.CreatedAt,&admin.UpdatedAt)

	if err == sql.ErrNoRows {
		hashedPass,err := CO.HashPassword(a.Password)

		if err != nil {
			return err
		}
		insertStmt,err := db.Prepare("INSERT INTO admins (email,password) VALUES (?,?)")

		if err != nil{
			return err
		}

		defer insertStmt.Close()

		_,err = insertStmt.Exec(a.Email,hashedPass)

		if err != nil{
			return err
		}

		return nil
			
	}else if err != nil{
		return err
	}else{
		err = errors.New("Admin already exists")
		return err
	}

	return nil
}

/*
	Check if the admin is logged in
*/
func AdminIsLoggedIn(token string) bool{
	isLogged := A.CheckAdmin(token)

	return isLogged
}

func AdminLogout(token string) error{
	err := A.DeleteAdminAuth(token)

	return err
}