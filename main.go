package main

import (
	"fmt"
	"net/http"

	authcontroller "ftdmt/controllers"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/addmultitrx", authcontroller.IndexMulti)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/auth/google", authcontroller.Logingoogle)
	http.HandleFunc("/auth/google/callback", authcontroller.Callbackgoogle)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)

	fmt.Println("Server jalan di: http://localhost:8899")
	http.ListenAndServe(":8899", nil)
}
