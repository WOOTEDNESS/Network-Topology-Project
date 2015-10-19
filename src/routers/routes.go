package routers

import (
	"controllers"
	"net/http"
)

func Init() {
  	//testing out Go's file server stuff 
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/", fs)
	
	
	// routes the homepage to the browser
	http.HandleFunc("/home",controllers.HomeController)

} 