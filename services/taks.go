package tasks

import (
	"fmt"
	"github.com/ilyes-rhdi/to-do-list/models"
	"github.com/ilyes-rhdi/to-do-list/pkg"
	"container/list"
)
const maxSize = 100000
const sizeBlock = 100
var id = 0
func CreateTask(MS models.Ms,title string, description string,size int ) models.Tache {
	if maxSize/sizeBlock < (MS.nbBlock + (size / sizeBlock) + 1) {
		fmt.Println("Max size reached")
	}
	newTask := models.Tache{
		ID:          pkg.Getid(), 
		Title:       title,
		Description: description,
	}
	fmt.Println("Task created:", newTask)
	id ++ 
	return newTask
}
func DeleteLastTask(){

}
func InitMS(MS int, array bool, size int) {
	var tasks any
	var t string

	if array {
		tasks = make([]models.Tache, 0, size) // slice vide avec capacité "size"
		t = "array"
	} else {
		l := list.New()
		tasks = l
		t = "list"
	}

	ms := models.Ms{
		Type:    t,
		Size:    size,
		Tasks:   tasks,
		nbBlock: 0,
	}

	fmt.Println("MS initialized:", ms)
}
