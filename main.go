package main

import (
	"WebProject/handler"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	handler.AddData()

	router.HandleFunc("/people",handler.CreatePerson).Methods("POST")
	router.HandleFunc("/people/all", handler.GetAllPeople).Methods("GET")
	router.HandleFunc("/people/find/{id}", handler.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", handler.UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{id}", handler.DeletePerson).Methods("DELETE")

	//corsObj := handlers.AllowedOrigins([]string{"*"})


	/*	cors := cors.New(cors.Options{
		AllowedOrigins:         []string{"*"},
		AllowOriginRequestFunc: func(r *http.Request, origin string) bool { return true },
		AllowedMethods:         []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:         []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:         []string{"Link"},
		AllowCredentials:       true,
		OptionsPassthrough:     true,
		MaxAge:                 3599, // Maximum value not ignored by any of major browsers
	})*/

	//router.Use(cors.Handler)
	// cors解决跨域
	//1.浏览器从一个域名的网页去请求另一个域名的资源时，域名、端口、协议任一不同，都是跨域
	//2.只要协议、域名、端口有任何一个不同，都被当作是不同的域，之间的请求就是跨域操作。

	//headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	//methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	//log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(router)))


	fmt.Println("Starting server on port 5000...")
	//http.ListenAndServe(":5000" , router)
	http.ListenAndServe(":5000",
		handlers.CORS(handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "HEAD","DELETE","PUT"}),
			handlers.AllowedHeaders([]string{"Content-Type"}))(router))

}