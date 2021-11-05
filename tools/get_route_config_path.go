package tools

import (
	"fmt"
	"os"
)

func LoadConfig() {
	path, _ := os.Getwd()
	fmt.Println(path)
}
