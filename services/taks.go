package tasks
import (
	"fmt"
	""
)
var id = 0
func createTask(title string, description string) model.Tache {
	newTask := model.Tache{
		ID:          id+1, 
		Title:       title,
		Description: description,
	}
	fmt.Println("Task created:", newTask)
	id ++ 
	return newTask
}