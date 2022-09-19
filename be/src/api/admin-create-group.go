package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	// middleware package
	"api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

type Groupnames struct {
	Name string `json:"user_group"`
}

// createGroup adds the specified group to database
// func (grp Groupnames) createGroup(res http.ResponseWriter, req *http.Request) {
// 	// var newGroup Groupnames

// 	// // check if groupname exist before creating
// 	// checkGroupname := "SELECT user_group FROM groupnames WHERE user_group = ?"

// 	// // get result
// 	// result := db.QueryRow(checkGroupname, newGroup.Name)

// 	_, err := db.Exec("INSERT INTO groupnames (user_group) VALUES (?)", grp.Name)
// 	check(err)
// 	fmt.Println("Inserted Successfully")
// }

func createGroup(c *gin.Context) {
	var newGroup Groupnames

	// call BindJSON to bind the received JSON to newGroup
	if err := c.BindJSON(&newGroup); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	// MySQL database connection
	db, err := sql.Open("mysql", "root:paassword@/c3_database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// insert newGroup
	// _, err := db.Exec("INSERT INTO groupnames (user_group) VALUES (?)", newGroup.Name)

	// if err != nil {
	// 	fmt.Println(err)
	// 	middleware.ErrorHandler(c, http.StatusBadRequest, "Unabled to create new group")
	// 	return
	// }

	// fmt.Println("Inserted Successfully")
	c.IndentedJSON(http.StatusCreated, gin.H{"code": 200, "message": "New group has created successfully"})
}

func main() {
	// Get a database handle
	// var err error
	// db, err = sql.Open("mysql", "root:password@/c3_database")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer db.Close()

	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	fmt.Println(pingErr)
	// 	return
	// }

	// fmt.Println("Database connected!")

	// http.HandleFunc("/admin-create-group", createGroup)

	// log.Fatal(http.ListenAndServe(":4000", nil))

	router := gin.Default()
	router.POST("/admin-create-group", createGroup)

	router.Run("localhost:4000")
}
