package pkg
import(
	"github.com/ilyes-rhdi/to-do-list/models"
	"container/list"

)
func Getid(MS models.Ms) int {
    var id int

    if MS.Type == "array" {
        tasks := MS.Tasks.([]models.Tache)
        id = tasks[len(tasks)-1].ID
    } else {

        tasks := MS.Tasks.(*list.List)
        last := tasks.Back()
        if last != nil {
            id = last.Value.(models.Tache).ID
        }
    }

    return id
}
