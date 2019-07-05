package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/ysn2233/kafka-prompt/kaprompt"
)

func main() {
	cmptr := kaprompt.NewCompleter()
	es := kaprompt.NewExecutorSession()
	// fmt.Println(config)
	fmt.Println("Welcome to kafka-promot!")
	fmt.Println("Please use ctrl + d / exit / quit to exit the program")
	p := prompt.New(
		es.Executor,
		cmptr.Complete,
		kaprompt.CustomOptions()...,
	)

	p.Run()
}
