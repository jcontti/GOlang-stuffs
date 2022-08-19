package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// function to connect with our Database, that will return a connection type  *sql.DB
func connectDB() (connection *sql.DB) {
	Driver := "mysql"
	User := "root"
	Pass := ""
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

	log.Println("Server is running...")

	// Starts the server ( http://localhost:8080/ )8080 indicate the server port
	http.ListenAndServe(":8080", nil)
}

// When we call our function with HandleFunc we will have 2 parameters of reception
func Home(w http.ResponseWriter, r *http.Request) {
	// with the first parameter I can give a response, with the second I will receive all the response, can be Post method, Get method, etc.
	// fmt.Fprintf(w, "Hi IT people") // we will write to the "ResponseWriter" which means our - internet browser
	// fmt.Fprintf(w, "\nThis is a new line")

	// use the database connection:
	connEstablished := connectDB()
	// insert new values to the DB
	insertReg, err := connEstablished.Prepare("INSERT INTO employees(name,email) VALUES('oscar','oscarcito@gmail.com') ")
	if err != nil {
		panic(err.Error())
	}
	insertReg.Exec()

	// show our template in the internet browser
	myTemplate.ExecuteTemplate(w, "home", nil)

}

func Create(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "create", nil) // call the template "create.html"
}

/* NOTES:
- we will need to download a mysql driver to use the database: go get -u github.com/go-sql-driver/mysql
   then we need to import that driver to our project:  _"github.com/go-sql-driver/mysql"
*/
