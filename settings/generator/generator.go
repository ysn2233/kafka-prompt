package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	generate(os.Args[1], os.Args[2])
}

func generate(settings string, filename string) {
	config := NewConfig(settings)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	f.Truncate(0)
	f.WriteString("package kaprompt\n\n")
	f.WriteString("import \"github.com/c-bata/go-prompt\"\n\n")
	f.WriteString("var(\n")
	f.WriteString("\tcommandSuggest = []prompt.Suggest{\n")

	for _, s := range config.CommandsSuggest {
		f.WriteString(fmt.Sprintf("\t\t{Text: \"%s\", Description: \"%s\"},\n", s.Text, s.Description))
	}
	f.WriteString("\t}\n")

	f.WriteString("\toptionSuggest = map[string][]prompt.Suggest{\n")
	for c, o := range config.OptionSuggest {
		f.WriteString(fmt.Sprintf("\t\t\"%s\": {\n", c))
		for _, s := range o {
			f.WriteString(fmt.Sprintf("\t\t\t{Text: \"%s\", Description: \"%s\"},\n", s.Text, s.Description))
		}
		f.WriteString("\t\t},\n")
	}
	f.WriteString("\t}\n")

	f.WriteString("\tbsCommand = map[string]bool{\n")
	for k, v := range config.BsCommand {
		f.WriteString(fmt.Sprintf("\t\t\"%s\": %v,\n", k, v))
	}
	f.WriteString("\t}\n")

	f.WriteString("\tblCommand = map[string]bool{\n")
	for k, v := range config.BlCommand {
		f.WriteString(fmt.Sprintf("\t\t\"%s\": %v,\n", k, v))
	}
	f.WriteString("\t}\n")

	f.WriteString("\tzkCommand = map[string]bool{\n")
	for k, v := range config.ZkCommand {
		f.WriteString(fmt.Sprintf("\t\t\"%s\": %v,\n", k, v))
	}
	f.WriteString("\t}\n")

	f.WriteString(")")
}
