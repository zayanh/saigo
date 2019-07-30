package main
import (
	"fmt"
	"os"
	c "github.com/zayanh/saigo/exercise-001-corpus/count"
)

func main() {
	// define command line parameter flags
	params := os.Args
	if len(params) < 2 {
		fmt.Println("no input file provided. Exiting")
		return
	}
	filename := params[1]
	//TODO: Make sure path is valid

	fmt.Println(c.Count(filename))
}
