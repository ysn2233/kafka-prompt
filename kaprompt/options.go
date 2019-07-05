package kaprompt

import (
	"log"

	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
)

func CustomOptions() []prompt.Option {
	histories, err := LoadHistory()
	if err != nil {
		log.Printf("Load history fails! %s", err)
	}
	options := []prompt.Option{
		prompt.OptionTitle("kafka-prompt: easy-use kafka command client"),
		prompt.OptionPrefix("kafka-prompt λ > "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionSuggestionBGColor(prompt.Purple),
		prompt.OptionDescriptionBGColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.Blue),
		prompt.OptionSelectedDescriptionBGColor(prompt.Purple),
		prompt.OptionSelectedSuggestionTextColor(prompt.Red),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
		prompt.OptionHistory(histories),
	}
	return options
}
