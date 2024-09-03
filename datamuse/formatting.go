package datamuse

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	maxDefsAmount = 10
)

// HTML style
func bold(str string) string {
	return fmt.Sprintf("<b>%s</b>", str)
}

// HTML style
func italic(str string) string {
	return fmt.Sprintf("<i>%s</i>", str)
}

func fullPOS(str string) string {
	switch str {
	case "n":
		return "noun"
	case "adj":
		return "adjective"
	case "adv":
		return "adverb"
	case "v":
		return "verb"
	default:
		return str
	}
}

func formatDefinition(def string) string {
	pattern := `^(\w+)(\s*\(([^)]+)\))?\s*(.*)`

	re, err := regexp.Compile(pattern)
	if err != nil {
		return ""
	}

	match := re.FindStringSubmatch(def)
	if match == nil {
		return ""
	}

	partOfSpeech := fullPOS(match[1])
	additionalInfo := match[3]
	remainingText := match[4]

	result := "➢  "
	result += remainingText
	if partOfSpeech != "" {
		result += strings.Repeat(" ", 4) + italic(partOfSpeech)
	}
	if additionalInfo != "" {
		result += fmt.Sprintf("  (%s)", italic(additionalInfo))
	}
	return result
}

func formatDefs(defs []string) string {
	var result strings.Builder
	for i, def := range defs {
		result.WriteString(formatDefinition(def))
		if i == len(defs)-1 || i == maxDefsAmount {
			break
		}
		result.WriteString("\n")
	}
	return result.String()
}

func (entry Entry) formatAsDescription() string {
	result := fmt.Sprintf(
		"<b>Word:</b> %s\n"+
			"<b>Pronunciation:</b> %s\n"+
			"<b>Definitions:</b>\n%s",
		entry.Word, entry.Tag.getPronunciation(), formatDefs(entry.Defs),
	)
	return result
}

func (resp ResponseData) FormatAsDescription() string {
	var result strings.Builder
	for _, entry := range resp.Entries {
		result.WriteString(entry.formatAsDescription())
	}
	return result.String()
}

func (resp ResponseData) FormatAsWordlist() string {
	var result strings.Builder
	for i, entry := range resp.Entries {
		result.WriteString(entry.Word)
		if i != len(resp.Entries)-1 {
			result.WriteString(", ")
		}
	}
	return result.String()
}
