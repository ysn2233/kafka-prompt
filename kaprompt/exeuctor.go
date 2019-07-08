package kaprompt

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

type ExecutorSession struct {
	BsServer  string
	Zookeeper string
	BkList    string
	mu        sync.Mutex
}

func NewExecutorSession() ExecutorSession {
	return ExecutorSession{}
}

func (es *ExecutorSession) Executor(cmd string) {
	cmd = strings.TrimSpace(cmd)
	savedCmd := cmd
	if cmd == "" {
		return
	} else if cmd == "exit" || cmd == "quit" {
		fmt.Println("Bye-bye~!")
		os.Exit(0)
		return
	} else {
		fmt.Println(es)
		args := regexp.MustCompile("\\s+").Split(cmd, -1)
		executable := args[0]
		_, err := exec.LookPath(executable)
		if err != nil {
			args[0] = args[0] + ".sh"
			cmd = strings.Replace(cmd, executable, executable+".sh", 1)
		}

		cmd = es.addServerFlag(args, executable, cmd)

		command := exec.Command("/bin/sh", "-c", cmd)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err := command.Run(); err != nil {
			log.Printf("Error occurs: %s\n", err.Error())
		} else {
			go es.updateServers(cmd)
		}

		go Persist(savedCmd)
	}
}

func (es *ExecutorSession) addServerFlag(args []string, executable string, cmd string) string {
	relativeExe := strings.Split(executable, "/")
	trimedExe := strings.ReplaceAll(relativeExe[len(relativeExe)-1], ".sh", "")

	var commandSlice []string
	if _, ok := bsCommand[trimedExe]; !strings.Contains(cmd, "--bootstrap-server") && ok {
		commandSlice = append(commandSlice, args[0], "--bootstrap-server", es.BsServer)
		cmd = strings.Join(append(commandSlice, args[1:]...), " ")
	} else if _, ok := blCommand[trimedExe]; !strings.Contains(cmd, "--broker-list") && ok {
		commandSlice = append(commandSlice, args[0], "--broker-list", es.BkList)
		cmd = strings.Join(append(commandSlice, args[1:]...), " ")
	} else if _, ok := zkCommand[trimedExe]; !strings.Contains(cmd, "--zookeeper") && ok {
		commandSlice = append(commandSlice, args[0], "--zookeeper", es.Zookeeper)
		cmd = strings.Join(append(commandSlice, args[1:]...), " ")
	}
	return cmd
}

func (es *ExecutorSession) updateServers(cmd string) {
	es.mu.Lock()
	defer es.mu.Unlock()
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
