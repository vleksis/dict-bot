package datamuse

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type ResponseData struct {
	Entries []Entry
	Error   error
}

type Entry struct {
	Word  string   `json:"word"`
	Score int      `json:"score"`
	Tag   Tags     `json:"tags"`
	Defs  []string `json:"defs"`
}

type Tags []string

func (t Tags) getPronunciation() string {
	for _, tag := range t {
		if len(tag) > 9 && tag[:9] == "ipa_pron:" {
			return tag[9:]
		}
	}
	return ""
}

func MakeRequest(req *RequestConfig) ResponseData {
	url := baseURL + strings.Join(req.params, "&")
	if len(req.metaFlags) > 0 {
		url += "&md=" + strings.Join(req.metaFlags, "")
	}
	resp, err := http.Get(url)
	if err != nil {
		return ResponseData{nil, err}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseData{nil, err}
	}
	log.Println(body)
	var result ResponseData
	result.Error = json.Unmarshal(body, &result.Entries)
	return result
}

func (resp ResponseData) String() string {
	if resp.Error != nil {
		return resp.Error.Error()
	}
	var result strings.Builder
	for i, entry := range resp.Entries {
		if i > 0 {
			result.WriteString(", ")
		}
		result.WriteString(entry.Word)
		result.WriteString(":\n")
		result.WriteString(entry.Tag.getPronunciation())
		result.WriteString("\n")
		result.WriteString(strings.Join(entry.Defs, "\n"))
	}
	return result.String()
}
