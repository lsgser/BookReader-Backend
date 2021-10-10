package recommended

import(
	"encoding/json"
	"net/http"
	"strconv"
	CO "example/BookReader-Backend/config"
	"github.com/julienschmidt/httprouter"
	A "example/BookReader-Backend/auth"
)

func ShowRecommendedByBook(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)

	if params.ByName("isbn") == "" {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Book was not provided"}`))
		return
	}

	book,err := GetRecommendedByBook(params.ByName("isbn"))
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return	
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}
}

func ShowRecommendedByModule(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)

	if params.ByName("m") == "" {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Module was not provided"}`))
		return
	}

	module,err := strconv.ParseInt(params.ByName("m"),10,64)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	recommend,err := GetRecommendedByModule(module)
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return	
	}

	err = json.NewEncoder(w).Encode(recommend)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}
}

func ShowRecommendedByModuleAndBook(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)

	if params.ByName("m") == "" || params.ByName("isbn") == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Module was not provided"}`))
		return
	}

	module,err := strconv.ParseInt(params.ByName("m"),10,64)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	recommend,err := GetRecommendedByModuleAndBook(params.ByName("isbn"),module)
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return	
	}

	err = json.NewEncoder(w).Encode(recommend)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}
}

func ShowRecommendedByUser(w http.ResponseWriter, req * http.Request,params httprouter.Params){
	CO.AddSafeHeaders(&w)
	if params.ByName("u") == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"User was not provided"}`))
		return
	}

	user := params.ByName("u")
	books,err := GetRecommendedByUser(user)
	
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(books)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}
}

func AddRecommended(w http.ResponseWriter , req *http.Request , _ httprouter.Params){
	CO.AddSafeHeaders(&w)
	recommended := NewSaveRecommended()
	body := req.Body
	defer body.Close()
	err := json.NewDecoder(body).Decode(recommended)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	if A.CheckAdmin(recommended.Token){
		if recommended.Module == 0 || recommended.ISBN == ""{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"Fill in all the required data."}`))
			return		
		}

		err = recommended.SaveRecommended()

		if err != nil{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"`+err.Error()+`"}`))
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