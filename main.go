package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s  He is %d and he has money"+
		"equal %d", u.name,
		u.age, u.money)
}
func (u *User) setNewName(newName string) { //силку передаю за допомогою *
	u.name = newName
}
func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Boob", 25, -50, 4.2, 0.7}
	// bob.setNewName("andrew")
	// fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}
func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contacts page")
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)

	http.ListenAndServe(":8080", nil)
}
func main() {
	// var bob User = {	// }
	// bob := User{name:"bob",	age: 25,money: -50,avg_grades: 4.2,happiness: 0.5}

	handleRequest()

}
