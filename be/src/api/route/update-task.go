package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"fmt"
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
	//var PlanColor sql.NullString
	var AppAcronym string = c.Query("AppAcronym")

	// call BindJSON to bind the received JSON to task
	if err := c.BindJSON(&task); err != nil {
		checkError(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	// check if taskname exists
	result := middleware.SelectTaskName(task.TaskName, AppAcronym)
	fmt.Println(AppAcronym)
	fmt.Println(task.TaskName)

	switch err := result.Scan(&task.TaskName); err {
	
	// task name does not exist
	case sql.ErrNoRows:
		c.JSON(200, gin.H{"code": 200, "message": "Does not exist"})
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
	var TaskNotes sql.NullString
	var AppAcronym = c.Query("AppAcronym")
	var taskNotesAuditString string

	var selectTaskNotesQuery = "SELECT task_notes FROM task WHERE task_name = ? AND task_app_acronym = ?"

	// new task notes 
	if (!middleware.CheckLength(task.TaskNotes)) {
		var TaskNotesDate, TaskNotesTime sql.NullString

		result := db.QueryRow(selectTaskNotesQuery, task.TaskName, AppAcronym)

		switch err := result.Scan(&TaskNotes); err {

		// no existing task notes
		case sql.ErrNoRows:
			// insert row into task notes table
			_, err = middleware.InsertCreateTaskNotes(task.TaskName, task.TaskNotes, task.TaskOwner, task.TaskState)

			if (err != nil) {
				panic(err)
			}
			
			// format new task notes  
			rows := middleware.SelectTaskNotesTimestamp(task.TaskName) 
			rows.Scan(&TaskNotesDate, &TaskNotesTime) 
			taskNotesAuditString = TaskNotesDate.String + " " + TaskNotesTime.String + "\n" + "Task Owner: " + task.TaskOwner + ", Task State: " + task.TaskState  + "\n" + task.TaskNotes + " \n";                                                     
			task.TaskNotes = taskNotesAuditString
			return 
			
		// existing task notes
		case nil:
			// insert row into task notes table
			_, err = middleware.InsertCreateTaskNotes(task.TaskName, task.TaskNotes, task.TaskOwner, task.TaskState)

			if (err != nil) {
				panic(err)
			}

			// format new task notes
			// concat with existing task notes
			rows := middleware.SelectTaskNotesTimestamp(task.TaskName) 
			rows.Scan(&TaskNotesDate, &TaskNotesTime) 
			taskNotesAuditString = TaskNotesDate.String + " " + TaskNotesTime.String + "\n" + "Task Owner: " + task.TaskOwner + ", Task State: " + task.TaskState  + "\n" + task.TaskNotes + " \n";                                                     
			task.TaskNotes = taskNotesAuditString + TaskNotes.String
			return 
		}
	}
}

// insert task notes
func InsertTaskNotes(task models.Task, c *gin.Context) {
	// do all the checking
	// insert task notes
}

// check if there is a plan
// 1. yes: insert with plan (plan name)
// 2. no: insert without plan (null)
func updateTaskTable(task models.Task) {
	var TaskPlan *string = nil

	if !middleware.CheckLength(task.TaskPlan) {
		_, err := middleware.InsertTask(task.TaskAppAcronym, task.TaskID, task.TaskName, task.TaskDescription, task.TaskNotes, task.TaskPlan, task.TaskColor, task.TaskState, task.TaskCreator, task.TaskOwner)
		checkError(err)
	} else {
		_, err := middleware.InsertTaskWithoutPlan(task.TaskAppAcronym, task.TaskID, task.TaskName, task.TaskDescription, task.TaskNotes, TaskPlan, task.TaskColor, task.TaskState, task.TaskCreator, task.TaskOwner)
		checkError(err)
	}
}