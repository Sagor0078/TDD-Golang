package ReadingFiles

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("testdata/input.txt")
	if err != nil {
		panic(err)
	}
	output := Process(data)
	fmt.Println(string(output))
}
