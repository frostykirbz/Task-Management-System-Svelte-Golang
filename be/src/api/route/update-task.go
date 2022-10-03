package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TO BE DONE BY BEATRICE
// Update Task (with/without plan)
// Validation:
// - check task plan - if there is selected plan (check plan color - update color based on selected plan [if plan is empty, update plan color to empty string])
// - check task notes if there is task notes (insert in tasknotes table and update tasknotes in task table)/there is no task notes (dont need insert in tasknotes table and update tasknotes in task table)
func UpdateTask(c *gin.Context) {
	var task models.Task
	var AppAcronym string = c.Query("AppAcronym")

	// call BindJSON to bind the received JSON to task
	if err := c.BindJSON(&task); err != nil {
		checkError(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	// check if taskname exists
	result := middleware.SelectTaskName(task.TaskName, AppAcronym)
	fmt.Println(task.TaskAppAcronym)
	fmt.Println(task.TaskName)

	switch err := result.Scan(&task.TaskName); err {

	// task name does not exist
	case sql.ErrNoRows:
		middleware.ErrorHandler(c, 400, "Task name does not exist")
		return

	// task name exists
	case nil:
		// plan color check
		checkTaskPlan(task, c)

		// task notes check
		checkTaskNotes(task, c)

		// update task with/without plan

		c.JSON(200, gin.H{"code": 200, "message": "Exists"})
		return
	}
}

// func validatePermit(task models.Task, c *gin.Context) {
// 	// var PermitOpen sql.NullString
// 	// var PermitToDo sql.NullString
// 	// var PermitDoing sql.NullString
// 	// var PermitDone sql.NullString

// 	// select query to store all the user under these permits (string/array)
// 	// for loop to loop through if user is within any of the permit (checkgroup)
// }

// check if there is plan color
func checkTaskPlan(task models.Task, c *gin.Context) {
	var PlanColor sql.NullString

	result := middleware.SelectPlanColor(task.TaskPlan)

	switch err := result.Scan(&PlanColor); err {
	case sql.ErrNoRows:
		task.TaskColor = ""
		return

	case nil:
		task.TaskColor = PlanColor.String
		return
	}
}

// check if there is task notes
func checkTaskNotes(task models.Task, c *gin.Context) {
	var TaskNotes, TaskNotesDate, TaskNotesTime, TaskOwner, TaskState sql.NullString
	var taskNotesAuditString string

	// new task notes
	if !middleware.CheckLength(task.TaskNotes) {

		var selectTaskQuery = "SELECT task_owner, task_state FROM task WHERE task_name = ? AND task_app_acronym = ?"

		result := db.QueryRow(selectTaskQuery, task.TaskName, task.TaskAppAcronym)

		switch err := result.Scan(&TaskOwner, &TaskState); err {

		case sql.ErrNoRows:
			middleware.ErrorHandler(c, 400, "Task owner/state does not exist")
			return

		case nil:
			task.TaskOwner = TaskOwner.String
			task.TaskState = TaskState.String

			// insert row into task notes table
			_, err = middleware.InsertCreateTaskNotes(task.TaskName, task.TaskNotes, task.TaskOwner, task.TaskState)

			if err != nil {
				panic(err)
			}

			// format new task notes
			// concat with existing task notes
			var tasknotesTimestamp = `SELECT DATE_FORMAT(last_updated, "%d/%m/%Y") as formattedDate, TIME_FORMAT(last_updated, "%H:%i:%s") as formattedTime, task_note, task_owner, task_state FROM task_notes WHERE task_name = ? ORDER BY last_updated DESC;`
			rows, err := db.Query(tasknotesTimestamp, task.TaskName)
			if err != nil {
				log.Fatal(err)
			}

			for rows.Next() {
				if err := rows.Scan(&TaskNotesDate, &TaskNotesTime, &TaskNotes, &TaskOwner, &TaskState); err != nil {
					log.Fatal(err)
				}

				taskNotesAuditString += TaskNotesDate.String + " " + TaskNotesTime.String + "\n" + "Task Owner: " + TaskOwner.String + ", Task State: " + TaskState.String + "\n" + TaskNotes.String + " \n\n"
			}

			task.TaskNotes = taskNotesAuditString
			fmt.Println(task.TaskNotes)
			updateTaskTable(task)
		}
	} else {
		// update task
		updateTaskTable(task)
	}
}

// check if there is a plan
// 1. yes: insert with plan (plan name)
// 2. no: insert without plan (null)
func updateTaskTable(task models.Task) {
	var TaskPlan *string = nil

	var updateTask = "UPDATE task SET task_notes = ?, task_plan = ?, task_color = ?,  task_owner = ? WHERE task_name = ? AND task_app_acronym = ?"

	if !middleware.CheckLength(task.TaskPlan) {
		_, err := db.Exec(updateTask, task.TaskNotes, task.TaskPlan, task.TaskColor, task.TaskOwner, task.TaskName, task.TaskAppAcronym)
		if err != nil {
			panic(err)
		}
	} else {
		_, err := db.Exec(updateTask, task.TaskNotes, TaskPlan, task.TaskColor, task.TaskOwner, task.TaskName, task.TaskAppAcronym)
		if err != nil {
			panic(err)
		}
	}
}
