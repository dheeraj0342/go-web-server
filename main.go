package main

import(
	"fmt"
	"log"
	"net/http"
)
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path !="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	} 
	if r.Method != "GET"{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w,"Hello World")
}
func formHandler(w http.ResponseWriter, r * http.Request){
	if err := r.ParseForm(); err != nil{
		http.Error(w,"Cannot parse form",http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w,"post req successful")
	name := r.FormValue("name")
	email:= r.FormValue("email")
	password := r.FormValue("password")
	fmt.Fprintf(w,"Name: %s\n",name)
	fmt.Fprintf(w,"Email: %s\n",email)
	fmt.Fprintf(w,"Password: %s\n",password)
}
func main(){

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/form",formHandler)
	fmt.Println("Starting server on port :3000")
	if err := http.ListenAndServe(":3000",nil);err != nil{
		log.Fatal(err)
	}
}