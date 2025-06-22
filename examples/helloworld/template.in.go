{{/*
Go "Hello, World!" template

Expected Data:
- name: string
*/}}
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, {{.name}}!")
}
