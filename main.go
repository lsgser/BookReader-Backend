package main

import(
	"log"
	"os"
	"github.com/joho/godotenv"
	"net/http"
	"github.com/rs/cors"
	R "./routes"
)

func init() {
	godotenv.Load()
}

func main(){
	log.Println("Running on port :",os.Getenv("PORT"))
	router := R.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		//AllowCredentials:true,
		//Debug: true,
		AllowedMethods: []string{"GET","POST","OPTIONS","DELETE","PUT","PATCH"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token,Authorization"},
	})
	//log.Println(c)
	//handler := cors.Default().Handler(router)
	handler := c.Handler(router)
	log.Fatalln(http.ListenAndServe(":"+os.Getenv("PORT"),handler))	
}