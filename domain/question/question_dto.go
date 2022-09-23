package question

// Question ...
type Question struct {
	Id       int    `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}
