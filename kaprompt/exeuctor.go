package kaprompt

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type ExecutorSession struct {
	BsServer  string
	Zookeeper string
	BkList    string
}

func NewExecutorSession() ExecutorSession {
	return ExecutorSession{}
}

func (es *ExecutorSession) Executor(cmd string) {
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return
	} else if cmd == "exit" || cmd == "quit" {
		fmt.Println("Bye-bye~!")
		os.Exit(0)
		return
	} else {
		command := exec.Command("/bin/sh", "-c", cmd)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err := command.Run(); err != nil {
			log.Printf("Error occurs: %s\n", err.Error())
		} else {
			re := regexp.MustCompile(`--bootstrap-server (.*:[0-9]*)`)
			group := re.FindStringSubmatch(cmd)
			if len(group) == 2 {
				es.BsServer = group[1]
			}

			re = regexp.MustCompile(`--zookeeper (.*:[0-9]*)`)
			group = re.FindStringSubmatch(cmd)
			if len(group) == 2 {
				es.Zookeeper = group[1]
			}

			re = regexp.MustCompile(`--broker-list (.*:[0-9]*)`)
			group = re.FindStringSubmatch(cmd)
			if len(group) == 2 {
				es.BkList = group[1]
			}
		}

		err := Persist(cmd)
		if err != nil {
			log.Panic(err)
		}
	}
}
