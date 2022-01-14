package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s  He is %d and he has money equal %d", u.name,
		u.age, u.money)
}
func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Boob", 25, -50, 4.2, 0.7}
	bob.name = "volodymyr"
	fmt.Fprintf(w, "username is:"+bob.name)
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
