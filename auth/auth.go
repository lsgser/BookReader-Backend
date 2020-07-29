package auth

import(
	CO "../config"
	"crypto/sha1"
)

type AuthToken struct{
	UserID int64 `json:"user_id,omitempty"`
	AToken string `json:"token,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`	
}

func CheckToken(token string,tokenType int) (bool,int64){
	//User
	Token := AuthToken{}
	data := []byte(token)
	hashToken := sha1.Sum(data)

	/*
		Run an SQL query to check for the token 
		if the token exists return true,and the user_id else
		return false and 0
	*/
	db,err := CO.GetDB()

	if err != nil{
		return false,0
	}

	if(tokenType == 1){
		stmt,err := db.Prepare("SELECT user_id FROM login_tokens WHERE token=?")
	}else if(tokenType == 2){
		stmt,err := db.Prepare("SELECT user_id FROM admin_login_tokens WHERE token=?")
	}
	
	if err != nil{
		return false,0	
	}
	
	defer stmt.Close()

	err = stmt.QueryRow(string(hashToken)).Scan(&Token.UserID)
	if err != nil {
		return false,0
	}

	return true,Token.UserID 
}

/*
	Checks if the user exists and also
	checks if the user is logged in
*/
func CheckStudent(student string) (bool,int64){
	return false,0
}