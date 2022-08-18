package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const templateText = `
Output 0:{{title.Name1}} 
Output 1:{{title.Name2}}
Output 2:{{.Name3 | title }}
`

func main() {
	funcMap := template.FuncMap{"title": strings.Title}
	tpl, _ := template.New("go_tour").Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "c",
		"Name2": "java",
		"Name3": "go",
	}
	_ = tpl.Execute(os.Stdout, data)
	fmt.Println("\ntemplate exec over...")
}
