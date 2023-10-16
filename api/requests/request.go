package requests

type Request struct {
	Directory string `json:"directory"`
	Word      string `json:"word"`
	Case      bool   `json:"case"`
	Whole     bool   `json:"whole"`
}
