package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// function to connect with our Database, that will return a connection type  *sql.DB
func connectDB() (connection *sql.DB) {
	Driver := "mysql"
	User := "root"
	Pass := "rootroot"
	dbName := "my_system_db"

	// if the connection fails with an error, that wouuld be saved in "err"
	connection, err := sql.Open(Driver, User+":"+Pass+"@tcp(127.0.0.1)/"+dbName)
	if err != nil {
		// if an error exist, show me the error.
		panic(err.Error())
	}
	return connection
}

// We will use templates
var myTemplate = template.Must(template.ParseGlob("templates/*")) // this help us to search information in folders about our templates
// "templates/*" means let me look into all files inside templates folder

func main() {

	http.HandleFunc("/", Home)         // "Home is the name of our function which will be call"
	http.HandleFunc("/create", Create) // when the user goes to /create url, call the function "Create"
	http.HandleFunc("/insert", Insert) // our create form send info to our insert page/function
	http.HandleFunc("/delete", Delete) // delete rows from table
	http.HandleFunc("/edit", Edit)     // edit values
	http.HandleFunc("/update", Update) // update values

	log.Println("Server is running...")

	// Starts the server ( http://localhost:8080/ )8080 indicate the server port
	http.ListenAndServe(":8080", nil)
}

// We need to create a struct to hold the select of the employees
type Employee struct {
	Id    int
	Name  string
	Email string
}

// When we call our function with HandleFunc we will have 2 parameters of reception
func Home(w http.ResponseWriter, r *http.Request) {
	// with the first parameter I can give a response, with the second I will receive all the response, can be Post method, Get method, etc.
	// fmt.Fprintf(w, "Hi IT people") // we will write to the "ResponseWriter" which means our - internet browser
	// fmt.Fprintf(w, "\nThis is a new line")

	// use the database connection:
	connEstablished := connectDB()
	// Get values from the DB
	getReg, err := connEstablished.Query("SELECT * FROM employees")
	if err != nil {
		panic(err.Error())
	}

	employee := Employee{}
	arrayEmployee := []Employee{}

	// reading the values of the select and assigning them to variables
	for getReg.Next() {
		var id int
		var name, email string
		err = getReg.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}

		employee.Id = id
		employee.Name = name
		employee.Email = email

		arrayEmployee = append(arrayEmployee, employee)
	}

	// show our template in the internet browser, filled with the table Employees, we send our array to be consumed by the home.html
	myTemplate.ExecuteTemplate(w, "home", arrayEmployee)

}

func Create(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "create", nil) // call the template "create.html"
}

// this function will received the data from the POST method of the create function
func Insert(w http.ResponseWriter, r *http.Request) {
	// if there's POST data, so we'll catch them
	if r.Method == "POST" {
		name := r.FormValue("name") // here we have to put the id/name from the form "create"
		email := r.FormValue("email")

		connEstablished := connectDB()                                                              // use the database connection:
		insertReg, err := connEstablished.Prepare("INSERT INTO employees(name,email) VALUES(?,?) ") // insert new values to the DB

		if err != nil {
			panic(err.Error())
		}

		insertReg.Exec(name, email) // the Exec, will pass that 2 parameters values to the VALUES (?,?) sql syntax...in the INSERT

		// redirect to home
		http.Redirect(w, r, "/", 301)
	}
}

// Function to delete rows from the table employees
func Delete(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id") //  get the request that is passed as GET which is ID

	connEstablished := connectDB()
	delReg, err := connEstablished.Prepare("DELETE FROM employees WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	delReg.Exec(idEmployee)

	http.Redirect(w, r, "/", 301)

}

// Function to delete rows from the table employees
func Edit(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")

	connEstablished := connectDB()
	editReg, err := connEstablished.Query("SELECT * FROM employees WHERE id=?", idEmployee)

	if err != nil {
		panic(err.Error())
	}

	employee := Employee{}

	for editReg.Next() {
		var id int
		var name, email string
		err = editReg.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}

		employee.Id = id
		employee.Name = name
		employee.Email = email
	}

	fmt.Println(employee)

	myTemplate.ExecuteTemplate(w, "edit", employee) // call the template "edit.html" and send the employee array

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")

		connEstablished := connectDB()
		UpdateReg, err := connEstablished.Prepare("UPDATE employees SET name=?,email=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		UpdateReg.Exec(name, email, id)

		http.Redirect(w, r, "/", 301)
	}
}

/* NOTES:
- we will need to download a mysql driver to use the database: go get -u github.com/go-sql-driver/mysql
   then we need to import that driver to our project:  _"github.com/go-sql-driver/mysql"
*/
