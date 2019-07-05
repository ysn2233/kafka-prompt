package kaprompt

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
)

type Completer struct {
}

func NewCompleter() *Completer {
	return &Completer{}
}

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")
	w := d.GetWordBeforeCursor()
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	fileCompleter := completer.FilePathCompleter{
		IgnoreCase: true,
		Filter:     nil,
	}

	if len(args) <= 1 {
		s = commandSuggest
		s = prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
		s = append(fileCompleter.Complete(d), s...)
		return s
	}

	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(w, "-") {
		return c.optionCompleter(args, strings.HasPrefix(w, "--"), d)
	}

	// If not starting a word yet, do not activate the filePathCompleter
	if d.GetCharRelativeToCursor(0) == ' ' {
		return s
	}

	return fileCompleter.Complete(d)
}

func (c *Completer) optionCompleter(args []string, hasDoubleDash bool, d prompt.Document) []prompt.Suggest {
	numArgs := len(args)
	if numArgs <= 1 {
		return []prompt.Suggest{}
	}
	cmd := args[0]
	s, ok := optionSuggest[cmd]
	if !ok {
		binPath := strings.Split(cmd, string(os.PathSeparator))
		cmd = binPath[len(binPath)-1]
		cmd = strings.TrimSuffix(cmd, ".sh")
		s = optionSuggest[cmd]
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
