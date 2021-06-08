package main
import
(
	"testing"
)

func init(){
	//connect http tp application
}

func main(){
	//call other test function and print out the result
}

func Search(searchInput string) string{
	//send http to search function and input the "searchInput" and return the search output as string
}

func TestSearchWithNoCommon(t *testing.T){
	searchInput := "asdasd"
	searchOutput := Search(searchInput)
	ExpectedResult := "No result"
	if searchIOutput != ExpectedResult{
		t.Error("Searching for ",searchInput," for no common didn't get the result of ",searchOutput)
	}
}

func TestSearchWithNoCommon(t *testing.T){
	searchInput := "su"
	searchOutput := Search(searchInput)
	ExpectedResult := ""
	if searchIOutput != ExpectedResult{
		t.Error("Searching for ",searchInput," for common didn't get the result of ",searchOutput)
	}
}

