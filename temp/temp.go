package temp

import (
	"fmt"
	"os"
)

func Temp() {
	directory, _ := os.Getwd()
	fmt.Println(directory)
}
