package main

import (
	"fmt"
	"net/http"
	"github.com/Mikhailli/WWW"
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	//HandleRequest()
	var serverName = "Server=localhost;database=CloudioPlayer"
	db, err := sql.Open("mssql", serverName)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	var t = 2
	all, err := db.Query("SELECT TOP 1 Id FROM Audios")
	if err != nil {
		panic(err)
	}
	for all.Next() {
		all.Scan(&t)
		fmt.Print(t)
	}
}

func HandleRequest() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contacts/", ContactsPage)
	http.ListenAndServe(":8080", nil)
}

func HomePage(page http.ResponseWriter, request *http.Request) {

	//var bob = User{"Bob", 25, -50, 4.2, 0.8, []string {"football", "skate"}}

	//template, _ := template.ParseFiles("templates/HomePage.html")
	//template.Execute(page, bob)

	db, err := sql.Open("mssql", "server=localhost;database=CloudioPlayer")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Print("Подключено к MsSql")
}

func ContactsPage(page http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(page, "Contacts page")
}

func (user User) GetAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money"+
		"equal: %d", user.Name, user.Age, user.Money)
}

func (user *User) SetNewName(newName string) {
	user.Name = newName
}

type User struct {
	Name                     string
	Age                      uint16
	Money                    int16
	AverageGrades, Happiness float64
	Hobbies                  []string
}
