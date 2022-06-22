package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "01/02/2006"
	t, _ := time.Parse(layout, layout)
	fmt.Print(t.Year())
}
