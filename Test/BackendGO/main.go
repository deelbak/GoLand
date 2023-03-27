package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)
type User struct{
	Name string
	Age uint16
	Balans int16
	Rating int16
}

var database *sql.DB



func (u User) getAllInfo() string {

	con := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", "localhost", 5432, "pp2example_user", "passw0rd", "pp2example")
	db, err := sql.Open("postgres", con)
	database = db
	if err != nil{
		panic(err)
	}
	fmt.Println("Connected to db")
	rows, err := database.Query("select * from newUsers")
	
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
	
}

func login() func(){
	fmt.Println("1)login\n" , "2)register\n")
	var s string
	fmt.Scan(&s)
	if(s == "1"){
		fmt.Println("Username:")
		var name string 
		fmt.Scan(&name)
		row := database.QueryRow("select name, age from newUsers where name = ?", name)
		if row != nil{
			fmt.Println("Password:")
			var pass string
			fmt.Scan(&pass)
			if(row.age == pass){

			}
		}else{
			fmt.Println("You must register!")
			return login()
		}
	}
}