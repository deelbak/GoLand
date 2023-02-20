package main

import (
	"fmt"
	"net/http"
	"html/template"
)
type User struct{
	Name string
	Age uint16
	Balans int16
	Rating float64
	Listofbooks []string 
}


func (u User) getAllInfo() string {
	return fmt.Sprintf("<b>UserName is: %s. He is %d and his Balans :%d $.</b>", u.Name, u.Age, u.Balans)
}

func(u *User) setNewBalans(newBalans uint16){
	u.Balans = int16(newBalans)
}

func home_page(w http.ResponseWriter, r *http.Request){
	bob:=User{"Bob", 20, 0, 4.50, []string{"Harry Potter", "Internet psyhology", "Algorithms and data structures"}}
	bob.setNewBalans(1785)
	// fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}
func contacts_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<b>Contats</b>")
}
func handlRequest(){
	fmt.Println("proccess:")
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":1010", nil)
}
func main(){
	handlRequest()
}