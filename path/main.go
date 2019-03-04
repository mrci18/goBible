package main

import (
	"fmt"
	"path"
)

func main() {
	dir, file := path.Split("css/my.css")
	fmt.Println("Dir:", dir)
	fmt.Println("File", file)
}
