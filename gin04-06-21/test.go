package main
import
(
	"testing"
)

func init(){

}

func main(){

}

func Search(searchInput string) string{

}
func TestSearchWithNoCommon(t *testing.T){
	searchInput := "asdasd"
	searchOutput := Search(searchInput)
	if searchInput != searchOutput{
		t.Error("Searching for ",searchInput," didn't get the result of ",searchOutput)
	}
}
