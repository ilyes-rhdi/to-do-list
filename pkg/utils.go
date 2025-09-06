package pkg
import(
	"github.com/ilyes-rhdi/to-do-list/models"

)
func Getid(MS models.Ms) int {
	if Ms.Type == "array" {
		id = MS.Tasks[len(MS.Tasks)-1].ID
	}else{
		id = MS.Tasks.Back().Value.ID
	}	
	return id
}