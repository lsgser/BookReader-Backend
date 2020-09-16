package auth

import(
	CO "../config"
	"github.com/badoux/checkmail"
	"strings"
	"database/sql"
	"errors"
	"log"
	"unicode/utf8"
)

/*
	Checks if the token is valid
	i.e checks if the user is logged in
*/
func CheckUser(token string) bool{
	db,err := CO.GetDB()
	if err != nil{
		return false
	}

	hashedToken := CO.HashData(token)

	stmt,err := db.Prepare("SELECT user_id FROM login_tokens WHERE token = ?")

	if err != nil{
		return false
	}

	defer stmt.Close()

	var user_id int64
	err = stmt.QueryRow(hashedToken).Scan(&user_id)

	if err != nil{
		return false
	}

	return true
}

/*
	Checks if the token is valid
	i.e checks if the admin is logged in
*/
func CheckAdmin(token string) bool{
	db,err := CO.GetDB()
	if err != nil{
		return false
	}

	hashedToken := CO.HashData(token)
	log.Println("Token: %s",hashedToken)
	stmt,err := db.Prepare("SELECT user_id FROM admin_login_tokens WHERE token = ?")

	if err != nil{
		return false
	}

	defer stmt.Close()

	var user_id int64
	err = stmt.QueryRow(hashedToken).Scan(&user_id)

	if err != nil{
		return false
	}
	log.Printf("Correct Check")
	return true
}

/*
	Authenticate an admin user
	when they log in
*/
func AuthAdmin(email string,password string)(string,error){
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return "",err
	}

	password = strings.TrimSpace(password)

	if password == ""{
		err = errors.New("Enter your password")
		return "",err
	}

	if err != nil{
		return "",err
	}

	if checkmail.ValidateFormat(email) != nil{
		err = errors.New("Invalid email format")
		return "",err
	}

	stmt,err := db.Prepare("SELECT * FROM admins WHERE email = ?")

	if err != nil{
		return "",err
	}

	defer stmt.Close()

	var (
		id int64
		mail string
		pass string
		created_at string
		updated_at string
		token string
	) 

	err = stmt.QueryRow(email).Scan(&id,&mail,&pass,&created_at,&updated_at)

	if err != nil{
		err = errors.New("This admin does not exist")
		return "",err
	}

	err = CO.CheckPassword(pass,password)

	if err != nil{
		err = errors.New("Wrong password")
		return "",err
	}

	/*
		We use the user id to check if the admin data exists
		in the tokens table,
		if it exists we remove it and then create a new token
	*/
	authStmt,err := db.Prepare("SELECT token FROM admin_login_tokens WHERE user_id = ?")

	if err != nil{
		return "",err
	}

	defer authStmt.Close()

	err = authStmt.QueryRow(id).Scan(&token)

	switch{
	case err == sql.ErrNoRows:
		//Do nothing
	case err != nil:
		return "",err
	default:
		//Delete the row and insert a new token and data
		delAdmin,err := db.Prepare("DELETE FROM admin_login_tokens WHERE user_id = ?")
		if err != nil{
			return "",err
		}
		_,err = delAdmin.Exec(id)
		if err != nil{
			return "",err
		}
	}
	/*
		Add new auth data
		and create a token
	*/
	authQuery,err := db.Prepare("INSERT INTO admin_login_tokens (user_id,token) VALUES(?,?)")
	if err != nil{
		return "",err
	}

	/*
		Create a token
	*/
	token = CO.GenerateToken(32)
	log.Println("%T",token)
	if utf8.ValidString(token){
		log.Println("Valid")
	}else{
		log.Println("Invalid")
	}
	hashedToken := CO.HashData(token)
	if utf8.ValidString(hashedToken){
		log.Println("Valid hash")
	}else{
		log.Println("Invalid hash")
		log.Println(hashedToken)
	}
	_,err = authQuery.Exec(id,hashedToken)

	if err != nil{
		return "",err
	}

	return token,nil
}

/*
	DeleteAdminAuth helps
	with logging out the admin
	user
*/
func DeleteAdminAuth(token string) error{
	db,err := CO.GetDB()
	if err != nil{
		err = errors.New("DB connection error")
		return err
	}

	hashedToken := CO.HashData(token)

	stmt,err := db.Prepare("SELECT user_id FROM admin_login_tokens WHERE token = ?")

	if err != nil{
		return err
	}

	defer stmt.Close()

	var user_id int64
	err = stmt.QueryRow(hashedToken).Scan(&user_id)

	if err == sql.ErrNoRows{
		err = errors.New("Admin already logged out")
		return err
	}else if err != nil{
		return err
	}

	logoutStmt,err := db.Prepare("DELETE FROM admin_login_tokens WHERE token = ?")
	
	if err != nil{
		return err
	}

	_,err = logoutStmt.Exec(hashedToken)
	
	if err != nil{
		return err	
	}

	return err		
}