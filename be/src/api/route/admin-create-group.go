package route

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	// middleware package
	"backend/api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

type Groupnames struct {
	// json tag to de-serialize json body
	Name string `json:"user_group"`
}

func AdminCreateGroup(context *gin.Context) {
	var newGroup Groupnames

	// call BindJSON to bind the received JSON to newGroup
	if err := context.BindJSON(&newGroup); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(context, 200, "Bad Request")
		return
	}

	// remove spaces in groupname
	groupname := strings.TrimSpace(newGroup.Name)

	// check if groupname field is empty
	minLength := middleware.CheckLength(groupname)
	if minLength {
		middleware.ErrorHandler(context, 200, "Groupname should not be empty")
		return
	}

	// check for existing groupname before creating
	checkGroupname := "SELECT user_group FROM groupnames WHERE user_group = ?"

	// return first result (single row result)
	result := db.QueryRow(checkGroupname, groupname)

	// Scan: scanning and reading input (texts given in standard input)
	switch err := result.Scan(&groupname); err {

	// New Group
	case sql.ErrNoRows:
		// insert new group
		_, err := db.Exec("INSERT INTO Groupnames (user_group) VALUES (?)", groupname)

		if err != nil {
			fmt.Println(err)
			middleware.ErrorHandler(context, 200, "Unable to create new group")
			return
		}

		context.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "New group has created successfully"})
		return

	// Existing groupname
	case nil:
		middleware.ErrorHandler(context, 200, "Existing Groupname")
		return

	// Invalid Field
	default:
		middleware.ErrorHandler(context, 200, "Invalid Field")
		return
	}
}
