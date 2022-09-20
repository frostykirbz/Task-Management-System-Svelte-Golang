package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// middleware package
	"api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

 var db *sql.DB

type Groupnames struct {
	// json tag to de-serialize json body
	Name string `json:"user_group"`
}

func createGroup(context *gin.Context) {
	var newGroup Groupnames
	var groupname string

	// call BindJSON to bind the received JSON to newGroup
	if err := context.BindJSON(&newGroup); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(context, http.StatusBadRequest, "Bad Request")
		return
	}

	// check if groupname field has whitespace
	whiteSpace := middleware.CheckWhiteSpace(newGroup.Name)
	if whiteSpace == true {
		middleware.ErrorHandler(context, 400, "Groupname should not contain whitespace")
		return
	}

	// check if groupname field is empty
	minLength := middleware.CheckLength(newGroup.Name)
	if minLength == true {
		middleware.ErrorHandler(context, 400, "Groupname should not be empty")
		return
	}

	// check for existing groupname before creating
	checkGroupname := "SELECT user_group FROM groupnames WHERE user_group = ?"

	// return first result (single row result)
	result := db.QueryRow(checkGroupname, newGroup.Name)

	// Scan: scanning and reading input (texts given in standard input)
	switch err := result.Scan(&groupname); err {
		
	// New Group
	case sql.ErrNoRows:
		// insert new group
		_, err := db.Exec("INSERT INTO Groupnames (user_group) VALUES (?)", newGroup.Name)

		if err != nil {
			fmt.Println(err)
			middleware.ErrorHandler(context, http.StatusBadRequest, "Unable to create new group")
			return
		}

		fmt.Println(newGroup)
		context.IndentedJSON(http.StatusCreated, gin.H{"code": 200, "message": "New group has created successfully"})

	// Existing groupname
	case nil: 
		middleware.ErrorHandler(context, http.StatusBadRequest, "Existing Groupname")
	
	// Invalid Field
	default:
		middleware.ErrorHandler(context, http.StatusBadRequest, "Invalid Field")
	}
}

func main() {
	connectionToMySQL()
	defer db.Close()

	fmt.Println("Database connected!")

	// http.HandleFunc("/admin-create-group", createGroup)

	// log.Fatal(http.ListenAndServe(":4000", nil))

	router := gin.Default()
	router.POST("/admin-create-group", createGroup)
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "TEST",
		})
	})
	router.Run("localhost:4000")
}

func connectionToMySQL(){
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", "root:password@/c3_database")
	if err != nil {
		fmt.Println(err)
		return
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
		return
	}
}