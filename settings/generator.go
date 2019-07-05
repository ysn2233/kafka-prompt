package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ysn2233/kafka-prompt/kaprompt"
)

func main() {
	generate("../kaprompt/suggests.go")
}

func generate(filename string) {
	config := kaprompt.NewConfig("kafka-command.json")
	f, err := os.OpenFile(filename, os.O_WRONLY, 0644)
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

	f.WriteString(")")
}
