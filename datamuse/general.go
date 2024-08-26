package datamuse

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	baseURL = "https://api.datamuse.com/words?"
)

type RequestConfig struct {
	params    []string
	metaFlags []string
}

func NewEmptyRequest() *RequestConfig {
	req := RequestConfig{nil, nil}
	return &req
}

func NewCertainWordRequest(word string) *RequestConfig {
	var req RequestConfig
	req.params = append(req.params, "sp="+word)
	req.params = append(req.params, "qe=sp")
	req.params = append(req.params, "max=1")
	return &req
}

func (req *RequestConfig) AddMeaningConstraint(word string) *RequestConfig {
	req.params = append(req.params, "ml="+word)
	return req
}

func (req *RequestConfig) AddSynonymConstraint(word string) *RequestConfig {
	req.params = append(req.params, "rel_syn="+word)
	return req
}

func (req *RequestConfig) AddAntonymConstraint(word string) *RequestConfig {
	req.params = append(req.params, "rel_ant="+word)
	return req
}

func (req *RequestConfig) AddHypernymConstraint(word string) *RequestConfig {
	req.params = append(req.params, "rel_spc="+word)
	return req
}

func (req *RequestConfig) AddHyponymConstraint(word string) *RequestConfig {
	req.params = append(req.params, "rel_gen="+word)
	return req
}

func (req *RequestConfig) AddHolonymConstraint(word string) *RequestConfig {
	req.params = append(req.params, "rel_com="+word)
	return req
}

func (req *RequestConfig) AddMeronymConstraint(word string) *RequestConfig {
	req.params = append(req.params, "rel_par="+word)
	return req
}

func (req *RequestConfig) AddDefinitionInfo() *RequestConfig {
	req.metaFlags = append(req.metaFlags, "d")
	return req
}

func (req *RequestConfig) AddPronunciationInfo() *RequestConfig {
	req.metaFlags = append(req.metaFlags, "r")
	req.params = append(req.params, "ipa=1")
	return req
}

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
