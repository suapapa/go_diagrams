package diagrams

type Result struct {
	Msg string `json:"msg,omitempty"`
	Err string `json:"err,omitempty"`
	Img string `json:"img,omitempty"`
}
