package datamuse

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
