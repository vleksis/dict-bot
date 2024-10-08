package telegram

import "dict-bot/datamuse"

type botCommand struct {
	name             string
	shortDescription string
	longDescription  string // for /help command
	handler          func(args string) string
}

var (
	availableCommands = [...]botCommand{
		{
			name:             "lookup",
			shortDescription: "gives general information about the word",
			handler: func(args string) string {
				if args == "" {
					return "invalid command. You need to provide a word"
				}
				req := datamuse.NewCertainWordRequest(args).
					AddDefinitionInfo().
					AddPronunciationInfo()
				resp := datamuse.MakeRequest(req)
				return resp.FormatAsDescription()
			},
		},
		{
			name:             "synonyms",
			shortDescription: "get synonyms to the following word",
			handler: func(args string) string {
				if args == "" {
					return "invalid command. You need to provide a word"
				}
				req := datamuse.NewEmptyRequest().
					AddSynonymConstraint(args)
				resp := datamuse.MakeRequest(req)
				return resp.FormatAsWordlist()
			},
		},
		{
			name:             "antonyms",
			shortDescription: "get antonyms to the following word",
			handler: func(args string) string {
				if args == "" {
					return "invalid command. You need to provide a word"
				}
				req := datamuse.NewEmptyRequest().
					AddAntonymConstraint(args)
				resp := datamuse.MakeRequest(req)
				return resp.FormatAsWordlist()
			},
		},
		{
			name:             "means",
			shortDescription: "return results that have a meaning related to the following word",
			handler: func(args string) string {
				if args == "" {
					return "invalid command. You need to provide a word"
				}
				req := datamuse.NewEmptyRequest().
					AddMeaningConstraint(args)
				resp := datamuse.MakeRequest(req)
				return resp.FormatAsWordlist()
			},
		},
		{
			name:             "hypernym",
			shortDescription: "/hypernym <word> returns the hypernym to the word",
			handler: func(args string) string {
				if args == "" {
					return "invalid command. You need to provide a word"
				}
				req := datamuse.NewEmptyRequest().
					AddHypernymConstraint(args)
				resp := datamuse.MakeRequest(req)
				return resp.FormatAsWordlist()
			},
		},
		{
			name:             "hyponym",
			shortDescription: "/hyponym <word> returns the hyponym to the word",
			handler: func(args string) string {
				if args == "" {
					return "invalid command. You need to provide a word"
				}
				req := datamuse.NewEmptyRequest().
					AddHyponymConstraint(args)
				resp := datamuse.MakeRequest(req)
				return resp.FormatAsWordlist()
			},
		},
		{
			name:             "holonym",
			shortDescription: "/holonym <word> returns the holonym to the word",
			handler: func(args string) string {
				if args == "" {
					return "invalid command. You need to provide a word"
				}
				req := datamuse.NewEmptyRequest().
					AddHolonymConstraint(args)
				resp := datamuse.MakeRequest(req)
				return resp.FormatAsWordlist()
			},
		},
		{
			name:             "meronym",
			shortDescription: "/meronym <word> returns the meronym to the word",
			handler: func(args string) string {
				if args == "" {
					return "invalid command. You need to provide a word"
				}
				req := datamuse.NewEmptyRequest().
					AddMeronymConstraint(args)
				resp := datamuse.MakeRequest(req)
				return resp.FormatAsWordlist()
			},
		},
		// TODO
		{
			name:             "help",
			shortDescription: "todo",
		},
	}
)
