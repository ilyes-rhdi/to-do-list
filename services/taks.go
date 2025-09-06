package tasks

import (
	"fmt"
	"github.com/ilyes-rhdi/to-do-list/models"
	"container/list"
)
const maxSize = 100000
const sizeBlock = 100
var id = 0
func CreateTask(title string, description string) models.Tache {
	newTask := models.Tache{
		ID:          utils.Getid(), 
		Title:       title,
		Description: description,
	}
	fmt.Println("Task created:", newTask)
	id ++ 
	return newTask
}
func DeleteLastTask(){

}
func InitMS(MS int,array bool){
	if array {
      tasks := make([]models.Tache, 0)
	}else{
      tasks := list.New()
	  tasks.Init()
	}
 
}