package schools

import(
	"errors"
	CO "../config"
	"strings"
)

/*
	School struct
*/
type School struct{
	ID int64 `json:"-"`
	School string `json:"school"`
	SchoolIcon string `json:"school_icon"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Token string `json:"token,omitempty"`
}

type EditForm struct{
	School string `json:"school,omitempty"`
	NewSchoolName string `json:"new_school_name,omitempty"`
	NewSchoolIcon string `json:"new_school_icon,omitempty"`
	Token string `json:"token,omitempty"`
}

//NewSchool returns a pointer to a School struct
func NewSchool() *School{
	return new(School)
}

func NewEditForm() *EditForm{
	return new(EditForm)
}

//GetSchools returns a slice of School that contains all schools in the database
func GetSchools() ([]School,error){
	schools := make([]School,0)
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return schools,err
	}

	rows,err := db.Query("SELECT school,school_icon FROM schools")

	if err != nil{
		return schools,err
	}

	defer rows.Close()

	for rows.Next(){
		s := School{}
		rows.Scan(&s.School,&s.SchoolIcon)
		schools = append(schools,s)
	}

	return schools,nil
}

//GetSchool returns a school based on the school string that's provided 
func GetSchool(school string) (School,error){
	s := School{}
	db,err := CO.GetDB()

	if err != nil{
		return s,err
	}

	stmt,err := db.Prepare("SELECT school,school_icon FROM schools WHERE school=?")

	if err != nil{
		return s,err
	}

	defer stmt.Close()
	err = stmt.QueryRow(school).Scan(&s.School,&s.SchoolIcon)

	if err != nil{
		return s,err
	}

	return s,nil
}

//SaveSchool is a method that will return an error if the school data cannot be saved
func (s *School) SaveSchool() (err error){
	db,err := CO.GetDB()

	if err != nil{
		return err
	}

	s.School = strings.ToUpper(strings.TrimSpace(s.School))

	if s.School != "" && s.SchoolIcon != ""{
		stmt,err := db.Prepare("INSERT INTO schools (school,school_icon) VALUES(?,?)")

		if err != nil{
			return err
		}

		_,err = stmt.Exec(s.School,s.SchoolIcon)

		if err != nil{
			return err
		}

	}else{
		err = errors.New("Fill in the institution name and also upload it's logo.")
		return err
	}

	return err	
}

//EditSchool is a method that will return an error if the school data is not edited
func (s *EditForm) EditSchool() (err error){
	school := EditForm{}

	s.School = strings.ToUpper(strings.TrimSpace(s.School))
	s.NewSchoolName =  strings.ToUpper(strings.TrimSpace(s.NewSchoolName))
	s.NewSchoolIcon = strings.TrimSpace(s.NewSchoolIcon)
	if s.Token == ""{
		err = errors.New("Invalid login data")
		return err
	}

	db,err := CO.GetDB()

	if err != nil{
		return err
	}

	if s.NewSchoolName == "" && s.NewSchoolIcon == ""{
		err = errors.New("Edit fields are both empty")
		return err
	}

	/*
		Check if the school exists
	*/
	if  s.School != "" {
		schoolStmt,err := db.Prepare("SELECT school FROM schools WHERE school = ?")

		if err != nil{
			return err
		}

		defer schoolStmt.Close()

		err = schoolStmt.QueryRow(s.School).Scan(&school.School)

		if err != nil{
			return err
		}

		school.School  = strings.ToUpper(strings.TrimSpace(school.School))
		if s.NewSchoolName != ""{
			if s.NewSchoolName != school.School{
				editNameStmt,err := db.Prepare("UPDATE schools SET school = ? WHERE school=?")

				if err != nil{
					return err
				}

				_,err = editNameStmt.Exec(s.NewSchoolName,school.School)

			}else{
				err = errors.New("Can't edit a name using the same institution name.")
			}
		}

		if s.NewSchoolIcon != ""{
			editIconStmt,err := db.Prepare("UPDATE schools SET school_icon = ? WHERE school=?")

			if err != nil{
				return err
			}

			_,err = editIconStmt.Exec(s.NewSchoolIcon,school.School)

			if err != nil{
				return err
			}
		}
	}else{
		err = errors.New("Institution not selected")
	}

	return err
}

//DeleteSchool is a function that will be used to delete a specific school
func DeleteSchool(school string,token string) (err error){
	db,err := CO.GetDB()

	school = strings.ToUpper(strings.TrimSpace(school))

	if err != nil{
		return err
	}

	if school != ""{
		delSchool,err := db.Prepare("DELETE FROM schools WHERE school = ?")

		if err != nil{
			return err
		}

		_,err = delSchool.Exec(school)
	}else{
		err = errors.New("Select an institution to delete")
	}

	return err
}