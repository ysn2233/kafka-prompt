package kaprompt

import (
	"bufio"
	"os"
	"os/user"
)

const (
	MAX_HISTORIES = 25
)

func initHistoryFile() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(user.HomeDir+"/.kaprompt", 0700)
	if err != nil {
		return "", err
	}
	return user.HomeDir + "/.kaprompt/history", nil

}

func Persist(record string) error {
	history_file, err := initHistoryFile()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(history_file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(record + "\n"); err != nil {
		return err
	}
	return nil
}

func LoadHistory() ([]string, error) {
	history_file, err := initHistoryFile()
	if err != nil {
		return []string{}, err
	}
	if _, err := os.Stat(history_file); err == nil {
		f, err := os.Open(history_file)
		if err != nil {
			return []string{}, err
		}
		defer f.Close()
		var histories []string
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			histories = append(histories, scanner.Text())
		}
		histories, err = clearHistory(histories, history_file)
		if err != nil {
			return histories, err
		}
		return histories, scanner.Err()
	} else {
		return []string{}, nil
	}
}

func clearHistory(histories []string, filename string) ([]string, error) {
	if len(histories) > MAX_HISTORIES {
		cleared := histories[len(histories)-MAX_HISTORIES:]
		f, err := os.OpenFile(filename, os.O_WRONLY, 0644)
		if err != nil {
			return histories, err
		}
		f.Truncate(0)
		for _, s := range cleared {
			f.WriteString(s + "\n")
		}
		return cleared, nil
	}
	return histories, nil
}
